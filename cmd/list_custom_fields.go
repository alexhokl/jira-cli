package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listCustomFieldsOptions struct {
	query        string
	project      string
	issueType    string
	showContexts bool
}

var listCustomFieldsOpts = listCustomFieldsOptions{}

var listCustomFieldsCmd = &cobra.Command{
	Use:   "custom-fields",
	Short: "List custom fields",
	Long: `List all custom fields available in the Jira instance.

Examples:
  # List all custom fields
  jira-cli list custom-fields

  # Search custom fields by name
  jira-cli list custom-fields --query "story"

  # List custom fields available for a specific project
  jira-cli list custom-fields --project PROJ

  # List custom fields available for a specific issue type (requires admin)
  jira-cli list custom-fields --issue-type "Story"

  # List custom fields for a project and issue type combination
  jira-cli list custom-fields --project PROJ --issue-type "Story"

  # Show issue type contexts for each custom field (requires admin)
  jira-cli list custom-fields --show-contexts`,
	RunE: runListCustomFields,
}

func init() {
	listCmd.AddCommand(listCustomFieldsCmd)

	flags := listCustomFieldsCmd.Flags()
	flags.StringVarP(&listCustomFieldsOpts.query, "query", "q", "", "Filter fields by name (partial match)")
	flags.StringVarP(&listCustomFieldsOpts.project, "project", "p", "", "Filter fields by project key")
	flags.StringVarP(&listCustomFieldsOpts.issueType, "issue-type", "t", "", "Filter fields by issue type name")
	flags.BoolVarP(&listCustomFieldsOpts.showContexts, "show-contexts", "c", false, "Show issue type mapping details")
}

func runListCustomFields(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	var customFields []customFieldDetails

	// Determine which API to use based on filters
	if listCustomFieldsOpts.project != "" && listCustomFieldsOpts.issueType != "" {
		// Use GetProjectFields API (most efficient for project+issue type combo)
		fields, err := getCustomFieldsForProjectAndIssueType(client, ctx,
			listCustomFieldsOpts.project, listCustomFieldsOpts.issueType, listCustomFieldsOpts.query)
		if err != nil {
			return err
		}
		customFields = fields
	} else if listCustomFieldsOpts.issueType != "" {
		// Filter by issue type only - requires fetching context mappings (admin permission)
		fields, err := getCustomFieldsByIssueType(client, ctx,
			listCustomFieldsOpts.issueType, listCustomFieldsOpts.query)
		if err != nil {
			return err
		}
		customFields = fields
	} else if listCustomFieldsOpts.project != "" || listCustomFieldsOpts.query != "" {
		// Use paginated API with filters
		fields, err := getCustomFieldsPaginated(client, ctx,
			listCustomFieldsOpts.project, listCustomFieldsOpts.query)
		if err != nil {
			return err
		}
		customFields = fields
	} else {
		// Get all custom fields using basic API
		fields, err := getAllCustomFields(client, ctx)
		if err != nil {
			return err
		}
		customFields = fields
	}

	if len(customFields) == 0 {
		fmt.Println("No custom fields found")
		return nil
	}

	// Fetch context mappings if requested
	if listCustomFieldsOpts.showContexts {
		for i := range customFields {
			issueTypes, err := getIssueTypesForField(client, ctx, customFields[i].id)
			if err != nil {
				// Check if it's a permission error
				if strings.Contains(err.Error(), "403") {
					customFields[i].issueTypes = "(requires admin)"
				} else {
					customFields[i].issueTypes = "(error)"
				}
			} else {
				customFields[i].issueTypes = issueTypes
			}
		}
	}

	// Sort by name
	sort.Slice(customFields, func(i, j int) bool {
		return strings.ToLower(customFields[i].name) < strings.ToLower(customFields[j].name)
	})

	// Display results
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)

	if listCustomFieldsOpts.showContexts {
		w.row(cyan("ID"), cyan("NAME"), cyan("TYPE"), cyan("SCHEMA"), cyan("ISSUE TYPES"))
		for _, f := range customFields {
			w.row(f.id, yellow(f.name), f.fieldType, f.schema, f.issueTypes)
		}
	} else {
		w.row(cyan("ID"), cyan("NAME"), cyan("TYPE"), cyan("SCHEMA"))
		for _, f := range customFields {
			w.row(f.id, yellow(f.name), f.fieldType, f.schema)
		}
	}
	w.flush()

	fmt.Printf("\nFound %d custom fields\n", len(customFields))

	return nil
}

type customFieldDetails struct {
	id         string
	name       string
	fieldType  string
	schema     string
	issueTypes string
}

