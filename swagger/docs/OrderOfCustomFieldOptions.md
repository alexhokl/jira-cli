# OrderOfCustomFieldOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The ID of the custom field option or cascading option to place the moved options after. Required if &#x60;position&#x60; isn&#39;t provided. | [optional] 
**CustomFieldOptionIds** | **[]string** | A list of IDs of custom field options to move. The order of the custom field option IDs in the list is the order they are given after the move. The list must contain custom field options or cascading options, but not both. | 
**Position** | Pointer to **string** | The position the custom field options should be moved to. Required if &#x60;after&#x60; isn&#39;t provided. | [optional] 

## Methods

### NewOrderOfCustomFieldOptions

`func NewOrderOfCustomFieldOptions(customFieldOptionIds []string, ) *OrderOfCustomFieldOptions`

NewOrderOfCustomFieldOptions instantiates a new OrderOfCustomFieldOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOfCustomFieldOptionsWithDefaults

`func NewOrderOfCustomFieldOptionsWithDefaults() *OrderOfCustomFieldOptions`

NewOrderOfCustomFieldOptionsWithDefaults instantiates a new OrderOfCustomFieldOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *OrderOfCustomFieldOptions) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *OrderOfCustomFieldOptions) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *OrderOfCustomFieldOptions) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *OrderOfCustomFieldOptions) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetCustomFieldOptionIds

`func (o *OrderOfCustomFieldOptions) GetCustomFieldOptionIds() []string`

GetCustomFieldOptionIds returns the CustomFieldOptionIds field if non-nil, zero value otherwise.

### GetCustomFieldOptionIdsOk

`func (o *OrderOfCustomFieldOptions) GetCustomFieldOptionIdsOk() (*[]string, bool)`

GetCustomFieldOptionIdsOk returns a tuple with the CustomFieldOptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldOptionIds

`func (o *OrderOfCustomFieldOptions) SetCustomFieldOptionIds(v []string)`

SetCustomFieldOptionIds sets CustomFieldOptionIds field to given value.


### GetPosition

`func (o *OrderOfCustomFieldOptions) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *OrderOfCustomFieldOptions) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *OrderOfCustomFieldOptions) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *OrderOfCustomFieldOptions) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


