# GetCustomFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldId** | **int64** | The custom field ID. | 
**Filter** | Pointer to **bool** | Allows filtering issues based on their values for the custom field. | [optional] 

## Methods

### NewGetCustomFieldResponse

`func NewGetCustomFieldResponse(customFieldId int64, ) *GetCustomFieldResponse`

NewGetCustomFieldResponse instantiates a new GetCustomFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomFieldResponseWithDefaults

`func NewGetCustomFieldResponseWithDefaults() *GetCustomFieldResponse`

NewGetCustomFieldResponseWithDefaults instantiates a new GetCustomFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldId

`func (o *GetCustomFieldResponse) GetCustomFieldId() int64`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *GetCustomFieldResponse) GetCustomFieldIdOk() (*int64, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *GetCustomFieldResponse) SetCustomFieldId(v int64)`

SetCustomFieldId sets CustomFieldId field to given value.


### GetFilter

`func (o *GetCustomFieldResponse) GetFilter() bool`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetCustomFieldResponse) GetFilterOk() (*bool, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetCustomFieldResponse) SetFilter(v bool)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetCustomFieldResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


