package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type updateWorkflowStatusPropertyOptions struct {
	workflowName string
	statusName   string
	property     string
	value        string
}

var updateWorkflowStatusPropertyOpts = updateWorkflowStatusPropertyOptions{}

var updateWorkflowStatusPropertyCmd = &cobra.Command{
	Use:   "workflow-status-property",
	Short: "Update a workflow status property",
	Long: `Update or add a property on a status within a workflow.

Workflow status properties are key-value pairs that can be used to configure
behavior for specific statuses within a workflow (e.g., jira.issue.editable).

Note: This command modifies the workflow. Only use on workflows that are editable.

Examples:
  # Set a status property
  jira-cli update workflow-status-property -n "My Workflow" -s "Done" -p "jira.issue.editable" -v "false"

  # Update an existing property
  jira-cli update workflow-status-property -n "My Workflow" -s "In Progress" -p "jira.issue.editable" -v "true"`,
	RunE: runUpdateWorkflowStatusProperty,
}

func init() {
	updateCmd.AddCommand(updateWorkflowStatusPropertyCmd)

	flags := updateWorkflowStatusPropertyCmd.Flags()
	flags.StringVarP(&updateWorkflowStatusPropertyOpts.workflowName, "name", "n", "", "Workflow name (required)")
	flags.StringVarP(&updateWorkflowStatusPropertyOpts.statusName, "status", "s", "", "Status name (required)")
	flags.StringVarP(&updateWorkflowStatusPropertyOpts.property, "property", "p", "", "Property key (required)")
	flags.StringVarP(&updateWorkflowStatusPropertyOpts.value, "value", "v", "", "Property value (required)")

	updateWorkflowStatusPropertyCmd.MarkFlagRequired("name")
	updateWorkflowStatusPropertyCmd.MarkFlagRequired("status")
	updateWorkflowStatusPropertyCmd.MarkFlagRequired("property")
	updateWorkflowStatusPropertyCmd.MarkFlagRequired("value")
}

func runUpdateWorkflowStatusProperty(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// Read the workflow
	readRequest := swagger.NewWorkflowReadRequest()
	readRequest.SetWorkflowNames([]string{updateWorkflowStatusPropertyOpts.workflowName})

	readResult, _, err := client.WorkflowsAPI.ReadWorkflows(ctx).
		WorkflowReadRequest(*readRequest).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get workflow: %w", err)
	}

	if len(readResult.GetWorkflows()) == 0 {
		return fmt.Errorf("workflow %q not found", updateWorkflowStatusPropertyOpts.workflowName)
	}

	workflow := readResult.GetWorkflows()[0]

	// Check if workflow is editable
	if !workflow.GetIsEditable() {
		return fmt.Errorf("workflow %q is not editable", updateWorkflowStatusPropertyOpts.workflowName)
	}

	// Build a map from status name to status reference
	statusNameToRef := make(map[string]string)
	for _, s := range readResult.GetStatuses() {
		statusNameToRef[strings.ToLower(s.GetName())] = s.GetStatusReference()
	}

	// Find the target status reference
	targetStatusRef := statusNameToRef[strings.ToLower(updateWorkflowStatusPropertyOpts.statusName)]
	if targetStatusRef == "" {
		return fmt.Errorf("status %q not found in workflow %q", updateWorkflowStatusPropertyOpts.statusName, updateWorkflowStatusPropertyOpts.workflowName)
	}

	// Build the status layout updates, modifying the target status property
	var statusLayoutUpdates []swagger.StatusLayoutUpdate
	statusFound := false

	for _, s := range workflow.GetStatuses() {
		props := s.GetProperties()
		if props == nil {
			props = make(map[string]string)
		}

		statusRef := s.GetStatusReference()

		if statusRef == targetStatusRef {
			statusFound = true
			// Update the property on this status
			props[updateWorkflowStatusPropertyOpts.property] = updateWorkflowStatusPropertyOpts.value
		}

		statusUpdate := swagger.NewStatusLayoutUpdate(props, statusRef)

		// Preserve layout if set (convert WorkflowStatusLayout to WorkflowLayout)
		if s.HasLayout() {
			oldLayout := s.GetLayout()
			newLayout := swagger.NewWorkflowLayout()
			if oldLayout.X.IsSet() {
				newLayout.SetX(oldLayout.GetX())
			}
			if oldLayout.Y.IsSet() {
				newLayout.SetY(oldLayout.GetY())
			}
			statusUpdate.SetLayout(*newLayout)
		}

		// Preserve approval configuration if set
		if s.HasApprovalConfiguration() {
			statusUpdate.SetApprovalConfiguration(s.GetApprovalConfiguration())
		}

		statusLayoutUpdates = append(statusLayoutUpdates, *statusUpdate)
	}

	if !statusFound {
		return fmt.Errorf("status %q not found in workflow statuses", updateWorkflowStatusPropertyOpts.statusName)
	}

	// Build transition updates from existing transitions
	var transitionUpdates []swagger.TransitionUpdateDTO

	for _, t := range workflow.GetTransitions() {
		transitionUpdate := swagger.NewTransitionUpdateDTO()

		if t.HasId() {
			transitionUpdate.SetId(t.GetId())
		}
		if t.HasName() {
			transitionUpdate.SetName(t.GetName())
		}
		if t.HasDescription() {
			transitionUpdate.SetDescription(t.GetDescription())
		}
		if t.HasType() {
			transitionUpdate.SetType(t.GetType())
		}
		if t.HasToStatusReference() {
			transitionUpdate.SetToStatusReference(t.GetToStatusReference())
		}
		if t.HasLinks() {
			transitionUpdate.SetLinks(t.GetLinks())
		}
		if t.HasProperties() {
			transitionUpdate.SetProperties(t.GetProperties())
		}
		if t.HasActions() {
			transitionUpdate.SetActions(t.GetActions())
		}
		if t.HasValidators() {
			transitionUpdate.SetValidators(t.GetValidators())
		}
		if t.HasTriggers() {
			transitionUpdate.SetTriggers(t.GetTriggers())
		}

		transitionUpdates = append(transitionUpdates, *transitionUpdate)
	}

	// Build the workflow update
	workflowUpdate := swagger.NewWorkflowUpdate(
		workflow.GetId(),
		statusLayoutUpdates,
		transitionUpdates,
		*workflow.Version,
	)

	if workflow.HasDescription() {
		workflowUpdate.SetDescription(workflow.GetDescription())
	}

	// Build the update request
	updateRequest := swagger.NewWorkflowUpdateRequest()
	updateRequest.SetWorkflows([]swagger.WorkflowUpdate{*workflowUpdate})

	_, _, err = client.WorkflowsAPI.UpdateWorkflows(ctx).
		WorkflowUpdateRequest(*updateRequest).
		Execute()
	if err != nil {
		var openAPIErr *swagger.GenericOpenAPIError
		if errors.As(err, &openAPIErr) {
			body := string(openAPIErr.Body())
			if body != "" {
				return fmt.Errorf("%s\n\n%s", err, body)
			}
		}
		return err
	}

	fmt.Printf("Updated property %q on status %q in workflow %q\n",
		updateWorkflowStatusPropertyOpts.property,
		updateWorkflowStatusPropertyOpts.statusName,
		updateWorkflowStatusPropertyOpts.workflowName)

	return nil
}
