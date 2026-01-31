package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listIssuesOptions struct {
	jql           string
	project       string
	status        string
	assignee      string
	reporter      string
	labels        []string
	issueType     string
	priority      string
	sprint        string
	component     string
	fixVersion    string
	createdAfter  string
	createdBefore string
	updatedAfter  string
	updatedBefore string
	orderBy       string
	maxResults    int32
}

var listIssuesOpts = listIssuesOptions{}

var listIssuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List issues using JQL search",
	Long: `List issues using JQL search.

You can use the --jql flag to specify a raw JQL query, or use the convenience
flags to build a JQL query. If --jql is specified, the other filter flags are ignored.

Examples:
  # List all issues in a project
  jira-cli list issues --project MYPROJ

  # List issues assigned to current user in a specific sprint
  jira-cli list issues --project MYPROJ --assignee currentUser() --sprint 123

  # List issues with a specific label
  jira-cli list issues --project MYPROJ --label bug --label critical

  # List issues using raw JQL
  jira-cli list issues --jql "project = MYPROJ AND status = 'In Progress'"

  # List issues created in the last 7 days
  jira-cli list issues --project MYPROJ --created-after -7d

  # List high priority bugs
  jira-cli list issues --project MYPROJ --type Bug --priority High`,
	RunE: runListIssues,
}

func init() {
	listCmd.AddCommand(listIssuesCmd)

	flags := listIssuesCmd.Flags()
	flags.StringVarP(&listIssuesOpts.jql, "jql", "q", "", "Raw JQL query (overrides other filter flags)")
	flags.StringVarP(&listIssuesOpts.project, "project", "p", "", "Filter by project key")
	flags.StringVarP(&listIssuesOpts.status, "status", "s", "", "Filter by status (e.g., 'In Progress', 'Done')")
	flags.StringVarP(&listIssuesOpts.assignee, "assignee", "a", "", "Filter by assignee (use 'currentUser()' for yourself, 'unassigned' for unassigned)")
	flags.StringVarP(&listIssuesOpts.reporter, "reporter", "r", "", "Filter by reporter")
	flags.StringSliceVarP(&listIssuesOpts.labels, "label", "l", nil, "Filter by label (can be specified multiple times)")
	flags.StringVarP(&listIssuesOpts.issueType, "type", "t", "", "Filter by issue type (e.g., Bug, Story, Task)")
	flags.StringVar(&listIssuesOpts.priority, "priority", "", "Filter by priority (e.g., High, Medium, Low)")
	flags.StringVar(&listIssuesOpts.sprint, "sprint", "", "Filter by sprint ID or name (use 'openSprints()' for open sprints)")
	flags.StringVar(&listIssuesOpts.component, "component", "", "Filter by component name")
	flags.StringVar(&listIssuesOpts.fixVersion, "fix-version", "", "Filter by fix version")
	flags.StringVar(&listIssuesOpts.createdAfter, "created-after", "", "Filter by created date (e.g., '2024-01-01', '-7d')")
	flags.StringVar(&listIssuesOpts.createdBefore, "created-before", "", "Filter by created date (e.g., '2024-12-31', '-1d')")
	flags.StringVar(&listIssuesOpts.updatedAfter, "updated-after", "", "Filter by updated date (e.g., '2024-01-01', '-7d')")
	flags.StringVar(&listIssuesOpts.updatedBefore, "updated-before", "", "Filter by updated date (e.g., '2024-12-31', '-1d')")
	flags.StringVarP(&listIssuesOpts.orderBy, "order-by", "o", "", "Order by field (e.g., 'created DESC', 'priority ASC')")
	flags.Int32VarP(&listIssuesOpts.maxResults, "max-results", "m", 0, "Maximum number of results to return (0 for all)")
}

