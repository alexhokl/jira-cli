package cmd

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type getIssueOptions struct {
	id              string
	descriptionOnly bool
	noImages        bool
	view            bool
}

var getIssueOpts = getIssueOptions{}

// getIssueCmd represents the list command
var getIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Get issue",
	Long: `Get details of a specific issue.

Examples:
  # Get issue details (images displayed by default in supported terminals)
  jira-cli get issue -i PROJ-123

  # Get only the description in markdown format
  jira-cli get issue -i PROJ-123 --description-only

  # Get issue without inline images
  jira-cli get issue -i PROJ-123 --no-images

  # Open the issue in the default browser
  jira-cli get issue -i PROJ-123 --view

Note: Images are displayed inline by default in terminals that support
the Kitty graphics protocol (Ghostty, Kitty, WezTerm).`,
	RunE: runGetIssue,
}

func init() {
	getCmd.AddCommand(getIssueCmd)

	flags := getIssueCmd.Flags()
	flags.StringVarP(&getIssueOpts.id, "id", "i", "", "Issue ID")
	flags.BoolVar(&getIssueOpts.descriptionOnly, "description-only", false, "Output only the description in markdown format")
	flags.BoolVar(&getIssueOpts.noImages, "no-images", false, "Do not display images inline")
	flags.BoolVarP(&getIssueOpts.view, "view", "v", false, "Open the issue in the default browser")

	getIssueCmd.MarkFlagRequired("id")
}

