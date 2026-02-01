package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listWorkflowSchemesCmd = &cobra.Command{
	Use:   "workflow-schemes",
	Short: "List workflow schemes",
	Long: `List all workflow schemes available in the Jira instance.

Workflow schemes define the mapping between issue types and workflows for projects.

Note: Requires 'Administer Jira' global permission.

Examples:
  # List all workflow schemes
  jira-cli list workflow-schemes`,
	RunE: runListWorkflowSchemes,
}

func init() {
	listCmd.AddCommand(listWorkflowSchemesCmd)
}

func runListWorkflowSchemes(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	var schemes []workflowSchemeInfo
	var startAt int64 = 0

	for {
		result, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).
			StartAt(startAt).
			MaxResults(50).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to get workflow schemes: %w", err)
		}

		for _, scheme := range result.GetValues() {
			schemes = append(schemes, workflowSchemeInfo{
				id:              scheme.GetId(),
				name:            scheme.GetName(),
				description:     scheme.GetDescription(),
				defaultWorkflow: scheme.GetDefaultWorkflow(),
				draft:           scheme.GetDraft(),
			})
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(schemes) == 0 {
		fmt.Println("No workflow schemes found")
		return nil
	}

	// Build a map of scheme ID to project keys
	schemeProjects := make(map[int64][]string)

	// Get all projects first
	projects, _, err := client.ProjectsAPI.GetAllProjects(ctx).Execute()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}

	// Build project ID to key map
	projectIdToKey := make(map[string]string)
	var projectIds []int64
	for _, p := range projects {
		projectIdToKey[p.GetId()] = p.GetKey()
		var id int64
		fmt.Sscanf(p.GetId(), "%d", &id)
		projectIds = append(projectIds, id)
	}

	// Get scheme-to-project mappings
	if len(projectIds) > 0 {
		batchSize := 50
		for i := 0; i < len(projectIds); i += batchSize {
			end := i + batchSize
			if end > len(projectIds) {
				end = len(projectIds)
			}
			batch := projectIds[i:end]

			assocResult, _, err := client.WorkflowSchemeProjectAssociationsAPI.
				GetWorkflowSchemeProjectAssociations(ctx).
				ProjectId(batch).
				Execute()
			if err != nil {
				// Non-fatal: continue without project mappings
				continue
			}

			for _, assoc := range assocResult.GetValues() {
				schemeId := assoc.WorkflowScheme.GetId()
				for _, projectId := range assoc.GetProjectIds() {
					if key, ok := projectIdToKey[projectId]; ok {
						schemeProjects[schemeId] = append(schemeProjects[schemeId], key)
					}
				}
			}
		}
	}

	// Add project info to schemes
	for i := range schemes {
		if projects, ok := schemeProjects[schemes[i].id]; ok {
			sort.Strings(projects)
			schemes[i].projects = strings.Join(projects, ", ")
		}
	}

	// Sort by name
	sort.Slice(schemes, func(i, j int) bool {
		return strings.ToLower(schemes[i].name) < strings.ToLower(schemes[j].name)
	})

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("DEFAULT WORKFLOW"), cyan("PROJECTS"), cyan("DESCRIPTION"))

	for _, s := range schemes {
		// Truncate description if too long
		desc := s.description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		// Truncate projects if too long
		projects := s.projects
		if len(projects) > 25 {
			projects = projects[:22] + "..."
		}
		// Truncate default workflow if too long
		defaultWf := s.defaultWorkflow
		if len(defaultWf) > 20 {
			defaultWf = defaultWf[:17] + "..."
		}
		w.row(strconv.FormatInt(s.id, 10), yellow(s.name), defaultWf, projects, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d workflow schemes\n", len(schemes))

	return nil
}

type workflowSchemeInfo struct {
	id              int64
	name            string
	description     string
	defaultWorkflow string
	draft           bool
	projects        string
}
