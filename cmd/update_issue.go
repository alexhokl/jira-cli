package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/alexhokl/jira-cli/swagger_software"
)

type updateIssueOptions struct {
	issueKey        string
	summary         string
	descriptionFile string
	priority        string
	assignee        string
	labels          string
	addLabels       []string
	deleteLabels    []string
	components      string
	dueDate         string
	sprint          string
	transition      string
	noNotify        bool
	linkIssue       string
	linkType        string
	customFields    []string
	parent          string
}

var updateIssueOpts = updateIssueOptions{}

var updateIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Update an issue",
	Long: `Update an issue's fields or transition its status.

Examples:
  # Update issue summary
  jira-cli update issue --id PROJ-123 --summary "New summary"

  # Update issue description from file
  jira-cli update issue --id PROJ-123 --description-file description.md

  # Update issue priority
  jira-cli update issue --id PROJ-123 --priority High

  # Update assignee
  jira-cli update issue --id PROJ-123 --assignee 5b10ac8d82e05b22cc7d4ef5

  # Update labels (replaces existing labels)
  jira-cli update issue --id PROJ-123 --labels "bug,critical"

  # Add labels (keeps existing labels)
  jira-cli update issue --id PROJ-123 --add-label bug --add-label critical

  # Delete labels (keeps other existing labels)
  jira-cli update issue --id PROJ-123 --delete-label obsolete

  # Update multiple fields
  jira-cli update issue --id PROJ-123 --summary "Updated" --priority High

  # Update without sending notifications
  jira-cli update issue --id PROJ-123 --summary "Updated" --no-notify

  # Link to another issue (this issue blocks PROJ-456)
  # Use 'jira-cli list link-types' to see available link type names
  jira-cli update issue --id PROJ-123 --link-issue PROJ-456 --link-type Blocks

  # Update a custom field (text field)
  jira-cli update issue --id PROJ-123 --custom-field "Team=Backend"

  # Update multiple custom fields
  jira-cli update issue --id PROJ-123 --custom-field "Team=Backend" --custom-field "Story Points=5"

  # Update a select list custom field (use the option value)
  jira-cli update issue --id PROJ-123 --custom-field "Environment=Production"

  # Move issue to a sprint (by ID or name)
  jira-cli update issue --id PROJ-123 --sprint 42
  jira-cli update issue --id PROJ-123 --sprint "Sprint 5"

  # Move issue to backlog (remove from sprint)
  jira-cli update issue --id PROJ-123 --sprint none

  # Assign a parent issue (make this issue a child/subtask of another issue)
  jira-cli update issue --id PROJ-123 --parent PROJ-100

  # Remove parent (unassign from parent issue)
  jira-cli update issue --id PROJ-123 --parent none

  # Clear a custom field value
  jira-cli update issue --id PROJ-123 --custom-field "Team="

  # Transition issue status (by name or ID)
  # Use 'jira-cli list issue-transitions --id PROJ-123' to see available transitions
  jira-cli update issue --id PROJ-123 --transition "In Progress"
  jira-cli update issue --id PROJ-123 --transition 21

  # Transition and update fields together
  jira-cli update issue --id PROJ-123 --transition "In Progress" --assignee me`,
	RunE: runUpdateIssue,
}

