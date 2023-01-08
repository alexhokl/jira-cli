# CreateCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldId** | **int64** | The custom field ID. | 
**Filter** | Pointer to **bool** | Allows filtering issues based on their values for the custom field. | [optional] 

## Methods

### NewCreateCustomFieldRequest

`func NewCreateCustomFieldRequest(customFieldId int64, ) *CreateCustomFieldRequest`

NewCreateCustomFieldRequest instantiates a new CreateCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomFieldRequestWithDefaults

`func NewCreateCustomFieldRequestWithDefaults() *CreateCustomFieldRequest`

NewCreateCustomFieldRequestWithDefaults instantiates a new CreateCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldId

`func (o *CreateCustomFieldRequest) GetCustomFieldId() int64`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *CreateCustomFieldRequest) GetCustomFieldIdOk() (*int64, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *CreateCustomFieldRequest) SetCustomFieldId(v int64)`

SetCustomFieldId sets CustomFieldId field to given value.


### GetFilter

`func (o *CreateCustomFieldRequest) GetFilter() bool`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateCustomFieldRequest) GetFilterOk() (*bool, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateCustomFieldRequest) SetFilter(v bool)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateCustomFieldRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


