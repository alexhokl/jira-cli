package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type updateIssueOptions struct {
	issueKey     string
	summary      string
	description  string
	priority     string
	assignee     string
	labels       string
	addLabels    []string
	deleteLabels []string
	components   string
	dueDate      string
	noNotify     bool
	linkIssue    string
	linkType     string
}

var updateIssueOpts = updateIssueOptions{}

var updateIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Update an issue",
	Long: `Update an issue's fields.

Note: Issue transitions (status changes) are not supported by this command.
Use the transition command to change issue status.

Examples:
  # Update issue summary
  jira-cli update issue --key PROJ-123 --summary "New summary"

  # Update issue description
  jira-cli update issue --key PROJ-123 --description "New description"

  # Update issue priority
  jira-cli update issue --key PROJ-123 --priority High

  # Update assignee
  jira-cli update issue --key PROJ-123 --assignee 5b10ac8d82e05b22cc7d4ef5

  # Update labels (replaces existing labels)
  jira-cli update issue --key PROJ-123 --labels "bug,critical"

  # Add labels (keeps existing labels)
  jira-cli update issue --key PROJ-123 --add-label bug --add-label critical

  # Delete labels (keeps other existing labels)
  jira-cli update issue --key PROJ-123 --delete-label obsolete

  # Update multiple fields
  jira-cli update issue --key PROJ-123 --summary "Updated" --priority High

  # Update without sending notifications
  jira-cli update issue --key PROJ-123 --summary "Updated" --no-notify

  # Link to another issue (this issue blocks PROJ-456)
  # Use 'jira-cli list link-types' to see available link type names
  jira-cli update issue --key PROJ-123 --link-issue PROJ-456 --link-type Blocks`,
	RunE: runUpdateIssue,
}

func init() {
	updateCmd.AddCommand(updateIssueCmd)

	flags := updateIssueCmd.Flags()
	flags.StringVarP(&updateIssueOpts.issueKey, "key", "k", "", "Issue key (e.g., PROJ-123) (required)")
	flags.StringVarP(&updateIssueOpts.summary, "summary", "s", "", "Issue summary/title")
	flags.StringVarP(&updateIssueOpts.description, "description", "d", "", "Issue description")
	flags.StringVar(&updateIssueOpts.priority, "priority", "", "Priority (e.g., Highest, High, Medium, Low, Lowest)")
	flags.StringVarP(&updateIssueOpts.assignee, "assignee", "a", "", "Assignee account ID (use 'none' to unassign)")
	flags.StringVarP(&updateIssueOpts.labels, "labels", "l", "", "Comma-separated labels (replaces existing)")
	flags.StringArrayVar(&updateIssueOpts.addLabels, "add-label", nil, "Add a label (keeps existing, can be specified multiple times)")
	flags.StringArrayVar(&updateIssueOpts.deleteLabels, "delete-label", nil, "Delete a label (keeps others, can be specified multiple times)")
	flags.StringVarP(&updateIssueOpts.components, "components", "c", "", "Comma-separated component names (replaces existing)")
	flags.StringVar(&updateIssueOpts.dueDate, "due-date", "", "Due date (format: 2006-01-02, use 'none' to clear)")
	flags.BoolVar(&updateIssueOpts.noNotify, "no-notify", false, "Don't send notification emails")
	flags.StringVar(&updateIssueOpts.linkIssue, "link-issue", "", "Issue key to link to (e.g., PROJ-456)")
	flags.StringVar(&updateIssueOpts.linkType, "link-type", "", "Link type name (use 'jira-cli list link-types' to see available types)")

	updateIssueCmd.MarkFlagRequired("key")
}

func runUpdateIssue(_ *cobra.Command, _ []string) error {
	// Check that at least one field is being updated
	if updateIssueOpts.summary == "" &&
		updateIssueOpts.description == "" &&
		updateIssueOpts.priority == "" &&
		updateIssueOpts.assignee == "" &&
		updateIssueOpts.labels == "" &&
		len(updateIssueOpts.addLabels) == 0 &&
		len(updateIssueOpts.deleteLabels) == 0 &&
		updateIssueOpts.components == "" &&
		updateIssueOpts.dueDate == "" &&
		updateIssueOpts.linkIssue == "" {
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

	fields := make(map[string]interface{})
	update := make(map[string][]swagger.FieldUpdateOperation)

	if updateIssueOpts.summary != "" {
		fields["summary"] = updateIssueOpts.summary
	}

	if updateIssueOpts.description != "" {
		fields["description"] = newADFDocument(updateIssueOpts.description)
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
			fields["assignee"] = map[string]interface{}{
				"id": updateIssueOpts.assignee,
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

	issueUpdateDetails := swagger.NewIssueUpdateDetails()
	if len(fields) > 0 {
		issueUpdateDetails.SetFields(fields)
	}
	if len(update) > 0 {
		issueUpdateDetails.SetUpdate(update)
	}

	client := newClient()
	ctx := getAuthContext()

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
			return err
		}

		fmt.Printf("Issue %s linked to %s with type '%s'\n", updateIssueOpts.issueKey, updateIssueOpts.linkIssue, updateIssueOpts.linkType)
	}

	return nil
}
