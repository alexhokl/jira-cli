# OrderOfIssueTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The ID of the issue type to place the moved issue types after. Required if &#x60;position&#x60; isn&#39;t provided. | [optional] 
**IssueTypeIds** | **[]string** | A list of the issue type IDs to move. The order of the issue type IDs in the list is the order they are given after the move. | 
**Position** | Pointer to **string** | The position the issue types should be moved to. Required if &#x60;after&#x60; isn&#39;t provided. | [optional] 

## Methods

### NewOrderOfIssueTypes

`func NewOrderOfIssueTypes(issueTypeIds []string, ) *OrderOfIssueTypes`

NewOrderOfIssueTypes instantiates a new OrderOfIssueTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOfIssueTypesWithDefaults

`func NewOrderOfIssueTypesWithDefaults() *OrderOfIssueTypes`

NewOrderOfIssueTypesWithDefaults instantiates a new OrderOfIssueTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *OrderOfIssueTypes) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *OrderOfIssueTypes) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *OrderOfIssueTypes) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *OrderOfIssueTypes) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *OrderOfIssueTypes) GetIssueTypeIds() []string`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *OrderOfIssueTypes) GetIssueTypeIdsOk() (*[]string, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *OrderOfIssueTypes) SetIssueTypeIds(v []string)`

SetIssueTypeIds sets IssueTypeIds field to given value.


### GetPosition

`func (o *OrderOfIssueTypes) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *OrderOfIssueTypes) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *OrderOfIssueTypes) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *OrderOfIssueTypes) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


