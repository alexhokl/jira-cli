package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type listBoardsOptions struct {
	name           string
	projectKeyOrId string
	boardType      string
	includePrivate bool
	orderByName    bool
	maxResults     int32
	format         string
}

var listBoardsOpts = listBoardsOptions{}

var listBoardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List all boards",
	RunE:  runListBoards,
}

func init() {
	listCmd.AddCommand(listBoardsCmd)

	flags := listBoardsCmd.Flags()
	flags.StringVarP(&listBoardsOpts.name, "name", "n", "", "Filter by board name (partial match)")
	flags.StringVarP(&listBoardsOpts.projectKeyOrId, "project", "p", "", "Filter by project key or ID")
	flags.StringVarP(&listBoardsOpts.boardType, "type", "t", "", "Filter by board type (scrum, kanban, simple)")
	flags.BoolVar(&listBoardsOpts.includePrivate, "include-private", false, "Include private boards")
	flags.BoolVar(&listBoardsOpts.orderByName, "order-by-name", false, "Order results by name")
	flags.Int32Var(&listBoardsOpts.maxResults, "limit", 0, "Maximum number of results to return (0 for all)")
	flags.StringVar(&listBoardsOpts.format, "format", "table", "Output format: table or json")
}

// Note: filterBoardsByType is defined in helper.go

func runListBoards(_ *cobra.Command, _ []string) error {
	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	var allBoards []swagger_software.GetAllBoards200ResponseValuesInner
	var startAt int64 = 0

	for {
		request := client.BoardAPI.GetAllBoards(ctx).StartAt(startAt)

		if listBoardsOpts.name != "" {
			request = request.Name(listBoardsOpts.name)
		}
		if listBoardsOpts.projectKeyOrId != "" {
			request = request.ProjectKeyOrId(listBoardsOpts.projectKeyOrId)
		}
		// Note: We intentionally do NOT use the generated Type_() method here
		// because the swagger-generated code incorrectly serializes the type parameter
		// as type[type]=scrum instead of type=scrum. We filter by type client-side instead.
		if listBoardsOpts.includePrivate {
			request = request.IncludePrivate(true)
		}
		if listBoardsOpts.orderByName {
			request = request.OrderBy("name")
		}

		result, _, err := request.Execute()
		if err != nil {
			return wrapAPIError(err)
		}

		allBoards = append(allBoards, result.GetValues()...)

		if listBoardsOpts.maxResults > 0 && int32(len(allBoards)) >= listBoardsOpts.maxResults {
			allBoards = allBoards[:listBoardsOpts.maxResults]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	// Filter by board type client-side (workaround for swagger-generated code bug)
	if listBoardsOpts.boardType != "" {
		allBoards = filterBoardsByType(allBoards, listBoardsOpts.boardType)
	}

	if len(allBoards) == 0 {
		fmt.Println("No boards found")
		return nil
	}

	// Build typed slice for output
	var items []boardOutput
	for _, board := range allBoards {
		id := board.GetId()
		name := board.GetName()
		boardType := board.GetType()
		projectKey := ""
		if board.HasLocation() {
			location := board.GetLocation()
			projectKey = location.GetProjectKey()
		}
		items = append(items, boardOutput{
			ID:         id,
			Name:       name,
			Type:       boardType,
			ProjectKey: projectKey,
		})
	}

	if listBoardsOpts.format == "json" {
		return printJSON(items)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for _, item := range items {
		if item.ProjectKey != "" {
			fmt.Printf("%s %s %s [%s]\n", yellow(fmt.Sprintf("%d", item.ID)), item.Name, cyan(item.Type), green(item.ProjectKey))
		} else {
			fmt.Printf("%s %s %s\n", yellow(fmt.Sprintf("%d", item.ID)), item.Name, cyan(item.Type))
		}
	}

	return nil
}

type boardOutput struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ProjectKey string `json:"projectKey"`
}
