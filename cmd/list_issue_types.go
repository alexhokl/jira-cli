package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listIssueTypesOptions struct {
	project string
}

var listIssueTypesOpts = listIssueTypesOptions{}

var listIssueTypesCmd = &cobra.Command{
	Use:   "issue-types",
	Short: "List issue types",
	Long: `List all issue types available in the Jira instance.

Examples:
  # List all issue types
  jira-cli list issue-types

  # List issue types available for a specific project
  jira-cli list issue-types --project PROJ`,
	RunE: runListIssueTypes,
}

func init() {
	listCmd.AddCommand(listIssueTypesCmd)

	flags := listIssueTypesCmd.Flags()
	flags.StringVarP(&listIssueTypesOpts.project, "project", "p", "", "Filter issue types by project key")
}

func runListIssueTypes(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	var issueTypes []issueTypeInfo

	if listIssueTypesOpts.project != "" {
		// Get project ID first
		project, _, err := client.ProjectsAPI.GetProject(ctx, listIssueTypesOpts.project).Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get project: %w", err))
		}

		projectId, err := parseProjectId(project.GetId())
		if err != nil {
			return err
		}

		// Get issue types for the project
		types, _, err := client.IssueTypesAPI.GetIssueTypesForProject(ctx).
			ProjectId(projectId).
			Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get issue types for project: %w", err))
		}

		for _, t := range types {
			issueTypes = append(issueTypes, issueTypeInfo{
				id:          t.GetId(),
				name:        t.GetName(),
				description: t.GetDescription(),
				subtask:     t.GetSubtask(),
				project:     listIssueTypesOpts.project,
			})
		}
	} else {
		// Get all issue types
		types, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get issue types: %w", err))
		}

		for _, t := range types {
			projectKey := ""
			if t.HasScope() {
				scope := t.GetScope()
				if scope.HasProject() {
					project := scope.GetProject()
					projectKey = project.GetKey()
				}
			}
			issueTypes = append(issueTypes, issueTypeInfo{
				id:          t.GetId(),
				name:        t.GetName(),
				description: t.GetDescription(),
				subtask:     t.GetSubtask(),
				project:     projectKey,
			})
		}
	}

	if len(issueTypes) == 0 {
		fmt.Println("No issue types found")
		return nil
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("PROJECT"), cyan("SUBTASK"), cyan("DESCRIPTION"))

	for _, t := range issueTypes {
		subtaskStr := "No"
		if t.subtask {
			subtaskStr = "Yes"
		}
		projectStr := t.project
		if projectStr == "" {
			projectStr = "(global)"
		}
		// Truncate description if too long
		desc := t.description
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		w.row(t.id, yellow(t.name), projectStr, subtaskStr, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d issue types\n", len(issueTypes))

	return nil
}

type issueTypeInfo struct {
	id          string
	name        string
	description string
	subtask     bool
	project     string
}

// parseProjectId converts project ID string to int64
func parseProjectId(id string) (int64, error) {
	var projectId int64
	_, err := fmt.Sscanf(id, "%d", &projectId)
	if err != nil {
		return 0, fmt.Errorf("failed to parse project ID: %w", err)
	}
	return projectId, nil
}
