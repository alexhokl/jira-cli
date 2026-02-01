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

type listWorkflowStatusPropertiesOptions struct {
	name string
}

var listWorkflowStatusPropertiesOpts = listWorkflowStatusPropertiesOptions{}

var listWorkflowStatusPropertiesCmd = &cobra.Command{
	Use:   "workflow-status-properties",
	Short: "List status properties in a workflow",
	Long: `List all status properties associated with statuses in a workflow.

Workflow status properties are key-value pairs that can be used to configure
behavior for specific statuses within a workflow.

Examples:
  # List status properties in a workflow
  jira-cli list workflow-status-properties -n "Software Simplified Workflow for Project PROJ"`,
	RunE: runListWorkflowStatusProperties,
}

func init() {
	listCmd.AddCommand(listWorkflowStatusPropertiesCmd)

	flags := listWorkflowStatusPropertiesCmd.Flags()
	flags.StringVarP(&listWorkflowStatusPropertiesOpts.name, "name", "n", "", "Workflow name")

	listWorkflowStatusPropertiesCmd.MarkFlagRequired("name")
}

func runListWorkflowStatusProperties(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// Build the request to get workflow by name
	request := swagger.NewWorkflowReadRequest()
	request.SetWorkflowNames([]string{listWorkflowStatusPropertiesOpts.name})

	result, _, err := client.WorkflowsAPI.ReadWorkflows(ctx).
		WorkflowReadRequest(*request).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get workflow: %w", err)
	}

	if len(result.GetWorkflows()) == 0 {
		return fmt.Errorf("workflow %q not found", listWorkflowStatusPropertiesOpts.name)
	}

	workflow := result.GetWorkflows()[0]

	// Build a map from statusReference to status name
	statusRefToName := make(map[string]string)
	for _, s := range result.GetStatuses() {
		statusRefToName[s.GetStatusReference()] = s.GetName()
	}

	// Collect status properties
	type statusPropertyInfo struct {
		statusName string
		key        string
		value      string
	}

	var properties []statusPropertyInfo

	for _, s := range workflow.GetStatuses() {
		statusRef := s.GetStatusReference()
		statusName := statusRefToName[statusRef]
		if statusName == "" {
			statusName = statusRef // fallback to reference if name not found
		}

		props := s.GetProperties()
		for key, value := range props {
			properties = append(properties, statusPropertyInfo{
				statusName: statusName,
				key:        key,
				value:      value,
			})
		}
	}

	if len(properties) == 0 {
		fmt.Printf("No status properties found in workflow %q\n", listWorkflowStatusPropertiesOpts.name)
		return nil
	}

	// Sort by status name, then by key
	sort.Slice(properties, func(i, j int) bool {
		if strings.ToLower(properties[i].statusName) != strings.ToLower(properties[j].statusName) {
			return strings.ToLower(properties[i].statusName) < strings.ToLower(properties[j].statusName)
		}
		return strings.ToLower(properties[i].key) < strings.ToLower(properties[j].key)
	})

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("Status properties in workflow %q:\n\n", listWorkflowStatusPropertiesOpts.name)

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("STATUS"), cyan("PROPERTY"), cyan("VALUE"))

	for _, p := range properties {
		// Truncate value if too long
		value := p.value
		if len(value) > 60 {
			value = value[:57] + "..."
		}
		w.row(yellow(p.statusName), p.key, value)
	}
	w.flush()

	fmt.Printf("\nFound %d properties\n", len(properties))

	return nil
}
