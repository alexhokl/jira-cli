# RedactionJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkRedactionResponse** | Pointer to [**BulkRedactionResponse**](BulkRedactionResponse.md) |  | [optional] 
**JobStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewRedactionJobStatusResponse

`func NewRedactionJobStatusResponse() *RedactionJobStatusResponse`

NewRedactionJobStatusResponse instantiates a new RedactionJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionJobStatusResponseWithDefaults

`func NewRedactionJobStatusResponseWithDefaults() *RedactionJobStatusResponse`

NewRedactionJobStatusResponseWithDefaults instantiates a new RedactionJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkRedactionResponse

`func (o *RedactionJobStatusResponse) GetBulkRedactionResponse() BulkRedactionResponse`

GetBulkRedactionResponse returns the BulkRedactionResponse field if non-nil, zero value otherwise.

### GetBulkRedactionResponseOk

`func (o *RedactionJobStatusResponse) GetBulkRedactionResponseOk() (*BulkRedactionResponse, bool)`

GetBulkRedactionResponseOk returns a tuple with the BulkRedactionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkRedactionResponse

`func (o *RedactionJobStatusResponse) SetBulkRedactionResponse(v BulkRedactionResponse)`

SetBulkRedactionResponse sets BulkRedactionResponse field to given value.

### HasBulkRedactionResponse

`func (o *RedactionJobStatusResponse) HasBulkRedactionResponse() bool`

HasBulkRedactionResponse returns a boolean if a field has been set.

### GetJobStatus

`func (o *RedactionJobStatusResponse) GetJobStatus() string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *RedactionJobStatusResponse) GetJobStatusOk() (*string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *RedactionJobStatusResponse) SetJobStatus(v string)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *RedactionJobStatusResponse) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


