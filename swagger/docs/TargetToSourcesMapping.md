# TargetToSourcesMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InferClassificationDefaults** | **bool** | If &#x60;true&#x60;, when issues are moved into this target group, they will adopt the target project&#39;s default classification, if they don&#39;t have a classification already. If they do have a classification, it will be kept the same even after the move. Leave &#x60;targetClassification&#x60; empty when using this.  If &#x60;false&#x60;, you must provide a &#x60;targetClassification&#x60; mapping for each classification associated with the selected issues.  [Benefit from data classification](https://support.atlassian.com/security-and-access-policies/docs/what-is-data-classification/) | 
**InferFieldDefaults** | **bool** | If &#x60;true&#x60;, values from the source issues will be retained for the mandatory fields in the field configuration of the destination project. The &#x60;targetMandatoryFields&#x60; property shouldn&#39;t be defined.  If &#x60;false&#x60;, the user is required to set values for mandatory fields present in the field configuration of the destination project. Provide input by defining the &#x60;targetMandatoryFields&#x60; property | 
**InferStatusDefaults** | **bool** | If &#x60;true&#x60;, the statuses of issues being moved in this target group that are not present in the target workflow will be changed to the default status of the target workflow (see below). Leave &#x60;targetStatus&#x60; empty when using this.  If &#x60;false&#x60;, you must provide a &#x60;targetStatus&#x60; for each status not present in the target workflow.  The default status in a workflow is referred to as the \&quot;initial status\&quot;. Each workflow has its own unique initial status. When an issue is created, it is automatically assigned to this initial status. Read more about configuring initial statuses: [Configure the initial status | Atlassian Support.](https://support.atlassian.com/jira-cloud-administration/docs/configure-the-initial-status/) | 
**InferSubtaskTypeDefault** | **bool** | When an issue is moved, its subtasks (if there are any) need to be moved with it. &#x60;inferSubtaskTypeDefault&#x60; helps with moving the subtasks by picking a random subtask type in the target project.  If &#x60;true&#x60;, subtasks will automatically move to the same project as their parent.  When they move:   *  Their &#x60;issueType&#x60; will be set to the default for subtasks in the target project.  *  Values for mandatory fields will be retained from the source issues  *  Specifying separate mapping for implicit subtasks won’t be allowed.  If &#x60;false&#x60;, you must manually move the subtasks. They will retain the parent which they had in the current project after being moved. | 
**IssueIdsOrKeys** | Pointer to **[]string** | List of issue IDs or keys to be moved. | [optional] 
**TargetClassification** | Pointer to [**[]TargetClassification**](TargetClassification.md) | List of the objects containing classifications in the source issues and their new values which need to be set during the bulk move operation.  It is mandatory to provide source classification to target classification mapping when the source classification is invalid for the target project and issue type.   *  **You should only define this property when &#x60;inferClassificationDefaults&#x60; is &#x60;false&#x60;.**  *  **In order to provide mapping for issues which don&#39;t have a classification, use &#x60;\&quot;-1\&quot;&#x60;.** | [optional] 
**TargetMandatoryFields** | Pointer to [**[]TargetMandatoryFields**](TargetMandatoryFields.md) | List of objects containing mandatory fields in the target field configuration and new values that need to be set during the bulk move operation.  The new values will only be applied if the field is mandatory in the target project and at least one issue from the source has that field empty, or if the field context is different in the target project (e.g. project-scoped version fields).  **You should only define this property when &#x60;inferFieldDefaults&#x60; is &#x60;false&#x60;.** | [optional] 
**TargetStatus** | Pointer to [**[]TargetStatus**](TargetStatus.md) | List of the objects containing statuses in the source workflow and their new values which need to be set during the bulk move operation.  The new values will only be applied if the source status is invalid for the target project and issue type.  It is mandatory to provide source status to target status mapping when the source status is invalid for the target project and issue type.  **You should only define this property when &#x60;inferStatusDefaults&#x60; is &#x60;false&#x60;.** | [optional] 

## Methods

### NewTargetToSourcesMapping

`func NewTargetToSourcesMapping(inferClassificationDefaults bool, inferFieldDefaults bool, inferStatusDefaults bool, inferSubtaskTypeDefault bool, ) *TargetToSourcesMapping`

NewTargetToSourcesMapping instantiates a new TargetToSourcesMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetToSourcesMappingWithDefaults

`func NewTargetToSourcesMappingWithDefaults() *TargetToSourcesMapping`

NewTargetToSourcesMappingWithDefaults instantiates a new TargetToSourcesMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInferClassificationDefaults

`func (o *TargetToSourcesMapping) GetInferClassificationDefaults() bool`

GetInferClassificationDefaults returns the InferClassificationDefaults field if non-nil, zero value otherwise.

### GetInferClassificationDefaultsOk

`func (o *TargetToSourcesMapping) GetInferClassificationDefaultsOk() (*bool, bool)`

GetInferClassificationDefaultsOk returns a tuple with the InferClassificationDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferClassificationDefaults

`func (o *TargetToSourcesMapping) SetInferClassificationDefaults(v bool)`

SetInferClassificationDefaults sets InferClassificationDefaults field to given value.


### GetInferFieldDefaults

`func (o *TargetToSourcesMapping) GetInferFieldDefaults() bool`

GetInferFieldDefaults returns the InferFieldDefaults field if non-nil, zero value otherwise.

### GetInferFieldDefaultsOk

`func (o *TargetToSourcesMapping) GetInferFieldDefaultsOk() (*bool, bool)`

GetInferFieldDefaultsOk returns a tuple with the InferFieldDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferFieldDefaults

