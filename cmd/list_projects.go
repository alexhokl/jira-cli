package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listProjectsOptions struct {
	query      string
	orderBy    string
	maxResults int32
	format     string
}

var listProjectsOpts = listProjectsOptions{}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"project"},
	Short:   "List all projects",
	RunE:    runListProjects,
}

func init() {
	listCmd.AddCommand(listProjectsCmd)

	flags := listProjectsCmd.Flags()
	flags.StringVarP(&listProjectsOpts.query, "query", "q", "", "Filter by project name or key (partial match)")
	flags.StringVarP(&listProjectsOpts.orderBy, "order-by", "o", "", "Order by field (e.g., name, key, -name, -key)")
	flags.Int32Var(&listProjectsOpts.maxResults, "limit", 0, "Maximum number of results to return (0 for all)")
	flags.StringVar(&listProjectsOpts.format, "format", "table", "Output format: table or json")
}

func runListProjects(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	var allProjects []swagger.Project
	var startAt int64 = 0

	for {
		request := client.ProjectsAPI.SearchProjects(ctx).StartAt(startAt)

		if listProjectsOpts.query != "" {
			request = request.Query(listProjectsOpts.query)
		}
		if listProjectsOpts.orderBy != "" {
			request = request.OrderBy(listProjectsOpts.orderBy)
		}

		result, _, err := request.Execute()
		if err != nil {
			return wrapAPIError(err)
		}

		allProjects = append(allProjects, result.GetValues()...)

		if listProjectsOpts.maxResults > 0 && int32(len(allProjects)) >= listProjectsOpts.maxResults {
			allProjects = allProjects[:listProjectsOpts.maxResults]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) == 0 {
		fmt.Println("No projects found")
		return nil
	}

	// Build typed slice for output
	var items []projectOutput
	for _, project := range allProjects {
		items = append(items, projectOutput{
			Key:  project.GetKey(),
			Name: project.GetName(),
			Type: project.GetProjectTypeKey(),
		})
	}

	if listProjectsOpts.format == "json" {
		return printJSON(items)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for _, item := range items {
		fmt.Printf("%s %s %s\n", yellow(item.Key), item.Name, cyan(item.Type))
	}

	return nil
}

type projectOutput struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}
