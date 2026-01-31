package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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
	client := newClient()
	issue, _, err := client.IssuesAPI.GetIssue(getAuthContext(), getIssueOpts.id).Execute()
	if err != nil {
		return err
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
