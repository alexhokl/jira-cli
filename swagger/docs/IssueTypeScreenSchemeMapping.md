# IssueTypeScreenSchemeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeId** | **string** | The ID of the issue type or *default*. Only issue types used in classic projects are accepted. An entry for *default* must be provided and defines the mapping for all issue types without a screen scheme. | 
**ScreenSchemeId** | **string** | The ID of the screen scheme. Only screen schemes used in classic projects are accepted. | 

## Methods

### NewIssueTypeScreenSchemeMapping

`func NewIssueTypeScreenSchemeMapping(issueTypeId string, screenSchemeId string, ) *IssueTypeScreenSchemeMapping`

NewIssueTypeScreenSchemeMapping instantiates a new IssueTypeScreenSchemeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeScreenSchemeMappingWithDefaults

`func NewIssueTypeScreenSchemeMappingWithDefaults() *IssueTypeScreenSchemeMapping`

NewIssueTypeScreenSchemeMappingWithDefaults instantiates a new IssueTypeScreenSchemeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeId

`func (o *IssueTypeScreenSchemeMapping) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *IssueTypeScreenSchemeMapping) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *IssueTypeScreenSchemeMapping) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.


### GetScreenSchemeId

`func (o *IssueTypeScreenSchemeMapping) GetScreenSchemeId() string`

GetScreenSchemeId returns the ScreenSchemeId field if non-nil, zero value otherwise.

### GetScreenSchemeIdOk

`func (o *IssueTypeScreenSchemeMapping) GetScreenSchemeIdOk() (*string, bool)`

GetScreenSchemeIdOk returns a tuple with the ScreenSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenSchemeId

`func (o *IssueTypeScreenSchemeMapping) SetScreenSchemeId(v string)`

SetScreenSchemeId sets ScreenSchemeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