func init() {
	updateCmd.AddCommand(updateIssueCmd)

	flags := updateIssueCmd.Flags()
	flags.StringVarP(&updateIssueOpts.issueKey, "id", "i", "", "Issue ID (e.g., PROJ-123) (required)")
	flags.StringVarP(&updateIssueOpts.summary, "summary", "s", "", "Issue summary/title")
	flags.StringVarP(&updateIssueOpts.descriptionFile, "description-file", "d", "", "Path to a file containing issue description in markdown")
	flags.StringVar(&updateIssueOpts.priority, "priority", "", "Priority (e.g., Highest, High, Medium, Low, Lowest)")
	flags.StringVarP(&updateIssueOpts.assignee, "assignee", "a", "", "Assignee account ID (use 'me' for yourself, 'none' to unassign)")
	flags.StringVarP(&updateIssueOpts.labels, "labels", "l", "", "Comma-separated labels (replaces existing)")
	flags.StringArrayVar(&updateIssueOpts.addLabels, "add-label", nil, "Add a label (keeps existing, can be specified multiple times)")
	flags.StringArrayVar(&updateIssueOpts.deleteLabels, "delete-label", nil, "Delete a label (keeps others, can be specified multiple times)")
	flags.StringVarP(&updateIssueOpts.components, "components", "c", "", "Comma-separated component names (replaces existing)")
	flags.StringVar(&updateIssueOpts.dueDate, "due-date", "", "Due date (format: 2006-01-02, use 'none' to clear)")
	flags.StringVar(&updateIssueOpts.sprint, "sprint", "", "Sprint ID or name (use 'none' to move to backlog)")
	flags.StringVarP(&updateIssueOpts.transition, "transition", "t", "", "Transition name or ID to change issue status (use 'jira-cli list issue-transitions' to see available transitions)")
	flags.BoolVar(&updateIssueOpts.noNotify, "no-notify", false, "Don't send notification emails")
	flags.StringVar(&updateIssueOpts.linkIssue, "link-issue", "", "Issue key to link to (e.g., PROJ-456)")
	flags.StringVar(&updateIssueOpts.linkType, "link-type", "", "Link type name (use 'jira-cli list link-types' to see available types)")
	flags.StringArrayVar(&updateIssueOpts.customFields, "custom-field", nil, "Custom field to update in format 'name=value' (can be specified multiple times)")
	flags.StringVar(&updateIssueOpts.parent, "parent", "", "Parent issue key to assign (use 'none' to remove parent)")

	updateIssueCmd.MarkFlagRequired("id")
}

