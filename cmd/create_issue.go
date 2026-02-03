package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type createIssueOptions struct {
	project         string
	issueType       string
	summary         string
	descriptionFile string
	priority        string
	assignee        string
	labels          string
	components      string
	parent          string
	dueDate         string
	customFields    []string
}

var createIssueOpts = createIssueOptions{}

var createIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Create an issue",
	Long: `Create a new issue in Jira.

Examples:
  # Create a basic issue
  jira-cli create issue -p PROJ -t Task -s "My task"

  # Create an issue with description from file
  jira-cli create issue -p PROJ -t Bug -s "Bug title" -d description.md

  # Create an issue with custom fields
  jira-cli create issue -p PROJ -t Story -s "My story" --custom-field "Team=Backend"

  # Create an issue with multiple custom fields
  jira-cli create issue -p PROJ -t Task -s "My task" --custom-field "Team=Backend" --custom-field "Story Points=5"

  # Create an issue with a select list custom field
  jira-cli create issue -p PROJ -t Task -s "My task" --custom-field "Environment=Production"`,
	RunE: runCreateIssue,
}

func init() {
	createCmd.AddCommand(createIssueCmd)

	flags := createIssueCmd.Flags()
	flags.StringVarP(&createIssueOpts.project, "project", "p", "", "Project key (required)")
	flags.StringVarP(&createIssueOpts.issueType, "type", "t", "", "Issue type (e.g., Bug, Task, Story) (required)")
	flags.StringVarP(&createIssueOpts.summary, "summary", "s", "", "Issue summary/title (required)")
	flags.StringVarP(&createIssueOpts.descriptionFile, "description-file", "d", "", "Path to a file containing issue description in markdown")
	flags.StringVar(&createIssueOpts.priority, "priority", "", "Priority (e.g., Highest, High, Medium, Low, Lowest)")
	flags.StringVarP(&createIssueOpts.assignee, "assignee", "a", "", "Assignee account ID (use 'me' for yourself)")
	flags.StringVarP(&createIssueOpts.labels, "labels", "l", "", "Comma-separated labels")
	flags.StringVarP(&createIssueOpts.components, "components", "c", "", "Comma-separated component names")
	flags.StringVar(&createIssueOpts.parent, "parent", "", "Parent issue key (for subtasks)")
	flags.StringVar(&createIssueOpts.dueDate, "due-date", "", "Due date (format: 2006-01-02)")
	flags.StringArrayVar(&createIssueOpts.customFields, "custom-field", nil, "Custom field in format 'name=value' (can be specified multiple times)")

	createIssueCmd.MarkFlagRequired("project")
	createIssueCmd.MarkFlagRequired("type")
	createIssueCmd.MarkFlagRequired("summary")
}

func runCreateIssue(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	description := ""
	if createIssueOpts.descriptionFile != "" {
		content, err := os.ReadFile(createIssueOpts.descriptionFile)
		if err != nil {
			return fmt.Errorf("failed to read description file: %w", err)
		}
		description = strings.TrimSpace(string(content))
	}

	fields := map[string]any{
		"project": map[string]any{
			"key": createIssueOpts.project,
		},
		"issuetype": map[string]any{
			"name": createIssueOpts.issueType,
		},
		"summary": createIssueOpts.summary,
	}

	if strings.TrimSpace(description) != "" {
		fields["description"] = newADFDocument(description)
	}

	if createIssueOpts.priority != "" {
		fields["priority"] = map[string]any{
			"name": createIssueOpts.priority,
		}
	}

	if createIssueOpts.assignee != "" {
		assigneeId, err := resolveUserAlias(createIssueOpts.assignee)
		if err != nil {
			return fmt.Errorf("failed to resolve assignee: %w", err)
		}
		fields["assignee"] = map[string]any{
			"id": assigneeId,
		}
	}

	if createIssueOpts.labels != "" {
		labelList := strings.Split(createIssueOpts.labels, ",")
		for i, label := range labelList {
			labelList[i] = strings.TrimSpace(label)
		}
		fields["labels"] = labelList
	}

	if createIssueOpts.components != "" {
		componentList := strings.Split(createIssueOpts.components, ",")
		components := make([]map[string]any, len(componentList))
		for i, comp := range componentList {
			components[i] = map[string]any{
				"name": strings.TrimSpace(comp),
			}
		}
		fields["components"] = components
	}

	if createIssueOpts.parent != "" {
		fields["parent"] = map[string]any{
			"key": createIssueOpts.parent,
		}
	}

	if createIssueOpts.dueDate != "" {
		fields["duedate"] = createIssueOpts.dueDate
	}

	// Handle custom fields
	if len(createIssueOpts.customFields) > 0 {
		fieldInfoMap, err := getCustomFieldInfoMap(client, ctx)
		if err != nil {
			return fmt.Errorf("failed to get custom field information: %w", err)
		}

		for _, cf := range createIssueOpts.customFields {
			// Parse the name=value format
			eqIdx := strings.Index(cf, "=")
			if eqIdx == -1 {
				return fmt.Errorf("invalid custom field format %q: expected 'name=value'", cf)
			}
			fieldName := strings.TrimSpace(cf[:eqIdx])
			fieldValue := cf[eqIdx+1:]

			// Find the field by name
			fieldInfo, found := findCustomFieldByName(fieldInfoMap, fieldName)
			if !found {
				return fmt.Errorf("custom field %q not found", fieldName)
			}

			// Enrich field info with allowed values for option-type fields
			enrichFieldInfoWithAllowedValues(client, ctx, &fieldInfo, createIssueOpts.project, createIssueOpts.issueType)

			// Convert the value to the appropriate format for the field type
			// Pass project key for sprint name resolution
			convertedValue, err := convertCustomFieldValue(fieldValue, fieldInfo, createIssueOpts.project)
			if err != nil {
				return fmt.Errorf("failed to convert value for custom field %q: %w", fieldName, err)
			}

			fields[fieldInfo.id] = convertedValue
		}
	}

	issueUpdateDetails := swagger.NewIssueUpdateDetails()
	issueUpdateDetails.SetFields(fields)

	createdIssue, _, err := client.IssuesAPI.CreateIssue(ctx).IssueUpdateDetails(*issueUpdateDetails).Execute()
	if err != nil {
		var openAPIErr *swagger.GenericOpenAPIError
		if errors.As(err, &openAPIErr) {
			body := string(openAPIErr.Body())
			if body != "" {
				return fmt.Errorf("%s\n\n%s", err, body)
			}
		}
		return err
	}

	fmt.Printf("Issue created: %s\n", createdIssue.GetKey())
	return nil
}

// Note: newADFDocument is defined in helper.go
