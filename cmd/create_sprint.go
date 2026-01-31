package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type createSprintOptions struct {
	boardId int64
	name    string
	start   string
	end     string
}

var createSprintOpts = createSprintOptions{}

var createSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Create a sprint",
	RunE:  runCreateSprint,
}

func init() {
	createCmd.AddCommand(createSprintCmd)

	flags := createSprintCmd.Flags()
	flags.Int64VarP(&createSprintOpts.boardId, "board-id", "b", 0, "Board ID (required)")
	flags.StringVarP(&createSprintOpts.name, "name", "n", "", "Sprint name (required)")
	flags.StringVarP(&createSprintOpts.start, "start", "s", "", "Start date (format: 2006-01-02)")
	flags.StringVarP(&createSprintOpts.end, "end", "e", "", "End date (format: 2006-01-02)")

	createSprintCmd.MarkFlagRequired("board-id")
	createSprintCmd.MarkFlagRequired("name")
}

func runCreateSprint(_ *cobra.Command, _ []string) error {
	request := swagger_software.NewCreateSprintRequest()
	request.SetOriginBoardId(createSprintOpts.boardId)
	request.SetName(createSprintOpts.name)

	if createSprintOpts.start != "" {
		startDate, err := parseDate(createSprintOpts.start)
		if err != nil {
			return fmt.Errorf("invalid start date: %w", err)
		}
		request.SetStartDate(startDate)
	}

	if createSprintOpts.end != "" {
		endDate, err := parseDate(createSprintOpts.end)
		if err != nil {
			return fmt.Errorf("invalid end date: %w", err)
		}
		request.SetEndDate(endDate)
	}

	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	_, err := client.SprintAPI.CreateSprint(ctx).CreateSprintRequest(*request).Execute()
	if err != nil {
		return err
	}

	fmt.Println("Sprint created")
	return nil
}
