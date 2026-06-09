package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type getCommentOptions struct {
	id          string
	commentID   string
	noImages    bool
	commentOnly bool
	format      string
}

var getCommentOpts = getCommentOptions{}

var getCommentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Get a comment of an issue",
	Long: `Get a specific comment of an issue.

Examples:
  # Get a comment by issue ID and comment ID
  jira-cli get comment --id PROJ-123 --comment-id 10001`,
	RunE: runGetComment,
}

func init() {
	getCmd.AddCommand(getCommentCmd)

	flags := getCommentCmd.Flags()
	flags.StringVarP(&getCommentOpts.id, "id", "i", "", "Issue ID (e.g. PROJ-123)")
	flags.StringVarP(&getCommentOpts.commentID, "comment-id", "c", "", "Comment ID")
	flags.BoolVar(&getCommentOpts.noImages, "no-images", false, "Do not display images inline")
	flags.BoolVar(&getCommentOpts.commentOnly, "comment-only", false, "Print comment body only, without ID, author, or date")
	flags.StringVar(&getCommentOpts.format, "format", "table", "Output format: table or json")

	getCommentCmd.MarkFlagRequired("id")
	getCommentCmd.MarkFlagRequired("comment-id")
}

func runGetComment(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	result, _, err := client.IssueCommentsAPI.GetComment(ctx, getCommentOpts.id, getCommentOpts.commentID).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	if getCommentOpts.format == "json" {
		authorObj := result.GetAuthor()
		author := authorObj.GetDisplayName()
		if author == "" {
			author = authorObj.GetEmailAddress()
		}
		if author == "" {
			author = authorObj.GetAccountId()
		}
		return printJSON(commentOutput{
			ID:      result.GetId(),
			Author:  author,
			Created: result.GetCreated().String(),
			Body:    result.Body,
		})
	}

	showImages := !getCommentOpts.noImages && supportsKittyGraphics()

	var attMaps *attachmentMaps
	if showImages {
		var attErr error
		attMaps, attErr = buildAttachmentMaps(client, ctx, getCommentOpts.id)
		if attErr != nil {
			fmt.Printf("Warning: could not fetch attachments: %v\n\n", attErr)
			showImages = false
		}
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	printCommentWithImages(result, yellow, cyan, showImages, attMaps, getCommentOpts.commentOnly)
	return nil
}
