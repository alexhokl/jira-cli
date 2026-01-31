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
}

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
		if listBoardsOpts.boardType != "" {
			request = request.Type_(map[string]interface{}{"type": listBoardsOpts.boardType})
		}
		if listBoardsOpts.includePrivate {
			request = request.IncludePrivate(true)
		}
		if listBoardsOpts.orderByName {
			request = request.OrderBy("name")
		}

		result, _, err := request.Execute()
		if err != nil {
			return err
		}

		allBoards = append(allBoards, result.GetValues()...)

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allBoards) == 0 {
		fmt.Println("No boards found")
		return nil
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for _, board := range allBoards {
		id := board.GetId()
		name := board.GetName()
		boardType := board.GetType()

		projectKey := ""
		if board.HasLocation() {
			location := board.GetLocation()
			projectKey = location.GetProjectKey()
		}

		if projectKey != "" {
			fmt.Printf("%s %s %s [%s]\n", yellow(fmt.Sprintf("%d", id)), name, cyan(boardType), green(projectKey))
		} else {
			fmt.Printf("%s %s %s\n", yellow(fmt.Sprintf("%d", id)), name, cyan(boardType))
		}
	}

	return nil
}
