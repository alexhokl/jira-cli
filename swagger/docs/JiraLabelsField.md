# JiraLabelsField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkEditMultiSelectFieldOption** | **string** |  | 
**FieldId** | **string** |  | 
**LabelProperties** | Pointer to [**[]JiraLabelPropertiesInputJackson1**](JiraLabelPropertiesInputJackson1.md) |  | [optional] 
**Labels** | [**[]JiraLabelsInput**](JiraLabelsInput.md) |  | 

## Methods

### NewJiraLabelsField

`func NewJiraLabelsField(bulkEditMultiSelectFieldOption string, fieldId string, labels []JiraLabelsInput, ) *JiraLabelsField`

NewJiraLabelsField instantiates a new JiraLabelsField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraLabelsFieldWithDefaults

`func NewJiraLabelsFieldWithDefaults() *JiraLabelsField`

NewJiraLabelsFieldWithDefaults instantiates a new JiraLabelsField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkEditMultiSelectFieldOption

`func (o *JiraLabelsField) GetBulkEditMultiSelectFieldOption() string`

GetBulkEditMultiSelectFieldOption returns the BulkEditMultiSelectFieldOption field if non-nil, zero value otherwise.

### GetBulkEditMultiSelectFieldOptionOk

`func (o *JiraLabelsField) GetBulkEditMultiSelectFieldOptionOk() (*string, bool)`

GetBulkEditMultiSelectFieldOptionOk returns a tuple with the BulkEditMultiSelectFieldOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkEditMultiSelectFieldOption

`func (o *JiraLabelsField) SetBulkEditMultiSelectFieldOption(v string)`

SetBulkEditMultiSelectFieldOption sets BulkEditMultiSelectFieldOption field to given value.


### GetFieldId

`func (o *JiraLabelsField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraLabelsField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraLabelsField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetLabelProperties

`func (o *JiraLabelsField) GetLabelProperties() []JiraLabelPropertiesInputJackson1`

GetLabelProperties returns the LabelProperties field if non-nil, zero value otherwise.

### GetLabelPropertiesOk

`func (o *JiraLabelsField) GetLabelPropertiesOk() (*[]JiraLabelPropertiesInputJackson1, bool)`

GetLabelPropertiesOk returns a tuple with the LabelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelProperties

`func (o *JiraLabelsField) SetLabelProperties(v []JiraLabelPropertiesInputJackson1)`

SetLabelProperties sets LabelProperties field to given value.

### HasLabelProperties

`func (o *JiraLabelsField) HasLabelProperties() bool`

HasLabelProperties returns a boolean if a field has been set.

### GetLabels

`func (o *JiraLabelsField) GetLabels() []JiraLabelsInput`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *JiraLabelsField) GetLabelsOk() (*[]JiraLabelsInput, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *JiraLabelsField) SetLabels(v []JiraLabelsInput)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