// getAllCustomFields fetches all custom fields using the basic GetFields API
func getAllCustomFields(client *swagger.APIClient, ctx context.Context) ([]customFieldDetails, error) {
	fields, _, err := client.IssueFieldsAPI.GetFields(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get fields: %w", err)
	}

	var customFields []customFieldDetails
	for _, field := range fields {
		if field.GetCustom() {
			fieldType := "unknown"
			schema := ""
			if s := field.Schema; s != nil {
				fieldType = s.GetType()
				if items := s.GetItems(); items != "" {
					fieldType = items + "[]"
				}
				// Extract schema type (e.g., "com.atlassian.jira.plugin.system.customfieldtypes:textfield" -> "textfield")
				if custom := s.GetCustom(); custom != "" {
					schema = extractSchemaName(custom)
				}
			}
			customFields = append(customFields, customFieldDetails{
				id:        field.GetId(),
				name:      field.GetName(),
				fieldType: fieldType,
				schema:    schema,
			})
		}
	}

	return customFields, nil
}

// getCustomFieldsPaginated fetches custom fields using paginated API with optional filters
func getCustomFieldsPaginated(client *swagger.APIClient, ctx context.Context, projectKey string, query string) ([]customFieldDetails, error) {
	request := client.IssueFieldsAPI.GetFieldsPaginated(ctx).
		Type_([]string{"custom"}).
		MaxResults(50)

	if query != "" {
		request = request.Query(query)
	}

	if projectKey != "" {
		// Resolve project key to ID
		project, _, err := client.ProjectsAPI.GetProject(ctx, projectKey).Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to get project: %w", err)
		}
		projectId, err := parseProjectId(project.GetId())
		if err != nil {
			return nil, err
		}
		request = request.ProjectIds([]int64{projectId})
	}

	var customFields []customFieldDetails
	var startAt int64 = 0

	for {
		result, _, err := request.StartAt(startAt).Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to get fields: %w", err)
		}

		for _, field := range result.GetValues() {
			fieldType := "unknown"
			schemaName := ""
			s := field.GetSchema()
			fieldType = s.GetType()
			if items := s.GetItems(); items != "" {
				fieldType = items + "[]"
			}
			if custom := s.GetCustom(); custom != "" {
				schemaName = extractSchemaName(custom)
			}

			customFields = append(customFields, customFieldDetails{
				id:        field.GetId(),
				name:      field.GetName(),
				fieldType: fieldType,
				schema:    schemaName,
			})
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	return customFields, nil
}

// getCustomFieldsForProjectAndIssueType fetches fields for a specific project and issue type
func getCustomFieldsForProjectAndIssueType(client *swagger.APIClient, ctx context.Context, projectKey string, issueTypeName string, query string) ([]customFieldDetails, error) {
	// Resolve project key to ID
	project, _, err := client.ProjectsAPI.GetProject(ctx, projectKey).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	projectId, err := parseProjectId(project.GetId())
	if err != nil {
		return nil, err
	}

	// Get issue types for the project and find the matching one
	issueTypes, _, err := client.IssueTypesAPI.GetIssueTypesForProject(ctx).
		ProjectId(projectId).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get issue types: %w", err)
	}

	var issueTypeId int64
	for _, it := range issueTypes {
		if strings.EqualFold(it.GetName(), issueTypeName) {
			_, err := fmt.Sscanf(it.GetId(), "%d", &issueTypeId)
			if err != nil {
				return nil, fmt.Errorf("failed to parse issue type ID: %w", err)
			}
			break
		}
	}

	if issueTypeId == 0 {
		return nil, fmt.Errorf("issue type %q not found in project %s", issueTypeName, projectKey)
	}

	// First, get all custom fields so we can look up names and types
	allFields, err := getAllCustomFields(client, ctx)
	if err != nil {
		return nil, err
	}

	// Build a map for quick lookup
	fieldInfoMap := make(map[string]customFieldDetails)
	for _, f := range allFields {
		fieldInfoMap[f.id] = f
	}

	// Get fields for project and issue type
	var customFields []customFieldDetails
	var startAt int64 = 0

	for {
		result, _, err := client.IssueFieldsAPI.GetProjectFields(ctx).
			ProjectId([]int64{projectId}).
			WorkTypeId([]int64{issueTypeId}).
			StartAt(startAt).
			MaxResults(50).
			Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to get project fields: %w", err)
		}

		for _, field := range result.GetValues() {
			fieldId := field.GetFieldId()

			// Filter to custom fields only (ID starts with customfield_)
			if !strings.HasPrefix(fieldId, "customfield_") {
				continue
			}

			// Look up field info from our map
			fieldInfo, ok := fieldInfoMap[fieldId]
			if !ok {
				// Field not in our map, use ID as name
				fieldInfo = customFieldDetails{
					id:        fieldId,
					name:      fieldId,
					fieldType: "unknown",
				}
			}

			// Apply query filter if provided
			if query != "" && !strings.Contains(strings.ToLower(fieldInfo.name), strings.ToLower(query)) {
				continue
			}

			customFields = append(customFields, fieldInfo)
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	return customFields, nil
}

