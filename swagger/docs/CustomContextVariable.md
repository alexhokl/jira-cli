# CustomContextVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of custom context variable. | 
**AccountId** | **string** | The account ID of the user. | 
**Id** | Pointer to **int64** | The issue ID. | [optional] 
**Key** | Pointer to **string** | The issue key. | [optional] 
**Value** | Pointer to **map[string]interface{}** | A JSON object containing custom content. | [optional] 

## Methods

### NewCustomContextVariable

`func NewCustomContextVariable(type_ string, accountId string, ) *CustomContextVariable`

NewCustomContextVariable instantiates a new CustomContextVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomContextVariableWithDefaults

`func NewCustomContextVariableWithDefaults() *CustomContextVariable`

NewCustomContextVariableWithDefaults instantiates a new CustomContextVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomContextVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomContextVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomContextVariable) SetType(v string)`

SetType sets Type field to given value.


### GetAccountId

`func (o *CustomContextVariable) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomContextVariable) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomContextVariable) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetId

`func (o *CustomContextVariable) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomContextVariable) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomContextVariable) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomContextVariable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CustomContextVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomContextVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomContextVariable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomContextVariable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *CustomContextVariable) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomContextVariable) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomContextVariable) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomContextVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


