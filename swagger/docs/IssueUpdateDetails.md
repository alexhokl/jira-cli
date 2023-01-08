# IssueUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **map[string]interface{}** | List of issue screen fields to update, specifying the sub-field to update and its value for each field. This field provides a straightforward option when setting a sub-field. When multiple sub-fields or other operations are required, use &#x60;update&#x60;. Fields included in here cannot be included in &#x60;update&#x60;. | [optional] 
**HistoryMetadata** | Pointer to [**HistoryMetadata**](HistoryMetadata.md) | Additional issue history details. | [optional] 
**Properties** | Pointer to [**[]EntityProperty**](EntityProperty.md) | Details of issue properties to be add or update. | [optional] 
**Transition** | Pointer to [**IssueTransition**](IssueTransition.md) | Details of a transition. Required when performing a transition, optional when creating or editing an issue. | [optional] 
**Update** | Pointer to [**map[string][]FieldUpdateOperation**](array.md) | A Map containing the field field name and a list of operations to perform on the issue screen field. Note that fields included in here cannot be included in &#x60;fields&#x60;. | [optional] 

## Methods

### NewIssueUpdateDetails

`func NewIssueUpdateDetails() *IssueUpdateDetails`

NewIssueUpdateDetails instantiates a new IssueUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueUpdateDetailsWithDefaults

`func NewIssueUpdateDetailsWithDefaults() *IssueUpdateDetails`

NewIssueUpdateDetailsWithDefaults instantiates a new IssueUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *IssueUpdateDetails) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IssueUpdateDetails) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IssueUpdateDetails) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *IssueUpdateDetails) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHistoryMetadata

`func (o *IssueUpdateDetails) GetHistoryMetadata() HistoryMetadata`

GetHistoryMetadata returns the HistoryMetadata field if non-nil, zero value otherwise.

### GetHistoryMetadataOk

`func (o *IssueUpdateDetails) GetHistoryMetadataOk() (*HistoryMetadata, bool)`

GetHistoryMetadataOk returns a tuple with the HistoryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryMetadata

`func (o *IssueUpdateDetails) SetHistoryMetadata(v HistoryMetadata)`

SetHistoryMetadata sets HistoryMetadata field to given value.

### HasHistoryMetadata

`func (o *IssueUpdateDetails) HasHistoryMetadata() bool`

HasHistoryMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *IssueUpdateDetails) GetProperties() []EntityProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueUpdateDetails) GetPropertiesOk() (*[]EntityProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueUpdateDetails) SetProperties(v []EntityProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueUpdateDetails) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTransition

`func (o *IssueUpdateDetails) GetTransition() IssueTransition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *IssueUpdateDetails) GetTransitionOk() (*IssueTransition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *IssueUpdateDetails) SetTransition(v IssueTransition)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *IssueUpdateDetails) HasTransition() bool`

HasTransition returns a boolean if a field has been set.

### GetUpdate

`func (o *IssueUpdateDetails) GetUpdate() map[string][]FieldUpdateOperation`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *IssueUpdateDetails) GetUpdateOk() (*map[string][]FieldUpdateOperation, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *IssueUpdateDetails) SetUpdate(v map[string][]FieldUpdateOperation)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *IssueUpdateDetails) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


