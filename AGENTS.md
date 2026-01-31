# AGENTS.md - Coding Agent Guidelines for jira-cli

This document provides guidelines for AI coding agents working in this repository.

## Project Overview

- **Language**: Go 1.19+
- **Type**: CLI application for interacting with Jira Cloud APIs
- **CLI Framework**: Cobra (github.com/spf13/cobra)
- **Configuration**: Viper (github.com/spf13/viper)
- **API Client**: Auto-generated Swagger client for Jira Cloud REST API v3
- **API Reference**: See `jira_apis.md` for a description of available Jira APIs

## Build, Test, and Lint Commands

```bash
# Install the CLI binary
task install

# Run all tests
task test
go test ./...

# Run a specific test by name
go test ./... -run TestFunctionName

# Run tests with coverage
task coverage

# Format and vet code
go fmt ./...
go vet ./...

# Tidy dependencies
go mod tidy
```

## Code Style Guidelines

### Import Organization

Imports should be grouped in this order, separated by blank lines:
1. Standard library packages
2. External/third-party packages
3. Internal project packages

```go
import (
    "context"
    "fmt"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    "github.com/alexhokl/jira-cli/swagger"
)
```

### Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Package names | lowercase, single word | `cmd`, `swagger` |
| Variables | camelCase | `cfgFile`, `apiKey` |
| Exported functions | PascalCase | `Execute()` |
| Unexported functions | camelCase | `newClient()`, `getAuthContext()` |
| Struct types | PascalCase | `Comment`, `CommentBody` |
| Options structs | camelCase + "Options" | `getIssueOptions` |
| Command variables | camelCase + "Cmd" | `getIssueCmd` |
| Options variables | camelCase + "Opts" | `getIssueOpts` |

### Command Structure Pattern

Follow this structure for new commands in `cmd/<action>_<subject>.go`:

```go
package cmd

type actionSubjectOptions struct {
    id string
}

var actionSubjectOpts = actionSubjectOptions{}

var actionSubjectCmd = &cobra.Command{
    Use:   "subject",
    Short: "Brief description",
    RunE:  runActionSubject,  // Use RunE for commands that can error
}

func init() {
    parentCmd.AddCommand(actionSubjectCmd)
    flags := actionSubjectCmd.Flags()
    flags.StringVarP(&actionSubjectOpts.id, "id", "i", "", "Description")
    actionSubjectCmd.MarkFlagRequired("id")
}

func runActionSubject(_ *cobra.Command, _ []string) error {
    client := newClient()
    ctx := getAuthContext()
    // Make API call, return error on failure
    return nil
}
```

### Error Handling

- Use `RunE` (not `Run`) for commands that can return errors
- Return errors directly; Cobra handles display
- Use `fmt.Errorf()` for creating error messages

### API Client Usage

Always use the helper functions from `cmd/helper.go`:

```go
ctx := getAuthContext()     // Get authenticated context
client := newClient()       // Create new API client
result, _, err := client.SomeApi.SomeMethod(ctx, params, opts)
```

## Project Structure

```
jira-cli/
├── main.go              # Entry point - calls cmd.Execute()
├── go.mod               # Go module definition
├── Taskfile.yml         # Build automation (go-task)
├── cmd/                 # CLI commands (Cobra-based)
│   ├── root.go          # Root command, config setup
│   ├── helper.go        # Shared utilities (auth, client)
│   ├── get.go           # Parent command for 'get' subcommands
│   ├── get_issue.go     # 'get issue' implementation
│   ├── add.go           # Parent command for 'add' subcommands
│   └── add_comment.go   # 'add comment' implementation
└── swagger/             # Auto-generated Jira API client (DO NOT EDIT)
```

### Configuration

The application reads configuration from `$HOME/.jira-cli.yml`:

```yaml
email: user@example.com
api_key: your_jira_api_key
organization: your_org_name
```

### Adding New Commands

1. Create new file: `cmd/<action>_<subject>.go`
2. Define options struct and package-level variable
3. Create command with `RunE` function
4. Register in `init()` under appropriate parent command
5. Implement using `newClient()` and `getAuthContext()` helpers

### Testing Conventions

- Test files: `*_test.go` alongside code they test
- Use Go's standard `testing` package
- Use table-driven tests where appropriate