func runUpdateIssue(_ *cobra.Command, _ []string) error {
	// Check that at least one field is being updated
	if updateIssueOpts.summary == "" &&
		updateIssueOpts.descriptionFile == "" &&
		updateIssueOpts.priority == "" &&
		updateIssueOpts.assignee == "" &&
		updateIssueOpts.labels == "" &&
		len(updateIssueOpts.addLabels) == 0 &&
		len(updateIssueOpts.deleteLabels) == 0 &&
		updateIssueOpts.components == "" &&
		updateIssueOpts.dueDate == "" &&
		updateIssueOpts.sprint == "" &&
		updateIssueOpts.transition == "" &&
		updateIssueOpts.linkIssue == "" &&
		updateIssueOpts.parent == "" &&
		len(updateIssueOpts.customFields) == 0 {
		return fmt.Errorf("at least one field must be specified to update")
	}

	// Check for conflicting label options
	if updateIssueOpts.labels != "" && (len(updateIssueOpts.addLabels) > 0 || len(updateIssueOpts.deleteLabels) > 0) {
		return fmt.Errorf("--labels cannot be used with --add-label or --delete-label")
	}

	// Check that link-type is provided when link-issue is specified
	if updateIssueOpts.linkIssue != "" && updateIssueOpts.linkType == "" {
		return fmt.Errorf("--link-type is required when --link-issue is specified")
	}

	client := newClient()
	ctx := getAuthContext()

	fields := make(map[string]interface{})
	update := make(map[string][]swagger.FieldUpdateOperation)

	if updateIssueOpts.summary != "" {
		fields["summary"] = updateIssueOpts.summary
	}

	if updateIssueOpts.descriptionFile != "" {
		content, err := os.ReadFile(updateIssueOpts.descriptionFile)
		if err != nil {
			return fmt.Errorf("failed to read description file: %w", err)
		}
		description := strings.TrimSpace(string(content))
		if description != "" {
			fields["description"] = convertMarkdownToADF(description)
		}
	}

	if updateIssueOpts.priority != "" {
		fields["priority"] = map[string]interface{}{
			"name": updateIssueOpts.priority,
		}
	}

	if updateIssueOpts.assignee != "" {
		if updateIssueOpts.assignee == "none" {
			fields["assignee"] = nil
		} else {
			assigneeId := updateIssueOpts.assignee
			if assigneeId == "me" {
				// Fetch current user's account ID
				currentUser, _, err := client.MyselfAPI.GetCurrentUser(ctx).Execute()
				if err != nil {
					return fmt.Errorf("failed to get current user: %w", err)
				}
				assigneeId = currentUser.GetAccountId()
			}
			fields["assignee"] = map[string]interface{}{
				"id": assigneeId,
			}
		}
	}

	if updateIssueOpts.labels != "" {
		labelList := strings.Split(updateIssueOpts.labels, ",")
		for i, label := range labelList {
			labelList[i] = strings.TrimSpace(label)
		}
		fields["labels"] = labelList
	}

	if len(updateIssueOpts.addLabels) > 0 {
		var labelOps []swagger.FieldUpdateOperation
		for _, label := range updateIssueOpts.addLabels {
			op := swagger.NewFieldUpdateOperation()
			op.SetAdd(strings.TrimSpace(label))
			labelOps = append(labelOps, *op)
		}
		update["labels"] = labelOps
	}

	if len(updateIssueOpts.deleteLabels) > 0 {
		labelOps := update["labels"]
		for _, label := range updateIssueOpts.deleteLabels {
			op := swagger.NewFieldUpdateOperation()
			op.SetRemove(strings.TrimSpace(label))
			labelOps = append(labelOps, *op)
		}
		update["labels"] = labelOps
	}

	if updateIssueOpts.components != "" {
		componentList := strings.Split(updateIssueOpts.components, ",")
		components := make([]map[string]interface{}, len(componentList))
		for i, comp := range componentList {
			components[i] = map[string]interface{}{
				"name": strings.TrimSpace(comp),
			}
		}
		fields["components"] = components
	}

	if updateIssueOpts.dueDate != "" {
		if updateIssueOpts.dueDate == "none" {
			fields["duedate"] = nil
		} else {
			fields["duedate"] = updateIssueOpts.dueDate
		}
	}

	if updateIssueOpts.parent != "" {
		if strings.EqualFold(updateIssueOpts.parent, "none") {
			// To remove parent, use update.parent with set.none = true
			parentOp := swagger.NewFieldUpdateOperation()
			parentOp.SetSet(map[string]interface{}{"none": true})
			update["parent"] = []swagger.FieldUpdateOperation{*parentOp}
		} else {
			fields["parent"] = map[string]interface{}{
				"key": updateIssueOpts.parent,
			}
		}
	}

	// Handle custom fields
	if len(updateIssueOpts.customFields) > 0 {
		// Get field metadata to resolve names to IDs and determine field types
		fieldInfoMap, err := getCustomFieldInfoMap(client, ctx)
		if err != nil {
			return fmt.Errorf("failed to get custom field information: %w", err)
		}

		// Extract project key from issue key (e.g., PROJ-123 -> PROJ)
		projectKey := extractProjectKeyFromIssueKey(updateIssueOpts.issueKey)

		// Get the issue to determine its type (needed for allowed values lookup)
		issue, _, err := client.IssuesAPI.GetIssue(ctx, updateIssueOpts.issueKey).Execute()
		var issueTypeName string
		if err == nil {
			if issueFields := issue.GetFields(); issueFields != nil {
				if issueType, ok := issueFields["issuetype"].(map[string]interface{}); ok {
					if name, ok := issueType["name"].(string); ok {
						issueTypeName = name
					}
				}
			}
		}

		for _, cf := range updateIssueOpts.customFields {
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
			if issueTypeName != "" {
				enrichFieldInfoWithAllowedValues(client, ctx, &fieldInfo, projectKey, issueTypeName)
			}

			// Convert the value to the appropriate format for the field type
			convertedValue, err := convertCustomFieldValue(fieldValue, fieldInfo, projectKey)
			if err != nil {
				return fmt.Errorf("failed to convert value for custom field %q: %w", fieldName, err)
			}

			fields[fieldInfo.id] = convertedValue
		}
	}

	issueUpdateDetails := swagger.NewIssueUpdateDetails()
	if len(fields) > 0 {
		issueUpdateDetails.SetFields(fields)
	}
	if len(update) > 0 {
		issueUpdateDetails.SetUpdate(update)
	}

	// Handle issue field updates if any fields are being updated
	hasFieldUpdates := len(fields) > 0 || len(update) > 0
	if hasFieldUpdates {
		request := client.IssuesAPI.EditIssue(ctx, updateIssueOpts.issueKey).
			IssueUpdateDetails(*issueUpdateDetails)

		if updateIssueOpts.noNotify {
			request = request.NotifyUsers(false)
		}

		_, _, err := request.Execute()
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

		fmt.Printf("Issue %s updated\n", updateIssueOpts.issueKey)
	}

	// Handle issue linking if specified
	if updateIssueOpts.linkIssue != "" {
		outwardIssue := swagger.NewLinkedIssue()
		outwardIssue.SetKey(updateIssueOpts.linkIssue)

		inwardIssue := swagger.NewLinkedIssue()
		inwardIssue.SetKey(updateIssueOpts.issueKey)

		linkType := swagger.NewIssueLinkType()
		linkType.SetName(updateIssueOpts.linkType)

		linkRequest := swagger.NewLinkIssueRequestJsonBean(*inwardIssue, *outwardIssue, *linkType)

		_, _, err := client.IssueLinksAPI.LinkIssues(ctx).
			LinkIssueRequestJsonBean(*linkRequest).
			Execute()
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

		fmt.Printf("Issue %s linked to %s with type '%s'\n", updateIssueOpts.issueKey, updateIssueOpts.linkIssue, updateIssueOpts.linkType)
	}

	// Handle sprint update if specified
	if updateIssueOpts.sprint != "" {
		if err := updateIssueSprint(updateIssueOpts.issueKey, updateIssueOpts.sprint); err != nil {
			return err
		}
	}

	// Handle transition if specified
	if updateIssueOpts.transition != "" {
		if err := transitionIssue(client, ctx, updateIssueOpts.issueKey, updateIssueOpts.transition); err != nil {
			return err
		}
	}

	return nil
}

