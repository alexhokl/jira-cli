package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listIssueTransitionsOptions struct {
	id string
}

var listIssueTransitionsOpts = listIssueTransitionsOptions{}

var listIssueTransitionsCmd = &cobra.Command{
	Use:   "issue-transitions",
	Short: "List available transitions for an issue",
	Long: `List all available transitions for a specific issue based on its current status.

This command shows the transitions that can be performed on the issue from its
current workflow status.

Examples:
  # List available transitions for an issue
  jira-cli list issue-transitions --id PROJ-123`,
	RunE: runListIssueTransitions,
}

func init() {
	listCmd.AddCommand(listIssueTransitionsCmd)

	flags := listIssueTransitionsCmd.Flags()
	flags.StringVarP(&listIssueTransitionsOpts.id, "id", "i", "", "Issue ID (e.g., PROJ-123) (required)")
	listIssueTransitionsCmd.MarkFlagRequired("id")
}

func runListIssueTransitions(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// First, get the issue to display current status
	issue, _, err := client.IssuesAPI.GetIssue(ctx, listIssueTransitionsOpts.id).Execute()
	if err != nil {
		return wrapAPIError(fmt.Errorf("failed to get issue: %w", err))
	}

	// Extract current status from issue fields
	currentStatus := "Unknown"
	if statusObj, ok := issue.Fields["status"].(map[string]any); ok {
		if name, ok := statusObj["name"].(string); ok {
			currentStatus = name
		}
	}

	// Get available transitions for the issue
	result, _, err := client.IssuesAPI.GetTransitions(ctx, listIssueTransitionsOpts.id).Execute()
	if err != nil {
		return wrapAPIError(fmt.Errorf("failed to get transitions: %w", err))
	}

	transitions := result.GetTransitions()

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	// Display issue and current status
	fmt.Printf("%s: %s\n", cyan("Issue"), yellow(listIssueTransitionsOpts.id))
	fmt.Printf("%s: %s\n", cyan("Current Status"), magenta(currentStatus))

	if len(transitions) == 0 {
		fmt.Println("\nNo transitions available from this status")
		return nil
	}

	fmt.Printf("\n%s\n", cyan("Available Transitions:"))

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("TO STATUS"), cyan("CATEGORY"))

	for _, transition := range transitions {
		toStatus := transition.GetTo()
		toStatusName := toStatus.GetName()
		category := ""
		if statusCategory, ok := toStatus.GetStatusCategoryOk(); ok && statusCategory != nil {
			category = statusCategory.GetName()
		}

		w.row(
			transition.GetId(),
			yellow(transition.GetName()),
			magenta(toStatusName),
			category,
		)
	}
	w.flush()

	fmt.Printf("\nFound %d available transitions\n", len(transitions))

	return nil
}
