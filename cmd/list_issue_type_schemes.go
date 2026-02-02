package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listIssueTypeSchemesCmd = &cobra.Command{
	Use:   "issue-type-schemes",
	Short: "List issue type schemes",
	Long: `List all issue type schemes available in the Jira instance.

Issue type schemes define which issue types are available for projects.

Note: Requires 'Administer Jira' global permission.

Examples:
  # List all issue type schemes
  jira-cli list issue-type-schemes`,
	RunE: runListIssueTypeSchemes,
}

func init() {
	listCmd.AddCommand(listIssueTypeSchemesCmd)
}

func runListIssueTypeSchemes(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	color.NoColor = noColor

	var schemes []issueTypeSchemeInfo
	var startAt int64 = 0

	for {
		result, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).
			StartAt(startAt).
			MaxResults(50).
			OrderBy("name").
			Execute()
		if err != nil {
			return wrapAPIError(fmt.Errorf("failed to get issue type schemes: %w", err))
		}

		for _, scheme := range result.GetValues() {
			isDefault := ""
			if scheme.GetIsDefault() {
				isDefault = "Yes"
			}
			schemes = append(schemes, issueTypeSchemeInfo{
				id:          scheme.GetId(),
				name:        scheme.GetName(),
				description: scheme.GetDescription(),
				isDefault:   isDefault,
			})
		}

		// Check if we've fetched all results
		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(schemes) == 0 {
		fmt.Println("No issue type schemes found")
		return nil
	}

	// Build a map of scheme ID to project keys
	schemeProjects := make(map[string][]string)

	// Get all projects first
	projects, _, err := client.ProjectsAPI.GetAllProjects(ctx).Execute()
	if err != nil {
		return wrapAPIError(fmt.Errorf("failed to get projects: %w", err))
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

	// Get scheme-to-project mappings in batches
	if len(projectIds) > 0 {
		batchSize := 50
		for i := 0; i < len(projectIds); i += batchSize {
			end := i + batchSize
			if end > len(projectIds) {
				end = len(projectIds)
			}
			batch := projectIds[i:end]

			var mappingStartAt int64 = 0
			for {
				mappingResult, _, err := client.IssueTypeSchemesAPI.GetIssueTypeSchemeForProjects(ctx).
					ProjectId(batch).
					StartAt(mappingStartAt).
					MaxResults(50).
					Execute()
				if err != nil {
					// Non-fatal: continue without project mappings
					break
				}

				for _, mapping := range mappingResult.GetValues() {
					schemeId := mapping.IssueTypeScheme.GetId()
					for _, projectId := range mapping.GetProjectIds() {
						if key, ok := projectIdToKey[projectId]; ok {
							schemeProjects[schemeId] = append(schemeProjects[schemeId], key)
						}
					}
				}

				if mappingResult.GetIsLast() || len(mappingResult.GetValues()) == 0 {
					break
				}
				mappingStartAt += int64(len(mappingResult.GetValues()))
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

	cyan := color.New(color.FgCyan).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row(cyan("ID"), cyan("NAME"), cyan("DEFAULT"), cyan("PROJECTS"), cyan("DESCRIPTION"))

	for _, s := range schemes {
		// Truncate description if too long
		desc := s.description
		if len(desc) > 40 {
			desc = desc[:37] + "..."
		}
		// Truncate projects if too long
		projects := s.projects
		if len(projects) > 30 {
			projects = projects[:27] + "..."
		}
		w.row(s.id, s.name, s.isDefault, projects, desc)
	}
	w.flush()

	fmt.Printf("\nFound %d issue type schemes\n", len(schemes))

	return nil
}

type issueTypeSchemeInfo struct {
	id          string
	name        string
	description string
	isDefault   string
	projects    string
}
