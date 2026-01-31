# jira-cli

This is a CLI application managing JIRA.

### Setup

This application uses API keys to authenticate JIRA APIs. See references for how
to get a key.

The default configuration is stored at `\$HOME/.jira-cli.yml`. Once you obtained
an API key, configure property `email` and `api_key` like the following.

```yaml
email: some_user@testing.com
api_key: key_from_jira
```

### Stub documentation

- [JIRA Cloud Platform REST API and Software Platform REST API](./jira_apis.md)

### References

- [Normal API
  token](https://id.atlassian.com/manage-profile/security/api-tokens)
- [Admin API
  key](https://support.atlassian.com/organization-administration/docs/manage-an-organization-with-the-admin-apis/)
- [API
  reference](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)
