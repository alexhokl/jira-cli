# IssueLayoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**IssueLayoutType** | Pointer to **string** | The issue layout type | [optional] 
**Items** | Pointer to [**[]IssueLayoutItemPayload**](IssueLayoutItemPayload.md) | The configuration of items in the issue layout | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewIssueLayoutPayload

`func NewIssueLayoutPayload() *IssueLayoutPayload`

NewIssueLayoutPayload instantiates a new IssueLayoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLayoutPayloadWithDefaults

`func NewIssueLayoutPayloadWithDefaults() *IssueLayoutPayload`

NewIssueLayoutPayloadWithDefaults instantiates a new IssueLayoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *IssueLayoutPayload) GetContainerId() ProjectCreateResourceIdentifier`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *IssueLayoutPayload) GetContainerIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *IssueLayoutPayload) SetContainerId(v ProjectCreateResourceIdentifier)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *IssueLayoutPayload) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetIssueLayoutType

`func (o *IssueLayoutPayload) GetIssueLayoutType() string`

GetIssueLayoutType returns the IssueLayoutType field if non-nil, zero value otherwise.

### GetIssueLayoutTypeOk

`func (o *IssueLayoutPayload) GetIssueLayoutTypeOk() (*string, bool)`

GetIssueLayoutTypeOk returns a tuple with the IssueLayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLayoutType

`func (o *IssueLayoutPayload) SetIssueLayoutType(v string)`

SetIssueLayoutType sets IssueLayoutType field to given value.

### HasIssueLayoutType

`func (o *IssueLayoutPayload) HasIssueLayoutType() bool`

HasIssueLayoutType returns a boolean if a field has been set.

### GetItems

`func (o *IssueLayoutPayload) GetItems() []IssueLayoutItemPayload`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IssueLayoutPayload) GetItemsOk() (*[]IssueLayoutItemPayload, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IssueLayoutPayload) SetItems(v []IssueLayoutItemPayload)`

SetItems sets Items field to given value.

### HasItems

`func (o *IssueLayoutPayload) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPcri

`func (o *IssueLayoutPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *IssueLayoutPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *IssueLayoutPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *IssueLayoutPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


