package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type getStatusOptions struct {
	id string
}

var getStatusOpts = getStatusOptions{}

var getStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status",
	Long: `Get details of a specific status by ID.

Examples:
  # Get status details
  jira-cli get status -i 10001`,
	RunE: runGetStatus,
}

func init() {
	getCmd.AddCommand(getStatusCmd)

	flags := getStatusCmd.Flags()
	flags.StringVarP(&getStatusOpts.id, "id", "i", "", "Status ID")

	getStatusCmd.MarkFlagRequired("id")
}

func runGetStatus(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	statuses, _, err := client.StatusAPI.GetStatusesById(ctx).
		Id([]string{getStatusOpts.id}).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get status: %w", err)
	}

	if len(statuses) == 0 {
		return fmt.Errorf("status with ID %s not found", getStatusOpts.id)
	}

	status := statuses[0]

	fmt.Printf("%s %s\n", cyan("ID:"), status.GetId())
	fmt.Printf("%s %s\n", cyan("Name:"), yellow(status.GetName()))
	fmt.Printf("%s %s\n", cyan("Category:"), status.GetStatusCategory())

	// Display scope information
	if status.Scope != nil {
		scope := status.Scope
		fmt.Printf("%s %s\n", cyan("Scope:"), scope.GetType())

		if scope.Project.Get() != nil {
			projectId := scope.Project.Get().GetId()
			// Try to resolve project ID to key
			project, _, err := client.ProjectsAPI.GetProject(ctx, projectId).Execute()
			if err == nil {
				fmt.Printf("%s %s\n", cyan("Project:"), project.GetKey())
			} else {
				fmt.Printf("%s %s\n", cyan("Project ID:"), projectId)
			}
		}
	} else {
		fmt.Printf("%s GLOBAL\n", cyan("Scope:"))
	}

	description := status.GetDescription()
	if description != "" {
		fmt.Printf("%s %s\n", cyan("Description:"), description)
	}

	return nil
}
