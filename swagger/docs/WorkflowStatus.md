# WorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the issue status. | 
**Name** | **string** | The name of the status in the workflow. | 
**Properties** | Pointer to **map[string]interface{}** | Additional properties that modify the behavior of issues in this status. Supports the properties &#x60;jira.issue.editable&#x60; and &#x60;issueEditable&#x60; (deprecated) that indicate whether issues are editable. | [optional] 

## Methods

### NewWorkflowStatus

`func NewWorkflowStatus(id string, name string, ) *WorkflowStatus`

NewWorkflowStatus instantiates a new WorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusWithDefaults

`func NewWorkflowStatusWithDefaults() *WorkflowStatus`

NewWorkflowStatusWithDefaults instantiates a new WorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStatus) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *WorkflowStatus) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowStatus) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowStatus) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowStatus) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


