# ProjectDataPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyContentBlocked** | Pointer to **bool** | Whether the project contains any content inaccessible to the requesting application. | [optional] [readonly] 

## Methods

### NewProjectDataPolicy

`func NewProjectDataPolicy() *ProjectDataPolicy`

NewProjectDataPolicy instantiates a new ProjectDataPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDataPolicyWithDefaults

`func NewProjectDataPolicyWithDefaults() *ProjectDataPolicy`

NewProjectDataPolicyWithDefaults instantiates a new ProjectDataPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyContentBlocked

`func (o *ProjectDataPolicy) GetAnyContentBlocked() bool`

GetAnyContentBlocked returns the AnyContentBlocked field if non-nil, zero value otherwise.

### GetAnyContentBlockedOk

`func (o *ProjectDataPolicy) GetAnyContentBlockedOk() (*bool, bool)`

GetAnyContentBlockedOk returns a tuple with the AnyContentBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyContentBlocked

`func (o *ProjectDataPolicy) SetAnyContentBlocked(v bool)`

SetAnyContentBlocked sets AnyContentBlocked field to given value.

### HasAnyContentBlocked

`func (o *ProjectDataPolicy) HasAnyContentBlocked() bool`

HasAnyContentBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