`func (o *TargetToSourcesMapping) SetInferFieldDefaults(v bool)`

SetInferFieldDefaults sets InferFieldDefaults field to given value.


### GetInferStatusDefaults

`func (o *TargetToSourcesMapping) GetInferStatusDefaults() bool`

GetInferStatusDefaults returns the InferStatusDefaults field if non-nil, zero value otherwise.

### GetInferStatusDefaultsOk

`func (o *TargetToSourcesMapping) GetInferStatusDefaultsOk() (*bool, bool)`

GetInferStatusDefaultsOk returns a tuple with the InferStatusDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferStatusDefaults

`func (o *TargetToSourcesMapping) SetInferStatusDefaults(v bool)`

SetInferStatusDefaults sets InferStatusDefaults field to given value.


### GetInferSubtaskTypeDefault

`func (o *TargetToSourcesMapping) GetInferSubtaskTypeDefault() bool`

GetInferSubtaskTypeDefault returns the InferSubtaskTypeDefault field if non-nil, zero value otherwise.

### GetInferSubtaskTypeDefaultOk

`func (o *TargetToSourcesMapping) GetInferSubtaskTypeDefaultOk() (*bool, bool)`

GetInferSubtaskTypeDefaultOk returns a tuple with the InferSubtaskTypeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferSubtaskTypeDefault

`func (o *TargetToSourcesMapping) SetInferSubtaskTypeDefault(v bool)`

SetInferSubtaskTypeDefault sets InferSubtaskTypeDefault field to given value.


### GetIssueIdsOrKeys

`func (o *TargetToSourcesMapping) GetIssueIdsOrKeys() []string`

GetIssueIdsOrKeys returns the IssueIdsOrKeys field if non-nil, zero value otherwise.

### GetIssueIdsOrKeysOk

`func (o *TargetToSourcesMapping) GetIssueIdsOrKeysOk() (*[]string, bool)`

GetIssueIdsOrKeysOk returns a tuple with the IssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdsOrKeys

`func (o *TargetToSourcesMapping) SetIssueIdsOrKeys(v []string)`

SetIssueIdsOrKeys sets IssueIdsOrKeys field to given value.

### HasIssueIdsOrKeys

`func (o *TargetToSourcesMapping) HasIssueIdsOrKeys() bool`

HasIssueIdsOrKeys returns a boolean if a field has been set.

### GetTargetClassification

`func (o *TargetToSourcesMapping) GetTargetClassification() []TargetClassification`

GetTargetClassification returns the TargetClassification field if non-nil, zero value otherwise.

### GetTargetClassificationOk

`func (o *TargetToSourcesMapping) GetTargetClassificationOk() (*[]TargetClassification, bool)`

GetTargetClassificationOk returns a tuple with the TargetClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClassification

`func (o *TargetToSourcesMapping) SetTargetClassification(v []TargetClassification)`

SetTargetClassification sets TargetClassification field to given value.

### HasTargetClassification

`func (o *TargetToSourcesMapping) HasTargetClassification() bool`

HasTargetClassification returns a boolean if a field has been set.

### SetTargetClassificationNil

`func (o *TargetToSourcesMapping) SetTargetClassificationNil(b bool)`

 SetTargetClassificationNil sets the value for TargetClassification to be an explicit nil

### UnsetTargetClassification
`func (o *TargetToSourcesMapping) UnsetTargetClassification()`

UnsetTargetClassification ensures that no value is present for TargetClassification, not even an explicit nil
### GetTargetMandatoryFields

`func (o *TargetToSourcesMapping) GetTargetMandatoryFields() []TargetMandatoryFields`

GetTargetMandatoryFields returns the TargetMandatoryFields field if non-nil, zero value otherwise.

### GetTargetMandatoryFieldsOk

`func (o *TargetToSourcesMapping) GetTargetMandatoryFieldsOk() (*[]TargetMandatoryFields, bool)`

GetTargetMandatoryFieldsOk returns a tuple with the TargetMandatoryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMandatoryFields

`func (o *TargetToSourcesMapping) SetTargetMandatoryFields(v []TargetMandatoryFields)`

SetTargetMandatoryFields sets TargetMandatoryFields field to given value.

### HasTargetMandatoryFields

`func (o *TargetToSourcesMapping) HasTargetMandatoryFields() bool`

HasTargetMandatoryFields returns a boolean if a field has been set.

### SetTargetMandatoryFieldsNil

`func (o *TargetToSourcesMapping) SetTargetMandatoryFieldsNil(b bool)`

 SetTargetMandatoryFieldsNil sets the value for TargetMandatoryFields to be an explicit nil

### UnsetTargetMandatoryFields
`func (o *TargetToSourcesMapping) UnsetTargetMandatoryFields()`

UnsetTargetMandatoryFields ensures that no value is present for TargetMandatoryFields, not even an explicit nil
### GetTargetStatus

`func (o *TargetToSourcesMapping) GetTargetStatus() []TargetStatus`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *TargetToSourcesMapping) GetTargetStatusOk() (*[]TargetStatus, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *TargetToSourcesMapping) SetTargetStatus(v []TargetStatus)`

SetTargetStatus sets TargetStatus field to given value.

### HasTargetStatus

`func (o *TargetToSourcesMapping) HasTargetStatus() bool`

HasTargetStatus returns a boolean if a field has been set.

### SetTargetStatusNil

`func (o *TargetToSourcesMapping) SetTargetStatusNil(b bool)`

 SetTargetStatusNil sets the value for TargetStatus to be an explicit nil

### UnsetTargetStatus
`func (o *TargetToSourcesMapping) UnsetTargetStatus()`

UnsetTargetStatus ensures that no value is present for TargetStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


