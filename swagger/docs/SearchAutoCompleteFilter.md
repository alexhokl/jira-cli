# SearchAutoCompleteFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeCollapsedFields** | Pointer to **bool** | Include collapsed fields for fields that have non-unique names. | [optional] [default to false]
**ProjectIds** | Pointer to **[]int64** | List of project IDs used to filter the visible field details returned. | [optional] 

## Methods

### NewSearchAutoCompleteFilter

`func NewSearchAutoCompleteFilter() *SearchAutoCompleteFilter`

NewSearchAutoCompleteFilter instantiates a new SearchAutoCompleteFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAutoCompleteFilterWithDefaults

`func NewSearchAutoCompleteFilterWithDefaults() *SearchAutoCompleteFilter`

NewSearchAutoCompleteFilterWithDefaults instantiates a new SearchAutoCompleteFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeCollapsedFields

`func (o *SearchAutoCompleteFilter) GetIncludeCollapsedFields() bool`

GetIncludeCollapsedFields returns the IncludeCollapsedFields field if non-nil, zero value otherwise.

### GetIncludeCollapsedFieldsOk

`func (o *SearchAutoCompleteFilter) GetIncludeCollapsedFieldsOk() (*bool, bool)`

GetIncludeCollapsedFieldsOk returns a tuple with the IncludeCollapsedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCollapsedFields

`func (o *SearchAutoCompleteFilter) SetIncludeCollapsedFields(v bool)`

SetIncludeCollapsedFields sets IncludeCollapsedFields field to given value.

### HasIncludeCollapsedFields

`func (o *SearchAutoCompleteFilter) HasIncludeCollapsedFields() bool`

HasIncludeCollapsedFields returns a boolean if a field has been set.

### GetProjectIds

`func (o *SearchAutoCompleteFilter) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SearchAutoCompleteFilter) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SearchAutoCompleteFilter) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SearchAutoCompleteFilter) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


