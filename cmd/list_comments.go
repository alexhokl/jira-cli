package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alexhokl/jira-cli/swagger"
)

type listCommentsOptions struct {
	id       string
	noImages bool
}

var listCommentsOpts = listCommentsOptions{}

var listCommentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "List comments of an issue",
	Long: `List comments of an issue.

Examples:
  # List comments for an issue (images displayed by default in supported terminals)
  jira-cli list comments -i PROJ-123

  # List comments without inline images
  jira-cli list comments -i PROJ-123 --no-images

Note: Images are displayed inline by default in terminals that support
the Kitty graphics protocol (Ghostty, Kitty, WezTerm).`,
	RunE: runListComments,
}

func init() {
	listCmd.AddCommand(listCommentsCmd)

	flags := listCommentsCmd.Flags()
	flags.StringVarP(&listCommentsOpts.id, "id", "i", "", "Issue ID")
	flags.BoolVar(&listCommentsOpts.noImages, "no-images", false, "Do not display images inline")

	listCommentsCmd.MarkFlagRequired("id")
}

func runListComments(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	result, _, err := client.IssueCommentsAPI.GetComments(ctx, listCommentsOpts.id).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	comments := result.GetComments()
	if len(comments) == 0 {
		fmt.Println("No comments found")
		return nil
	}

	// Check if we should show images (default: yes, unless --no-images or terminal doesn't support it)
	showImages := !listCommentsOpts.noImages && supportsKittyGraphics()

	// Build attachment maps if showing images
	var attMaps *attachmentMaps
	if showImages {
		attMaps, err = buildAttachmentMaps(client, ctx, listCommentsOpts.id)
		if err != nil {
			// Non-fatal: just warn and continue without images
			fmt.Printf("Warning: could not fetch attachments: %v\n\n", err)
			showImages = false
		}
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
		printCommentWithImages(&comment, yellow, cyan, showImages, attMaps)
	}

	return nil
}

// attachmentInfo holds information about an attachment for image display
type attachmentInfo struct {
	id       string
	mimeType string
	content  string // URL to download content
	filename string
}

// attachmentMaps holds maps for looking up attachments by different keys
type attachmentMaps struct {
	byID       map[string]attachmentInfo // keyed by attachment ID (numeric string)
	byFilename map[string]attachmentInfo // keyed by filename
}

// buildAttachmentMaps fetches issue attachments and builds maps for lookup by ID and filename
func buildAttachmentMaps(client *swagger.APIClient, ctx context.Context, issueID string) (*attachmentMaps, error) {
	// Get issue with attachment field
	issue, _, err := client.IssuesAPI.GetIssue(ctx, issueID).Fields([]string{"attachment"}).Execute()
	if err != nil {
		return nil, err
	}

	maps := &attachmentMaps{
		byID:       make(map[string]attachmentInfo),
		byFilename: make(map[string]attachmentInfo),
	}

	// Extract attachments from issue fields
	attachmentsField := issue.Fields["attachment"]
	if attachmentsField == nil {
		return maps, nil
	}

	attachments, ok := attachmentsField.([]any)
	if !ok {
		return maps, nil
	}

	for _, att := range attachments {
		attMap, ok := att.(map[string]any)
		if !ok {
			continue
		}

		info := attachmentInfo{}

		if id, ok := attMap["id"].(string); ok {
			info.id = id
		}
		if mimeType, ok := attMap["mimeType"].(string); ok {
			info.mimeType = mimeType
		}
		if content, ok := attMap["content"].(string); ok {
			info.content = content
		}
		if filename, ok := attMap["filename"].(string); ok {
			info.filename = filename
		}

		if info.id != "" {
			maps.byID[info.id] = info
		}
		if info.filename != "" {
			maps.byFilename[info.filename] = info
		}
	}

	return maps, nil
}

func printCommentWithImages(comment *swagger.Comment, yellow, cyan func(a ...interface{}) string, showImages bool, attMaps *attachmentMaps) {
	authorName := getAuthorDisplayName(comment.Author)
	createdDate := ""
	if comment.HasCreated() {
		createdDate = comment.GetCreated().Format("2006-01-02 15:04:05")
	}
	commentID := ""
	if comment.HasId() {
		commentID = comment.GetId()
	}

	fmt.Printf("[%s] %s - %s\n", commentID, yellow(authorName), cyan(createdDate))

	if showImages && attMaps != nil {
		// Use the enhanced version that can display images
		printADFWithImages(comment.Body, attMaps)
	} else {
		body := convertADFToMarkdown(comment.Body)
		if body != "" {
			fmt.Printf("%s\n", body)
		}
	}
}

// printADFWithImages converts ADF to markdown and displays images inline
func printADFWithImages(adf any, attMaps *attachmentMaps) {
	adfMap, ok := adf.(map[string]any)
	if !ok {
		return
	}

	content, ok := adfMap["content"].([]any)
	if !ok {
		return
	}

	var result strings.Builder
	for i, block := range content {
		if i > 0 {
			result.WriteString("\n")
		}
		convertBlockWithImages(block, &result, 0, attMaps)
	}

	output := result.String()
	if output != "" {
		fmt.Print(output)
	}
}

