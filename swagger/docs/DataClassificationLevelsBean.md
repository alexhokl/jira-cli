# DataClassificationLevelsBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifications** | Pointer to [**[]DataClassificationTagBean**](DataClassificationTagBean.md) | The data classifications. | [optional] 

## Methods

### NewDataClassificationLevelsBean

`func NewDataClassificationLevelsBean() *DataClassificationLevelsBean`

NewDataClassificationLevelsBean instantiates a new DataClassificationLevelsBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataClassificationLevelsBeanWithDefaults

`func NewDataClassificationLevelsBeanWithDefaults() *DataClassificationLevelsBean`

NewDataClassificationLevelsBeanWithDefaults instantiates a new DataClassificationLevelsBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifications

`func (o *DataClassificationLevelsBean) GetClassifications() []DataClassificationTagBean`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *DataClassificationLevelsBean) GetClassificationsOk() (*[]DataClassificationTagBean, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *DataClassificationLevelsBean) SetClassifications(v []DataClassificationTagBean)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *DataClassificationLevelsBean) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


