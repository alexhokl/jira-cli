# MappingsByWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewWorkflowId** | **string** | The ID of the new workflow. | 
**OldWorkflowId** | **string** | The ID of the old workflow. | 
**StatusMappings** | [**[]WorkflowAssociationStatusMapping**](WorkflowAssociationStatusMapping.md) | The list of status mappings. | 

## Methods

### NewMappingsByWorkflow

`func NewMappingsByWorkflow(newWorkflowId string, oldWorkflowId string, statusMappings []WorkflowAssociationStatusMapping, ) *MappingsByWorkflow`

NewMappingsByWorkflow instantiates a new MappingsByWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingsByWorkflowWithDefaults

`func NewMappingsByWorkflowWithDefaults() *MappingsByWorkflow`

NewMappingsByWorkflowWithDefaults instantiates a new MappingsByWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewWorkflowId

`func (o *MappingsByWorkflow) GetNewWorkflowId() string`

GetNewWorkflowId returns the NewWorkflowId field if non-nil, zero value otherwise.

### GetNewWorkflowIdOk

`func (o *MappingsByWorkflow) GetNewWorkflowIdOk() (*string, bool)`

GetNewWorkflowIdOk returns a tuple with the NewWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWorkflowId

`func (o *MappingsByWorkflow) SetNewWorkflowId(v string)`

SetNewWorkflowId sets NewWorkflowId field to given value.


### GetOldWorkflowId

`func (o *MappingsByWorkflow) GetOldWorkflowId() string`

GetOldWorkflowId returns the OldWorkflowId field if non-nil, zero value otherwise.

### GetOldWorkflowIdOk

`func (o *MappingsByWorkflow) GetOldWorkflowIdOk() (*string, bool)`

GetOldWorkflowIdOk returns a tuple with the OldWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldWorkflowId

`func (o *MappingsByWorkflow) SetOldWorkflowId(v string)`

SetOldWorkflowId sets OldWorkflowId field to given value.


### GetStatusMappings

`func (o *MappingsByWorkflow) GetStatusMappings() []WorkflowAssociationStatusMapping`

GetStatusMappings returns the StatusMappings field if non-nil, zero value otherwise.

### GetStatusMappingsOk

`func (o *MappingsByWorkflow) GetStatusMappingsOk() (*[]WorkflowAssociationStatusMapping, bool)`

GetStatusMappingsOk returns a tuple with the StatusMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappings

`func (o *MappingsByWorkflow) SetStatusMappings(v []WorkflowAssociationStatusMapping)`

SetStatusMappings sets StatusMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