// customFieldInfo holds information about a custom field for update operations
type customFieldInfo struct {
	id            string
	name          string
	schemaType    string        // The schema type (e.g., "string", "number", "option", "array")
	itemType      string        // For array types, the type of items
	custom        string        // The custom field type URI (e.g., "com.atlassian.jira.plugin.system.customfieldtypes:select")
	allowedValues []interface{} // The list of allowed values for option fields
}

// getCustomFieldInfoMap returns a map of field ID to customFieldInfo for all custom fields
func getCustomFieldInfoMap(client *swagger.APIClient, ctx context.Context) (map[string]customFieldInfo, error) {
	fields, _, err := client.IssueFieldsAPI.GetFields(ctx).Execute()
	if err != nil {
		return nil, err
	}

	result := make(map[string]customFieldInfo)
	for _, field := range fields {
		if !field.GetCustom() {
			continue
		}

		info := customFieldInfo{
			id:   field.GetId(),
			name: field.GetName(),
		}

		if schema := field.Schema; schema != nil {
			info.schemaType = schema.GetType()
			info.itemType = schema.GetItems()
			info.custom = schema.GetCustom()
		}

		result[field.GetId()] = info
	}

	return result, nil
}

// getFieldAllowedValues fetches the allowed values for a custom field from create metadata
func getFieldAllowedValues(client *swagger.APIClient, ctx context.Context, projectKey, issueTypeName, fieldId string) ([]interface{}, error) {
	// First, get the issue type ID
	issueTypes, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get issue types: %w", err)
	}

	var issueTypeId string
	for _, it := range issueTypes {
		if strings.EqualFold(it.GetName(), issueTypeName) {
			issueTypeId = it.GetId()
			break
		}
	}

	if issueTypeId == "" {
		// Issue type not found, return empty allowed values
		return nil, nil
	}

	// Fetch create metadata for this project and issue type
	var startAt int32 = 0
	const maxResults int32 = 50

	for {
		result, _, err := client.IssuesAPI.GetCreateIssueMetaIssueTypeId(ctx, projectKey, issueTypeId).
			StartAt(startAt).
			MaxResults(maxResults).
			Execute()
		if err != nil {
			// If we can't get metadata, return empty allowed values (non-fatal)
			return nil, nil
		}

		// Check both Fields and Results (API may return either)
		fields := result.GetFields()
		if len(fields) == 0 {
			fields = result.GetResults()
		}

		for _, field := range fields {
			if field.FieldId == fieldId {
				return field.GetAllowedValues(), nil
			}
		}

		// Check if we've fetched all results
		total := result.GetTotal()
		fetched := startAt + int32(len(fields))
		if int64(fetched) >= total || len(fields) == 0 {
			break
		}

		startAt += int32(len(fields))
	}

	return nil, nil
}

