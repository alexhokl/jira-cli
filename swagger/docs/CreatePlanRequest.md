# CreatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossProjectReleases** | Pointer to [**[]CreateCrossProjectReleaseRequest**](CreateCrossProjectReleaseRequest.md) | The cross-project releases to include in the plan. | [optional] 
**CustomFields** | Pointer to [**[]CreateCustomFieldRequest**](CreateCustomFieldRequest.md) | The custom fields for the plan. | [optional] 
**ExclusionRules** | Pointer to [**CreateExclusionRulesRequest**](CreateExclusionRulesRequest.md) | The exclusion rules for the plan. | [optional] 
**IssueSources** | [**[]CreateIssueSourceRequest**](CreateIssueSourceRequest.md) | The issue sources to include in the plan. | 
**LeadAccountId** | Pointer to **string** | The account ID of the plan lead. | [optional] 
**Name** | **string** | The plan name. | 
**Permissions** | Pointer to [**[]CreatePermissionRequest**](CreatePermissionRequest.md) | The permissions for the plan. | [optional] 
**Scheduling** | [**CreateSchedulingRequest**](CreateSchedulingRequest.md) | The scheduling settings for the plan. | 

## Methods

### NewCreatePlanRequest

`func NewCreatePlanRequest(issueSources []CreateIssueSourceRequest, name string, scheduling CreateSchedulingRequest, ) *CreatePlanRequest`

NewCreatePlanRequest instantiates a new CreatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanRequestWithDefaults

`func NewCreatePlanRequestWithDefaults() *CreatePlanRequest`

NewCreatePlanRequestWithDefaults instantiates a new CreatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossProjectReleases

`func (o *CreatePlanRequest) GetCrossProjectReleases() []CreateCrossProjectReleaseRequest`

GetCrossProjectReleases returns the CrossProjectReleases field if non-nil, zero value otherwise.

### GetCrossProjectReleasesOk

`func (o *CreatePlanRequest) GetCrossProjectReleasesOk() (*[]CreateCrossProjectReleaseRequest, bool)`

GetCrossProjectReleasesOk returns a tuple with the CrossProjectReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossProjectReleases

`func (o *CreatePlanRequest) SetCrossProjectReleases(v []CreateCrossProjectReleaseRequest)`

SetCrossProjectReleases sets CrossProjectReleases field to given value.

### HasCrossProjectReleases

`func (o *CreatePlanRequest) HasCrossProjectReleases() bool`

HasCrossProjectReleases returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreatePlanRequest) GetCustomFields() []CreateCustomFieldRequest`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreatePlanRequest) GetCustomFieldsOk() (*[]CreateCustomFieldRequest, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreatePlanRequest) SetCustomFields(v []CreateCustomFieldRequest)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreatePlanRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetExclusionRules

`func (o *CreatePlanRequest) GetExclusionRules() CreateExclusionRulesRequest`

GetExclusionRules returns the ExclusionRules field if non-nil, zero value otherwise.

### GetExclusionRulesOk

`func (o *CreatePlanRequest) GetExclusionRulesOk() (*CreateExclusionRulesRequest, bool)`

GetExclusionRulesOk returns a tuple with the ExclusionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRules

`func (o *CreatePlanRequest) SetExclusionRules(v CreateExclusionRulesRequest)`

SetExclusionRules sets ExclusionRules field to given value.

### HasExclusionRules

`func (o *CreatePlanRequest) HasExclusionRules() bool`

HasExclusionRules returns a boolean if a field has been set.

### GetIssueSources

`func (o *CreatePlanRequest) GetIssueSources() []CreateIssueSourceRequest`

GetIssueSources returns the IssueSources field if non-nil, zero value otherwise.

### GetIssueSourcesOk

`func (o *CreatePlanRequest) GetIssueSourcesOk() (*[]CreateIssueSourceRequest, bool)`

GetIssueSourcesOk returns a tuple with the IssueSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSources

`func (o *CreatePlanRequest) SetIssueSources(v []CreateIssueSourceRequest)`

SetIssueSources sets IssueSources field to given value.


### GetLeadAccountId

`func (o *CreatePlanRequest) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *CreatePlanRequest) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *CreatePlanRequest) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *CreatePlanRequest) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetName

`func (o *CreatePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CreatePlanRequest) GetPermissions() []CreatePermissionRequest`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreatePlanRequest) GetPermissionsOk() (*[]CreatePermissionRequest, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreatePlanRequest) SetPermissions(v []CreatePermissionRequest)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreatePlanRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetScheduling

`func (o *CreatePlanRequest) GetScheduling() CreateSchedulingRequest`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *CreatePlanRequest) GetSchedulingOk() (*CreateSchedulingRequest, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *CreatePlanRequest) SetScheduling(v CreateSchedulingRequest)`

SetScheduling sets Scheduling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


