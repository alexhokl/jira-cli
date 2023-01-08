# IssueTypeSchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIssueTypeId** | Pointer to **string** | The ID of the default issue type of the issue type scheme. This ID must be included in &#x60;issueTypeIds&#x60;. | [optional] 
**Description** | Pointer to **string** | The description of the issue type scheme. The maximum length is 4000 characters. | [optional] 
**IssueTypeIds** | **[]string** | The list of issue types IDs of the issue type scheme. At least one standard issue type ID is required. | 
**Name** | **string** | The name of the issue type scheme. The name must be unique. The maximum length is 255 characters. | 

## Methods

### NewIssueTypeSchemeDetails

`func NewIssueTypeSchemeDetails(issueTypeIds []string, name string, ) *IssueTypeSchemeDetails`

NewIssueTypeSchemeDetails instantiates a new IssueTypeSchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeSchemeDetailsWithDefaults

`func NewIssueTypeSchemeDetailsWithDefaults() *IssueTypeSchemeDetails`

NewIssueTypeSchemeDetailsWithDefaults instantiates a new IssueTypeSchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultIssueTypeId

`func (o *IssueTypeSchemeDetails) GetDefaultIssueTypeId() string`

GetDefaultIssueTypeId returns the DefaultIssueTypeId field if non-nil, zero value otherwise.

### GetDefaultIssueTypeIdOk

`func (o *IssueTypeSchemeDetails) GetDefaultIssueTypeIdOk() (*string, bool)`

GetDefaultIssueTypeIdOk returns a tuple with the DefaultIssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueTypeId

`func (o *IssueTypeSchemeDetails) SetDefaultIssueTypeId(v string)`

SetDefaultIssueTypeId sets DefaultIssueTypeId field to given value.

### HasDefaultIssueTypeId

`func (o *IssueTypeSchemeDetails) HasDefaultIssueTypeId() bool`

HasDefaultIssueTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeSchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeSchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeSchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeSchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *IssueTypeSchemeDetails) GetIssueTypeIds() []string`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *IssueTypeSchemeDetails) GetIssueTypeIdsOk() (*[]string, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *IssueTypeSchemeDetails) SetIssueTypeIds(v []string)`

SetIssueTypeIds sets IssueTypeIds field to given value.


### GetName

`func (o *IssueTypeSchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeSchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeSchemeDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


