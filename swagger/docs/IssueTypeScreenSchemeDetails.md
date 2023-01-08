# IssueTypeScreenSchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the issue type screen scheme. The maximum length is 255 characters. | [optional] 
**IssueTypeMappings** | [**[]IssueTypeScreenSchemeMapping**](IssueTypeScreenSchemeMapping.md) | The IDs of the screen schemes for the issue type IDs and *default*. A *default* entry is required to create an issue type screen scheme, it defines the mapping for all issue types without a screen scheme. | 
**Name** | **string** | The name of the issue type screen scheme. The name must be unique. The maximum length is 255 characters. | 

## Methods

### NewIssueTypeScreenSchemeDetails

`func NewIssueTypeScreenSchemeDetails(issueTypeMappings []IssueTypeScreenSchemeMapping, name string, ) *IssueTypeScreenSchemeDetails`

NewIssueTypeScreenSchemeDetails instantiates a new IssueTypeScreenSchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeScreenSchemeDetailsWithDefaults

`func NewIssueTypeScreenSchemeDetailsWithDefaults() *IssueTypeScreenSchemeDetails`

NewIssueTypeScreenSchemeDetailsWithDefaults instantiates a new IssueTypeScreenSchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueTypeScreenSchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeScreenSchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeScreenSchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeScreenSchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssueTypeMappings

`func (o *IssueTypeScreenSchemeDetails) GetIssueTypeMappings() []IssueTypeScreenSchemeMapping`

GetIssueTypeMappings returns the IssueTypeMappings field if non-nil, zero value otherwise.

### GetIssueTypeMappingsOk

`func (o *IssueTypeScreenSchemeDetails) GetIssueTypeMappingsOk() (*[]IssueTypeScreenSchemeMapping, bool)`

GetIssueTypeMappingsOk returns a tuple with the IssueTypeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeMappings

`func (o *IssueTypeScreenSchemeDetails) SetIssueTypeMappings(v []IssueTypeScreenSchemeMapping)`

SetIssueTypeMappings sets IssueTypeMappings field to given value.


### GetName

`func (o *IssueTypeScreenSchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeScreenSchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeScreenSchemeDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


