# ConvertedJQLQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueriesWithUnknownUsers** | Pointer to [**[]JQLQueryWithUnknownUsers**](JQLQueryWithUnknownUsers.md) | List of queries containing user information that could not be mapped to an existing user | [optional] 
**QueryStrings** | Pointer to **[]string** | The list of converted query strings with account IDs in place of user identifiers. | [optional] 

## Methods

### NewConvertedJQLQueries

`func NewConvertedJQLQueries() *ConvertedJQLQueries`

NewConvertedJQLQueries instantiates a new ConvertedJQLQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertedJQLQueriesWithDefaults

`func NewConvertedJQLQueriesWithDefaults() *ConvertedJQLQueries`

NewConvertedJQLQueriesWithDefaults instantiates a new ConvertedJQLQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueriesWithUnknownUsers

`func (o *ConvertedJQLQueries) GetQueriesWithUnknownUsers() []JQLQueryWithUnknownUsers`

GetQueriesWithUnknownUsers returns the QueriesWithUnknownUsers field if non-nil, zero value otherwise.

### GetQueriesWithUnknownUsersOk

`func (o *ConvertedJQLQueries) GetQueriesWithUnknownUsersOk() (*[]JQLQueryWithUnknownUsers, bool)`

GetQueriesWithUnknownUsersOk returns a tuple with the QueriesWithUnknownUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesWithUnknownUsers

`func (o *ConvertedJQLQueries) SetQueriesWithUnknownUsers(v []JQLQueryWithUnknownUsers)`

SetQueriesWithUnknownUsers sets QueriesWithUnknownUsers field to given value.

### HasQueriesWithUnknownUsers

`func (o *ConvertedJQLQueries) HasQueriesWithUnknownUsers() bool`

HasQueriesWithUnknownUsers returns a boolean if a field has been set.

### GetQueryStrings

`func (o *ConvertedJQLQueries) GetQueryStrings() []string`

GetQueryStrings returns the QueryStrings field if non-nil, zero value otherwise.

### GetQueryStringsOk

`func (o *ConvertedJQLQueries) GetQueryStringsOk() (*[]string, bool)`

GetQueryStringsOk returns a tuple with the QueryStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStrings

`func (o *ConvertedJQLQueries) SetQueryStrings(v []string)`

SetQueryStrings sets QueryStrings field to given value.

### HasQueryStrings

`func (o *ConvertedJQLQueries) HasQueryStrings() bool`

HasQueryStrings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


