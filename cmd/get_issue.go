package cmd

import (
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type getIssueOptions struct {
	id string
}

var getIssueOpts = getIssueOptions{}

// getIssueCmd represents the list command
var getIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Get issue",
	RunE:  runGetIssue,
}

func init() {
	getCmd.AddCommand(getIssueCmd)

	flags := getIssueCmd.Flags()
	flags.StringVarP(&getIssueOpts.id, "id", "i", "", "Issue ID")

	getIssueCmd.MarkFlagRequired("id")
}

func runGetIssue(_ *cobra.Command, _ []string) error {
	opts := &swagger.IssuesApiGetIssueOpts{}
	client := swagger.NewAPIClient(getConfiguration())
	issue, _, err := client.IssuesApi.GetIssue(getAuthContext(), getIssueOpts.id, opts)
	if err != nil {
		return err
	}

	summaryObject := issue.Fields["summary"]
	fmt.Printf("%s %s\n", issue.Key, summaryObject)

	return nil
}
