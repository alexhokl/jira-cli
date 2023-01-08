# OldToNewSecurityLevelMappingsBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewLevelId** | **string** | The new issue security level ID. Providing null will clear the assigned old level from issues. | 
**OldLevelId** | **string** | The old issue security level ID. Providing null will remap all issues without any assigned levels. | 

## Methods

### NewOldToNewSecurityLevelMappingsBean

`func NewOldToNewSecurityLevelMappingsBean(newLevelId string, oldLevelId string, ) *OldToNewSecurityLevelMappingsBean`

NewOldToNewSecurityLevelMappingsBean instantiates a new OldToNewSecurityLevelMappingsBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOldToNewSecurityLevelMappingsBeanWithDefaults

`func NewOldToNewSecurityLevelMappingsBeanWithDefaults() *OldToNewSecurityLevelMappingsBean`

NewOldToNewSecurityLevelMappingsBeanWithDefaults instantiates a new OldToNewSecurityLevelMappingsBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewLevelId

`func (o *OldToNewSecurityLevelMappingsBean) GetNewLevelId() string`

GetNewLevelId returns the NewLevelId field if non-nil, zero value otherwise.

### GetNewLevelIdOk

`func (o *OldToNewSecurityLevelMappingsBean) GetNewLevelIdOk() (*string, bool)`

GetNewLevelIdOk returns a tuple with the NewLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLevelId

`func (o *OldToNewSecurityLevelMappingsBean) SetNewLevelId(v string)`

SetNewLevelId sets NewLevelId field to given value.


### GetOldLevelId

`func (o *OldToNewSecurityLevelMappingsBean) GetOldLevelId() string`

GetOldLevelId returns the OldLevelId field if non-nil, zero value otherwise.

### GetOldLevelIdOk

`func (o *OldToNewSecurityLevelMappingsBean) GetOldLevelIdOk() (*string, bool)`

GetOldLevelIdOk returns a tuple with the OldLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldLevelId

`func (o *OldToNewSecurityLevelMappingsBean) SetOldLevelId(v string)`

SetOldLevelId sets OldLevelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


