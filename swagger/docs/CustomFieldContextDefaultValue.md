# CustomFieldContextDefaultValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CascadingOptionId** | Pointer to **string** | The ID of the default cascading option. | [optional] 
**ContextId** | **string** | The ID of the context. | 
**OptionId** | **string** | The ID of the default option. | 
**Type** | **string** |  | 
**OptionIds** | **[]string** | The list of IDs of the default options. | 
**AccountId** | **string** | The ID of the default user. | 
**UserFilter** | [**UserFilter**](UserFilter.md) |  | 
**AccountIds** | **[]string** | The IDs of the default users. | 
**GroupId** | **string** | The ID of the the default group. | 
**GroupIds** | **[]string** | The IDs of the default groups. | 
**Date** | Pointer to **string** | The default date in ISO format. Ignored if &#x60;useCurrent&#x60; is true. | [optional] 
**UseCurrent** | Pointer to **bool** | Whether to use the current date. | [optional] [default to false]
**DateTime** | Pointer to **string** | The default date-time in ISO format. Ignored if &#x60;useCurrent&#x60; is true. | [optional] 
**Url** | **string** | The default URL. | 
**ProjectId** | **string** | The ID of the default project. | 
**Number** | **float64** | The default floating-point number. | 
**Labels** | **[]string** | The default labels value. | 
**Text** | Pointer to **string** | The default text. The maximum length is 254 characters. | [optional] 
**VersionId** | **string** | The ID of the default version. | 
**VersionOrder** | Pointer to **string** | The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are &#x60;\&quot;releasedFirst\&quot;&#x60; and &#x60;\&quot;unreleasedFirst\&quot;&#x60;. | [optional] 
**VersionIds** | **[]string** | The IDs of the default versions. | 
**Values** | Pointer to **[]string** | List of string values. The maximum length for a value is 254 characters. | [optional] 
**Object** | Pointer to **map[string]interface{}** | The default JSON object. | [optional] 

## Methods

### NewCustomFieldContextDefaultValue

`func NewCustomFieldContextDefaultValue(contextId string, optionId string, type_ string, optionIds []string, accountId string, userFilter UserFilter, accountIds []string, groupId string, groupIds []string, url string, projectId string, number float64, labels []string, versionId string, versionIds []string, ) *CustomFieldContextDefaultValue`

NewCustomFieldContextDefaultValue instantiates a new CustomFieldContextDefaultValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueWithDefaults

`func NewCustomFieldContextDefaultValueWithDefaults() *CustomFieldContextDefaultValue`

NewCustomFieldContextDefaultValueWithDefaults instantiates a new CustomFieldContextDefaultValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadingOptionId

`func (o *CustomFieldContextDefaultValue) GetCascadingOptionId() string`

GetCascadingOptionId returns the CascadingOptionId field if non-nil, zero value otherwise.

### GetCascadingOptionIdOk

`func (o *CustomFieldContextDefaultValue) GetCascadingOptionIdOk() (*string, bool)`

GetCascadingOptionIdOk returns a tuple with the CascadingOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadingOptionId

`func (o *CustomFieldContextDefaultValue) SetCascadingOptionId(v string)`

SetCascadingOptionId sets CascadingOptionId field to given value.

### HasCascadingOptionId

`func (o *CustomFieldContextDefaultValue) HasCascadingOptionId() bool`

HasCascadingOptionId returns a boolean if a field has been set.

### GetContextId

`func (o *CustomFieldContextDefaultValue) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CustomFieldContextDefaultValue) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CustomFieldContextDefaultValue) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetOptionId

`func (o *CustomFieldContextDefaultValue) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *CustomFieldContextDefaultValue) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *CustomFieldContextDefaultValue) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.


### GetType

`func (o *CustomFieldContextDefaultValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValue) SetType(v string)`

SetType sets Type field to given value.


### GetOptionIds

`func (o *CustomFieldContextDefaultValue) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *CustomFieldContextDefaultValue) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *CustomFieldContextDefaultValue) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.


### GetAccountId

`func (o *CustomFieldContextDefaultValue) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomFieldContextDefaultValue) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomFieldContextDefaultValue) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetUserFilter

