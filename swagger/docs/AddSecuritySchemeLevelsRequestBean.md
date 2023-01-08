# AddSecuritySchemeLevelsRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Levels** | Pointer to [**[]SecuritySchemeLevelBean**](SecuritySchemeLevelBean.md) | The list of scheme levels which should be added to the security scheme. | [optional] 

## Methods

### NewAddSecuritySchemeLevelsRequestBean

`func NewAddSecuritySchemeLevelsRequestBean() *AddSecuritySchemeLevelsRequestBean`

NewAddSecuritySchemeLevelsRequestBean instantiates a new AddSecuritySchemeLevelsRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecuritySchemeLevelsRequestBeanWithDefaults

`func NewAddSecuritySchemeLevelsRequestBeanWithDefaults() *AddSecuritySchemeLevelsRequestBean`

NewAddSecuritySchemeLevelsRequestBeanWithDefaults instantiates a new AddSecuritySchemeLevelsRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevels

`func (o *AddSecuritySchemeLevelsRequestBean) GetLevels() []SecuritySchemeLevelBean`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *AddSecuritySchemeLevelsRequestBean) GetLevelsOk() (*[]SecuritySchemeLevelBean, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *AddSecuritySchemeLevelsRequestBean) SetLevels(v []SecuritySchemeLevelBean)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *AddSecuritySchemeLevelsRequestBean) HasLevels() bool`

HasLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


