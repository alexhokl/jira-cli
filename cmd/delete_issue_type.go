package cmd

import (
	"errors"
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type deleteIssueTypeOptions struct {
	id            string
	alternativeId string
}

var deleteIssueTypeOpts = deleteIssueTypeOptions{}

var deleteIssueTypeCmd = &cobra.Command{
	Use:   "issue-type",
	Short: "Delete an issue type",
	Long: `Delete an issue type by its ID.

If the issue type is in use by existing issues, you must specify an alternative
issue type to migrate those issues to using the --alternative flag.

Use 'jira-cli list issue-types' to find issue type IDs.

Note: Requires 'Administer Jira' global permission.

Examples:
  # Delete an issue type that is not in use
  jira-cli delete issue-type --id 10001

  # Delete an issue type and migrate existing issues to another type
  jira-cli delete issue-type --id 10001 --alternative 10002`,
	RunE: runDeleteIssueType,
}

func init() {
	deleteCmd.AddCommand(deleteIssueTypeCmd)

	flags := deleteIssueTypeCmd.Flags()
	flags.StringVarP(&deleteIssueTypeOpts.id, "id", "i", "", "Issue type ID (required)")
	flags.StringVarP(&deleteIssueTypeOpts.alternativeId, "alternative", "a", "", "Alternative issue type ID to migrate existing issues to")
	deleteIssueTypeCmd.MarkFlagRequired("id")
}

func runDeleteIssueType(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	req := client.IssueTypesAPI.DeleteIssueType(ctx, deleteIssueTypeOpts.id)

	if deleteIssueTypeOpts.alternativeId != "" {
		req = req.AlternativeIssueTypeId(deleteIssueTypeOpts.alternativeId)
	}

	_, err := req.Execute()
	if err != nil {
		// Check if it's an API error and provide helpful guidance
		var openAPIErr *swagger.GenericOpenAPIError
		if errors.As(err, &openAPIErr) {
			body := string(openAPIErr.Body())
			// If no alternative was provided, suggest getting alternatives
			if deleteIssueTypeOpts.alternativeId == "" {
				if body != "" {
					return fmt.Errorf("%s\n\n%s\n\nThe issue type may be in use. Try specifying an alternative issue type with --alternative flag", err, body)
				}
				return fmt.Errorf("%w\n\nThe issue type may be in use. Try specifying an alternative issue type with --alternative flag", err)
			}
			if body != "" {
				return fmt.Errorf("%s\n\n%s", err, body)
			}
		}
		return err
	}

	if deleteIssueTypeOpts.alternativeId != "" {
		fmt.Printf("Issue type %s has been deleted. Existing issues migrated to issue type %s.\n",
			deleteIssueTypeOpts.id, deleteIssueTypeOpts.alternativeId)
	} else {
		fmt.Printf("Issue type %s has been deleted\n", deleteIssueTypeOpts.id)
	}

	return nil
}
