package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listCustomFieldValuesOptions struct {
	project string
	name    string
}

var listCustomFieldValuesOpts = listCustomFieldValuesOptions{}

var listCustomFieldValuesCmd = &cobra.Command{
	Use:   "custom-field-values",
	Short: "List values of a custom field",
	Long: `List all available values (options) for a custom field.

This command is useful for viewing the allowed values of select list custom fields.

Examples:
  # List all values for a custom field by name
  jira-cli list custom-field-values --name "Priority Level"

  # List values for a custom field in a specific project context
  jira-cli list custom-field-values --name "Team" --project PROJ`,
	RunE: runListCustomFieldValues,
}

func init() {
	listCmd.AddCommand(listCustomFieldValuesCmd)

	flags := listCustomFieldValuesCmd.Flags()
	flags.StringVarP(&listCustomFieldValuesOpts.project, "project", "p", "", "Filter by project key")
	flags.StringVarP(&listCustomFieldValuesOpts.name, "name", "n", "", "Custom field name (required)")
	listCustomFieldValuesCmd.MarkFlagRequired("name")
}

func runListCustomFieldValues(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	// Find the custom field by name
	fieldId, err := findCustomFieldIdByName(client, ctx, listCustomFieldValuesOpts.name)
	if err != nil {
		return err
	}

	// Get contexts for the field
	contexts, err := getFieldContexts(client, ctx, fieldId, listCustomFieldValuesOpts.project)
	if err != nil {
		return err
	}

	if len(contexts) == 0 {
		fmt.Println("No contexts found for this custom field")
		return nil
	}

	// Get options for each context
	var allOptions []customFieldOptionValue
	for _, context := range contexts {
		contextId, err := parseContextId(context.GetId())
		if err != nil {
			continue
		}

		options, err := getOptionsForContext(client, ctx, fieldId, contextId)
		if err != nil {
			// Skip contexts that fail (might not have options)
			continue
		}

		for _, opt := range options {
			allOptions = append(allOptions, customFieldOptionValue{
				id:          opt.GetId(),
				value:       opt.GetValue(),
				disabled:    opt.GetDisabled(),
				contextName: context.GetName(),
			})
		}
	}

	if len(allOptions) == 0 {
		fmt.Println("No values found for this custom field")
		return nil
	}

	// Display results
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("VALUE"), cyan("STATUS"), cyan("CONTEXT"))

	for _, opt := range allOptions {
		status := "Active"
		if opt.disabled {
			status = red("Disabled")
		}
		w.row(opt.id, yellow(opt.value), status, opt.contextName)
	}
	w.flush()

	fmt.Printf("\nFound %d values\n", len(allOptions))

	return nil
}

type customFieldOptionValue struct {
	id          string
	value       string
	disabled    bool
	contextName string
}

// findCustomFieldIdByName finds a custom field ID by its name
func findCustomFieldIdByName(client *swagger.APIClient, ctx context.Context, name string) (string, error) {
	fields, _, err := client.IssueFieldsAPI.GetFields(ctx).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to get fields: %w", err)
	}

	for _, field := range fields {
		if field.GetCustom() && strings.EqualFold(field.GetName(), name) {
			return field.GetId(), nil
		}
	}

	return "", fmt.Errorf("custom field %q not found", name)
}

// getFieldContexts returns contexts for a custom field, optionally filtered by project
func getFieldContexts(client *swagger.APIClient, ctx context.Context, fieldId string, projectKey string) ([]swagger.CustomFieldContext, error) {
	var allContexts []swagger.CustomFieldContext
	var startAt int64 = 0

	for {
		request := client.IssueCustomFieldContextsAPI.GetContextsForField(ctx, fieldId).
			StartAt(startAt).
			MaxResults(50)

		result, resp, err := request.Execute()
		if err != nil {
			if resp != nil && resp.StatusCode == 403 {
				return nil, fmt.Errorf("permission denied: this operation requires Jira administrator permissions")
			}
			return nil, fmt.Errorf("failed to get contexts: %w", err)
		}

		allContexts = append(allContexts, result.GetValues()...)

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	// Filter by project if specified
	if projectKey != "" {
		// Get project ID
		project, _, err := client.ProjectsAPI.GetProject(ctx, projectKey).Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to get project: %w", err)
		}
		projectId, err := parseProjectId(project.GetId())
		if err != nil {
			return nil, err
		}

		// Get project context mappings
		projectContexts, err := getProjectContextMappings(client, ctx, fieldId, projectId)
		if err != nil {
			// If we can't get mappings, return global contexts only
			var globalContexts []swagger.CustomFieldContext
			for _, c := range allContexts {
				if c.GetIsGlobalContext() {
					globalContexts = append(globalContexts, c)
				}
			}
			return globalContexts, nil
		}

		// Filter to contexts that apply to this project
		contextIdSet := make(map[string]bool)
		for _, pc := range projectContexts {
			contextIdSet[pc.GetContextId()] = true
		}

		var filteredContexts []swagger.CustomFieldContext
		for _, c := range allContexts {
			if c.GetIsGlobalContext() || contextIdSet[c.GetId()] {
				filteredContexts = append(filteredContexts, c)
			}
		}
		return filteredContexts, nil
	}

	return allContexts, nil
}

// getProjectContextMappings returns project context mappings for a field filtered by project ID
func getProjectContextMappings(client *swagger.APIClient, ctx context.Context, fieldId string, projectId int64) ([]swagger.CustomFieldContextProjectMapping, error) {
	var allMappings []swagger.CustomFieldContextProjectMapping
	var startAt int64 = 0

	projectIdStr := fmt.Sprintf("%d", projectId)

	for {
		result, _, err := client.IssueCustomFieldContextsAPI.GetProjectContextMapping(ctx, fieldId).
			StartAt(startAt).
			MaxResults(50).
			Execute()
		if err != nil {
			return nil, err
		}

		// Filter mappings for the specified project
		for _, mapping := range result.GetValues() {
			if mapping.GetProjectId() == projectIdStr || mapping.GetIsGlobalContext() {
				allMappings = append(allMappings, mapping)
			}
		}

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	return allMappings, nil
}

// getOptionsForContext returns all options for a custom field context
func getOptionsForContext(client *swagger.APIClient, ctx context.Context, fieldId string, contextId int64) ([]swagger.CustomFieldContextOption, error) {
	var allOptions []swagger.CustomFieldContextOption
	var startAt int64 = 0

	for {
		result, _, err := client.IssueCustomFieldOptionsAPI.GetOptionsForContext(ctx, fieldId, contextId).
			StartAt(startAt).
			MaxResults(100).
			OnlyOptions(true).
			Execute()
		if err != nil {
			return nil, err
		}

		allOptions = append(allOptions, result.GetValues()...)

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	return allOptions, nil
}

// parseContextId parses a context ID string to int64
func parseContextId(id string) (int64, error) {
	var contextId int64
	_, err := fmt.Sscanf(id, "%d", &contextId)
	if err != nil {
		return 0, fmt.Errorf("failed to parse context ID: %w", err)
	}
	return contextId, nil
}
