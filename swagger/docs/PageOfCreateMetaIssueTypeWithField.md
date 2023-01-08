# PageOfCreateMetaIssueTypeWithField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]FieldCreateMetadata**](FieldCreateMetadata.md) | The collection of FieldCreateMetaBeans. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of items to return per page. | [optional] [readonly] 
**Results** | Pointer to [**[]FieldCreateMetadata**](FieldCreateMetadata.md) |  | [optional] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] [readonly] 
**Total** | Pointer to **int64** | The total number of items in all pages. | [optional] [readonly] 

## Methods

### NewPageOfCreateMetaIssueTypeWithField

`func NewPageOfCreateMetaIssueTypeWithField() *PageOfCreateMetaIssueTypeWithField`

NewPageOfCreateMetaIssueTypeWithField instantiates a new PageOfCreateMetaIssueTypeWithField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfCreateMetaIssueTypeWithFieldWithDefaults

`func NewPageOfCreateMetaIssueTypeWithFieldWithDefaults() *PageOfCreateMetaIssueTypeWithField`

NewPageOfCreateMetaIssueTypeWithFieldWithDefaults instantiates a new PageOfCreateMetaIssueTypeWithField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *PageOfCreateMetaIssueTypeWithField) GetFields() []FieldCreateMetadata`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PageOfCreateMetaIssueTypeWithField) GetFieldsOk() (*[]FieldCreateMetadata, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PageOfCreateMetaIssueTypeWithField) SetFields(v []FieldCreateMetadata)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PageOfCreateMetaIssueTypeWithField) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageOfCreateMetaIssueTypeWithField) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfCreateMetaIssueTypeWithField) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfCreateMetaIssueTypeWithField) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfCreateMetaIssueTypeWithField) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetResults

`func (o *PageOfCreateMetaIssueTypeWithField) GetResults() []FieldCreateMetadata`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PageOfCreateMetaIssueTypeWithField) GetResultsOk() (*[]FieldCreateMetadata, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PageOfCreateMetaIssueTypeWithField) SetResults(v []FieldCreateMetadata)`

SetResults sets Results field to given value.

### HasResults

`func (o *PageOfCreateMetaIssueTypeWithField) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfCreateMetaIssueTypeWithField) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfCreateMetaIssueTypeWithField) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfCreateMetaIssueTypeWithField) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfCreateMetaIssueTypeWithField) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfCreateMetaIssueTypeWithField) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfCreateMetaIssueTypeWithField) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfCreateMetaIssueTypeWithField) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfCreateMetaIssueTypeWithField) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