// enrichFieldInfoWithAllowedValues adds allowed values to a field info if it's an option-type field
func enrichFieldInfoWithAllowedValues(client *swagger.APIClient, ctx context.Context, info *customFieldInfo, projectKey, issueTypeName string) {
	fieldType := extractFieldType(info.custom)

	// Only fetch allowed values for option-type fields
	isOptionField := info.schemaType == "option" ||
		(info.schemaType == "array" && info.itemType == "option") ||
		fieldType == "select" || fieldType == "radiobuttons" ||
		fieldType == "multiselect" || fieldType == "multicheckboxes" ||
		fieldType == "cascadingselect"

	if !isOptionField {
		return
	}

	allowedValues, err := getFieldAllowedValues(client, ctx, projectKey, issueTypeName, info.id)
	if err == nil && len(allowedValues) > 0 {
		info.allowedValues = allowedValues
	}
}

// findCustomFieldByName searches for a custom field by name (case-insensitive)
func findCustomFieldByName(fieldMap map[string]customFieldInfo, name string) (customFieldInfo, bool) {
	for _, info := range fieldMap {
		if strings.EqualFold(info.name, name) {
			return info, true
		}
	}
	return customFieldInfo{}, false
}

// convertCustomFieldValue converts a string value to the appropriate format for the custom field type
// projectKey is optional and used for resolving sprint names to IDs
func convertCustomFieldValue(value string, field customFieldInfo, projectKey string) (interface{}, error) {
	// Empty value means clear the field
	if value == "" {
		return nil, nil
	}

	// Determine the field type from the custom field type URI
	fieldType := extractFieldType(field.custom)

	// Handle special field types first, regardless of schema type
	switch fieldType {
	case "gh-sprint":
		// Sprint field - requires sprint ID(s)
		// Parse as comma-separated values (could be IDs or names)
		sprintValues := parseCommaSeparatedValues(value)
		ids := make([]int, 0, len(sprintValues))
		for _, val := range sprintValues {
			// Try to parse as integer first
			id, err := strconv.Atoi(val)
			if err == nil {
				ids = append(ids, id)
				continue
			}

			// Not a number, try to look up sprint by name
			if projectKey == "" {
				return nil, fmt.Errorf("sprint %q is not a valid ID; to use sprint names, provide a project key, or use sprint IDs (run 'jira-cli list sprints --board-id <id>' to find IDs)", val)
			}

			sprintId, err := lookupSprintByName(projectKey, val)
			if err != nil {
				return nil, err
			}
			ids = append(ids, sprintId)
		}
		return ids, nil
	}

	switch field.schemaType {
	case "string":
		return value, nil

	case "number":
		// Try to parse as float
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number value: %w", err)
		}
		return f, nil

	case "option":
		// Single select - look up option ID from allowed values, fall back to value
		return formatOptionValue(value, field.allowedValues), nil

	case "array":
		// Handle array types based on item type
		switch field.itemType {
		case "option":
			// Multi-select - parse comma-separated values
			values := parseCommaSeparatedValues(value)
			options := make([]map[string]interface{}, len(values))
			for i, v := range values {
				options[i] = formatOptionValue(v, field.allowedValues)
			}
			return options, nil

		case "string":
			// Array of strings (like labels)
			return parseCommaSeparatedValues(value), nil

		case "user":
			// Multi-user picker - parse comma-separated account IDs (supports "me" alias)
			accountIds := parseCommaSeparatedValues(value)
			users := make([]map[string]interface{}, len(accountIds))
			for i, id := range accountIds {
				resolvedId, err := resolveUserAlias(id)
				if err != nil {
					return nil, err
				}
				users[i] = map[string]interface{}{
					"id": resolvedId,
				}
			}
			return users, nil

		default:
			// Unknown array type - try as comma-separated values
			return parseCommaSeparatedValues(value), nil
		}

	case "user":
		// Single user picker (supports "me" alias)
		resolvedId, err := resolveUserAlias(value)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"id": resolvedId,
		}, nil

	case "date":
		// Date field - expect format YYYY-MM-DD
		return value, nil

	case "datetime":
		// DateTime field - expect ISO format
		return value, nil

	default:
		// Handle based on custom field type if schema type is not conclusive
		switch fieldType {
		case "select", "radiobuttons":
			return formatOptionValue(value, field.allowedValues), nil

		case "multiselect", "multicheckboxes":
			values := parseCommaSeparatedValues(value)
			options := make([]map[string]interface{}, len(values))
			for i, v := range values {
				options[i] = formatOptionValue(v, field.allowedValues)
			}
			return options, nil

		case "userpicker":
			// Single user picker (supports "me" alias)
			resolvedId, err := resolveUserAlias(value)
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"id": resolvedId,
			}, nil

		case "multiuserpicker":
			// Multi-user picker (supports "me" alias)
			accountIds := parseCommaSeparatedValues(value)
			users := make([]map[string]interface{}, len(accountIds))
			for i, id := range accountIds {
				resolvedId, err := resolveUserAlias(id)
				if err != nil {
					return nil, err
				}
				users[i] = map[string]interface{}{
					"id": resolvedId,
				}
			}
			return users, nil

		case "cascadingselect":
			// Cascading select expects parent and optionally child value
			// Format: "Parent" or "Parent - Child"
			parts := strings.SplitN(value, " - ", 2)
			result := map[string]interface{}{
				"value": strings.TrimSpace(parts[0]),
			}
			if len(parts) > 1 {
				result["child"] = map[string]interface{}{
					"value": strings.TrimSpace(parts[1]),
				}
			}
			return result, nil

		case "float":
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number value: %w", err)
			}
			return f, nil

		case "datepicker":
			return value, nil

		case "datetime":
			return value, nil

		case "url":
			return value, nil

		case "textarea", "textfield":
			return value, nil

		default:
			// Default: return as-is
			return value, nil
		}
	}
}

