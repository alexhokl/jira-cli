package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listCommentsOptions struct {
	id string
}

var listCommentsOpts = listCommentsOptions{}

var listCommentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "List comments of an issue",
	RunE:  runListComments,
}

func init() {
	listCmd.AddCommand(listCommentsCmd)

	flags := listCommentsCmd.Flags()
	flags.StringVarP(&listCommentsOpts.id, "id", "i", "", "Issue ID")

	listCommentsCmd.MarkFlagRequired("id")
}

func runListComments(_ *cobra.Command, _ []string) error {
	client := newClient()
	result, _, err := client.IssueCommentsAPI.GetComments(getAuthContext(), listCommentsOpts.id).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	comments := result.GetComments()
	if len(comments) == 0 {
		fmt.Println("No comments found")
		return nil
	}

	// API returns comments in chronological order, reverse to get newest first
	for i, j := 0, len(comments)-1; i < j; i, j = i+1, j-1 {
		comments[i], comments[j] = comments[j], comments[i]
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for i, comment := range comments {
		if i > 0 {
			fmt.Println()
		}
		printComment(&comment, yellow, cyan)
	}

	return nil
}

func printComment(comment *swagger.Comment, yellow, cyan func(a ...interface{}) string) {
	authorName := getAuthorDisplayName(comment.Author)
	createdDate := ""
	if comment.HasCreated() {
		createdDate = comment.GetCreated().Format("2006-01-02 15:04:05")
	}

	fmt.Printf("%s - %s\n", yellow(authorName), cyan(createdDate))

	body := convertADFToMarkdown(comment.Body)
	if body != "" {
		fmt.Printf("%s\n", body)
	}
}

func getAuthorDisplayName(author *swagger.UserDetails) string {
	if author == nil {
		return "Unknown"
	}
	if author.HasDisplayName() {
		return author.GetDisplayName()
	}
	if author.HasEmailAddress() {
		return author.GetEmailAddress()
	}
	if author.HasAccountId() {
		return author.GetAccountId()
	}
	return "Unknown"
}
