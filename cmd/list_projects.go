package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listProjectsOptions struct {
	query   string
	orderBy string
}

var listProjectsOpts = listProjectsOptions{}

var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List all projects",
	RunE:  runListProjects,
}

func init() {
	listCmd.AddCommand(listProjectsCmd)

	flags := listProjectsCmd.Flags()
	flags.StringVarP(&listProjectsOpts.query, "query", "q", "", "Filter by project name or key (partial match)")
	flags.StringVarP(&listProjectsOpts.orderBy, "order-by", "o", "", "Order by field (e.g., name, key, -name, -key)")
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
			return err
		}

		allProjects = append(allProjects, result.GetValues()...)

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) == 0 {
		fmt.Println("No projects found")
		return nil
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for _, project := range allProjects {
		key := project.GetKey()
		name := project.GetName()
		projectType := project.GetProjectTypeKey()

		fmt.Printf("%s %s %s\n", yellow(key), name, cyan(projectType))
	}

	return nil
}