func runListIssues(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	jql := listIssuesOpts.jql
	if jql == "" {
		jql = buildJQL()
	}

	if jql == "" {
		return fmt.Errorf("at least one filter option is required (use --project, --jql, or other filter flags)")
	}

	var allIssues []swagger.IssueBean
	var nextPageToken string
	const pageSize int32 = 100 // Page size for API calls
	totalLimit := listIssuesOpts.maxResults

	for {
		request := client.IssueSearchAPI.SearchAndReconsileIssuesUsingJql(ctx).
			Jql(jql).
			MaxResults(pageSize).
			Fields([]string{"summary", "status", "assignee", "priority", "issuetype", "labels", "created", "updated"})

		if nextPageToken != "" {
			request = request.NextPageToken(nextPageToken)
		}

		result, _, err := request.Execute()
		if err != nil {
			return err
		}

		allIssues = append(allIssues, result.GetIssues()...)

		// If user specified a limit, stop when we have enough
		if totalLimit > 0 && int32(len(allIssues)) >= totalLimit {
			allIssues = allIssues[:totalLimit]
			break
		}

		// Check if this is the last page
		if result.GetIsLast() {
			break
		}

		// Get next page token for pagination
		nextPageToken = result.GetNextPageToken()
		if nextPageToken == "" {
			break
		}
	}

	if len(allIssues) == 0 {
		fmt.Println("No issues found")
		return nil
	}

	printIssues(allIssues)
	return nil
}

func buildJQL() string {
	var conditions []string

	if listIssuesOpts.project != "" {
		conditions = append(conditions, fmt.Sprintf("project = %s", quoteJQLValue(listIssuesOpts.project)))
	}

	if listIssuesOpts.status != "" {
		conditions = append(conditions, fmt.Sprintf("status = %s", quoteJQLValue(listIssuesOpts.status)))
	}

	if listIssuesOpts.assignee != "" {
		if listIssuesOpts.assignee == "unassigned" {
			conditions = append(conditions, "assignee IS EMPTY")
		} else if strings.HasSuffix(listIssuesOpts.assignee, "()") {
			// Function like currentUser()
			conditions = append(conditions, fmt.Sprintf("assignee = %s", listIssuesOpts.assignee))
		} else {
			conditions = append(conditions, fmt.Sprintf("assignee = %s", quoteJQLValue(listIssuesOpts.assignee)))
		}
	}

	if listIssuesOpts.reporter != "" {
		if strings.HasSuffix(listIssuesOpts.reporter, "()") {
			conditions = append(conditions, fmt.Sprintf("reporter = %s", listIssuesOpts.reporter))
		} else {
			conditions = append(conditions, fmt.Sprintf("reporter = %s", quoteJQLValue(listIssuesOpts.reporter)))
		}
	}

	if len(listIssuesOpts.labels) > 0 {
		for _, label := range listIssuesOpts.labels {
			conditions = append(conditions, fmt.Sprintf("labels = %s", quoteJQLValue(label)))
		}
	}

	if listIssuesOpts.issueType != "" {
		conditions = append(conditions, fmt.Sprintf("issuetype = %s", quoteJQLValue(listIssuesOpts.issueType)))
	}

	if listIssuesOpts.priority != "" {
		conditions = append(conditions, fmt.Sprintf("priority = %s", quoteJQLValue(listIssuesOpts.priority)))
	}

	if listIssuesOpts.sprint != "" {
		if strings.HasSuffix(listIssuesOpts.sprint, "()") {
			// Function like openSprints()
			conditions = append(conditions, fmt.Sprintf("sprint IN %s", listIssuesOpts.sprint))
		} else {
			conditions = append(conditions, fmt.Sprintf("sprint = %s", quoteJQLValue(listIssuesOpts.sprint)))
		}
	}

	if listIssuesOpts.component != "" {
		conditions = append(conditions, fmt.Sprintf("component = %s", quoteJQLValue(listIssuesOpts.component)))
	}

	if listIssuesOpts.fixVersion != "" {
		conditions = append(conditions, fmt.Sprintf("fixVersion = %s", quoteJQLValue(listIssuesOpts.fixVersion)))
	}

	if listIssuesOpts.createdAfter != "" {
		conditions = append(conditions, fmt.Sprintf("created >= %s", quoteJQLValue(listIssuesOpts.createdAfter)))
	}

	if listIssuesOpts.createdBefore != "" {
		conditions = append(conditions, fmt.Sprintf("created <= %s", quoteJQLValue(listIssuesOpts.createdBefore)))
	}

	if listIssuesOpts.updatedAfter != "" {
		conditions = append(conditions, fmt.Sprintf("updated >= %s", quoteJQLValue(listIssuesOpts.updatedAfter)))
	}

	if listIssuesOpts.updatedBefore != "" {
		conditions = append(conditions, fmt.Sprintf("updated <= %s", quoteJQLValue(listIssuesOpts.updatedBefore)))
	}

	jql := strings.Join(conditions, " AND ")

	if listIssuesOpts.orderBy != "" {
		jql = jql + " ORDER BY " + listIssuesOpts.orderBy
	} else {
		jql = jql + " ORDER BY status DESC"
	}

	return jql
}

