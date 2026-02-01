package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

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
			return fmt.Errorf("failed to get issue type schemes: %w", err)
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

	// Sort by name
	sort.Slice(schemes, func(i, j int) bool {
		return strings.ToLower(schemes[i].name) < strings.ToLower(schemes[j].name)
	})

	cyan := color.New(color.FgCyan).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, cyan("ID")+"\t"+cyan("NAME")+"\t"+cyan("DEFAULT")+"\t"+cyan("DESCRIPTION"))

	for _, s := range schemes {
		// Truncate description if too long
		desc := s.description
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", s.id, s.name, s.isDefault, desc)
	}
	w.Flush()

	fmt.Printf("\nFound %d issue type schemes\n", len(schemes))

	return nil
}

type issueTypeSchemeInfo struct {
	id          string
	name        string
	description string
	isDefault   string
}
