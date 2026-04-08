package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type getSprintOptions struct {
	id     int64
	format string
}

var getSprintOpts = getSprintOptions{}

var getSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Get sprint",
	Long: `Get details of a specific sprint.

Examples:
  # Get sprint details in text format
  jira-cli get sprint -i 42

  # Get sprint details in JSON format
  jira-cli get sprint -i 42 --format json`,
	RunE: runGetSprint,
}

func init() {
	getCmd.AddCommand(getSprintCmd)

	flags := getSprintCmd.Flags()
	flags.Int64VarP(&getSprintOpts.id, "id", "i", 0, "Sprint ID")
	flags.StringVar(&getSprintOpts.format, "format", "text", "Output format (text, json)")

	getSprintCmd.MarkFlagRequired("id")
}

func runGetSprint(_ *cobra.Command, _ []string) error {
	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	resp, err := client.SprintAPI.GetSprint(ctx, getSprintOpts.id).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var sprint swagger_software.SprintBean
	if err := json.Unmarshal(body, &sprint); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if getSprintOpts.format == "json" {
		output, err := json.MarshalIndent(sprint, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal sprint to JSON: %w", err)
		}
		fmt.Println(string(output))
		return nil
	}

	// Text output
	color.NoColor = noColor
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	stateColor := cyan
	state := sprint.GetState()
	switch state {
	case "active":
		stateColor = green
	case "closed":
		stateColor = red
	case "future":
		stateColor = cyan
	}

	fmt.Printf("%s %d\n", cyan("ID:"), sprint.GetId())
	fmt.Printf("%s %s\n", cyan("Name:"), yellow(sprint.GetName()))
	fmt.Printf("%s %s\n", cyan("State:"), stateColor(state))

	if sprint.HasOriginBoardId() {
		fmt.Printf("%s %d\n", cyan("Board ID:"), sprint.GetOriginBoardId())
	}

	if sprint.HasStartDate() {
		fmt.Printf("%s %s\n", cyan("Start Date:"), formatDateString(sprint.GetStartDate()))
	}
	if sprint.HasEndDate() {
		fmt.Printf("%s %s\n", cyan("End Date:"), formatDateString(sprint.GetEndDate()))
	}
	if sprint.HasCompleteDate() {
		fmt.Printf("%s %s\n", cyan("Complete Date:"), formatDateString(sprint.GetCompleteDate()))
	}
	if sprint.HasGoal() && sprint.GetGoal() != "" {
		fmt.Printf("%s %s\n", cyan("Goal:"), sprint.GetGoal())
	}

	return nil
}
