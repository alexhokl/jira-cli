# FieldConfigurationToIssueTypeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldConfigurationId** | **string** | The ID of the field configuration. | 
**IssueTypeId** | **string** | The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration. An issue type can be included only once in a request. | 

## Methods

### NewFieldConfigurationToIssueTypeMapping

`func NewFieldConfigurationToIssueTypeMapping(fieldConfigurationId string, issueTypeId string, ) *FieldConfigurationToIssueTypeMapping`

NewFieldConfigurationToIssueTypeMapping instantiates a new FieldConfigurationToIssueTypeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigurationToIssueTypeMappingWithDefaults

`func NewFieldConfigurationToIssueTypeMappingWithDefaults() *FieldConfigurationToIssueTypeMapping`

NewFieldConfigurationToIssueTypeMappingWithDefaults instantiates a new FieldConfigurationToIssueTypeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldConfigurationId

`func (o *FieldConfigurationToIssueTypeMapping) GetFieldConfigurationId() string`

GetFieldConfigurationId returns the FieldConfigurationId field if non-nil, zero value otherwise.

### GetFieldConfigurationIdOk

`func (o *FieldConfigurationToIssueTypeMapping) GetFieldConfigurationIdOk() (*string, bool)`

GetFieldConfigurationIdOk returns a tuple with the FieldConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigurationId

`func (o *FieldConfigurationToIssueTypeMapping) SetFieldConfigurationId(v string)`

SetFieldConfigurationId sets FieldConfigurationId field to given value.


### GetIssueTypeId

`func (o *FieldConfigurationToIssueTypeMapping) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *FieldConfigurationToIssueTypeMapping) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *FieldConfigurationToIssueTypeMapping) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


