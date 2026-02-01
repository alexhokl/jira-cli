package cmd

import (
	"fmt"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type getWorkflowOptions struct {
	name string
}

var getWorkflowOpts = getWorkflowOptions{}

var getWorkflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Get workflow details",
	Long: `Get details of a specific workflow including its statuses and transitions.

Examples:
  # Get workflow by name
  jira-cli get workflow --name "Software Simplified Workflow"`,
	RunE: runGetWorkflow,
}

func init() {
	getCmd.AddCommand(getWorkflowCmd)

	flags := getWorkflowCmd.Flags()
	flags.StringVarP(&getWorkflowOpts.name, "name", "n", "", "Workflow name (required)")
	getWorkflowCmd.MarkFlagRequired("name")
}

func runGetWorkflow(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// Build the request to read workflow by name
	request := swagger.NewWorkflowReadRequest()
	request.WorkflowNames = []string{getWorkflowOpts.name}

	result, _, err := client.WorkflowsAPI.ReadWorkflows(ctx).
		WorkflowReadRequest(*request).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get workflow: %w", err)
	}

	workflows := result.GetWorkflows()
	if len(workflows) == 0 {
		return fmt.Errorf("workflow %q not found", getWorkflowOpts.name)
	}

	workflow := workflows[0]

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	// Display workflow header
	fmt.Printf("%s\n", yellow(workflow.GetName()))
	fmt.Println(strings.Repeat("-", len(workflow.GetName())))

	if desc := workflow.GetDescription(); desc != "" {
		fmt.Printf("\n%s\n", desc)
	}

	// Build a map of status references to status details
	statusMap := make(map[string]swagger.JiraWorkflowStatus)
	for _, status := range result.GetStatuses() {
		if ref := status.GetStatusReference(); ref != "" {
			statusMap[ref] = status
		}
	}

	// Display statuses
	statuses := workflow.GetStatuses()
	if len(statuses) > 0 {
		fmt.Printf("\n%s\n", cyan("Statuses:"))
		for _, statusRef := range statuses {
			ref := statusRef.GetStatusReference()
			if status, ok := statusMap[ref]; ok {
				category := status.GetStatusCategory()
				categoryStr := ""
				if category != "" {
					categoryStr = fmt.Sprintf(" [%s]", magenta(category))
				}
				fmt.Printf("  %s%s\n", yellow(status.GetName()), categoryStr)
			} else {
				fmt.Printf("  %s\n", ref)
			}
		}
	}

	// Display transitions
	transitions := workflow.GetTransitions()
	if len(transitions) > 0 {
		fmt.Printf("\n%s\n", cyan("Transitions:"))
		for _, transition := range transitions {
			toRef := transition.GetToStatusReference()
			toStatusName := toRef
			if status, ok := statusMap[toRef]; ok {
				toStatusName = status.GetName()
			}

			// Get from statuses
			links := transition.GetLinks()
			var fromStatuses []string
			for _, link := range links {
				fromRef := link.GetFromStatusReference()
				if status, ok := statusMap[fromRef]; ok {
					fromStatuses = append(fromStatuses, status.GetName())
				} else if fromRef != "" {
					fromStatuses = append(fromStatuses, fromRef)
				}
			}

			fromStr := "Any Status"
			if len(fromStatuses) > 0 {
				fromStr = strings.Join(fromStatuses, ", ")
			}

			fmt.Printf("  %s: %s → %s\n", yellow(transition.GetName()), fromStr, magenta(toStatusName))
		}
	}

	return nil
}
