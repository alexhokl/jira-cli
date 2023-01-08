# IssueContextVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The issue ID. | [optional] 
**Key** | Pointer to **string** | The issue key. | [optional] 
**Type** | **string** | Type of custom context variable. | 

## Methods

### NewIssueContextVariable

`func NewIssueContextVariable(type_ string, ) *IssueContextVariable`

NewIssueContextVariable instantiates a new IssueContextVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueContextVariableWithDefaults

`func NewIssueContextVariableWithDefaults() *IssueContextVariable`

NewIssueContextVariableWithDefaults instantiates a new IssueContextVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueContextVariable) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueContextVariable) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueContextVariable) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueContextVariable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *IssueContextVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueContextVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueContextVariable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueContextVariable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *IssueContextVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueContextVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueContextVariable) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


