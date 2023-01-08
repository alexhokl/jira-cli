# SanitizedJqlQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | The account ID of the user for whom sanitization was performed. | [optional] 
**Errors** | Pointer to [**ErrorCollection**](ErrorCollection.md) | The list of errors. | [optional] 
**InitialQuery** | Pointer to **string** | The initial query. | [optional] 
**SanitizedQuery** | Pointer to **NullableString** | The sanitized query, if there were no errors. | [optional] 

## Methods

### NewSanitizedJqlQuery

`func NewSanitizedJqlQuery() *SanitizedJqlQuery`

NewSanitizedJqlQuery instantiates a new SanitizedJqlQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSanitizedJqlQueryWithDefaults

`func NewSanitizedJqlQueryWithDefaults() *SanitizedJqlQuery`

NewSanitizedJqlQueryWithDefaults instantiates a new SanitizedJqlQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SanitizedJqlQuery) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SanitizedJqlQuery) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SanitizedJqlQuery) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SanitizedJqlQuery) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *SanitizedJqlQuery) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *SanitizedJqlQuery) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetErrors

`func (o *SanitizedJqlQuery) GetErrors() ErrorCollection`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SanitizedJqlQuery) GetErrorsOk() (*ErrorCollection, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SanitizedJqlQuery) SetErrors(v ErrorCollection)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SanitizedJqlQuery) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInitialQuery

`func (o *SanitizedJqlQuery) GetInitialQuery() string`

GetInitialQuery returns the InitialQuery field if non-nil, zero value otherwise.

### GetInitialQueryOk

`func (o *SanitizedJqlQuery) GetInitialQueryOk() (*string, bool)`

GetInitialQueryOk returns a tuple with the InitialQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialQuery

`func (o *SanitizedJqlQuery) SetInitialQuery(v string)`

SetInitialQuery sets InitialQuery field to given value.

### HasInitialQuery

`func (o *SanitizedJqlQuery) HasInitialQuery() bool`

HasInitialQuery returns a boolean if a field has been set.

### GetSanitizedQuery

`func (o *SanitizedJqlQuery) GetSanitizedQuery() string`

GetSanitizedQuery returns the SanitizedQuery field if non-nil, zero value otherwise.

### GetSanitizedQueryOk

`func (o *SanitizedJqlQuery) GetSanitizedQueryOk() (*string, bool)`

GetSanitizedQueryOk returns a tuple with the SanitizedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitizedQuery

`func (o *SanitizedJqlQuery) SetSanitizedQuery(v string)`

SetSanitizedQuery sets SanitizedQuery field to given value.

### HasSanitizedQuery

`func (o *SanitizedJqlQuery) HasSanitizedQuery() bool`

HasSanitizedQuery returns a boolean if a field has been set.

### SetSanitizedQueryNil

`func (o *SanitizedJqlQuery) SetSanitizedQueryNil(b bool)`

 SetSanitizedQueryNil sets the value for SanitizedQuery to be an explicit nil

### UnsetSanitizedQuery
`func (o *SanitizedJqlQuery) UnsetSanitizedQuery()`

UnsetSanitizedQuery ensures that no value is present for SanitizedQuery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


