# ServerInformation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | **string** | The base URL of the Jira instance. | [optional] [default to null]
**BuildDate** | [**time.Time**](time.Time.md) | The timestamp when the Jira version was built. | [optional] [default to null]
**BuildNumber** | **int32** | The build number of the Jira version. | [optional] [default to null]
**DeploymentType** | **string** | The type of server deployment. This is always returned as *Cloud*. | [optional] [default to null]
**HealthChecks** | [**[]HealthCheckResult**](HealthCheckResult.md) | Jira instance health check results. Deprecated and no longer returned. | [optional] [default to null]
**ScmInfo** | **string** | The unique identifier of the Jira version. | [optional] [default to null]
**ServerTime** | [**time.Time**](time.Time.md) | The time in Jira when this request was responded to. | [optional] [default to null]
**ServerTitle** | **string** | The name of the Jira instance. | [optional] [default to null]
**Version** | **string** | The version of Jira. | [optional] [default to null]
**VersionNumbers** | **[]int32** | The major, minor, and revision version numbers of the Jira version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

