package cmd

import (
	"errors"
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type createStatusOptions struct {
	name        string
	category    string
	description string
	project     string
}

var createStatusOpts = createStatusOptions{}

var createStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Create a status",
	Long: `Create a new status.

You must provide the status name and category. The description is optional.

Valid status categories are:
  - TODO        (To Do)
  - IN_PROGRESS (In Progress)
  - DONE        (Done)

By default, statuses are created with GLOBAL scope (for company-managed projects).
Use --project to create a project-scoped status (for team-managed projects).

Note: Requires 'Administer projects' or 'Administer Jira' permission.

Examples:
  # Create a global status
  jira-cli create status --name "In Review" --category IN_PROGRESS

  # Create a status with description
  jira-cli create status --name "In Review" --category IN_PROGRESS --description "Code review in progress"

  # Create a project-scoped status (for team-managed projects)
  jira-cli create status --name "In Review" --category IN_PROGRESS --project PROJ`,
	RunE: runCreateStatus,
}

func init() {
	createCmd.AddCommand(createStatusCmd)

	flags := createStatusCmd.Flags()
	flags.StringVarP(&createStatusOpts.name, "name", "n", "", "Status name (required)")
	flags.StringVarP(&createStatusOpts.category, "category", "c", "", "Status category: TODO, IN_PROGRESS, DONE (required)")
	flags.StringVarP(&createStatusOpts.description, "description", "d", "", "Status description")
	flags.StringVarP(&createStatusOpts.project, "project", "p", "", "Project key (for project-scoped status in team-managed projects)")
	createStatusCmd.MarkFlagRequired("name")
	createStatusCmd.MarkFlagRequired("category")
}

func runCreateStatus(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// Validate category
	validCategories := map[string]bool{
		"TODO":        true,
		"IN_PROGRESS": true,
		"DONE":        true,
	}
	if !validCategories[createStatusOpts.category] {
		return fmt.Errorf("invalid category %q. Valid categories are: TODO, IN_PROGRESS, DONE", createStatusOpts.category)
	}

	// Create the status object
	statusCreate := swagger.NewStatusCreate(createStatusOpts.name, createStatusOpts.category)
	if createStatusOpts.description != "" {
		statusCreate.SetDescription(createStatusOpts.description)
	}

	// Create the scope
	var scope *swagger.StatusScope
	if createStatusOpts.project != "" {
		// Get project ID
		project, _, err := client.ProjectsAPI.GetProject(ctx, createStatusOpts.project).Execute()
		if err != nil {
			return fmt.Errorf("failed to get project: %w", err)
		}

		projectId := swagger.NewProjectId(project.GetId())
		scope = swagger.NewStatusScope("PROJECT")
		scope.SetProject(*projectId)
	} else {
		scope = swagger.NewStatusScope("GLOBAL")
	}

	createRequest := swagger.NewStatusCreateRequest(*scope, []swagger.StatusCreate{*statusCreate})

	createdStatuses, _, err := client.StatusAPI.CreateStatuses(ctx).
		StatusCreateRequest(*createRequest).
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

	if len(createdStatuses) > 0 {
		status := createdStatuses[0]
		fmt.Printf("Status %q has been created with ID %s\n", status.GetName(), status.GetId())
	} else {
		fmt.Println("Status has been created")
	}

	return nil
}