func quoteJQLValue(value string) string {
	// If value contains spaces or special characters, quote it
	if strings.ContainsAny(value, " '\"") || strings.Contains(value, "-") {
		// Escape any existing quotes
		escaped := strings.ReplaceAll(value, "\"", "\\\"")
		return fmt.Sprintf("\"%s\"", escaped)
	}
	return value
}

func printIssues(issues []swagger.IssueBean) {
	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	for _, issue := range issues {
		key := issue.GetKey()
		fields := issue.GetFields()

		summary := ""
		if summaryObj, ok := fields["summary"]; ok && summaryObj != nil {
			summary = fmt.Sprintf("%v", summaryObj)
		}

		status := ""
		if statusObj, ok := fields["status"]; ok && statusObj != nil {
			if statusMap, ok := statusObj.(map[string]interface{}); ok {
				if name, ok := statusMap["name"].(string); ok {
					status = name
				}
			}
		}

		assignee := "Unassigned"
		if assigneeObj, ok := fields["assignee"]; ok && assigneeObj != nil {
			if assigneeMap, ok := assigneeObj.(map[string]interface{}); ok {
				if displayName, ok := assigneeMap["displayName"].(string); ok {
					assignee = displayName
				}
			}
		}

		priority := ""
		if priorityObj, ok := fields["priority"]; ok && priorityObj != nil {
			if priorityMap, ok := priorityObj.(map[string]interface{}); ok {
				if name, ok := priorityMap["name"].(string); ok {
					priority = name
				}
			}
		}

		issueType := ""
		if typeObj, ok := fields["issuetype"]; ok && typeObj != nil {
			if typeMap, ok := typeObj.(map[string]interface{}); ok {
				if name, ok := typeMap["name"].(string); ok {
					issueType = name
				}
			}
		}

		// Color the status based on common status categories
		statusColored := cyan(status)
		statusLower := strings.ToLower(status)
		if strings.Contains(statusLower, "done") || strings.Contains(statusLower, "closed") || strings.Contains(statusLower, "resolved") {
			statusColored = green(status)
		} else if strings.Contains(statusLower, "progress") || strings.Contains(statusLower, "review") {
			statusColored = yellow(status)
		} else if strings.Contains(statusLower, "block") {
			statusColored = red(status)
		}

		// Color priority
		priorityColored := priority
		priorityLower := strings.ToLower(priority)
		if strings.Contains(priorityLower, "high") || strings.Contains(priorityLower, "critical") || strings.Contains(priorityLower, "blocker") {
			priorityColored = red(priority)
		} else if strings.Contains(priorityLower, "medium") {
			priorityColored = yellow(priority)
		} else if strings.Contains(priorityLower, "low") || strings.Contains(priorityLower, "minor") {
			priorityColored = green(priority)
		}

		// Format output: KEY [Type] Summary (Status) - Assignee [Priority]
		output := fmt.Sprintf("%s", yellow(key))
		if issueType != "" {
			output += fmt.Sprintf(" [%s]", magenta(issueType))
		}
		output += fmt.Sprintf(" %s", summary)
		if status != "" {
			output += fmt.Sprintf(" (%s)", statusColored)
		}
		output += fmt.Sprintf(" - %s", cyan(assignee))
		if priority != "" {
			output += fmt.Sprintf(" [%s]", priorityColored)
		}

		fmt.Println(output)
	}
}
