package cmd

import (
	"errors"
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type deleteStatusOptions struct {
	id string
}

var deleteStatusOpts = deleteStatusOptions{}

var deleteStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Delete a status",
	Long: `Delete a status by its ID.

The status must not be in use by any workflow. If the status is in use,
you must first remove it from all workflows before deleting.

Use 'jira-cli list status' to find status IDs.

Note: Requires 'Administer projects' or 'Administer Jira' permission.

Examples:
  # Delete a status
  jira-cli delete status --id 10001`,
	RunE: runDeleteStatus,
}

func init() {
	deleteCmd.AddCommand(deleteStatusCmd)

	flags := deleteStatusCmd.Flags()
	flags.StringVarP(&deleteStatusOpts.id, "id", "i", "", "Status ID (required)")
	deleteStatusCmd.MarkFlagRequired("id")
}

func runDeleteStatus(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	_, _, err := client.StatusAPI.DeleteStatusesById(ctx).
		Id([]string{deleteStatusOpts.id}).
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

	fmt.Printf("Status %s has been deleted\n", deleteStatusOpts.id)

	return nil
}
