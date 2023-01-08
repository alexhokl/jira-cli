# UserContextVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the user. | 
**Type** | **string** | Type of custom context variable. | 

## Methods

### NewUserContextVariable

`func NewUserContextVariable(accountId string, type_ string, ) *UserContextVariable`

NewUserContextVariable instantiates a new UserContextVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserContextVariableWithDefaults

`func NewUserContextVariableWithDefaults() *UserContextVariable`

NewUserContextVariableWithDefaults instantiates a new UserContextVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserContextVariable) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserContextVariable) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserContextVariable) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *UserContextVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserContextVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserContextVariable) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


