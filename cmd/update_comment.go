package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type updateCommentOptions struct {
	issueID   string
	commentID string
}

var updateCommentOpts = updateCommentOptions{}

var updateCommentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Update a comment",
	Long: `Update an existing comment on an issue.

The command fetches the existing comment from Jira and opens it in your editor
for modification. Markdown formatting is preserved.

Examples:
  # Update a comment (opens $EDITOR with existing content)
  jira-cli update comment --issue PROJ-123 --comment-id 10001

Note: Use 'jira-cli list comments --id PROJ-123' to find comment IDs.`,
	RunE: runUpdateComment,
}

func init() {
	updateCmd.AddCommand(updateCommentCmd)

	flags := updateCommentCmd.Flags()
	flags.StringVarP(&updateCommentOpts.issueID, "issue", "i", "", "Issue ID (e.g., PROJ-123)")
	flags.StringVarP(&updateCommentOpts.commentID, "comment-id", "c", "", "Comment ID")

	updateCommentCmd.MarkFlagRequired("issue")
	updateCommentCmd.MarkFlagRequired("comment-id")
}

func runUpdateComment(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// Fetch the existing comment from Jira
	existingComment, _, err := client.IssueCommentsAPI.GetComment(
		ctx,
		updateCommentOpts.issueID,
		updateCommentOpts.commentID,
	).Execute()
	if err != nil {
		return fmt.Errorf("failed to fetch existing comment: %w", wrapAPIError(err))
	}

	// Convert ADF to markdown to preserve formatting
	existingMarkdown := convertADFToMarkdown(existingComment.Body)

	// Open editor with existing content
	message, err := getMessageFromEditorWithContent(existingMarkdown)
	if err != nil {
		return err
	}

	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("comment message cannot be empty")
	}

	body := newComment(message)
	_, _, err = client.IssueCommentsAPI.UpdateComment(
		ctx,
		updateCommentOpts.issueID,
		updateCommentOpts.commentID,
	).Comment(*body).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	fmt.Println("Comment updated")
	return nil
}

// getMessageFromEditorWithContent opens the user's editor with pre-populated content
// and returns the edited message.
func getMessageFromEditorWithContent(initialContent string) (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "", fmt.Errorf("EDITOR environment variable is not set")
	}

	tmpFile, err := os.CreateTemp("", "jira-comment-*.md")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpFileName := tmpFile.Name()

	// Write the initial content to the temp file
	if _, err := tmpFile.WriteString(initialContent); err != nil {
		tmpFile.Close()
		os.Remove(tmpFileName)
		return "", fmt.Errorf("failed to write initial content: %w", err)
	}
	tmpFile.Close()
	defer os.Remove(tmpFileName)

	cmd := exec.Command(editor, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run editor: %w", err)
	}

	content, err := os.ReadFile(tmpFileName)
	if err != nil {
		return "", fmt.Errorf("failed to read temporary file: %w", err)
	}

	return strings.TrimSpace(string(content)), nil
}
