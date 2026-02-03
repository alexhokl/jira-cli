package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type bulkUpdateCustomFieldValueOptions struct {
	name    string
	oldVal  string
	newVal  string
	project string
	dryRun  bool
}

var bulkUpdateCustomFieldValueOpts = bulkUpdateCustomFieldValueOptions{}

var bulkUpdateCustomFieldValueCmd = &cobra.Command{
	Use:   "custom-field-value",
	Short: "Update a custom field value across all matching issues",
	Long: `Search for all issues with a specific custom field value and replace it with a new value.

This command will:
1. Search for all issues that have the specified custom field with the old value
2. For each issue, update the custom field to the new value

Examples:
  # Update custom field 'Team' from 'Platform' to 'Infrastructure'
  jira-cli bulk-update custom-field-value --name Team --old Platform --new Infrastructure

  # Update custom field only in a specific project
  jira-cli bulk-update custom-field-value --name Team --old Platform --new Infrastructure --project MYPROJ

  # Preview changes without applying them
  jira-cli bulk-update custom-field-value --name Team --old Platform --new Infrastructure --dry-run

  # Update using custom field ID notation
  jira-cli bulk-update custom-field-value --name "cf[10001]" --old "Old Value" --new "New Value"`,
	RunE: runBulkUpdateCustomFieldValue,
}

func init() {
	bulkUpdateCmd.AddCommand(bulkUpdateCustomFieldValueCmd)

	flags := bulkUpdateCustomFieldValueCmd.Flags()
	flags.StringVar(&bulkUpdateCustomFieldValueOpts.name, "name", "", "Custom field name (required)")
	flags.StringVar(&bulkUpdateCustomFieldValueOpts.oldVal, "old", "", "Old value to search for (required)")
	flags.StringVar(&bulkUpdateCustomFieldValueOpts.newVal, "new", "", "New value to replace with (required)")
	flags.StringVarP(&bulkUpdateCustomFieldValueOpts.project, "project", "p", "", "Filter by project key")
	flags.BoolVar(&bulkUpdateCustomFieldValueOpts.dryRun, "dry-run", false, "Preview changes without applying them")

	bulkUpdateCustomFieldValueCmd.MarkFlagRequired("name")
	bulkUpdateCustomFieldValueCmd.MarkFlagRequired("old")
	bulkUpdateCustomFieldValueCmd.MarkFlagRequired("new")
}