// extractFieldType extracts the short field type from a custom field type URI
// e.g., "com.atlassian.jira.plugin.system.customfieldtypes:select" -> "select"
func extractFieldType(customType string) string {
	if idx := strings.LastIndex(customType, ":"); idx != -1 {
		return customType[idx+1:]
	}
	return customType
}

// parseCommaSeparatedValues splits a string by commas and trims whitespace
func parseCommaSeparatedValues(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// Note: extractProjectKeyFromIssueKey is defined in helper.go

// lookupOptionId looks up an option's ID from allowed values by matching the value name
// Returns the ID if found, or empty string if not found
func lookupOptionId(allowedValues []interface{}, valueName string) string {
	for _, av := range allowedValues {
		if avMap, ok := av.(map[string]interface{}); ok {
			// Check if the "value" field matches (case-insensitive)
			if v, ok := avMap["value"].(string); ok {
				if strings.EqualFold(v, valueName) {
					// Return the ID if present
					if id, ok := avMap["id"].(string); ok {
						return id
					}
					// Some fields use numeric IDs
					if id, ok := avMap["id"].(float64); ok {
						return strconv.FormatFloat(id, 'f', 0, 64)
					}
				}
			}
			// Also check "name" field for some field types
			if v, ok := avMap["name"].(string); ok {
				if strings.EqualFold(v, valueName) {
					if id, ok := avMap["id"].(string); ok {
						return id
					}
					if id, ok := avMap["id"].(float64); ok {
						return strconv.FormatFloat(id, 'f', 0, 64)
					}
				}
			}
		}
	}
	return ""
}

// formatOptionValue formats an option value for the API
// If allowed values are provided and contain the value, use the ID
// Otherwise, fall back to using the value name directly
func formatOptionValue(valueName string, allowedValues []interface{}) map[string]interface{} {
	if len(allowedValues) > 0 {
		if id := lookupOptionId(allowedValues, valueName); id != "" {
			return map[string]interface{}{
				"id": id,
			}
		}
	}
	// Fall back to value-based lookup
	return map[string]interface{}{
		"value": valueName,
	}
}

// updateIssueSprint moves an issue to a sprint or to the backlog
func updateIssueSprint(issueKey, sprint string) error {
	softwareClient := newSoftwareClient()
	softwareCtx := getSoftwareAuthContext()

	// Extract project key from issue key for sprint name lookup
	projectKey := extractProjectKeyFromIssueKey(issueKey)

	// Handle "none" or "backlog" to move issue to backlog
	if strings.EqualFold(sprint, "none") || strings.EqualFold(sprint, "backlog") {
		moveRequest := swagger_software.NewMoveIssuesToBacklogRequest()
		moveRequest.SetIssues([]string{issueKey})

		_, err := softwareClient.BacklogAPI.MoveIssuesToBacklog(softwareCtx).
			MoveIssuesToBacklogRequest(*moveRequest).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to move issue to backlog: %w", err)
		}

		fmt.Printf("Issue %s moved to backlog\n", issueKey)
		return nil
	}

	// Try to parse sprint as ID first
	sprintId, err := strconv.ParseInt(sprint, 10, 64)
	if err != nil {
		// Not a number, try to look up sprint by name
		id, err := lookupSprintByName(projectKey, sprint)
		if err != nil {
			return fmt.Errorf("failed to resolve sprint: %w", err)
		}
		sprintId = int64(id)
	}

	// Move issue to the sprint
	moveRequest := swagger_software.NewMoveIssuesToBacklogForBoardRequest()
	moveRequest.SetIssues([]string{issueKey})

	_, err = softwareClient.SprintAPI.MoveIssuesToSprintAndRank(softwareCtx, sprintId).
		MoveIssuesToBacklogForBoardRequest(*moveRequest).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to move issue to sprint: %w", err)
	}

	fmt.Printf("Issue %s moved to sprint %d\n", issueKey, sprintId)
	return nil
}

