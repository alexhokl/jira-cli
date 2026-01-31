# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**IssueId** | Pointer to **int64** |  | [optional] 
**IssueKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Entry) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Entry) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Entry) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Entry) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetIssueId

`func (o *Entry) GetIssueId() int64`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Entry) GetIssueIdOk() (*int64, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Entry) SetIssueId(v int64)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *Entry) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetIssueKey

`func (o *Entry) GetIssueKey() string`

GetIssueKey returns the IssueKey field if non-nil, zero value otherwise.

### GetIssueKeyOk

`func (o *Entry) GetIssueKeyOk() (*string, bool)`

GetIssueKeyOk returns a tuple with the IssueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKey

`func (o *Entry) SetIssueKey(v string)`

SetIssueKey sets IssueKey field to given value.

### HasIssueKey

`func (o *Entry) HasIssueKey() bool`

HasIssueKey returns a boolean if a field has been set.

### GetStatus

`func (o *Entry) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Entry) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Entry) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Entry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


