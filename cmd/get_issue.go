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

	summaryObject := issue.Fields["summary"]
	summary := fmt.Sprintf("%v", summaryObject)

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()

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

	return nil
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
