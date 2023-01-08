# GetPlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossProjectReleases** | Pointer to [**[]GetCrossProjectReleaseResponse**](GetCrossProjectReleaseResponse.md) | The cross-project releases included in the plan. | [optional] 
**CustomFields** | Pointer to [**[]GetCustomFieldResponse**](GetCustomFieldResponse.md) | The custom fields for the plan. | [optional] 
**ExclusionRules** | Pointer to [**GetExclusionRulesResponse**](GetExclusionRulesResponse.md) | The exclusion rules for the plan. | [optional] 
**Id** | **int64** | The plan ID. | 
**IssueSources** | Pointer to [**[]GetIssueSourceResponse**](GetIssueSourceResponse.md) | The issue sources included in the plan. | [optional] 
**LastSaved** | Pointer to **string** | The date when the plan was last saved in UTC. | [optional] 
**LeadAccountId** | Pointer to **string** | The account ID of the plan lead. | [optional] 
**Name** | Pointer to **string** | The plan name. | [optional] 
**Permissions** | Pointer to [**[]GetPermissionResponse**](GetPermissionResponse.md) | The permissions for the plan. | [optional] 
**Scheduling** | [**GetSchedulingResponse**](GetSchedulingResponse.md) | The scheduling settings for the plan. | 
**Status** | **string** | The plan status. This is \&quot;Active\&quot;, \&quot;Trashed\&quot; or \&quot;Archived\&quot;. | 

## Methods

### NewGetPlanResponse

`func NewGetPlanResponse(id int64, scheduling GetSchedulingResponse, status string, ) *GetPlanResponse`

NewGetPlanResponse instantiates a new GetPlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlanResponseWithDefaults

`func NewGetPlanResponseWithDefaults() *GetPlanResponse`

NewGetPlanResponseWithDefaults instantiates a new GetPlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossProjectReleases

`func (o *GetPlanResponse) GetCrossProjectReleases() []GetCrossProjectReleaseResponse`

GetCrossProjectReleases returns the CrossProjectReleases field if non-nil, zero value otherwise.

### GetCrossProjectReleasesOk

`func (o *GetPlanResponse) GetCrossProjectReleasesOk() (*[]GetCrossProjectReleaseResponse, bool)`

GetCrossProjectReleasesOk returns a tuple with the CrossProjectReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossProjectReleases

`func (o *GetPlanResponse) SetCrossProjectReleases(v []GetCrossProjectReleaseResponse)`

SetCrossProjectReleases sets CrossProjectReleases field to given value.

### HasCrossProjectReleases

`func (o *GetPlanResponse) HasCrossProjectReleases() bool`

HasCrossProjectReleases returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetPlanResponse) GetCustomFields() []GetCustomFieldResponse`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetPlanResponse) GetCustomFieldsOk() (*[]GetCustomFieldResponse, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetPlanResponse) SetCustomFields(v []GetCustomFieldResponse)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetPlanResponse) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetExclusionRules

`func (o *GetPlanResponse) GetExclusionRules() GetExclusionRulesResponse`

GetExclusionRules returns the ExclusionRules field if non-nil, zero value otherwise.

### GetExclusionRulesOk

`func (o *GetPlanResponse) GetExclusionRulesOk() (*GetExclusionRulesResponse, bool)`

GetExclusionRulesOk returns a tuple with the ExclusionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRules

`func (o *GetPlanResponse) SetExclusionRules(v GetExclusionRulesResponse)`

SetExclusionRules sets ExclusionRules field to given value.

### HasExclusionRules

`func (o *GetPlanResponse) HasExclusionRules() bool`

HasExclusionRules returns a boolean if a field has been set.

### GetId

`func (o *GetPlanResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlanResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlanResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueSources

`func (o *GetPlanResponse) GetIssueSources() []GetIssueSourceResponse`

GetIssueSources returns the IssueSources field if non-nil, zero value otherwise.

### GetIssueSourcesOk

`func (o *GetPlanResponse) GetIssueSourcesOk() (*[]GetIssueSourceResponse, bool)`

GetIssueSourcesOk returns a tuple with the IssueSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSources

`func (o *GetPlanResponse) SetIssueSources(v []GetIssueSourceResponse)`

SetIssueSources sets IssueSources field to given value.

### HasIssueSources

`func (o *GetPlanResponse) HasIssueSources() bool`

HasIssueSources returns a boolean if a field has been set.

### GetLastSaved

`func (o *GetPlanResponse) GetLastSaved() string`

GetLastSaved returns the LastSaved field if non-nil, zero value otherwise.

### GetLastSavedOk

`func (o *GetPlanResponse) GetLastSavedOk() (*string, bool)`

GetLastSavedOk returns a tuple with the LastSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaved

`func (o *GetPlanResponse) SetLastSaved(v string)`

SetLastSaved sets LastSaved field to given value.

### HasLastSaved

`func (o *GetPlanResponse) HasLastSaved() bool`

HasLastSaved returns a boolean if a field has been set.

### GetLeadAccountId

`func (o *GetPlanResponse) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *GetPlanResponse) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *GetPlanResponse) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *GetPlanResponse) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetName

`func (o *GetPlanResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPlanResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPlanResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPlanResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *GetPlanResponse) GetPermissions() []GetPermissionResponse`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetPlanResponse) GetPermissionsOk() (*[]GetPermissionResponse, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetPlanResponse) SetPermissions(v []GetPermissionResponse)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetPlanResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetScheduling

`func (o *GetPlanResponse) GetScheduling() GetSchedulingResponse`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *GetPlanResponse) GetSchedulingOk() (*GetSchedulingResponse, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *GetPlanResponse) SetScheduling(v GetSchedulingResponse)`

SetScheduling sets Scheduling field to given value.


### GetStatus

`func (o *GetPlanResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlanResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlanResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


