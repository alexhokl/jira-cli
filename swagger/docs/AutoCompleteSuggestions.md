# AutoCompleteSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]AutoCompleteSuggestion**](AutoCompleteSuggestion.md) | The list of suggested item. | [optional] 

## Methods

### NewAutoCompleteSuggestions

`func NewAutoCompleteSuggestions() *AutoCompleteSuggestions`

NewAutoCompleteSuggestions instantiates a new AutoCompleteSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCompleteSuggestionsWithDefaults

`func NewAutoCompleteSuggestionsWithDefaults() *AutoCompleteSuggestions`

NewAutoCompleteSuggestionsWithDefaults instantiates a new AutoCompleteSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AutoCompleteSuggestions) GetResults() []AutoCompleteSuggestion`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AutoCompleteSuggestions) GetResultsOk() (*[]AutoCompleteSuggestion, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AutoCompleteSuggestions) SetResults(v []AutoCompleteSuggestion)`

SetResults sets Results field to given value.

### HasResults

`func (o *AutoCompleteSuggestions) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


