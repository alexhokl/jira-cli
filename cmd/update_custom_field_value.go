package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type updateCustomFieldValueOptions struct {
	fieldName string
	oldValue  string
	newValue  string
	enable    bool
	enableSet bool // tracks if --enable flag was explicitly set
}

var updateCustomFieldValueOpts = updateCustomFieldValueOptions{}

var updateCustomFieldValueCmd = &cobra.Command{
	Use:   "custom-field-value",
	Short: "Update a custom field option value",
	Long: `Update the value of an option in a custom field's option list.

This command modifies the options (allowed values) of a custom field itself,
not the value of a custom field on an issue.

This is useful for renaming options or enabling/disabling them in select list custom fields.

Note: This operation requires Jira administrator permissions.

Examples:
  # Rename an option value in a custom field
  jira-cli update custom-field-value --name "Environment" --old "Prod" --new "Production"

  # Update a team name option
  jira-cli update custom-field-value --name "Team" --old "Backend Team" --new "Platform Team"

  # Disable an option (users can no longer select it)
  jira-cli update custom-field-value --name "Environment" --old "Legacy" --enable=false

  # Re-enable a disabled option
  jira-cli update custom-field-value --name "Environment" --old "Legacy" --enable=true

  # Rename and enable an option at the same time
  jira-cli update custom-field-value --name "Environment" --old "Old" --new "Current" --enable=true`,
	RunE: runUpdateCustomFieldValue,
}

func init() {
	updateCmd.AddCommand(updateCustomFieldValueCmd)

	flags := updateCustomFieldValueCmd.Flags()
	flags.StringVarP(&updateCustomFieldValueOpts.fieldName, "name", "n", "", "Custom field name (required)")
	flags.StringVarP(&updateCustomFieldValueOpts.oldValue, "old", "o", "", "Current option value (required)")
	flags.StringVarP(&updateCustomFieldValueOpts.newValue, "new", "w", "", "New option value")
	flags.BoolVarP(&updateCustomFieldValueOpts.enable, "enable", "e", false, "Enable (true) or disable (false) the option")

	updateCustomFieldValueCmd.MarkFlagRequired("name")
	updateCustomFieldValueCmd.MarkFlagRequired("old")
}

func runUpdateCustomFieldValue(cmd *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// Check if --enable flag was explicitly set
	enableSet := cmd.Flags().Changed("enable")

	// Validate that at least one of --new or --enable is provided
	if updateCustomFieldValueOpts.newValue == "" && !enableSet {
		return fmt.Errorf("at least one of --new or --enable must be specified")
	}

	// Find the custom field by name
	fieldId, err := findCustomFieldIdByName(client, ctx, updateCustomFieldValueOpts.fieldName)
	if err != nil {
		return err
	}

	// Get contexts for the field
	contexts, err := getFieldContexts(client, ctx, fieldId, "")
	if err != nil {
		return err
	}

	if len(contexts) == 0 {
		return fmt.Errorf("no contexts found for custom field %q", updateCustomFieldValueOpts.fieldName)
	}

	// Find the option in any context and update it
	var updated bool
	var lastErr error

	for _, context := range contexts {
		contextId, err := parseContextId(context.GetId())
		if err != nil {
			continue
		}

		// Get options for this context
		options, err := getOptionsForContext(client, ctx, fieldId, contextId)
		if err != nil {
			lastErr = err
			continue
		}

		// Find the option with the old value
		var optionId string
		for _, opt := range options {
			if strings.EqualFold(opt.GetValue(), updateCustomFieldValueOpts.oldValue) {
				optionId = opt.GetId()
				break
			}
		}

		if optionId == "" {
			// Option not found in this context, try next
			continue
		}

		// Update the option
		optionUpdate := swagger.NewCustomFieldOptionUpdate(optionId)
		if updateCustomFieldValueOpts.newValue != "" {
			optionUpdate.SetValue(updateCustomFieldValueOpts.newValue)
		}
		if enableSet {
			// disabled is the opposite of enable
			optionUpdate.SetDisabled(!updateCustomFieldValueOpts.enable)
		}

		updateRequest := swagger.NewBulkCustomFieldOptionUpdateRequest()
		updateRequest.SetOptions([]swagger.CustomFieldOptionUpdate{*optionUpdate})

		_, resp, err := client.IssueCustomFieldOptionsAPI.UpdateCustomFieldOption(ctx, fieldId, contextId).
			BulkCustomFieldOptionUpdateRequest(*updateRequest).
			Execute()
		if err != nil {
			if resp != nil && resp.StatusCode == 403 {
				return fmt.Errorf("permission denied: this operation requires Jira administrator permissions")
			}
			lastErr = wrapAPIError(fmt.Errorf("failed to update option: %w", err))
			continue
		}

		// Build output message
		var changes []string
		if updateCustomFieldValueOpts.newValue != "" {
			changes = append(changes, fmt.Sprintf("renamed %q -> %q", updateCustomFieldValueOpts.oldValue, updateCustomFieldValueOpts.newValue))
		}
		if enableSet {
			if updateCustomFieldValueOpts.enable {
				changes = append(changes, "enabled")
			} else {
				changes = append(changes, "disabled")
			}
		}

		fmt.Printf("Updated option in custom field %q (context: %s): %s\n",
			updateCustomFieldValueOpts.fieldName,
			context.GetName(),
			strings.Join(changes, ", "))
		updated = true
	}

	if !updated {
		if lastErr != nil {
			return lastErr
		}
		return fmt.Errorf("option %q not found in any context of custom field %q",
			updateCustomFieldValueOpts.oldValue,
			updateCustomFieldValueOpts.fieldName)
	}

	return nil
}

// listMatchingCustomFields returns custom fields that match the given query (for help/debugging)
func listMatchingCustomFields(fieldInfoMap map[string]customFieldInfo, query string) []customFieldInfo {
	var matches []customFieldInfo
	queryLower := strings.ToLower(query)

	for _, info := range fieldInfoMap {
		if strings.Contains(strings.ToLower(info.name), queryLower) {
			matches = append(matches, info)
		}
	}

	return matches
}
