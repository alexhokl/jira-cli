# ParsedJqlQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** | The list of syntax or validation errors. | [optional] 
**Query** | **string** | The JQL query that was parsed and validated. | 
**Structure** | Pointer to [**JqlQuery**](JqlQuery.md) | The syntax tree of the query. Empty if the query was invalid. | [optional] 
**Warnings** | Pointer to **[]string** | The list of warning messages | [optional] 

## Methods

### NewParsedJqlQuery

`func NewParsedJqlQuery(query string, ) *ParsedJqlQuery`

NewParsedJqlQuery instantiates a new ParsedJqlQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedJqlQueryWithDefaults

`func NewParsedJqlQueryWithDefaults() *ParsedJqlQuery`

NewParsedJqlQueryWithDefaults instantiates a new ParsedJqlQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ParsedJqlQuery) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ParsedJqlQuery) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ParsedJqlQuery) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ParsedJqlQuery) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQuery

`func (o *ParsedJqlQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ParsedJqlQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ParsedJqlQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetStructure

`func (o *ParsedJqlQuery) GetStructure() JqlQuery`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *ParsedJqlQuery) GetStructureOk() (*JqlQuery, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *ParsedJqlQuery) SetStructure(v JqlQuery)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *ParsedJqlQuery) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetWarnings

`func (o *ParsedJqlQuery) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ParsedJqlQuery) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ParsedJqlQuery) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ParsedJqlQuery) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


