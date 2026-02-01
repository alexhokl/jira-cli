package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type bulkUpdateRenameLabelOptions struct {
	oldLabel string
	newLabel string
	project  string
	dryRun   bool
}

var bulkUpdateRenameLabelOpts = bulkUpdateRenameLabelOptions{}

var bulkUpdateRenameLabelCmd = &cobra.Command{
	Use:   "rename-label",
	Short: "Rename a label across all issues",
	Long: `Search for all issues with a specific label and replace it with a new label.

This command will:
1. Search for all issues that have the old label
2. For each issue, remove the old label and add the new label

Examples:
  # Rename label 'bug' to 'defect' across all issues
  jira-cli bulk-update rename-label --old bug --new defect

  # Rename label only in a specific project
  jira-cli bulk-update rename-label --old bug --new defect --project MYPROJ

  # Preview changes without applying them
  jira-cli bulk-update rename-label --old bug --new defect --dry-run`,
	RunE: runBulkUpdateRenameLabel,
}

func init() {
	bulkUpdateCmd.AddCommand(bulkUpdateRenameLabelCmd)

	flags := bulkUpdateRenameLabelCmd.Flags()
	flags.StringVar(&bulkUpdateRenameLabelOpts.oldLabel, "old", "", "Old label name to search for (required)")
	flags.StringVar(&bulkUpdateRenameLabelOpts.newLabel, "new", "", "New label name to replace with (required)")
	flags.StringVarP(&bulkUpdateRenameLabelOpts.project, "project", "p", "", "Filter by project key")
	flags.BoolVar(&bulkUpdateRenameLabelOpts.dryRun, "dry-run", false, "Preview changes without applying them")

	bulkUpdateRenameLabelCmd.MarkFlagRequired("old")
	bulkUpdateRenameLabelCmd.MarkFlagRequired("new")
}

func runBulkUpdateRenameLabel(_ *cobra.Command, _ []string) error {
	oldLabel := strings.TrimSpace(bulkUpdateRenameLabelOpts.oldLabel)
	newLabel := strings.TrimSpace(bulkUpdateRenameLabelOpts.newLabel)

	if oldLabel == "" {
		return fmt.Errorf("old label cannot be empty")
	}
	if newLabel == "" {
		return fmt.Errorf("new label cannot be empty")
	}
	if oldLabel == newLabel {
		return fmt.Errorf("old and new labels are the same")
	}

	client := newClient()
	ctx := getAuthContext()

	// Build JQL to find issues with the old label
	jql := fmt.Sprintf("labels = %s", quoteJQLValue(oldLabel))
	if bulkUpdateRenameLabelOpts.project != "" {
		jql = fmt.Sprintf("project = %s AND %s", quoteJQLValue(bulkUpdateRenameLabelOpts.project), jql)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	searchMsg := fmt.Sprintf("Searching for issues with label '%s'", cyan(oldLabel))
	if bulkUpdateRenameLabelOpts.project != "" {
		searchMsg += fmt.Sprintf(" in project '%s'", cyan(bulkUpdateRenameLabelOpts.project))
	}
	fmt.Println(searchMsg + "...")

	// Collect all issues with the old label
	var allIssues []swagger.IssueBean
	var nextPageToken string
	const pageSize int32 = 100

	for {
		request := client.IssueSearchAPI.SearchAndReconsileIssuesUsingJql(ctx).
			Jql(jql).
			MaxResults(pageSize).
			Fields([]string{"summary", "labels"})

		if nextPageToken != "" {
			request = request.NextPageToken(nextPageToken)
		}

		result, _, err := request.Execute()
		if err != nil {
			return fmt.Errorf("failed to search issues: %w", err)
		}

		allIssues = append(allIssues, result.GetIssues()...)

		if result.GetIsLast() {
			break
		}

		nextPageToken = result.GetNextPageToken()
		if nextPageToken == "" {
			break
		}
	}

	if len(allIssues) == 0 {
		fmt.Printf("No issues found with label '%s'\n", oldLabel)
		return nil
	}

	fmt.Printf("Found %d issue(s) with label '%s'\n\n", len(allIssues), cyan(oldLabel))

	if bulkUpdateRenameLabelOpts.dryRun {
		fmt.Println("DRY RUN - No changes will be made")
		fmt.Println("The following issues would be updated:")
		for _, issue := range allIssues {
			key := issue.GetKey()
			fields := issue.GetFields()
			summary := ""
			if summaryObj, ok := fields["summary"]; ok && summaryObj != nil {
				summary = fmt.Sprintf("%v", summaryObj)
			}
			fmt.Printf("  %s %s\n", yellow(key), summary)
		}
		fmt.Printf("\nTotal: %d issue(s) would have label '%s' replaced with '%s'\n",
			len(allIssues), cyan(oldLabel), green(newLabel))
		return nil
	}

	// Update each issue: remove old label, add new label
	successCount := 0
	failCount := 0

	for _, issue := range allIssues {
		key := issue.GetKey()
		fields := issue.GetFields()
		summary := ""
		if summaryObj, ok := fields["summary"]; ok && summaryObj != nil {
			summary = fmt.Sprintf("%v", summaryObj)
		}

		// Create update operations: remove old label, add new label
		var labelOps []swagger.FieldUpdateOperation

		removeOp := swagger.NewFieldUpdateOperation()
		removeOp.SetRemove(oldLabel)
		labelOps = append(labelOps, *removeOp)

		addOp := swagger.NewFieldUpdateOperation()
		addOp.SetAdd(newLabel)
		labelOps = append(labelOps, *addOp)

		update := map[string][]swagger.FieldUpdateOperation{
			"labels": labelOps,
		}

		issueUpdateDetails := swagger.NewIssueUpdateDetails()
		issueUpdateDetails.SetUpdate(update)

		_, _, err := client.IssuesAPI.EditIssue(ctx, key).
			IssueUpdateDetails(*issueUpdateDetails).
			Execute()

		if err != nil {
			fmt.Printf("  %s %s - %s\n", yellow(key), summary, color.RedString("FAILED: %v", err))
			failCount++
		} else {
			fmt.Printf("  %s %s - %s\n", yellow(key), summary, green("OK"))
			successCount++
		}
	}

	fmt.Printf("\nCompleted: %d succeeded, %d failed\n", successCount, failCount)
	fmt.Printf("Label '%s' has been renamed to '%s'\n", cyan(oldLabel), green(newLabel))

	return nil
}
