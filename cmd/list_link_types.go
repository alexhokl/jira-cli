package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listLinkTypesOptions struct {
	format string
}

var listLinkTypesOpts = listLinkTypesOptions{}

var listLinkTypesCmd = &cobra.Command{
	Use:   "link-types",
	Short: "List all issue link types",
	Long: `List all available issue link types in the Jira instance.

Each link type has a name and two descriptions:
- Inward: describes the relationship from the target's perspective
- Outward: describes the relationship from the source's perspective

For example, a "Blocks" link type might have:
- Outward: "blocks" (this issue blocks the linked issue)
- Inward: "is blocked by" (the linked issue is blocked by this issue)

Examples:
  # List all link types
  jira-cli list link-types`,
	RunE: runListLinkTypes,
}

func init() {
	listCmd.AddCommand(listLinkTypesCmd)
	listLinkTypesCmd.Flags().StringVar(&listLinkTypesOpts.format, "format", "table", "Output format: table or json")
}

func runListLinkTypes(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	result, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	linkTypes := result.GetIssueLinkTypes()

	if len(linkTypes) == 0 {
		fmt.Println("No link types found")
		return nil
	}

	// Build typed slice for output
	var items []linkTypeOutput
	for _, lt := range linkTypes {
		items = append(items, linkTypeOutput{
			Name:    lt.GetName(),
			Inward:  lt.GetInward(),
			Outward: lt.GetOutward(),
		})
	}

	if listLinkTypesOpts.format == "json" {
		return printJSON(items)
	}

	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for _, item := range items {
		fmt.Printf("%s (inward: %s, outward: %s)\n",
			yellow(item.Name),
			cyan(item.Inward),
			cyan(item.Outward),
		)
	}

	return nil
}

type linkTypeOutput struct {
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}