func runGetIssue(_ *cobra.Command, _ []string) error {
	// Handle --view flag: open issue in browser and exit
	if getIssueOpts.view {
		url := getIssueURL(getIssueOpts.id)
		fmt.Printf("Opening %s in browser...\n", url)
		return openBrowser(url)
	}

	client := newClient()
	ctx := getAuthContext()
	issue, _, err := client.IssuesAPI.GetIssue(ctx, getIssueOpts.id).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	// Check if we should show images (default: yes, unless --no-images or terminal doesn't support it)
	showImages := !getIssueOpts.noImages && supportsKittyGraphics()

	// Build attachment maps if showing images
	var attMaps *attachmentMaps
	if showImages {
		attMaps, err = buildAttachmentMaps(client, ctx, getIssueOpts.id)
		if err != nil {
			// Non-fatal: just warn and continue without images
			fmt.Printf("Warning: could not fetch attachments: %v\n\n", err)
			showImages = false
		}
	}

	// Handle description-only mode
	if getIssueOpts.descriptionOnly {
		descriptionObject := issue.Fields["description"]
		if descriptionObject != nil {
			if showImages && attMaps != nil {
				printADFWithImages(descriptionObject, attMaps)
			} else {
				markdown := convertADFToMarkdown(descriptionObject)
				fmt.Print(markdown)
			}
		}
		return nil
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	summaryObject := issue.Fields["summary"]
	summary := fmt.Sprintf("%v", summaryObject)

	summaryLine := fmt.Sprintf("%s %s", *issue.Key, yellow(summary))
	fmt.Println(summaryLine)

	// Calculate dashes based on plain text length (without color codes)
	plainSummaryLine := fmt.Sprintf("%s %s", *issue.Key, summary)
	fmt.Println(strings.Repeat("-", len(plainSummaryLine)))

	// Display status
	status := extractStatusName(issue.Fields["status"])
	if status != "" {
		fmt.Printf("%s %s\n", cyan("Status:"), magenta(status))
	}

	descriptionObject := issue.Fields["description"]
	if descriptionObject != nil {
		if showImages && attMaps != nil {
			fmt.Println()
			printADFWithImages(descriptionObject, attMaps)
		} else {
			description := convertADFToMarkdown(descriptionObject)
			if description != "" {
				fmt.Printf("\n%s", description)
			}
		}
	}

	// Display assignee and reporter
	assignee := extractUserDisplayName(issue.Fields["assignee"])
	reporter := extractUserDisplayName(issue.Fields["reporter"])
	if assignee != "" || reporter != "" {
		fmt.Println()
		if assignee != "" {
			fmt.Printf("%s %s\n", cyan("Assignee:"), yellow(assignee))
		}
		if reporter != "" {
			fmt.Printf("%s %s\n", cyan("Reporter:"), yellow(reporter))
		}
	}

	// Display labels
	labelsObject := issue.Fields["labels"]
	if labelsObject != nil {
		labels := extractLabels(labelsObject)
		if len(labels) > 0 {
			fmt.Printf("\n%s %s\n", cyan("Labels:"), yellow(strings.Join(labels, ", ")))
		}
	}

	// Display custom fields
	fieldNameMap, err := getFieldNameMap(client, getAuthContext())
	if err == nil {
		customFields := extractIssueCustomFields(issue.Fields, fieldNameMap)
		if len(customFields) > 0 {
			fmt.Printf("\n%s\n", cyan("Custom Fields:"))
			for _, cf := range customFields {
				fmt.Printf("  %s: %s\n", yellow(cf.name), cf.value)
			}
		}
	}

	// Display linked issues
	issueLinksObject := issue.Fields["issuelinks"]
	if issueLinksObject != nil {
		links := extractIssueLinks(issueLinksObject, *issue.Key)
		if len(links) > 0 {
			fmt.Printf("\n%s\n", cyan("Linked Issues:"))
			for _, link := range links {
				fmt.Printf("  %s %s %s\n", magenta(link.linkType), yellow(link.key), link.summary)
			}
		}
	}

	// Display subtasks/child issues
	subtasksObject := issue.Fields["subtasks"]
	if subtasksObject != nil {
		subtasks := extractSubtasks(subtasksObject)
		if len(subtasks) > 0 {
			fmt.Printf("\n%s\n", cyan("Sub-tasks:"))
			for _, subtask := range subtasks {
				statusStr := ""
				if subtask.status != "" {
					statusStr = fmt.Sprintf(" [%s]", magenta(subtask.status))
				}
				fmt.Printf("  %s %s%s\n", yellow(subtask.key), subtask.summary, statusStr)
			}
		}
	}

	return nil
}

type issueLink struct {
	linkType string
	key      string
	summary  string
}

type subtaskInfo struct {
	key     string
	summary string
	status  string
}

// extractLabels extracts labels from the labels field
func extractLabels(labelsObj any) []string {
	labelsArray, ok := labelsObj.([]any)
	if !ok {
		return nil
	}

	var labels []string
	for _, labelObj := range labelsArray {
		if label, ok := labelObj.(string); ok {
			labels = append(labels, label)
		}
	}

	return labels
}

// extractStatusName extracts the status name from the status field
func extractStatusName(statusObj any) string {
	if statusObj == nil {
		return ""
	}

	statusMap, ok := statusObj.(map[string]any)
	if !ok {
		return ""
	}

	if name, ok := statusMap["name"].(string); ok {
		return name
	}

	return ""
}

// extractUserDisplayName extracts the display name from a user field (assignee, reporter, etc.)
func extractUserDisplayName(userObj any) string {
	if userObj == nil {
		return ""
	}

	userMap, ok := userObj.(map[string]any)
	if !ok {
		return ""
	}

	if displayName, ok := userMap["displayName"].(string); ok {
		return displayName
	}

	// Fallback to emailAddress if displayName is not available
	if email, ok := userMap["emailAddress"].(string); ok {
		return email
	}

	return ""
}

// extractIssueLinks extracts linked issues from the issuelinks field
func extractIssueLinks(issueLinksObj any, currentKey string) []issueLink {
	linksArray, ok := issueLinksObj.([]any)
	if !ok {
		return nil
	}

	var links []issueLink
	for _, linkObj := range linksArray {
		linkMap, ok := linkObj.(map[string]any)
		if !ok {
			continue
		}

		var link issueLink

		// Get link type
		typeObj, ok := linkMap["type"].(map[string]any)
		if !ok {
			continue
		}

		// Determine if this is an inward or outward link
		if inwardIssue, ok := linkMap["inwardIssue"].(map[string]any); ok {
			// Current issue is the outward issue, linked issue is inward
			if inward, ok := typeObj["inward"].(string); ok {
				link.linkType = inward
			}
			if key, ok := inwardIssue["key"].(string); ok {
				link.key = key
			}
			if fields, ok := inwardIssue["fields"].(map[string]any); ok {
				if summary, ok := fields["summary"].(string); ok {
					link.summary = summary
				}
			}
		} else if outwardIssue, ok := linkMap["outwardIssue"].(map[string]any); ok {
			// Current issue is the inward issue, linked issue is outward
			if outward, ok := typeObj["outward"].(string); ok {
				link.linkType = outward
			}
			if key, ok := outwardIssue["key"].(string); ok {
				link.key = key
			}
			if fields, ok := outwardIssue["fields"].(map[string]any); ok {
				if summary, ok := fields["summary"].(string); ok {
					link.summary = summary
				}
			}
		}

		if link.key != "" && link.key != currentKey {
			links = append(links, link)
		}
	}

	return links
}

// extractSubtasks extracts subtasks/child issues from the subtasks field
func extractSubtasks(subtasksObj any) []subtaskInfo {
	subtasksArray, ok := subtasksObj.([]any)
	if !ok {
		return nil
	}

	var subtasks []subtaskInfo
	for _, subtaskObj := range subtasksArray {
		subtaskMap, ok := subtaskObj.(map[string]any)
		if !ok {
			continue
		}

		var subtask subtaskInfo

		if key, ok := subtaskMap["key"].(string); ok {
			subtask.key = key
		}

		if fields, ok := subtaskMap["fields"].(map[string]any); ok {
			if summary, ok := fields["summary"].(string); ok {
				subtask.summary = summary
			}
			if status, ok := fields["status"].(map[string]any); ok {
				if name, ok := status["name"].(string); ok {
					subtask.status = name
				}
			}
		}

		if subtask.key != "" {
			subtasks = append(subtasks, subtask)
		}
	}

	return subtasks
}

type customFieldValue struct {
	name  string
	value string
}

// getFieldNameMap fetches all fields and returns a map of custom field ID -> name
func getFieldNameMap(client *swagger.APIClient, ctx context.Context) (map[string]string, error) {
	fields, _, err := client.IssueFieldsAPI.GetFields(ctx).Execute()
	if err != nil {
		return nil, err
	}

	nameMap := make(map[string]string)
	for _, field := range fields {
		if field.GetCustom() {
			nameMap[field.GetId()] = field.GetName()
		}
	}
	return nameMap, nil
}

// extractIssueCustomFields extracts custom field values from issue fields
func extractIssueCustomFields(fields map[string]any, nameMap map[string]string) []customFieldValue {
	var result []customFieldValue

	for fieldId, value := range fields {
		if !strings.HasPrefix(fieldId, "customfield_") {
			continue
		}
		if value == nil {
			continue
		}

		name := nameMap[fieldId]
		if name == "" {
			name = fieldId
		}

		valueStr := formatCustomFieldValue(value)
		if valueStr != "" {
			result = append(result, customFieldValue{name: name, value: valueStr})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].name < result[j].name
	})

	return result
}

