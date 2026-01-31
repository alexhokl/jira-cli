package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type createCommentOptions struct {
	id      string
	message string
}

var createCommentOpts = createCommentOptions{}

var createCommentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Create a comment",
	RunE:  runCreateComment,
}

func init() {
	createCmd.AddCommand(createCommentCmd)

	flags := createCommentCmd.Flags()
	flags.StringVarP(&createCommentOpts.id, "id", "i", "", "Issue ID")
	flags.StringVarP(&createCommentOpts.message, "message", "m", "", "Message")

	createCommentCmd.MarkFlagRequired("id")
}

func runCreateComment(_ *cobra.Command, _ []string) error {
	message := createCommentOpts.message
	if strings.TrimSpace(message) == "" {
		var err error
		message, err = getMessageFromEditor()
		if err != nil {
			return err
		}
	}
	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("comment message cannot be empty")
	}

	body := newComment(message)
	client := newClient()
	_, _, err := client.IssueCommentsAPI.AddComment(getAuthContext(), createCommentOpts.id).Comment(*body).Execute()
	if err != nil {
		return err
	}
	fmt.Println("Comment created")
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
