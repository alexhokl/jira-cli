# ChangedValueBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedFrom** | Pointer to **string** | The value of the field before the change. | [optional] [readonly] 
**ChangedTo** | Pointer to **string** | The value of the field after the change. | [optional] [readonly] 
**FieldName** | Pointer to **string** | The name of the field changed. | [optional] [readonly] 

## Methods

### NewChangedValueBean

`func NewChangedValueBean() *ChangedValueBean`

NewChangedValueBean instantiates a new ChangedValueBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangedValueBeanWithDefaults

`func NewChangedValueBeanWithDefaults() *ChangedValueBean`

NewChangedValueBeanWithDefaults instantiates a new ChangedValueBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedFrom

`func (o *ChangedValueBean) GetChangedFrom() string`

GetChangedFrom returns the ChangedFrom field if non-nil, zero value otherwise.

### GetChangedFromOk

`func (o *ChangedValueBean) GetChangedFromOk() (*string, bool)`

GetChangedFromOk returns a tuple with the ChangedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedFrom

`func (o *ChangedValueBean) SetChangedFrom(v string)`

SetChangedFrom sets ChangedFrom field to given value.

### HasChangedFrom

`func (o *ChangedValueBean) HasChangedFrom() bool`

HasChangedFrom returns a boolean if a field has been set.

### GetChangedTo

`func (o *ChangedValueBean) GetChangedTo() string`

GetChangedTo returns the ChangedTo field if non-nil, zero value otherwise.

### GetChangedToOk

`func (o *ChangedValueBean) GetChangedToOk() (*string, bool)`

GetChangedToOk returns a tuple with the ChangedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedTo

`func (o *ChangedValueBean) SetChangedTo(v string)`

SetChangedTo sets ChangedTo field to given value.

### HasChangedTo

`func (o *ChangedValueBean) HasChangedTo() bool`

HasChangedTo returns a boolean if a field has been set.

### GetFieldName

`func (o *ChangedValueBean) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ChangedValueBean) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ChangedValueBean) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ChangedValueBean) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


