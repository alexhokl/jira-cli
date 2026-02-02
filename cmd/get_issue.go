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
	id               string
	showCustomFields bool
}

var getIssueOpts = getIssueOptions{}

// getIssueCmd represents the list command
var getIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Get issue",
	Long: `Get details of a specific issue.

Examples:
  # Get issue details
  jira-cli get issue -i PROJ-123

  # Get issue details including custom fields
  jira-cli get issue -i PROJ-123 --show-custom-fields`,
	RunE: runGetIssue,
}

func init() {
	getCmd.AddCommand(getIssueCmd)

	flags := getIssueCmd.Flags()
	flags.StringVarP(&getIssueOpts.id, "id", "i", "", "Issue ID")
	flags.BoolVarP(&getIssueOpts.showCustomFields, "show-custom-fields", "c", false, "Show custom field values")

	getIssueCmd.MarkFlagRequired("id")
}

func runGetIssue(_ *cobra.Command, _ []string) error {
	client := newClient()
	issue, _, err := client.IssuesAPI.GetIssue(getAuthContext(), getIssueOpts.id).Execute()
	if err != nil {
		return wrapAPIError(err)
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

	descriptionObject := issue.Fields["description"]
	if descriptionObject != nil {
		description := extractTextFromADF(descriptionObject)
		if description != "" {
			fmt.Printf("\n%s\n", description)
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

	// Display custom fields (if flag is set)
	if getIssueOpts.showCustomFields {
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