// convertBlockWithImages converts a single ADF block to markdown, displaying images inline
func convertBlockWithImages(block any, result *strings.Builder, depth int, attMaps *attachmentMaps) {
	blockMap, ok := block.(map[string]any)
	if !ok {
		return
	}

	blockType, _ := blockMap["type"].(string)

	switch blockType {
	case "mediaSingle", "mediaGroup":
		// Handle media elements - try to display images
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				if mediaMap, ok := item.(map[string]any); ok {
					if mediaMap["type"] == "media" {
						displayMediaElement(mediaMap, result, attMaps)
					}
				}
			}
		}

	case "paragraph":
		convertInlineContent(blockMap, result)
		result.WriteString("\n")

	case "heading":
		level := 1
		if attrs, ok := blockMap["attrs"].(map[string]any); ok {
			if l, ok := attrs["level"].(float64); ok {
				level = int(l)
			}
		}
		result.WriteString(strings.Repeat("#", level) + " ")
		convertInlineContent(blockMap, result)
		result.WriteString("\n")

	case "bulletList":
		convertListItemsWithImages(blockMap, result, depth, false, attMaps)

	case "orderedList":
		convertListItemsWithImages(blockMap, result, depth, true, attMaps)

	case "listItem":
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				convertBlockWithImages(item, result, depth, attMaps)
			}
		}

	case "codeBlock":
		language := ""
		if attrs, ok := blockMap["attrs"].(map[string]any); ok {
			if lang, ok := attrs["language"].(string); ok {
				language = lang
			}
		}
		result.WriteString("```" + language + "\n")
		convertInlineContent(blockMap, result)
		result.WriteString("\n```\n")

	case "blockquote":
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				result.WriteString("> ")
				convertBlockWithImages(item, result, depth, attMaps)
			}
		}

	case "rule":
		result.WriteString("---\n")

	case "table":
		convertTable(blockMap, result)

	case "panel":
		panelType := "info"
		if attrs, ok := blockMap["attrs"].(map[string]any); ok {
			if pt, ok := attrs["panelType"].(string); ok {
				panelType = pt
			}
		}
		result.WriteString(fmt.Sprintf("> **%s**\n", strings.ToUpper(panelType)))
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				result.WriteString("> ")
				convertBlockWithImages(item, result, depth, attMaps)
			}
		}

	default:
		// For unknown block types, try to extract text content
		convertInlineContent(blockMap, result)
		if blockMap["content"] != nil {
			result.WriteString("\n")
		}
	}
}

// displayMediaElement handles a media element, displaying images inline if possible
func displayMediaElement(mediaMap map[string]any, result *strings.Builder, attMaps *attachmentMaps) {
	attrs, ok := mediaMap["attrs"].(map[string]any)
	if !ok {
		result.WriteString("![media](media)\n")
		return
	}

	// Get the attachment ID and alt text from the media element
	// Note: The media element's "id" is a UUID, not the attachment's numeric ID
	// We need to look up by filename (from alt attribute) instead
	alt, _ := attrs["alt"].(string)

	if alt == "" {
		result.WriteString("![media](media)\n")
		return
	}

	// Look up attachment info by filename (alt contains the filename)
	info, found := attMaps.byFilename[alt]
	if !found || !isImageMimeType(info.mimeType) {
		// Not an image or not found, just print placeholder
		result.WriteString(fmt.Sprintf("![%s](media)\n", alt))
		return
	}

	// Flush any pending text before displaying image
	if result.Len() > 0 {
		fmt.Print(result.String())
		result.Reset()
	}

	// Fetch and display the image
	if err := fetchAndDisplayImage(info.content); err != nil {
		// Failed to display image, print placeholder
		result.WriteString(fmt.Sprintf("![%s](media) (failed to load)\n", alt))
	} else {
		// Print the image URL below the rendered image
		fmt.Printf("%s\n", info.content)
	}
}

// fetchAndDisplayImage downloads an image from a URL and displays it using Kitty graphics
func fetchAndDisplayImage(imageURL string) error {
	// Create authenticated HTTP request
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(email, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch image: status %d", resp.StatusCode)
	}

	// Read image data
	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read image data: %w", err)
	}

	// Display using Kitty graphics protocol
	return displayImageKitty(imageData)
}

// convertListItemsWithImages converts list items to markdown, handling images
func convertListItemsWithImages(blockMap map[string]any, result *strings.Builder, depth int, ordered bool, attMaps *attachmentMaps) {
	content, ok := blockMap["content"].([]any)
	if !ok {
		return
	}

	indent := strings.Repeat("  ", depth)

	for i, item := range content {
		itemMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		// Write list marker
		if ordered {
			result.WriteString(fmt.Sprintf("%s%d. ", indent, i+1))
		} else {
			result.WriteString(indent + "- ")
		}

		// Process list item content
		if itemContent, ok := itemMap["content"].([]any); ok {
			for j, contentItem := range itemContent {
				contentMap, ok := contentItem.(map[string]any)
				if !ok {
					continue
				}

				contentType, _ := contentMap["type"].(string)

				if contentType == "paragraph" {
					convertInlineContent(contentMap, result)
					if j == 0 {
						result.WriteString("\n")
					}
				} else if contentType == "bulletList" {
					if j > 0 {
						result.WriteString("\n")
					}
					convertListItemsWithImages(contentMap, result, depth+1, false, attMaps)
				} else if contentType == "orderedList" {
					if j > 0 {
						result.WriteString("\n")
					}
					convertListItemsWithImages(contentMap, result, depth+1, true, attMaps)
				} else if contentType == "mediaSingle" || contentType == "mediaGroup" {
					result.WriteString("\n")
					convertBlockWithImages(contentMap, result, depth+1, attMaps)
				}
			}
		}
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
