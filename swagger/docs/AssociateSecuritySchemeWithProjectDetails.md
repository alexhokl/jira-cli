# AssociateSecuritySchemeWithProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldToNewSecurityLevelMappings** | Pointer to [**[]OldToNewSecurityLevelMappingsBean**](OldToNewSecurityLevelMappingsBean.md) | The list of scheme levels which should be remapped to new levels of the issue security scheme. | [optional] 
**ProjectId** | **string** | The ID of the project. | 
**SchemeId** | **string** | The ID of the issue security scheme. Providing null will clear the association with the issue security scheme. | 

## Methods

### NewAssociateSecuritySchemeWithProjectDetails

`func NewAssociateSecuritySchemeWithProjectDetails(projectId string, schemeId string, ) *AssociateSecuritySchemeWithProjectDetails`

NewAssociateSecuritySchemeWithProjectDetails instantiates a new AssociateSecuritySchemeWithProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateSecuritySchemeWithProjectDetailsWithDefaults

`func NewAssociateSecuritySchemeWithProjectDetailsWithDefaults() *AssociateSecuritySchemeWithProjectDetails`

NewAssociateSecuritySchemeWithProjectDetailsWithDefaults instantiates a new AssociateSecuritySchemeWithProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldToNewSecurityLevelMappings

`func (o *AssociateSecuritySchemeWithProjectDetails) GetOldToNewSecurityLevelMappings() []OldToNewSecurityLevelMappingsBean`

GetOldToNewSecurityLevelMappings returns the OldToNewSecurityLevelMappings field if non-nil, zero value otherwise.

### GetOldToNewSecurityLevelMappingsOk

`func (o *AssociateSecuritySchemeWithProjectDetails) GetOldToNewSecurityLevelMappingsOk() (*[]OldToNewSecurityLevelMappingsBean, bool)`

GetOldToNewSecurityLevelMappingsOk returns a tuple with the OldToNewSecurityLevelMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldToNewSecurityLevelMappings

`func (o *AssociateSecuritySchemeWithProjectDetails) SetOldToNewSecurityLevelMappings(v []OldToNewSecurityLevelMappingsBean)`

SetOldToNewSecurityLevelMappings sets OldToNewSecurityLevelMappings field to given value.

### HasOldToNewSecurityLevelMappings

`func (o *AssociateSecuritySchemeWithProjectDetails) HasOldToNewSecurityLevelMappings() bool`

HasOldToNewSecurityLevelMappings returns a boolean if a field has been set.

### GetProjectId

`func (o *AssociateSecuritySchemeWithProjectDetails) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssociateSecuritySchemeWithProjectDetails) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssociateSecuritySchemeWithProjectDetails) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSchemeId

`func (o *AssociateSecuritySchemeWithProjectDetails) GetSchemeId() string`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *AssociateSecuritySchemeWithProjectDetails) GetSchemeIdOk() (*string, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *AssociateSecuritySchemeWithProjectDetails) SetSchemeId(v string)`

SetSchemeId sets SchemeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


