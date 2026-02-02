package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listWorkflowsOptions struct {
	query    string
	isActive *bool
}

var listWorkflowsOpts = listWorkflowsOptions{}

var listWorkflowsCmd = &cobra.Command{
	Use:   "workflows",
	Short: "List workflows",
	Long: `List all workflows available in the Jira instance.

Examples:
  # List all workflows
  jira-cli list workflows

  # Search workflows by name
  jira-cli list workflows --query "scrum"

  # List only active workflows
  jira-cli list workflows --active

  # List only inactive workflows
  jira-cli list workflows --inactive`,
	RunE: runListWorkflows,
}

func init() {
	listCmd.AddCommand(listWorkflowsCmd)

	flags := listWorkflowsCmd.Flags()
	flags.StringVarP(&listWorkflowsOpts.query, "query", "q", "", "Filter workflows by name (partial match)")
	flags.Bool("active", false, "Show only active workflows")
	flags.Bool("inactive", false, "Show only inactive workflows")
}

func runListWorkflows(cmd *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// Handle active/inactive flags
	activeFlag, _ := cmd.Flags().GetBool("active")
	inactiveFlag, _ := cmd.Flags().GetBool("inactive")

	var workflows []workflowInfo
	var startAt int64 = 0

	for {
		request := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).
			StartAt(startAt).
			MaxResults(50).
			OrderBy("name")

		if listWorkflowsOpts.query != "" {
			request = request.QueryString(listWorkflowsOpts.query)
		}

		if activeFlag {
			request = request.IsActive(true)
		} else if inactiveFlag {
			request = request.IsActive(false)
		}

		result, _, err := request.Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get workflows: %w", err))
		}

		for _, wf := range result.GetValues() {
			id := wf.GetId()
			workflows = append(workflows, workflowInfo{
				name:        id.GetName(),
				description: wf.GetDescription(),
				isDefault:   wf.GetIsDefault(),
			})
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(workflows) == 0 {
		fmt.Println("No workflows found")
		return nil
	}

	// Sort by name
	sort.Slice(workflows, func(i, j int) bool {
		return strings.ToLower(workflows[i].name) < strings.ToLower(workflows[j].name)
	})

	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("NAME"), cyan("DEFAULT"), cyan("DESCRIPTION"))

	for _, wf := range workflows {
		defaultStr := ""
		if wf.isDefault {
			defaultStr = "Yes"
		}
		// Truncate description if too long
		desc := wf.description
		if len(desc) > 60 {
			desc = desc[:57] + "..."
		}
		w.row(wf.name, defaultStr, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d workflows\n", len(workflows))

	return nil
}

type workflowInfo struct {
	name        string
	description string
	isDefault   bool
}
