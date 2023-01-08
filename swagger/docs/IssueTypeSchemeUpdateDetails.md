# IssueTypeSchemeUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIssueTypeId** | Pointer to **string** | The ID of the default issue type of the issue type scheme. | [optional] 
**Description** | Pointer to **string** | The description of the issue type scheme. The maximum length is 4000 characters. | [optional] 
**Name** | Pointer to **string** | The name of the issue type scheme. The name must be unique. The maximum length is 255 characters. | [optional] 

## Methods

### NewIssueTypeSchemeUpdateDetails

`func NewIssueTypeSchemeUpdateDetails() *IssueTypeSchemeUpdateDetails`

NewIssueTypeSchemeUpdateDetails instantiates a new IssueTypeSchemeUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeSchemeUpdateDetailsWithDefaults

`func NewIssueTypeSchemeUpdateDetailsWithDefaults() *IssueTypeSchemeUpdateDetails`

NewIssueTypeSchemeUpdateDetailsWithDefaults instantiates a new IssueTypeSchemeUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultIssueTypeId

`func (o *IssueTypeSchemeUpdateDetails) GetDefaultIssueTypeId() string`

GetDefaultIssueTypeId returns the DefaultIssueTypeId field if non-nil, zero value otherwise.

### GetDefaultIssueTypeIdOk

`func (o *IssueTypeSchemeUpdateDetails) GetDefaultIssueTypeIdOk() (*string, bool)`

GetDefaultIssueTypeIdOk returns a tuple with the DefaultIssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueTypeId

`func (o *IssueTypeSchemeUpdateDetails) SetDefaultIssueTypeId(v string)`

SetDefaultIssueTypeId sets DefaultIssueTypeId field to given value.

### HasDefaultIssueTypeId

`func (o *IssueTypeSchemeUpdateDetails) HasDefaultIssueTypeId() bool`

HasDefaultIssueTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeSchemeUpdateDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeSchemeUpdateDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeSchemeUpdateDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeSchemeUpdateDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeSchemeUpdateDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeSchemeUpdateDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeSchemeUpdateDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeSchemeUpdateDetails) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


