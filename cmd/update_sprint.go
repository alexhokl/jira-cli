package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type updateSprintOptions struct {
	id    int64
	name  string
	start string
	end   string
}

var updateSprintOpts = updateSprintOptions{}

var updateSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Update a sprint",
	RunE:  runUpdateSprint,
}

func init() {
	updateCmd.AddCommand(updateSprintCmd)

	flags := updateSprintCmd.Flags()
	flags.Int64VarP(&updateSprintOpts.id, "id", "i", 0, "Sprint ID (required)")
	flags.StringVarP(&updateSprintOpts.name, "name", "n", "", "Sprint name")
	flags.StringVarP(&updateSprintOpts.start, "start", "s", "", "Start date (format: 2006-01-02)")
	flags.StringVarP(&updateSprintOpts.end, "end", "e", "", "End date (format: 2006-01-02)")

	updateSprintCmd.MarkFlagRequired("id")
}

func runUpdateSprint(_ *cobra.Command, _ []string) error {
	if updateSprintOpts.name == "" && updateSprintOpts.start == "" && updateSprintOpts.end == "" {
		return fmt.Errorf("at least one of --name, --start, or --end must be specified")
	}

	request := swagger_software.NewUpdateSprintRequest()

	if updateSprintOpts.name != "" {
		request.SetName(updateSprintOpts.name)
	}

	if updateSprintOpts.start != "" {
		startDate, err := parseDate(updateSprintOpts.start)
		if err != nil {
			return fmt.Errorf("invalid start date: %w", err)
		}
		request.SetStartDate(startDate)
	}

	if updateSprintOpts.end != "" {
		endDate, err := parseDate(updateSprintOpts.end)
		if err != nil {
			return fmt.Errorf("invalid end date: %w", err)
		}
		request.SetEndDate(endDate)
	}

	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	_, err := client.SprintAPI.PartiallyUpdateSprint(ctx, updateSprintOpts.id).UpdateSprintRequest(*request).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	fmt.Println("Sprint updated")
	return nil
}

// parseDate parses a date string in yyyy-MM-dd format and returns it in ISO 8601 format
// using the local timezone
func parseDate(dateStr string) (string, error) {
	t, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		return "", err
	}
	return t.Format(time.RFC3339), nil
}
