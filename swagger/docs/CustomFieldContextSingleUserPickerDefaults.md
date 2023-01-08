# CustomFieldContextSingleUserPickerDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the default user. | 
**ContextId** | **string** | The ID of the context. | 
**Type** | **string** |  | 
**UserFilter** | [**UserFilter**](UserFilter.md) |  | 

## Methods

### NewCustomFieldContextSingleUserPickerDefaults

`func NewCustomFieldContextSingleUserPickerDefaults(accountId string, contextId string, type_ string, userFilter UserFilter, ) *CustomFieldContextSingleUserPickerDefaults`

NewCustomFieldContextSingleUserPickerDefaults instantiates a new CustomFieldContextSingleUserPickerDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextSingleUserPickerDefaultsWithDefaults

`func NewCustomFieldContextSingleUserPickerDefaultsWithDefaults() *CustomFieldContextSingleUserPickerDefaults`

NewCustomFieldContextSingleUserPickerDefaultsWithDefaults instantiates a new CustomFieldContextSingleUserPickerDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CustomFieldContextSingleUserPickerDefaults) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomFieldContextSingleUserPickerDefaults) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomFieldContextSingleUserPickerDefaults) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetContextId

`func (o *CustomFieldContextSingleUserPickerDefaults) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CustomFieldContextSingleUserPickerDefaults) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CustomFieldContextSingleUserPickerDefaults) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetType

`func (o *CustomFieldContextSingleUserPickerDefaults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextSingleUserPickerDefaults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextSingleUserPickerDefaults) SetType(v string)`

SetType sets Type field to given value.


### GetUserFilter

`func (o *CustomFieldContextSingleUserPickerDefaults) GetUserFilter() UserFilter`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *CustomFieldContextSingleUserPickerDefaults) GetUserFilterOk() (*UserFilter, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *CustomFieldContextSingleUserPickerDefaults) SetUserFilter(v UserFilter)`

SetUserFilter sets UserFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


