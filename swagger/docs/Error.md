# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**IssueIdsOrKeys** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Error) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Error) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Error) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Error) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIssueIdsOrKeys

`func (o *Error) GetIssueIdsOrKeys() []string`

GetIssueIdsOrKeys returns the IssueIdsOrKeys field if non-nil, zero value otherwise.

### GetIssueIdsOrKeysOk

`func (o *Error) GetIssueIdsOrKeysOk() (*[]string, bool)`

GetIssueIdsOrKeysOk returns a tuple with the IssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdsOrKeys

`func (o *Error) SetIssueIdsOrKeys(v []string)`

SetIssueIdsOrKeys sets IssueIdsOrKeys field to given value.

### HasIssueIdsOrKeys

`func (o *Error) HasIssueIdsOrKeys() bool`

HasIssueIdsOrKeys returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


