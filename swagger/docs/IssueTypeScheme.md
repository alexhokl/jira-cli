# IssueTypeScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIssueTypeId** | Pointer to **string** | The ID of the default issue type of the issue type scheme. | [optional] 
**Description** | Pointer to **string** | The description of the issue type scheme. | [optional] 
**Id** | **string** | The ID of the issue type scheme. | 
**IsDefault** | Pointer to **bool** | Whether the issue type scheme is the default. | [optional] 
**Name** | **string** | The name of the issue type scheme. | 

## Methods

### NewIssueTypeScheme

`func NewIssueTypeScheme(id string, name string, ) *IssueTypeScheme`

NewIssueTypeScheme instantiates a new IssueTypeScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeSchemeWithDefaults

`func NewIssueTypeSchemeWithDefaults() *IssueTypeScheme`

NewIssueTypeSchemeWithDefaults instantiates a new IssueTypeScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultIssueTypeId

`func (o *IssueTypeScheme) GetDefaultIssueTypeId() string`

GetDefaultIssueTypeId returns the DefaultIssueTypeId field if non-nil, zero value otherwise.

### GetDefaultIssueTypeIdOk

`func (o *IssueTypeScheme) GetDefaultIssueTypeIdOk() (*string, bool)`

GetDefaultIssueTypeIdOk returns a tuple with the DefaultIssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueTypeId

`func (o *IssueTypeScheme) SetDefaultIssueTypeId(v string)`

SetDefaultIssueTypeId sets DefaultIssueTypeId field to given value.

### HasDefaultIssueTypeId

`func (o *IssueTypeScheme) HasDefaultIssueTypeId() bool`

HasDefaultIssueTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IssueTypeScheme) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeScheme) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeScheme) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *IssueTypeScheme) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IssueTypeScheme) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IssueTypeScheme) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IssueTypeScheme) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeScheme) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