// formatCustomFieldValue converts various custom field value types to string
func formatCustomFieldValue(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d", int64(v))
		}
		return fmt.Sprintf("%v", v)
	case bool:
		return fmt.Sprintf("%v", v)
	case map[string]any:
		if name, ok := v["displayName"].(string); ok {
			return name
		}
		if val, ok := v["value"].(string); ok {
			return val
		}
		if name, ok := v["name"].(string); ok {
			return name
		}
		return fmt.Sprintf("%v", v)
	case []any:
		var values []string
		for _, item := range v {
			if m, ok := item.(map[string]any); ok {
				if val, ok := m["value"].(string); ok {
					values = append(values, val)
				} else if name, ok := m["displayName"].(string); ok {
					values = append(values, name)
				} else if name, ok := m["name"].(string); ok {
					values = append(values, name)
				}
			} else if s, ok := item.(string); ok {
				values = append(values, s)
			}
		}
		if len(values) > 0 {
			return strings.Join(values, ", ")
		}
		return ""
	default:
		return fmt.Sprintf("%v", value)
	}
}

// extractTextFromADF extracts plain text from Atlassian Document Format (ADF)
func extractTextFromADF(adf any) string {
	adfMap, ok := adf.(map[string]any)
	if !ok {
		return ""
	}

	content, ok := adfMap["content"].([]any)
	if !ok {
		return ""
	}

	var texts []string
	for _, block := range content {
		text := extractTextFromBlock(block)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, "\n")
}

