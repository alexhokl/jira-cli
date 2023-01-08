# ServerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The base URL of the Jira instance. | [optional] 
**BuildDate** | Pointer to **time.Time** | The timestamp when the Jira version was built. | [optional] 
**BuildNumber** | Pointer to **int32** | The build number of the Jira version. | [optional] 
**DeploymentType** | Pointer to **string** | The type of server deployment. This is always returned as *Cloud*. | [optional] 
**DisplayUrl** | Pointer to **string** | The display URL of the Jira instance. | [optional] 
**DisplayUrlConfluence** | Pointer to **string** | The display URL of Confluence. | [optional] 
**DisplayUrlServicedeskHelpCenter** | Pointer to **string** | The display URL of the Servicedesk Help Center. | [optional] 
**HealthChecks** | Pointer to [**[]HealthCheckResult**](HealthCheckResult.md) | Jira instance health check results. Deprecated and no longer returned. | [optional] 
**ScmInfo** | Pointer to **string** | The unique identifier of the Jira version. | [optional] 
**ServerTime** | Pointer to **time.Time** | The time in Jira when this request was responded to. | [optional] 
**ServerTimeZone** | Pointer to **string** | The default timezone of the Jira server. In a format known as Olson Time Zones, IANA Time Zones or TZ Database Time Zones. | [optional] 
**ServerTitle** | Pointer to **string** | The name of the Jira instance. | [optional] 
**Version** | Pointer to **string** | The version of Jira. | [optional] 
**VersionNumbers** | Pointer to **[]int32** | The major, minor, and revision version numbers of the Jira version. | [optional] 

## Methods

### NewServerInformation

`func NewServerInformation() *ServerInformation`

NewServerInformation instantiates a new ServerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInformationWithDefaults

`func NewServerInformationWithDefaults() *ServerInformation`

NewServerInformationWithDefaults instantiates a new ServerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *ServerInformation) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ServerInformation) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ServerInformation) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ServerInformation) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetBuildDate

`func (o *ServerInformation) GetBuildDate() time.Time`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *ServerInformation) GetBuildDateOk() (*time.Time, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *ServerInformation) SetBuildDate(v time.Time)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *ServerInformation) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetBuildNumber

`func (o *ServerInformation) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *ServerInformation) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *ServerInformation) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *ServerInformation) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetDeploymentType

`func (o *ServerInformation) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *ServerInformation) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *ServerInformation) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *ServerInformation) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDisplayUrl

