package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger"
)

type listUsersOptions struct {
	maxResults int32
	state      string
	kind       string
}

var listUsersOpts = listUsersOptions{}

var listUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all users",
	Long: `List all users in the Jira instance.

Returns active users, inactive users, and previously deleted users that have an Atlassian account.

Examples:
  # List all users
  jira-cli list users

  # List only active users
  jira-cli list users --state active

  # List only inactive users
  jira-cli list users --state inactive

  # List only regular Atlassian users
  jira-cli list users --kind atlassian

  # List only app/service accounts
  jira-cli list users --kind app

  # Combine filters
  jira-cli list users --state active --kind atlassian

  # List first 100 users
  jira-cli list users --max-results 100`,
	RunE: runListUsers,
}

func init() {
	listCmd.AddCommand(listUsersCmd)

	flags := listUsersCmd.Flags()
	flags.Int32VarP(&listUsersOpts.maxResults, "max-results", "m", 0, "Maximum number of results to return (0 for all)")
	flags.StringVarP(&listUsersOpts.state, "state", "s", "", "Filter by user state (active, inactive)")
	flags.StringVarP(&listUsersOpts.kind, "kind", "k", "", "Filter by account type (atlassian, app, customer)")
}

func runListUsers(_ *cobra.Command, _ []string) error {
	client := newClient()
	ctx := getAuthContext()

	// Validate state filter
	var filterActive *bool
	if listUsersOpts.state != "" {
		switch listUsersOpts.state {
		case "active":
			active := true
			filterActive = &active
		case "inactive":
			active := false
			filterActive = &active
		default:
			return fmt.Errorf("invalid state: %s (must be 'active' or 'inactive')", listUsersOpts.state)
		}
	}

	// Validate kind filter
	var filterKind string
	if listUsersOpts.kind != "" {
		switch listUsersOpts.kind {
		case "atlassian", "app", "customer":
			filterKind = listUsersOpts.kind
		default:
			return fmt.Errorf("invalid kind: %s (must be 'atlassian', 'app', or 'customer')", listUsersOpts.kind)
		}
	}

	var allUsers []swagger.User
	var startAt int32 = 0
	const pageSize int32 = 100

	for {
		result, _, err := client.UsersAPI.GetAllUsers(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			return wrapAPIError(err)
		}

		// Filter users
		for _, user := range result {
			if filterActive != nil && user.GetActive() != *filterActive {
				continue
			}
			if filterKind != "" && user.GetAccountType() != filterKind {
				continue
			}
			allUsers = append(allUsers, user)
		}

		// If user specified a limit, stop when we have enough
		if listUsersOpts.maxResults > 0 && int32(len(allUsers)) >= listUsersOpts.maxResults {
			allUsers = allUsers[:listUsersOpts.maxResults]
			break
		}

		// If we got fewer results than requested, we've reached the end
		if int32(len(result)) < pageSize {
			break
		}

		startAt += int32(len(result))
	}

	if len(allUsers) == 0 {
		fmt.Println("No users found")
		return nil
	}

	printUsers(allUsers)
	return nil
}

func printUsers(users []swagger.User) {
	color.NoColor = noColor
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for _, user := range users {
		displayName := user.GetDisplayName()
		email := user.GetEmailAddress()
		accountType := user.GetAccountType()
		active := user.GetActive()

		// Format status
		status := green("active")
		if !active {
			status = red("inactive")
		}

		// Format output
		output := yellow(displayName)
		if email != "" {
			output += fmt.Sprintf(" <%s>", cyan(email))
		}
		output += fmt.Sprintf(" [%s]", accountType)
		output += fmt.Sprintf(" (%s)", status)

		fmt.Println(output)
	}
}
