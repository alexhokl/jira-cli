# DefaultLevelValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLevelId** | **string** | The ID of the issue security level to set as default for the specified scheme. Providing null will reset the default level. | 
**IssueSecuritySchemeId** | **string** | The ID of the issue security scheme to set default level for. | 

## Methods

### NewDefaultLevelValue

`func NewDefaultLevelValue(defaultLevelId string, issueSecuritySchemeId string, ) *DefaultLevelValue`

NewDefaultLevelValue instantiates a new DefaultLevelValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultLevelValueWithDefaults

`func NewDefaultLevelValueWithDefaults() *DefaultLevelValue`

NewDefaultLevelValueWithDefaults instantiates a new DefaultLevelValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLevelId

`func (o *DefaultLevelValue) GetDefaultLevelId() string`

GetDefaultLevelId returns the DefaultLevelId field if non-nil, zero value otherwise.

### GetDefaultLevelIdOk

`func (o *DefaultLevelValue) GetDefaultLevelIdOk() (*string, bool)`

GetDefaultLevelIdOk returns a tuple with the DefaultLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLevelId

`func (o *DefaultLevelValue) SetDefaultLevelId(v string)`

SetDefaultLevelId sets DefaultLevelId field to given value.


### GetIssueSecuritySchemeId

`func (o *DefaultLevelValue) GetIssueSecuritySchemeId() string`

GetIssueSecuritySchemeId returns the IssueSecuritySchemeId field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeIdOk

`func (o *DefaultLevelValue) GetIssueSecuritySchemeIdOk() (*string, bool)`

GetIssueSecuritySchemeIdOk returns a tuple with the IssueSecuritySchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecuritySchemeId

`func (o *DefaultLevelValue) SetIssueSecuritySchemeId(v string)`

SetIssueSecuritySchemeId sets IssueSecuritySchemeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


