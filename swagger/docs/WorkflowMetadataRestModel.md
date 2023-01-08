# WorkflowMetadataRestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the workflow. | 
**Id** | **string** | The ID of the workflow. | 
**Name** | **string** | The name of the workflow. | 
**Version** | [**DocumentVersion**](DocumentVersion.md) |  | 

## Methods

### NewWorkflowMetadataRestModel

`func NewWorkflowMetadataRestModel(description string, id string, name string, version DocumentVersion, ) *WorkflowMetadataRestModel`

NewWorkflowMetadataRestModel instantiates a new WorkflowMetadataRestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMetadataRestModelWithDefaults

`func NewWorkflowMetadataRestModelWithDefaults() *WorkflowMetadataRestModel`

NewWorkflowMetadataRestModelWithDefaults instantiates a new WorkflowMetadataRestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowMetadataRestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowMetadataRestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowMetadataRestModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *WorkflowMetadataRestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowMetadataRestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowMetadataRestModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowMetadataRestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowMetadataRestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowMetadataRestModel) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *WorkflowMetadataRestModel) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowMetadataRestModel) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowMetadataRestModel) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


