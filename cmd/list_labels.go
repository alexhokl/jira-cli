package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listLabelsOptions struct {
	maxResults int32
}

var listLabelsOpts = listLabelsOptions{}

var listLabelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "List all labels",
	Long: `List all labels used in the Jira instance.

Examples:
  # List all labels
  jira-cli list labels

  # List first 100 labels
  jira-cli list labels --max-results 100`,
	RunE: runListLabels,
}

func init() {
	listCmd.AddCommand(listLabelsCmd)

	flags := listLabelsCmd.Flags()
	flags.Int32VarP(&listLabelsOpts.maxResults, "max-results", "m", 0, "Maximum number of results to return (0 for all)")
}

func runListLabels(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	var allLabels []string
	var startAt int64 = 0
	const pageSize int32 = 1000

	for {
		result, _, err := client.LabelsAPI.GetAllLabels(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			return err
		}

		allLabels = append(allLabels, result.GetValues()...)

		// If user specified a limit, stop when we have enough
		if listLabelsOpts.maxResults > 0 && int32(len(allLabels)) >= listLabelsOpts.maxResults {
			allLabels = allLabels[:listLabelsOpts.maxResults]
			break
		}

		// Check if this is the last page
		if result.GetIsLast() {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allLabels) == 0 {
		fmt.Println("No labels found")
		return nil
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()

	for _, label := range allLabels {
		fmt.Println(yellow(label))
	}

	return nil
}
