# JQLQueryWithUnknownUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertedQuery** | Pointer to **string** | The converted query, with accountIDs instead of user identifiers, or &#39;unknown&#39; for users that could not be found | [optional] 
**OriginalQuery** | Pointer to **string** | The original query, for reference | [optional] 

## Methods

### NewJQLQueryWithUnknownUsers

`func NewJQLQueryWithUnknownUsers() *JQLQueryWithUnknownUsers`

NewJQLQueryWithUnknownUsers instantiates a new JQLQueryWithUnknownUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJQLQueryWithUnknownUsersWithDefaults

`func NewJQLQueryWithUnknownUsersWithDefaults() *JQLQueryWithUnknownUsers`

NewJQLQueryWithUnknownUsersWithDefaults instantiates a new JQLQueryWithUnknownUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertedQuery

`func (o *JQLQueryWithUnknownUsers) GetConvertedQuery() string`

GetConvertedQuery returns the ConvertedQuery field if non-nil, zero value otherwise.

### GetConvertedQueryOk

`func (o *JQLQueryWithUnknownUsers) GetConvertedQueryOk() (*string, bool)`

GetConvertedQueryOk returns a tuple with the ConvertedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedQuery

`func (o *JQLQueryWithUnknownUsers) SetConvertedQuery(v string)`

SetConvertedQuery sets ConvertedQuery field to given value.

### HasConvertedQuery

`func (o *JQLQueryWithUnknownUsers) HasConvertedQuery() bool`

HasConvertedQuery returns a boolean if a field has been set.

### GetOriginalQuery

`func (o *JQLQueryWithUnknownUsers) GetOriginalQuery() string`

GetOriginalQuery returns the OriginalQuery field if non-nil, zero value otherwise.

### GetOriginalQueryOk

`func (o *JQLQueryWithUnknownUsers) GetOriginalQueryOk() (*string, bool)`

GetOriginalQueryOk returns a tuple with the OriginalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQuery

`func (o *JQLQueryWithUnknownUsers) SetOriginalQuery(v string)`

SetOriginalQuery sets OriginalQuery field to given value.

### HasOriginalQuery

`func (o *JQLQueryWithUnknownUsers) HasOriginalQuery() bool`

HasOriginalQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


