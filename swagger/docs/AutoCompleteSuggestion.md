# AutoCompleteSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of a suggested item. If &#x60;fieldValue&#x60; or &#x60;predicateValue&#x60; are provided, the matching text is highlighted with the HTML bold tag. | [optional] 
**Value** | Pointer to **string** | The value of a suggested item. | [optional] 

## Methods

### NewAutoCompleteSuggestion

`func NewAutoCompleteSuggestion() *AutoCompleteSuggestion`

NewAutoCompleteSuggestion instantiates a new AutoCompleteSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCompleteSuggestionWithDefaults

`func NewAutoCompleteSuggestionWithDefaults() *AutoCompleteSuggestion`

NewAutoCompleteSuggestionWithDefaults instantiates a new AutoCompleteSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutoCompleteSuggestion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoCompleteSuggestion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoCompleteSuggestion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoCompleteSuggestion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *AutoCompleteSuggestion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutoCompleteSuggestion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutoCompleteSuggestion) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AutoCompleteSuggestion) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


