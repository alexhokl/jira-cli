# Operations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkGroups** | Pointer to [**[]LinkGroup**](LinkGroup.md) | Details of the link groups defining issue operations. | [optional] [readonly] 

## Methods

### NewOperations

`func NewOperations() *Operations`

NewOperations instantiates a new Operations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationsWithDefaults

`func NewOperationsWithDefaults() *Operations`

NewOperationsWithDefaults instantiates a new Operations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkGroups

`func (o *Operations) GetLinkGroups() []LinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *Operations) GetLinkGroupsOk() (*[]LinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *Operations) SetLinkGroups(v []LinkGroup)`

SetLinkGroups sets LinkGroups field to given value.

### HasLinkGroups

`func (o *Operations) HasLinkGroups() bool`

HasLinkGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


