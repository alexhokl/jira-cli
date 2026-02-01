package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listWorkflowStatusOptions struct {
	name string
}

var listWorkflowStatusOpts = listWorkflowStatusOptions{}

var listWorkflowStatusCmd = &cobra.Command{
	Use:   "workflow-status",
	Short: "List statuses in a workflow",
	Long: `List all statuses associated with a workflow by workflow name.

Examples:
  # List statuses in a workflow
  jira-cli list workflow-status -n "Software Simplified Workflow for Project PROJ"`,
	RunE: runListWorkflowStatus,
}

func init() {
	listCmd.AddCommand(listWorkflowStatusCmd)

	flags := listWorkflowStatusCmd.Flags()
	flags.StringVarP(&listWorkflowStatusOpts.name, "name", "n", "", "Workflow name")

	listWorkflowStatusCmd.MarkFlagRequired("name")
}

func runListWorkflowStatus(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// Build the request to get workflow by name
	request := swagger.NewWorkflowReadRequest()
	request.SetWorkflowNames([]string{listWorkflowStatusOpts.name})

	result, _, err := client.WorkflowsAPI.ReadWorkflows(ctx).
		WorkflowReadRequest(*request).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get workflow: %w", err)
	}

	if len(result.GetWorkflows()) == 0 {
		return fmt.Errorf("workflow %q not found", listWorkflowStatusOpts.name)
	}

	statuses := result.GetStatuses()
	if len(statuses) == 0 {
		fmt.Println("No statuses found in this workflow")
		return nil
	}

	// Sort statuses by name
	sort.Slice(statuses, func(i, j int) bool {
		return strings.ToLower(statuses[i].GetName()) < strings.ToLower(statuses[j].GetName())
	})

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("Statuses in workflow %q:\n\n", listWorkflowStatusOpts.name)

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("CATEGORY"), cyan("DESCRIPTION"))

	for _, s := range statuses {
		// Truncate description if too long
		desc := s.GetDescription()
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		w.row(s.GetId(), yellow(s.GetName()), s.GetStatusCategory(), desc)
	}
	w.flush()

	fmt.Printf("\nFound %d statuses\n", len(statuses))

	return nil
}
