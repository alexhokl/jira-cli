package cmd

import (
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/cobra"
)

type addCommentOptions struct {
	id      string
	message string
}

var addCommentOpts = addCommentOptions{}

// addCommentCmd represents the comment command
var addCommentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a comment",
	RunE:  runAddComment,
}

func init() {
	addCmd.AddCommand(addCommentCmd)

	flags := addCommentCmd.Flags()
	flags.StringVarP(&addCommentOpts.id, "id", "i", "", "Issue ID")
	flags.StringVarP(&addCommentOpts.message, "message", "m", "", "Message")

	addCommentCmd.MarkFlagRequired("id")
}

func runAddComment(_ *cobra.Command, _ []string) error {
	body := newComment(addCommentOpts.message)
	client := newClient()
	_, _, err := client.IssueCommentsAPI.AddComment(getAuthContext(), addCommentOpts.id).Comment(*body).Execute()
	if err != nil {
		return err
	}
	fmt.Println("Comment added")
	return nil
}

func newComment(comment string) *swagger.Comment {
	body := map[string]interface{}{
		"type":    "doc",
		"version": 1,
		"content": []map[string]interface{}{
			{
				"type": "paragraph",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": comment,
					},
				},
			},
		},
	}
	return &swagger.Comment{
		Body: body,
	}
}
