# CreateWorkflowTransitionRulesDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**CreateWorkflowCondition**](CreateWorkflowCondition.md) | The workflow conditions. | [optional] 
**PostFunctions** | Pointer to [**[]CreateWorkflowTransitionRule**](CreateWorkflowTransitionRule.md) | The workflow post functions.  **Note:** The default post functions are always added to the *initial* transition, as in:      \&quot;postFunctions\&quot;: [         {             \&quot;type\&quot;: \&quot;IssueCreateFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;IssueReindexFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;FireIssueEventFunction\&quot;,             \&quot;configuration\&quot;: {                 \&quot;event\&quot;: {                     \&quot;id\&quot;: \&quot;1\&quot;,                     \&quot;name\&quot;: \&quot;issue_created\&quot;                 }             }         }     ]  **Note:** The default post functions are always added to the *global* and *directed* transitions, as in:      \&quot;postFunctions\&quot;: [         {             \&quot;type\&quot;: \&quot;UpdateIssueStatusFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;CreateCommentFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;GenerateChangeHistoryFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;IssueReindexFunction\&quot;         },         {             \&quot;type\&quot;: \&quot;FireIssueEventFunction\&quot;,             \&quot;configuration\&quot;: {                 \&quot;event\&quot;: {                     \&quot;id\&quot;: \&quot;13\&quot;,                     \&quot;name\&quot;: \&quot;issue_generic\&quot;                 }             }         }     ] | [optional] 
**Validators** | Pointer to [**[]CreateWorkflowTransitionRule**](CreateWorkflowTransitionRule.md) | The workflow validators.  **Note:** The default permission validator is always added to the *initial* transition, as in:      \&quot;validators\&quot;: [         {             \&quot;type\&quot;: \&quot;PermissionValidator\&quot;,             \&quot;configuration\&quot;: {                 \&quot;permissionKey\&quot;: \&quot;CREATE_ISSUES\&quot;             }         }     ] | [optional] 

## Methods

### NewCreateWorkflowTransitionRulesDetails

`func NewCreateWorkflowTransitionRulesDetails() *CreateWorkflowTransitionRulesDetails`

NewCreateWorkflowTransitionRulesDetails instantiates a new CreateWorkflowTransitionRulesDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowTransitionRulesDetailsWithDefaults

`func NewCreateWorkflowTransitionRulesDetailsWithDefaults() *CreateWorkflowTransitionRulesDetails`

NewCreateWorkflowTransitionRulesDetailsWithDefaults instantiates a new CreateWorkflowTransitionRulesDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CreateWorkflowTransitionRulesDetails) GetConditions() CreateWorkflowCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateWorkflowTransitionRulesDetails) GetConditionsOk() (*CreateWorkflowCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateWorkflowTransitionRulesDetails) SetConditions(v CreateWorkflowCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateWorkflowTransitionRulesDetails) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPostFunctions

`func (o *CreateWorkflowTransitionRulesDetails) GetPostFunctions() []CreateWorkflowTransitionRule`

GetPostFunctions returns the PostFunctions field if non-nil, zero value otherwise.

### GetPostFunctionsOk

`func (o *CreateWorkflowTransitionRulesDetails) GetPostFunctionsOk() (*[]CreateWorkflowTransitionRule, bool)`

GetPostFunctionsOk returns a tuple with the PostFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFunctions

`func (o *CreateWorkflowTransitionRulesDetails) SetPostFunctions(v []CreateWorkflowTransitionRule)`

SetPostFunctions sets PostFunctions field to given value.

### HasPostFunctions

`func (o *CreateWorkflowTransitionRulesDetails) HasPostFunctions() bool`

HasPostFunctions returns a boolean if a field has been set.

### GetValidators

`func (o *CreateWorkflowTransitionRulesDetails) GetValidators() []CreateWorkflowTransitionRule`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *CreateWorkflowTransitionRulesDetails) GetValidatorsOk() (*[]CreateWorkflowTransitionRule, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *CreateWorkflowTransitionRulesDetails) SetValidators(v []CreateWorkflowTransitionRule)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *CreateWorkflowTransitionRulesDetails) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


