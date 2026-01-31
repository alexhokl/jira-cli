package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type deleteProjectOptions struct {
	projectKey string
	enableUndo bool
}

var deleteProjectOpts = deleteProjectOptions{}

var deleteProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Delete a project",
	Long: `Delete a project by its key or ID.

By default, the project is moved to the recycle bin where it can be restored.
Use --permanent to permanently delete the project without the ability to restore.

Note: You cannot delete archived projects. Restore them first using the Jira UI.

Examples:
  # Delete a project (moves to recycle bin)
  jira-cli delete project --key MYPROJ

  # Permanently delete a project
  jira-cli delete project --key MYPROJ --permanent`,
	RunE: runDeleteProject,
}

func init() {
	deleteCmd.AddCommand(deleteProjectCmd)

	flags := deleteProjectCmd.Flags()
	flags.StringVarP(&deleteProjectOpts.projectKey, "key", "k", "", "Project key or ID (required)")
	flags.BoolVar(&deleteProjectOpts.enableUndo, "permanent", false, "Permanently delete (skip recycle bin)")
	deleteProjectCmd.MarkFlagRequired("key")
}

func runDeleteProject(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// enableUndo is the inverse of permanent flag
	enableUndo := !deleteProjectOpts.enableUndo

	_, err := client.ProjectsAPI.DeleteProject(ctx, deleteProjectOpts.projectKey).
		EnableUndo(enableUndo).
		Execute()
	if err != nil {
		return err
	}

	if enableUndo {
		fmt.Printf("Project %s has been deleted and moved to the recycle bin\n", deleteProjectOpts.projectKey)
	} else {
		fmt.Printf("Project %s has been permanently deleted\n", deleteProjectOpts.projectKey)
	}

	return nil
}