func runBulkUpdateCustomFieldValue(_ *cobra.Command, _ []string) error {
	fieldName := strings.TrimSpace(bulkUpdateCustomFieldValueOpts.name)
	oldValue := bulkUpdateCustomFieldValueOpts.oldVal
	newValue := bulkUpdateCustomFieldValueOpts.newVal

	if fieldName == "" {
		return fmt.Errorf("custom field name cannot be empty")
	}
	if oldValue == newValue {
		return fmt.Errorf("old and new values are the same")
	}

	client := newClient()
	ctx := getAuthContext()

	// Get custom field info to resolve name to ID and get field type
	fieldInfoMap, err := getCustomFieldInfoMap(client, ctx)
	if err != nil {
		return fmt.Errorf("failed to get custom field information: %w", err)
	}

	fieldInfo, found := findCustomFieldByName(fieldInfoMap, fieldName)
	if !found {
		// Check if it's already in cf[xxxxx] notation
		if !strings.HasPrefix(fieldName, "cf[") {
			return fmt.Errorf("custom field %q not found", fieldName)
		}
	}

	// Build JQL to find issues with the old value
	// Use the field name directly in JQL (similar to list issues --custom-field)
	jqlFieldName := fieldName
	if strings.Contains(fieldName, " ") {
		jqlFieldName = fmt.Sprintf("\"%s\"", fieldName)
	}

	jql := fmt.Sprintf("%s = %s", jqlFieldName, quoteJQLValue(oldValue))
	if bulkUpdateCustomFieldValueOpts.project != "" {
		jql = fmt.Sprintf("project = %s AND %s", quoteJQLValue(bulkUpdateCustomFieldValueOpts.project), jql)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	searchMsg := fmt.Sprintf("Searching for issues with '%s' = '%s'", cyan(fieldName), cyan(oldValue))
	if bulkUpdateCustomFieldValueOpts.project != "" {
		searchMsg += fmt.Sprintf(" in project '%s'", cyan(bulkUpdateCustomFieldValueOpts.project))
	}
	fmt.Println(searchMsg + "...")

	// Collect all issues with the old value
	var allIssues []swagger.IssueBean
	var nextPageToken string
	const pageSize int32 = 100

	for {
		request := client.IssueSearchAPI.SearchAndReconsileIssuesUsingJql(ctx).
			Jql(jql).
			MaxResults(pageSize).
			Fields([]string{"summary", "issuetype"})

		if nextPageToken != "" {
			request = request.NextPageToken(nextPageToken)
		}

		result, _, err := request.Execute()
		if err != nil {
			return fmt.Errorf("failed to search issues: %w", err)
		}

		allIssues = append(allIssues, result.GetIssues()...)

		if result.GetIsLast() {
			break
		}

		nextPageToken = result.GetNextPageToken()
		if nextPageToken == "" {
			break
		}
	}

	if len(allIssues) == 0 {
		fmt.Printf("No issues found with '%s' = '%s'\n", fieldName, oldValue)
		return nil
	}

	fmt.Printf("Found %d issue(s) with '%s' = '%s'\n\n", len(allIssues), cyan(fieldName), cyan(oldValue))

	if bulkUpdateCustomFieldValueOpts.dryRun {
		fmt.Println("DRY RUN - No changes will be made")
		fmt.Println("The following issues would be updated:")
		for _, issue := range allIssues {
			key := issue.GetKey()
			fields := issue.GetFields()
			summary := ""
			if summaryObj, ok := fields["summary"]; ok && summaryObj != nil {
				summary = fmt.Sprintf("%v", summaryObj)
			}
			fmt.Printf("  %s %s\n", yellow(key), summary)
		}
		fmt.Printf("\nTotal: %d issue(s) would have '%s' changed from '%s' to '%s'\n",
			len(allIssues), cyan(fieldName), cyan(oldValue), green(newValue))
		return nil
	}

	// Update each issue with the new value
	successCount := 0
	failCount := 0

	for _, issue := range allIssues {
		key := issue.GetKey()
		issueFields := issue.GetFields()
		summary := ""
		if summaryObj, ok := issueFields["summary"]; ok && summaryObj != nil {
			summary = fmt.Sprintf("%v", summaryObj)
		}

		// Get issue type for allowed values lookup
		issueTypeName := ""
		if typeObj, ok := issueFields["issuetype"]; ok && typeObj != nil {
			if typeMap, ok := typeObj.(map[string]interface{}); ok {
				if name, ok := typeMap["name"].(string); ok {
					issueTypeName = name
				}
			}
		}

		// Extract project key from issue key
		projectKey := extractProjectKeyFromIssueKey(key)

		// Enrich field info with allowed values for option-type fields
		if issueTypeName != "" && found {
			enrichFieldInfoWithAllowedValues(client, ctx, &fieldInfo, projectKey, issueTypeName)
		}

		// Convert the new value to the appropriate format for the field type
		var convertedValue interface{}
		if found {
			var err error
			convertedValue, err = convertCustomFieldValue(newValue, fieldInfo, projectKey)
			if err != nil {
				fmt.Printf("  %s %s - %s\n", yellow(key), summary, color.RedString("FAILED: %v", err))
				failCount++
				continue
			}
		} else {
			// For cf[xxxxx] notation without resolution, use value as-is
			convertedValue = newValue
		}

		// Build the update fields
		updateFields := make(map[string]interface{})
		if found {
			updateFields[fieldInfo.id] = convertedValue
		} else {
			// Use the field name directly (e.g., cf[10001])
			updateFields[fieldName] = convertedValue
		}

		issueUpdateDetails := swagger.NewIssueUpdateDetails()
		issueUpdateDetails.SetFields(updateFields)

		_, _, err := client.IssuesAPI.EditIssue(ctx, key).
			IssueUpdateDetails(*issueUpdateDetails).
			Execute()

		if err != nil {
			fmt.Printf("  %s %s - %s\n", yellow(key), summary, color.RedString("FAILED: %v", err))
			failCount++
		} else {
			fmt.Printf("  %s %s - %s\n", yellow(key), summary, green("OK"))
			successCount++
		}
	}

	fmt.Printf("\nCompleted: %d succeeded, %d failed\n", successCount, failCount)
	fmt.Printf("Custom field '%s' has been updated from '%s' to '%s'\n", cyan(fieldName), cyan(oldValue), green(newValue))

	return nil
}
