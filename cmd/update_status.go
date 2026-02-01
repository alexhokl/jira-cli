package cmd

import (
	"errors"
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type updateStatusOptions struct {
	id          string
	name        string
	category    string
	description string
}

var updateStatusOpts = updateStatusOptions{}

var updateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Update a status",
	Long: `Update a status by its ID.

You must provide the status ID and at least one field to update (name, category, or description).

Valid status categories are:
  - TODO        (To Do)
  - IN_PROGRESS (In Progress)
  - DONE        (Done)

Use 'jira-cli list status' to find status IDs and current values.

Note: Requires 'Administer projects' or 'Administer Jira' permission.

Examples:
  # Update only the status name
  jira-cli update status --id 10001 --name "In Review"

  # Update only the category
  jira-cli update status --id 10001 --category DONE

  # Update only the description
  jira-cli update status --id 10001 --description "Code review in progress"

  # Update multiple fields
  jira-cli update status --id 10001 --name "In Review" --category IN_PROGRESS --description "Under review"`,
	RunE: runUpdateStatus,
}

func init() {
	updateCmd.AddCommand(updateStatusCmd)

	flags := updateStatusCmd.Flags()
	flags.StringVarP(&updateStatusOpts.id, "id", "i", "", "Status ID (required)")
	flags.StringVarP(&updateStatusOpts.name, "name", "n", "", "New status name")
	flags.StringVarP(&updateStatusOpts.category, "category", "c", "", "Status category: TODO, IN_PROGRESS, DONE")
	flags.StringVarP(&updateStatusOpts.description, "description", "d", "", "Status description")
	updateStatusCmd.MarkFlagRequired("id")
}

func runUpdateStatus(_ *cobra.Command, _ []string) error {
	// Check that at least one update field is provided
	if updateStatusOpts.name == "" && updateStatusOpts.category == "" && updateStatusOpts.description == "" {
		return fmt.Errorf("at least one of --name, --category, or --description must be provided")
	}

	client := newClient()
	ctx := getAuthContext()

	// Validate category if provided
	if updateStatusOpts.category != "" {
		validCategories := map[string]bool{
			"TODO":        true,
			"IN_PROGRESS": true,
			"DONE":        true,
		}
		if !validCategories[updateStatusOpts.category] {
			return fmt.Errorf("invalid category %q. Valid categories are: TODO, IN_PROGRESS, DONE", updateStatusOpts.category)
		}
	}

	// Fetch current status to get existing values
	statuses, _, err := client.StatusAPI.GetStatusesById(ctx).
		Id([]string{updateStatusOpts.id}).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get status: %w", err)
	}

	if len(statuses) == 0 {
		return fmt.Errorf("status with ID %s not found", updateStatusOpts.id)
	}

	currentStatus := statuses[0]

	// Use current values as defaults, override with provided options
	name := currentStatus.GetName()
	if updateStatusOpts.name != "" {
		name = updateStatusOpts.name
	}

	category := currentStatus.GetStatusCategory()
	if updateStatusOpts.category != "" {
		category = updateStatusOpts.category
	}

	description := currentStatus.GetDescription()
	if updateStatusOpts.description != "" {
		description = updateStatusOpts.description
	}

	statusUpdate := swagger.NewStatusUpdate(
		updateStatusOpts.id,
		name,
		category,
	)
	statusUpdate.SetDescription(description)

	updateRequest := swagger.NewStatusUpdateRequest([]swagger.StatusUpdate{*statusUpdate})

	_, _, err = client.StatusAPI.UpdateStatuses(ctx).
		StatusUpdateRequest(*updateRequest).
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

	fmt.Printf("Status %s has been updated\n", updateStatusOpts.id)

	return nil
}