`func (o *ServerInformation) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ServerInformation) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ServerInformation) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *ServerInformation) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplayUrlConfluence

`func (o *ServerInformation) GetDisplayUrlConfluence() string`

GetDisplayUrlConfluence returns the DisplayUrlConfluence field if non-nil, zero value otherwise.

### GetDisplayUrlConfluenceOk

`func (o *ServerInformation) GetDisplayUrlConfluenceOk() (*string, bool)`

GetDisplayUrlConfluenceOk returns a tuple with the DisplayUrlConfluence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrlConfluence

`func (o *ServerInformation) SetDisplayUrlConfluence(v string)`

SetDisplayUrlConfluence sets DisplayUrlConfluence field to given value.

### HasDisplayUrlConfluence

`func (o *ServerInformation) HasDisplayUrlConfluence() bool`

HasDisplayUrlConfluence returns a boolean if a field has been set.

### GetDisplayUrlServicedeskHelpCenter

`func (o *ServerInformation) GetDisplayUrlServicedeskHelpCenter() string`

GetDisplayUrlServicedeskHelpCenter returns the DisplayUrlServicedeskHelpCenter field if non-nil, zero value otherwise.

### GetDisplayUrlServicedeskHelpCenterOk

`func (o *ServerInformation) GetDisplayUrlServicedeskHelpCenterOk() (*string, bool)`

GetDisplayUrlServicedeskHelpCenterOk returns a tuple with the DisplayUrlServicedeskHelpCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrlServicedeskHelpCenter

`func (o *ServerInformation) SetDisplayUrlServicedeskHelpCenter(v string)`

SetDisplayUrlServicedeskHelpCenter sets DisplayUrlServicedeskHelpCenter field to given value.

### HasDisplayUrlServicedeskHelpCenter

`func (o *ServerInformation) HasDisplayUrlServicedeskHelpCenter() bool`

HasDisplayUrlServicedeskHelpCenter returns a boolean if a field has been set.

### GetHealthChecks

`func (o *ServerInformation) GetHealthChecks() []HealthCheckResult`

GetHealthChecks returns the HealthChecks field if non-nil, zero value otherwise.

### GetHealthChecksOk

`func (o *ServerInformation) GetHealthChecksOk() (*[]HealthCheckResult, bool)`

GetHealthChecksOk returns a tuple with the HealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChecks

`func (o *ServerInformation) SetHealthChecks(v []HealthCheckResult)`

SetHealthChecks sets HealthChecks field to given value.

### HasHealthChecks

`func (o *ServerInformation) HasHealthChecks() bool`

HasHealthChecks returns a boolean if a field has been set.

### GetScmInfo

`func (o *ServerInformation) GetScmInfo() string`

GetScmInfo returns the ScmInfo field if non-nil, zero value otherwise.

### GetScmInfoOk

`func (o *ServerInformation) GetScmInfoOk() (*string, bool)`

GetScmInfoOk returns a tuple with the ScmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmInfo

`func (o *ServerInformation) SetScmInfo(v string)`

SetScmInfo sets ScmInfo field to given value.

### HasScmInfo

`func (o *ServerInformation) HasScmInfo() bool`

HasScmInfo returns a boolean if a field has been set.

### GetServerTime

`func (o *ServerInformation) GetServerTime() time.Time`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ServerInformation) GetServerTimeOk() (*time.Time, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ServerInformation) SetServerTime(v time.Time)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ServerInformation) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetServerTimeZone

`func (o *ServerInformation) GetServerTimeZone() string`

GetServerTimeZone returns the ServerTimeZone field if non-nil, zero value otherwise.

### GetServerTimeZoneOk

`func (o *ServerInformation) GetServerTimeZoneOk() (*string, bool)`

GetServerTimeZoneOk returns a tuple with the ServerTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimeZone

`func (o *ServerInformation) SetServerTimeZone(v string)`

SetServerTimeZone sets ServerTimeZone field to given value.

### HasServerTimeZone

`func (o *ServerInformation) HasServerTimeZone() bool`

HasServerTimeZone returns a boolean if a field has been set.

### GetServerTitle

`func (o *ServerInformation) GetServerTitle() string`

GetServerTitle returns the ServerTitle field if non-nil, zero value otherwise.

### GetServerTitleOk

`func (o *ServerInformation) GetServerTitleOk() (*string, bool)`

GetServerTitleOk returns a tuple with the ServerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTitle

`func (o *ServerInformation) SetServerTitle(v string)`

SetServerTitle sets ServerTitle field to given value.

### HasServerTitle

`func (o *ServerInformation) HasServerTitle() bool`

HasServerTitle returns a boolean if a field has been set.

### GetVersion

`func (o *ServerInformation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerInformation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerInformation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerInformation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionNumbers

`func (o *ServerInformation) GetVersionNumbers() []int32`

GetVersionNumbers returns the VersionNumbers field if non-nil, zero value otherwise.

### GetVersionNumbersOk

`func (o *ServerInformation) GetVersionNumbersOk() (*[]int32, bool)`

GetVersionNumbersOk returns a tuple with the VersionNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumbers

`func (o *ServerInformation) SetVersionNumbers(v []int32)`

SetVersionNumbers sets VersionNumbers field to given value.

### HasVersionNumbers

`func (o *ServerInformation) HasVersionNumbers() bool`

HasVersionNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


