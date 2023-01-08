# Locale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** | The locale code. The Java the locale format is used: a two character language code (ISO 639), an underscore, and two letter country code (ISO 3166). For example, en\\_US represents a locale of English (United States). Required on create. | [optional] 

## Methods

### NewLocale

`func NewLocale() *Locale`

NewLocale instantiates a new Locale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocaleWithDefaults

`func NewLocaleWithDefaults() *Locale`

NewLocaleWithDefaults instantiates a new Locale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *Locale) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Locale) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Locale) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Locale) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