`func (o *CustomFieldContextDefaultValue) GetUserFilter() UserFilter`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *CustomFieldContextDefaultValue) GetUserFilterOk() (*UserFilter, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *CustomFieldContextDefaultValue) SetUserFilter(v UserFilter)`

SetUserFilter sets UserFilter field to given value.


### GetAccountIds

`func (o *CustomFieldContextDefaultValue) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *CustomFieldContextDefaultValue) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *CustomFieldContextDefaultValue) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.


### GetGroupId

`func (o *CustomFieldContextDefaultValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CustomFieldContextDefaultValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CustomFieldContextDefaultValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupIds

`func (o *CustomFieldContextDefaultValue) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CustomFieldContextDefaultValue) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CustomFieldContextDefaultValue) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### GetDate

`func (o *CustomFieldContextDefaultValue) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CustomFieldContextDefaultValue) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CustomFieldContextDefaultValue) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CustomFieldContextDefaultValue) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUseCurrent

`func (o *CustomFieldContextDefaultValue) GetUseCurrent() bool`

GetUseCurrent returns the UseCurrent field if non-nil, zero value otherwise.

### GetUseCurrentOk

`func (o *CustomFieldContextDefaultValue) GetUseCurrentOk() (*bool, bool)`

GetUseCurrentOk returns a tuple with the UseCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrent

`func (o *CustomFieldContextDefaultValue) SetUseCurrent(v bool)`

SetUseCurrent sets UseCurrent field to given value.

### HasUseCurrent

`func (o *CustomFieldContextDefaultValue) HasUseCurrent() bool`

HasUseCurrent returns a boolean if a field has been set.

### GetDateTime

`func (o *CustomFieldContextDefaultValue) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *CustomFieldContextDefaultValue) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *CustomFieldContextDefaultValue) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *CustomFieldContextDefaultValue) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetUrl

`func (o *CustomFieldContextDefaultValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomFieldContextDefaultValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomFieldContextDefaultValue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProjectId

`func (o *CustomFieldContextDefaultValue) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CustomFieldContextDefaultValue) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CustomFieldContextDefaultValue) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetNumber

`func (o *CustomFieldContextDefaultValue) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CustomFieldContextDefaultValue) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CustomFieldContextDefaultValue) SetNumber(v float64)`

SetNumber sets Number field to given value.


### GetLabels

`func (o *CustomFieldContextDefaultValue) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CustomFieldContextDefaultValue) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CustomFieldContextDefaultValue) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetText

`func (o *CustomFieldContextDefaultValue) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CustomFieldContextDefaultValue) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CustomFieldContextDefaultValue) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CustomFieldContextDefaultValue) HasText() bool`

HasText returns a boolean if a field has been set.

### GetVersionId

`func (o *CustomFieldContextDefaultValue) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *CustomFieldContextDefaultValue) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *CustomFieldContextDefaultValue) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionOrder

`func (o *CustomFieldContextDefaultValue) GetVersionOrder() string`

GetVersionOrder returns the VersionOrder field if non-nil, zero value otherwise.

### GetVersionOrderOk

`func (o *CustomFieldContextDefaultValue) GetVersionOrderOk() (*string, bool)`

GetVersionOrderOk returns a tuple with the VersionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrder

`func (o *CustomFieldContextDefaultValue) SetVersionOrder(v string)`

SetVersionOrder sets VersionOrder field to given value.

### HasVersionOrder

`func (o *CustomFieldContextDefaultValue) HasVersionOrder() bool`

HasVersionOrder returns a boolean if a field has been set.

### GetVersionIds

`func (o *CustomFieldContextDefaultValue) GetVersionIds() []string`

GetVersionIds returns the VersionIds field if non-nil, zero value otherwise.

### GetVersionIdsOk

`func (o *CustomFieldContextDefaultValue) GetVersionIdsOk() (*[]string, bool)`

GetVersionIdsOk returns a tuple with the VersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIds

`func (o *CustomFieldContextDefaultValue) SetVersionIds(v []string)`

SetVersionIds sets VersionIds field to given value.


### GetValues

`func (o *CustomFieldContextDefaultValue) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomFieldContextDefaultValue) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomFieldContextDefaultValue) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CustomFieldContextDefaultValue) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetObject

`func (o *CustomFieldContextDefaultValue) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomFieldContextDefaultValue) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomFieldContextDefaultValue) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomFieldContextDefaultValue) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