// getCustomFieldsByIssueType fetches custom fields filtered by issue type using context mappings
// This requires Jira administrator permissions
func getCustomFieldsByIssueType(client *swagger.APIClient, ctx context.Context, issueTypeName string, query string) ([]customFieldDetails, error) {
	// First get all issue types to find the ID
	allIssueTypes, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get issue types: %w", err)
	}

	var targetIssueTypeId string
	for _, it := range allIssueTypes {
		if strings.EqualFold(it.GetName(), issueTypeName) {
			targetIssueTypeId = it.GetId()
			break
		}
	}

	if targetIssueTypeId == "" {
		return nil, fmt.Errorf("issue type %q not found", issueTypeName)
	}

	// Get all custom fields
	allFields, err := getAllCustomFields(client, ctx)
	if err != nil {
		return nil, err
	}

	// Apply query filter if provided
	if query != "" {
		var filtered []customFieldDetails
		for _, f := range allFields {
			if strings.Contains(strings.ToLower(f.name), strings.ToLower(query)) {
				filtered = append(filtered, f)
			}
		}
		allFields = filtered
	}

	// For each field, check if it applies to the target issue type
	var result []customFieldDetails
	for _, field := range allFields {
		appliesToIssueType, err := fieldAppliesToIssueType(client, ctx, field.id, targetIssueTypeId)
		if err != nil {
			// Check for permission error
			if isPermissionError(err) {
				return nil, fmt.Errorf("the --issue-type filter requires Jira administrator permissions.\nTry using --project instead, or contact your Jira administrator")
			}
			// Skip fields that error for other reasons
			continue
		}

		if appliesToIssueType {
			result = append(result, field)
		}
	}

	return result, nil
}

// fieldAppliesToIssueType checks if a custom field applies to a specific issue type
func fieldAppliesToIssueType(client *swagger.APIClient, ctx context.Context, fieldId string, issueTypeId string) (bool, error) {
	// Get issue type mappings for this field's contexts
	result, resp, err := client.IssueCustomFieldContextsAPI.GetIssueTypeMappingsForContexts(ctx, fieldId).
		MaxResults(1000).
		Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusForbidden {
			return false, fmt.Errorf("403 Forbidden")
		}
		return false, err
	}

	mappings := result.GetValues()

	// If no mappings, field applies to all issue types
	if len(mappings) == 0 {
		return true, nil
	}

	// Check if any mapping includes our target issue type or is global (isAnyIssueType)
	for _, mapping := range mappings {
		if mapping.GetIsAnyIssueType() {
			return true, nil
		}
		if mapping.GetIssueTypeId() == issueTypeId {
			return true, nil
		}
	}

	return false, nil
}

// getIssueTypesForField returns a string describing which issue types a field applies to
func getIssueTypesForField(client *swagger.APIClient, ctx context.Context, fieldId string) (string, error) {
	// Get issue type mappings for this field
	result, resp, err := client.IssueCustomFieldContextsAPI.GetIssueTypeMappingsForContexts(ctx, fieldId).
		MaxResults(1000).
		Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusForbidden {
			return "", fmt.Errorf("403 Forbidden")
		}
		return "", err
	}

	mappings := result.GetValues()

	// If no mappings, field applies to all issue types
	if len(mappings) == 0 {
		return "All issue types", nil
	}

	// Check if any context applies to all issue types
	for _, mapping := range mappings {
		if mapping.GetIsAnyIssueType() {
			return "All issue types", nil
		}
	}

	// Get all issue types to map IDs to names
	allIssueTypes, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to get issue types: %w", err)
	}

	issueTypeNames := make(map[string]string)
	for _, it := range allIssueTypes {
		issueTypeNames[it.GetId()] = it.GetName()
	}

	// Collect unique issue type names
	seen := make(map[string]bool)
	var names []string
	for _, mapping := range mappings {
		typeId := mapping.GetIssueTypeId()
		if typeId != "" && !seen[typeId] {
			seen[typeId] = true
			if name, ok := issueTypeNames[typeId]; ok {
				names = append(names, name)
			} else {
				names = append(names, typeId)
			}
		}
	}

	if len(names) == 0 {
		return "None", nil
	}

	return strings.Join(names, ", "), nil
}

// isPermissionError checks if an error is a permission/authorization error
func isPermissionError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "403") || strings.Contains(errStr, "Forbidden") || strings.Contains(errStr, "forbidden")
}

// extractSchemaName extracts the short schema name from a full custom field type URI
// e.g., "com.atlassian.jira.plugin.system.customfieldtypes:textfield" -> "textfield"
func extractSchemaName(fullSchema string) string {
	if idx := strings.LastIndex(fullSchema, ":"); idx != -1 {
		return fullSchema[idx+1:]
	}
	return fullSchema
}
