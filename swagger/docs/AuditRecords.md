# AuditRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The requested or default limit on the number of audit items to be returned. | [optional] [readonly] 
**Offset** | Pointer to **int32** | The number of audit items skipped before the first item in this list. | [optional] [readonly] 
**Records** | Pointer to [**[]AuditRecordBean**](AuditRecordBean.md) | The list of audit items. | [optional] [readonly] 
**Total** | Pointer to **int64** | The total number of audit items returned. | [optional] [readonly] 

## Methods

### NewAuditRecords

`func NewAuditRecords() *AuditRecords`

NewAuditRecords instantiates a new AuditRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditRecordsWithDefaults

`func NewAuditRecordsWithDefaults() *AuditRecords`

NewAuditRecordsWithDefaults instantiates a new AuditRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AuditRecords) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AuditRecords) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AuditRecords) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AuditRecords) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *AuditRecords) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AuditRecords) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AuditRecords) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AuditRecords) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRecords

`func (o *AuditRecords) GetRecords() []AuditRecordBean`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *AuditRecords) GetRecordsOk() (*[]AuditRecordBean, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *AuditRecords) SetRecords(v []AuditRecordBean)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *AuditRecords) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *AuditRecords) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditRecords) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditRecords) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuditRecords) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