// transitionIssue transitions an issue to a new status
func transitionIssue(client *swagger.APIClient, ctx context.Context, issueKey, transition string) error {
	// Get available transitions for the issue
	result, _, err := client.IssuesAPI.GetTransitions(ctx, issueKey).Execute()
	if err != nil {
		return fmt.Errorf("failed to get available transitions: %w", err)
	}

	transitions := result.GetTransitions()
	if len(transitions) == 0 {
		return fmt.Errorf("no transitions available for issue %s", issueKey)
	}

	// Find matching transition by ID or name (case-insensitive)
	var matchedTransition *swagger.IssueTransition
	for i, t := range transitions {
		if t.GetId() == transition || strings.EqualFold(t.GetName(), transition) {
			matchedTransition = &transitions[i]
			break
		}
	}

	if matchedTransition == nil {
		// Build list of available transitions for error message
		available := make([]string, len(transitions))
		for i, t := range transitions {
			available[i] = fmt.Sprintf("%s (ID: %s)", t.GetName(), t.GetId())
		}
		return fmt.Errorf("transition %q not found; available transitions: %s", transition, strings.Join(available, ", "))
	}

	// Execute the transition
	transitionRequest := swagger.NewIssueTransition()
	transitionRequest.SetId(matchedTransition.GetId())

	issueUpdateDetails := swagger.NewIssueUpdateDetails()
	issueUpdateDetails.SetTransition(*transitionRequest)

	_, _, err = client.IssuesAPI.DoTransition(ctx, issueKey).
		IssueUpdateDetails(*issueUpdateDetails).
		Execute()
	if err != nil {
		var openAPIErr *swagger.GenericOpenAPIError
		if errors.As(err, &openAPIErr) {
			body := string(openAPIErr.Body())
			if body != "" {
				return fmt.Errorf("failed to transition issue: %s\n\n%s", err, body)
			}
		}
		return fmt.Errorf("failed to transition issue: %w", err)
	}

	toStatus := matchedTransition.GetTo()
	fmt.Printf("Issue %s transitioned to '%s'\n", issueKey, toStatus.GetName())
	return nil
}
