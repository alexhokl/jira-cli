package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type bulkUpdateDeleteLabelOptions struct {
	name    string
	project string
	dryRun  bool
}

var bulkUpdateDeleteLabelOpts = bulkUpdateDeleteLabelOptions{}

var bulkUpdateDeleteLabelCmd = &cobra.Command{
	Use:   "delete-label",
	Short: "Delete a label from all issues",
	Long: `Search for all issues with a specific label and remove it.

This command will:
1. Search for all issues that have the specified label
2. For each issue, remove the label while keeping all other labels

Examples:
  # Delete label 'obsolete' from all issues
  jira-cli bulk-update delete-label --name obsolete

  # Delete label only in a specific project
  jira-cli bulk-update delete-label --name obsolete --project MYPROJ

  # Preview changes without applying them
  jira-cli bulk-update delete-label --name obsolete --dry-run`,
	RunE: runBulkUpdateDeleteLabel,
}

func init() {
	bulkUpdateCmd.AddCommand(bulkUpdateDeleteLabelCmd)

	flags := bulkUpdateDeleteLabelCmd.Flags()
	flags.StringVar(&bulkUpdateDeleteLabelOpts.name, "name", "", "Label name to delete (required)")
	flags.StringVarP(&bulkUpdateDeleteLabelOpts.project, "project", "p", "", "Filter by project key")
	flags.BoolVar(&bulkUpdateDeleteLabelOpts.dryRun, "dry-run", false, "Preview changes without applying them")

	bulkUpdateDeleteLabelCmd.MarkFlagRequired("name")
}

func runBulkUpdateDeleteLabel(_ *cobra.Command, _ []string) error {
	labelName := strings.TrimSpace(bulkUpdateDeleteLabelOpts.name)

	if labelName == "" {
		return fmt.Errorf("label name cannot be empty")
	}

	client := newClient()
	ctx := getAuthContext()

	// Build JQL to find issues with the label
	jql := fmt.Sprintf("labels = %s", quoteJQLValue(labelName))
	if bulkUpdateDeleteLabelOpts.project != "" {
		jql = fmt.Sprintf("project = %s AND %s", quoteJQLValue(bulkUpdateDeleteLabelOpts.project), jql)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	searchMsg := fmt.Sprintf("Searching for issues with label '%s'", cyan(labelName))
	if bulkUpdateDeleteLabelOpts.project != "" {
		searchMsg += fmt.Sprintf(" in project '%s'", cyan(bulkUpdateDeleteLabelOpts.project))
	}
	fmt.Println(searchMsg + "...")

	// Collect all issues with the label
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
		fmt.Printf("No issues found with label '%s'\n", labelName)
		return nil
	}

	fmt.Printf("Found %d issue(s) with label '%s'\n\n", len(allIssues), cyan(labelName))

	if bulkUpdateDeleteLabelOpts.dryRun {
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
		fmt.Printf("\nTotal: %d issue(s) would have label '%s' removed\n",
			len(allIssues), cyan(labelName))
		return nil
	}

	// Update each issue: remove the label
	successCount := 0
	failCount := 0

	for _, issue := range allIssues {
		key := issue.GetKey()
		fields := issue.GetFields()
		summary := ""
		if summaryObj, ok := fields["summary"]; ok && summaryObj != nil {
			summary = fmt.Sprintf("%v", summaryObj)
		}

		// Create update operation: remove the label
		var labelOps []swagger.FieldUpdateOperation

		removeOp := swagger.NewFieldUpdateOperation()
		removeOp.SetRemove(labelName)
		labelOps = append(labelOps, *removeOp)

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
	fmt.Printf("Label '%s' has been removed from %d issue(s)\n", cyan(labelName), successCount)

	return nil
}
