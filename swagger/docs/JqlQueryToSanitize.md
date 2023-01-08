# JqlQueryToSanitize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**Query** | **string** | The query to sanitize. | 

## Methods

### NewJqlQueryToSanitize

`func NewJqlQueryToSanitize(query string, ) *JqlQueryToSanitize`

NewJqlQueryToSanitize instantiates a new JqlQueryToSanitize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryToSanitizeWithDefaults

`func NewJqlQueryToSanitizeWithDefaults() *JqlQueryToSanitize`

NewJqlQueryToSanitizeWithDefaults instantiates a new JqlQueryToSanitize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *JqlQueryToSanitize) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *JqlQueryToSanitize) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *JqlQueryToSanitize) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *JqlQueryToSanitize) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *JqlQueryToSanitize) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *JqlQueryToSanitize) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetQuery

`func (o *JqlQueryToSanitize) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *JqlQueryToSanitize) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *JqlQueryToSanitize) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


