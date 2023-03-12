package cmd

import (
	"context"
	"fmt"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/spf13/viper"
)

func getAuthContext() context.Context {
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	auth := swagger.BasicAuth{
		UserName: email,
		Password: apiKey,
	}

	ctx := context.WithValue(context.Background(), swagger.ContextBasicAuth, auth)
	return ctx
}

func getConfiguration() *swagger.Configuration {
	configuration := swagger.NewConfiguration()
	configuration.BasePath = fmt.Sprintf("https://%s.atlassian.net", viper.GetString("organization"))
	return configuration
}
