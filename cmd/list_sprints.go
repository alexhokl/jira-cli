package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type listSprintsOptions struct {
	boardId int64
	state   []string
}

var listSprintsOpts = listSprintsOptions{}

var listSprintsCmd = &cobra.Command{
	Use:   "sprints",
	Short: "List all sprints for a board",
	RunE:  runListSprints,
}

func init() {
	listCmd.AddCommand(listSprintsCmd)

	flags := listSprintsCmd.Flags()
	flags.Int64VarP(&listSprintsOpts.boardId, "board-id", "b", 0, "Board ID (required)")
	flags.StringSliceVarP(&listSprintsOpts.state, "state", "s", nil, "Filter by sprint state (future, active, closed); can be specified multiple times")

	listSprintsCmd.MarkFlagRequired("board-id")
}

// sprintsResponse represents the paginated response from GetAllSprints API
type sprintsResponse struct {
	MaxResults int                           `json:"maxResults"`
	StartAt    int                           `json:"startAt"`
	IsLast     bool                          `json:"isLast"`
	Values     []swagger_software.SprintBean `json:"values"`
}

func runListSprints(_ *cobra.Command, _ []string) error {
	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	var allSprints []swagger_software.SprintBean
	var startAt int64 = 0

	// Build a set of allowed states for filtering
	stateFilter := make(map[string]bool)
	for _, s := range listSprintsOpts.state {
		stateFilter[strings.TrimSpace(s)] = true
	}

	for {
		request := client.BoardAPI.GetAllSprints(ctx, listSprintsOpts.boardId).StartAt(startAt)

		resp, err := request.Execute()
		if err != nil {
			return wrapAPIError(err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		var result sprintsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		// Filter sprints by state client-side
		// (The swagger-generated State() method doesn't work correctly)
		for _, sprint := range result.Values {
			if len(stateFilter) == 0 || stateFilter[sprint.GetState()] {
				allSprints = append(allSprints, sprint)
			}
		}

		if result.IsLast {
			break
		}

		startAt = int64(result.StartAt + result.MaxResults)
	}

	if len(allSprints) == 0 {
		fmt.Println("No sprints found")
		return nil
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for _, sprint := range allSprints {
		id := sprint.GetId()
		name := sprint.GetName()
		state := sprint.GetState()

		stateColor := cyan
		switch state {
		case "active":
			stateColor = green
		case "closed":
			stateColor = red
		case "future":
			stateColor = cyan
		}

		startDate := ""
		if sprint.HasStartDate() {
			startDate = formatDateString(sprint.GetStartDate())
		}
		endDate := ""
		if sprint.HasEndDate() {
			endDate = formatDateString(sprint.GetEndDate())
		}

		if startDate != "" && endDate != "" {
			fmt.Printf("%s %s %s (%s - %s)\n", yellow(fmt.Sprintf("%d", id)), name, stateColor(state), startDate, endDate)
		} else {
			fmt.Printf("%s %s %s\n", yellow(fmt.Sprintf("%d", id)), name, stateColor(state))
		}
	}

	return nil
}

// formatDateString parses a date string and returns it in yyyy-MM-dd format
// in the local timezone of the CLI client
func formatDateString(dateStr string) string {
	// Try parsing common date formats from Jira API
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05.000-0700",
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			// Convert to local timezone before formatting
			return t.Local().Format("2006-01-02")
		}
	}

	// If parsing fails, return the original string
	return dateStr
}