// extractTextFromBlock recursively extracts text from an ADF block
func extractTextFromBlock(block interface{}) string {
	blockMap, ok := block.(map[string]interface{})
	if !ok {
		return ""
	}

	// If this block has direct text content
	if text, ok := blockMap["text"].(string); ok {
		return text
	}

	// If this block has nested content, recursively extract
	content, ok := blockMap["content"].([]interface{})
	if !ok {
		return ""
	}

	var texts []string
	for _, item := range content {
		text := extractTextFromBlock(item)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, "")
}

// convertADFToMarkdown converts Atlassian Document Format (ADF) to markdown.
// This preserves the markdown formatting used in JIRA.
func convertADFToMarkdown(adf any) string {
	adfMap, ok := adf.(map[string]any)
	if !ok {
		return ""
	}

	content, ok := adfMap["content"].([]any)
	if !ok {
		return ""
	}

	var result strings.Builder
	for i, block := range content {
		if i > 0 {
			result.WriteString("\n")
		}
		convertBlockToMarkdown(block, &result, 0)
	}

	return result.String()
}

// convertBlockToMarkdown converts a single ADF block to markdown
func convertBlockToMarkdown(block any, result *strings.Builder, depth int) {
	blockMap, ok := block.(map[string]any)
	if !ok {
		return
	}

	blockType, _ := blockMap["type"].(string)

	switch blockType {
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
		convertListItems(blockMap, result, depth, false)

	case "orderedList":
		convertListItems(blockMap, result, depth, true)

	case "listItem":
		// List items are handled by convertListItems
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				convertBlockToMarkdown(item, result, depth)
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
				convertBlockToMarkdown(item, result, depth)
			}
		}

	case "rule":
		result.WriteString("---\n")

	case "table":
		convertTable(blockMap, result)

	case "mediaSingle", "mediaGroup":
		// Media elements - output placeholder
		if content, ok := blockMap["content"].([]any); ok {
			for _, item := range content {
				if mediaMap, ok := item.(map[string]any); ok {
					if mediaMap["type"] == "media" {
						if attrs, ok := mediaMap["attrs"].(map[string]any); ok {
							if alt, ok := attrs["alt"].(string); ok && alt != "" {
								result.WriteString(fmt.Sprintf("![%s](media)\n", alt))
							} else {
								result.WriteString("![media](media)\n")
							}
						}
					}
				}
			}
		}

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
				convertBlockToMarkdown(item, result, depth)
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

// convertInlineContent converts inline content (text, marks) to markdown
func convertInlineContent(blockMap map[string]any, result *strings.Builder) {
	content, ok := blockMap["content"].([]any)
	if !ok {
		return
	}

	for _, item := range content {
		itemMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		itemType, _ := itemMap["type"].(string)

		switch itemType {
		case "text":
			text, _ := itemMap["text"].(string)
			text = applyMarks(text, itemMap)
			result.WriteString(text)

		case "hardBreak":
			result.WriteString("\n")

		case "emoji":
			if attrs, ok := itemMap["attrs"].(map[string]any); ok {
				if shortName, ok := attrs["shortName"].(string); ok {
					result.WriteString(shortName)
				}
			}

		case "mention":
			if attrs, ok := itemMap["attrs"].(map[string]any); ok {
				if text, ok := attrs["text"].(string); ok {
					result.WriteString(text)
				}
			}

		case "inlineCard":
			if attrs, ok := itemMap["attrs"].(map[string]any); ok {
				if url, ok := attrs["url"].(string); ok {
					result.WriteString(url)
				}
			}
		}
	}
}

