# UiModificationContextDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the UI modification context. | [optional] [readonly] 
**IsAvailable** | Pointer to **bool** | Whether a context is available. For example, when a project is deleted the context becomes unavailable. | [optional] [readonly] 
**IssueTypeId** | Pointer to **string** | The issue type ID of the context. Null is treated as a wildcard, meaning the UI modification will be applied to all issue types. Each UI modification context can have a maximum of one wildcard. | [optional] 
**PortalId** | Pointer to **string** | The portal ID of the context. Only required for Jira Service Management request create portal view (&#x60;JSMRequestCreate&#x60;). | [optional] 
**ProjectId** | Pointer to **string** | The project ID of the context. Null is treated as a wildcard, meaning the UI modification will be applied to all projects. Each UI modification context can have a maximum of one wildcard. | [optional] 
**RequestTypeId** | Pointer to **string** | The request type ID of the context. Only required for Jira Service Management request create portal view (&#x60;JSMRequestCreate&#x60;). | [optional] 
**ViewType** | Pointer to **string** | The view type of the context.   Supported values:   *  &#x60;GIC&#x60; \\- Jira global issue create  *  &#x60;IssueView&#x60; \\- Jira issue view  *  &#x60;IssueTransition&#x60; \\- Jira issue transition  *  &#x60;JSMRequestCreate&#x60; \\- Jira Service Management request create portal view  For Jira view types (&#x60;GIC&#x60;, &#x60;IssueView&#x60;, &#x60;IssueTransition&#x60;), null is treated as a wildcard, meaning the UI modification will be applied to all view types. Each Jira context can have a maximum of one wildcard.      Wildcards are not applicable for JSM contexts. | [optional] 

## Methods

### NewUiModificationContextDetails

`func NewUiModificationContextDetails() *UiModificationContextDetails`

NewUiModificationContextDetails instantiates a new UiModificationContextDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiModificationContextDetailsWithDefaults

`func NewUiModificationContextDetailsWithDefaults() *UiModificationContextDetails`

NewUiModificationContextDetailsWithDefaults instantiates a new UiModificationContextDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UiModificationContextDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiModificationContextDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiModificationContextDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UiModificationContextDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAvailable

`func (o *UiModificationContextDetails) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *UiModificationContextDetails) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *UiModificationContextDetails) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *UiModificationContextDetails) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetIssueTypeId

`func (o *UiModificationContextDetails) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *UiModificationContextDetails) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *UiModificationContextDetails) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *UiModificationContextDetails) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.

### GetPortalId

`func (o *UiModificationContextDetails) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *UiModificationContextDetails) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *UiModificationContextDetails) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *UiModificationContextDetails) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetProjectId

`func (o *UiModificationContextDetails) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UiModificationContextDetails) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UiModificationContextDetails) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UiModificationContextDetails) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequestTypeId

`func (o *UiModificationContextDetails) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *UiModificationContextDetails) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *UiModificationContextDetails) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.

### HasRequestTypeId

`func (o *UiModificationContextDetails) HasRequestTypeId() bool`

HasRequestTypeId returns a boolean if a field has been set.

### GetViewType

`func (o *UiModificationContextDetails) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *UiModificationContextDetails) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *UiModificationContextDetails) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *UiModificationContextDetails) HasViewType() bool`

HasViewType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


