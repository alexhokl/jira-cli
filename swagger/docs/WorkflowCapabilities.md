# WorkflowCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectRules** | Pointer to [**[]AvailableWorkflowConnectRule**](AvailableWorkflowConnectRule.md) | The Connect provided ecosystem rules available. | [optional] 
**EditorScope** | Pointer to **string** | The scope of the workflow capabilities. &#x60;GLOBAL&#x60; for company-managed projects and &#x60;PROJECT&#x60; for team-managed projects. | [optional] 
**ForgeRules** | Pointer to [**[]AvailableWorkflowForgeRule**](AvailableWorkflowForgeRule.md) | The Forge provided ecosystem rules available. | [optional] 
**ProjectTypes** | Pointer to **[]string** | The types of projects that this capability set is available for. | [optional] 
**SystemRules** | Pointer to [**[]AvailableWorkflowSystemRule**](AvailableWorkflowSystemRule.md) | The Atlassian provided system rules available. | [optional] 
**TriggerRules** | Pointer to [**[]AvailableWorkflowTriggers**](AvailableWorkflowTriggers.md) | The trigger rules available. | [optional] 

## Methods

### NewWorkflowCapabilities

`func NewWorkflowCapabilities() *WorkflowCapabilities`

NewWorkflowCapabilities instantiates a new WorkflowCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCapabilitiesWithDefaults

`func NewWorkflowCapabilitiesWithDefaults() *WorkflowCapabilities`

NewWorkflowCapabilitiesWithDefaults instantiates a new WorkflowCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectRules

`func (o *WorkflowCapabilities) GetConnectRules() []AvailableWorkflowConnectRule`

GetConnectRules returns the ConnectRules field if non-nil, zero value otherwise.

### GetConnectRulesOk

`func (o *WorkflowCapabilities) GetConnectRulesOk() (*[]AvailableWorkflowConnectRule, bool)`

GetConnectRulesOk returns a tuple with the ConnectRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectRules

`func (o *WorkflowCapabilities) SetConnectRules(v []AvailableWorkflowConnectRule)`

SetConnectRules sets ConnectRules field to given value.

### HasConnectRules

`func (o *WorkflowCapabilities) HasConnectRules() bool`

HasConnectRules returns a boolean if a field has been set.

### GetEditorScope

`func (o *WorkflowCapabilities) GetEditorScope() string`

GetEditorScope returns the EditorScope field if non-nil, zero value otherwise.

### GetEditorScopeOk

`func (o *WorkflowCapabilities) GetEditorScopeOk() (*string, bool)`

GetEditorScopeOk returns a tuple with the EditorScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorScope

`func (o *WorkflowCapabilities) SetEditorScope(v string)`

SetEditorScope sets EditorScope field to given value.

### HasEditorScope

`func (o *WorkflowCapabilities) HasEditorScope() bool`

HasEditorScope returns a boolean if a field has been set.

### GetForgeRules

`func (o *WorkflowCapabilities) GetForgeRules() []AvailableWorkflowForgeRule`

GetForgeRules returns the ForgeRules field if non-nil, zero value otherwise.

### GetForgeRulesOk

`func (o *WorkflowCapabilities) GetForgeRulesOk() (*[]AvailableWorkflowForgeRule, bool)`

GetForgeRulesOk returns a tuple with the ForgeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgeRules

`func (o *WorkflowCapabilities) SetForgeRules(v []AvailableWorkflowForgeRule)`

SetForgeRules sets ForgeRules field to given value.

### HasForgeRules

`func (o *WorkflowCapabilities) HasForgeRules() bool`

HasForgeRules returns a boolean if a field has been set.

### GetProjectTypes

`func (o *WorkflowCapabilities) GetProjectTypes() []string`

GetProjectTypes returns the ProjectTypes field if non-nil, zero value otherwise.

### GetProjectTypesOk

`func (o *WorkflowCapabilities) GetProjectTypesOk() (*[]string, bool)`

GetProjectTypesOk returns a tuple with the ProjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypes

`func (o *WorkflowCapabilities) SetProjectTypes(v []string)`

SetProjectTypes sets ProjectTypes field to given value.

### HasProjectTypes

`func (o *WorkflowCapabilities) HasProjectTypes() bool`

HasProjectTypes returns a boolean if a field has been set.

### GetSystemRules

`func (o *WorkflowCapabilities) GetSystemRules() []AvailableWorkflowSystemRule`

GetSystemRules returns the SystemRules field if non-nil, zero value otherwise.

### GetSystemRulesOk

`func (o *WorkflowCapabilities) GetSystemRulesOk() (*[]AvailableWorkflowSystemRule, bool)`

GetSystemRulesOk returns a tuple with the SystemRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRules

`func (o *WorkflowCapabilities) SetSystemRules(v []AvailableWorkflowSystemRule)`

SetSystemRules sets SystemRules field to given value.

### HasSystemRules

`func (o *WorkflowCapabilities) HasSystemRules() bool`

HasSystemRules returns a boolean if a field has been set.

### GetTriggerRules

`func (o *WorkflowCapabilities) GetTriggerRules() []AvailableWorkflowTriggers`

GetTriggerRules returns the TriggerRules field if non-nil, zero value otherwise.

### GetTriggerRulesOk

`func (o *WorkflowCapabilities) GetTriggerRulesOk() (*[]AvailableWorkflowTriggers, bool)`

GetTriggerRulesOk returns a tuple with the TriggerRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRules

`func (o *WorkflowCapabilities) SetTriggerRules(v []AvailableWorkflowTriggers)`

SetTriggerRules sets TriggerRules field to given value.

### HasTriggerRules

`func (o *WorkflowCapabilities) HasTriggerRules() bool`

HasTriggerRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


