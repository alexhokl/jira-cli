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
	name   string
	format string
}

var listWorkflowStatusOpts = listWorkflowStatusOptions{}

var listWorkflowStatusCmd = &cobra.Command{
	Use:     "workflow-status",
	Aliases: []string{"workflow-statuses"},
	Short:   "List statuses in a workflow",
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
	flags.StringVar(&listWorkflowStatusOpts.format, "format", "table", "Output format: table or json")

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

	// Build typed slice for output
	var items []workflowStatusInfo
	for _, s := range statuses {
		items = append(items, workflowStatusInfo{
			ID:          s.GetId(),
			Name:        s.GetName(),
			Category:    s.GetStatusCategory(),
			Description: s.GetDescription(),
		})
	}

	// Sort statuses by name
	sort.Slice(items, func(i, j int) bool {
		return strings.ToLower(items[i].Name) < strings.ToLower(items[j].Name)
	})

	if listWorkflowStatusOpts.format == "json" {
		return printJSON(items)
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("Statuses in workflow %q:\n\n", listWorkflowStatusOpts.name)

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("CATEGORY"), cyan("DESCRIPTION"))

	for _, s := range items {
		// Truncate description if too long
		desc := s.Description
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		w.row(s.ID, yellow(s.Name), s.Category, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d statuses\n", len(items))

	return nil
}

type workflowStatusInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