// applyMarks applies text formatting marks (bold, italic, etc.)
func applyMarks(text string, itemMap map[string]any) string {
	marks, ok := itemMap["marks"].([]any)
	if !ok {
		return text
	}

	// Track which marks to apply
	var hasBold, hasItalic, hasStrike, hasCode, hasUnderline bool
	var linkURL string

	for _, mark := range marks {
		markMap, ok := mark.(map[string]any)
		if !ok {
			continue
		}

		markType, _ := markMap["type"].(string)

		switch markType {
		case "strong":
			hasBold = true
		case "em":
			hasItalic = true
		case "strike":
			hasStrike = true
		case "code":
			hasCode = true
		case "underline":
			hasUnderline = true
		case "link":
			if attrs, ok := markMap["attrs"].(map[string]any); ok {
				if href, ok := attrs["href"].(string); ok {
					linkURL = href
				}
			}
		}
	}

	// Apply marks (order matters for proper nesting)
	if hasCode {
		text = "`" + text + "`"
	}
	if hasStrike {
		text = "~~" + text + "~~"
	}
	if hasItalic {
		text = "*" + text + "*"
	}
	if hasBold {
		text = "**" + text + "**"
	}
	if hasUnderline {
		// Markdown doesn't have native underline, use HTML
		text = "<u>" + text + "</u>"
	}
	if linkURL != "" {
		text = "[" + text + "](" + linkURL + ")"
	}

	return text
}

// convertListItems converts list items to markdown
func convertListItems(blockMap map[string]any, result *strings.Builder, depth int, ordered bool) {
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
					convertListItems(contentMap, result, depth+1, false)
				} else if contentType == "orderedList" {
					if j > 0 {
						result.WriteString("\n")
					}
					convertListItems(contentMap, result, depth+1, true)
				}
			}
		}
	}
}

// convertTable converts ADF table to markdown table
func convertTable(blockMap map[string]any, result *strings.Builder) {
	content, ok := blockMap["content"].([]any)
	if !ok {
		return
	}

	var rows [][]string
	var maxCols int

	// Extract all cells
	for _, row := range content {
		rowMap, ok := row.(map[string]any)
		if !ok {
			continue
		}

		if rowMap["type"] != "tableRow" {
			continue
		}

		rowContent, ok := rowMap["content"].([]any)
		if !ok {
			continue
		}

		var cells []string
		for _, cell := range rowContent {
			cellMap, ok := cell.(map[string]any)
			if !ok {
				continue
			}

			var cellText strings.Builder
			if cellContent, ok := cellMap["content"].([]any); ok {
				for _, item := range cellContent {
					itemMap, ok := item.(map[string]any)
					if !ok {
						continue
					}
					if itemMap["type"] == "paragraph" {
						convertInlineContent(itemMap, &cellText)
					}
				}
			}
			cells = append(cells, cellText.String())
		}

		if len(cells) > maxCols {
			maxCols = len(cells)
		}
		rows = append(rows, cells)
	}

	if len(rows) == 0 {
		return
	}

	// Write header row
	if len(rows) > 0 {
		result.WriteString("| " + strings.Join(rows[0], " | ") + " |\n")
		// Write separator
		sep := make([]string, len(rows[0]))
		for i := range sep {
			sep[i] = "---"
		}
		result.WriteString("| " + strings.Join(sep, " | ") + " |\n")
	}

	// Write data rows
	for i := 1; i < len(rows); i++ {
		result.WriteString("| " + strings.Join(rows[i], " | ") + " |\n")
	}
}
