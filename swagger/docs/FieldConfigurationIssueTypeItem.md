# FieldConfigurationIssueTypeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldConfigurationId** | **string** | The ID of the field configuration. | 
**FieldConfigurationSchemeId** | **string** | The ID of the field configuration scheme. | 
**IssueTypeId** | **string** | The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration. | 

## Methods

### NewFieldConfigurationIssueTypeItem

`func NewFieldConfigurationIssueTypeItem(fieldConfigurationId string, fieldConfigurationSchemeId string, issueTypeId string, ) *FieldConfigurationIssueTypeItem`

NewFieldConfigurationIssueTypeItem instantiates a new FieldConfigurationIssueTypeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigurationIssueTypeItemWithDefaults

`func NewFieldConfigurationIssueTypeItemWithDefaults() *FieldConfigurationIssueTypeItem`

NewFieldConfigurationIssueTypeItemWithDefaults instantiates a new FieldConfigurationIssueTypeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldConfigurationId

`func (o *FieldConfigurationIssueTypeItem) GetFieldConfigurationId() string`

GetFieldConfigurationId returns the FieldConfigurationId field if non-nil, zero value otherwise.

### GetFieldConfigurationIdOk

`func (o *FieldConfigurationIssueTypeItem) GetFieldConfigurationIdOk() (*string, bool)`

GetFieldConfigurationIdOk returns a tuple with the FieldConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigurationId

`func (o *FieldConfigurationIssueTypeItem) SetFieldConfigurationId(v string)`

SetFieldConfigurationId sets FieldConfigurationId field to given value.


### GetFieldConfigurationSchemeId

`func (o *FieldConfigurationIssueTypeItem) GetFieldConfigurationSchemeId() string`

GetFieldConfigurationSchemeId returns the FieldConfigurationSchemeId field if non-nil, zero value otherwise.

### GetFieldConfigurationSchemeIdOk

`func (o *FieldConfigurationIssueTypeItem) GetFieldConfigurationSchemeIdOk() (*string, bool)`

GetFieldConfigurationSchemeIdOk returns a tuple with the FieldConfigurationSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigurationSchemeId

`func (o *FieldConfigurationIssueTypeItem) SetFieldConfigurationSchemeId(v string)`

SetFieldConfigurationSchemeId sets FieldConfigurationSchemeId field to given value.


### GetIssueTypeId

`func (o *FieldConfigurationIssueTypeItem) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *FieldConfigurationIssueTypeItem) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *FieldConfigurationIssueTypeItem) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


