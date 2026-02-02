package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listStatusOptions struct {
	query    string
	category string
	project  string
}

var listStatusOpts = listStatusOptions{}

var listStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "List statuses",
	Long: `List all statuses available in the Jira instance.

Examples:
  # List all statuses
  jira-cli list status

  # Search statuses by name
  jira-cli list status --query "progress"

  # Filter statuses by category (TODO, IN_PROGRESS, DONE)
  jira-cli list status --category IN_PROGRESS

  # List statuses for a specific project
  jira-cli list status --project PROJ`,
	RunE: runListStatus,
}

func init() {
	listCmd.AddCommand(listStatusCmd)

	flags := listStatusCmd.Flags()
	flags.StringVarP(&listStatusOpts.query, "query", "q", "", "Filter statuses by name (partial match)")
	flags.StringVarP(&listStatusOpts.category, "category", "c", "", "Filter by category (TODO, IN_PROGRESS, DONE)")
	flags.StringVarP(&listStatusOpts.project, "project", "p", "", "Filter statuses by project key")
}

func runListStatus(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	var statuses []statusInfo
	var startAt int64 = 0

	// If project is specified, get project ID
	var projectId string
	if listStatusOpts.project != "" {
		project, _, err := client.ProjectsAPI.GetProject(ctx, listStatusOpts.project).Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get project: %w", err))
		}
		projectId = project.GetId()
	}

	// Build project ID to key map for resolving project scopes
	projectIdToKey := make(map[string]string)
	projects, _, err := client.ProjectsAPI.GetAllProjects(ctx).Execute()
	if err == nil {
		for _, p := range projects {
			projectIdToKey[p.GetId()] = p.GetKey()
		}
	}

	for {
		request := client.StatusAPI.Search(ctx).
			StartAt(startAt).
			MaxResults(50)

		if listStatusOpts.query != "" {
			request = request.SearchString(listStatusOpts.query)
		}

		if listStatusOpts.category != "" {
			request = request.StatusCategory(listStatusOpts.category)
		}

		if projectId != "" {
			request = request.ProjectId(projectId)
		}

		result, _, err := request.Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get statuses: %w", err))
		}

		for _, s := range result.GetValues() {
			scope := "GLOBAL"
			projectKey := ""
			if s.Scope != nil {
				scope = s.Scope.GetType()
				if s.Scope.Project.Get() != nil {
					projId := s.Scope.Project.Get().GetId()
					if key, ok := projectIdToKey[projId]; ok {
						projectKey = key
					} else {
						projectKey = projId // fallback to ID if key not found
					}
				}
			}

			statuses = append(statuses, statusInfo{
				id:          s.GetId(),
				name:        s.GetName(),
				category:    s.GetStatusCategory(),
				scope:       scope,
				project:     projectKey,
				description: s.GetDescription(),
			})
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(statuses) == 0 {
		fmt.Println("No statuses found")
		return nil
	}

	// Sort by name
	sort.Slice(statuses, func(i, j int) bool {
		return strings.ToLower(statuses[i].name) < strings.ToLower(statuses[j].name)
	})

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("CATEGORY"), cyan("SCOPE"), cyan("PROJECT"), cyan("DESCRIPTION"))

	for _, s := range statuses {
		// Truncate description if too long
		desc := s.description
		if len(desc) > 40 {
			desc = desc[:37] + "..."
		}
		w.row(s.id, yellow(s.name), s.category, s.scope, s.project, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d statuses\n", len(statuses))

	return nil
}

type statusInfo struct {
	id          string
	name        string
	category    string
	scope       string
	project     string
	description string
}
