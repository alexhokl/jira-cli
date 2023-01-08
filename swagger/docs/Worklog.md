# Worklog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**UserDetails**](UserDetails.md) | Details of the user who created the worklog. | [optional] [readonly] 
**Comment** | Pointer to **interface{}** | A comment about the worklog in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/). Optional when creating or updating a worklog. | [optional] 
**Created** | Pointer to **time.Time** | The datetime on which the worklog was created. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the worklog record. | [optional] [readonly] 
**IssueId** | Pointer to **string** | The ID of the issue this worklog is for. | [optional] [readonly] 
**Properties** | Pointer to [**[]EntityProperty**](EntityProperty.md) | Details of properties for the worklog. Optional when creating or updating a worklog. | [optional] 
**Self** | Pointer to **string** | The URL of the worklog item. | [optional] [readonly] 
**Started** | Pointer to **time.Time** | The datetime on which the worklog effort was started. Required when creating a worklog. Optional when updating a worklog. | [optional] 
**TimeSpent** | Pointer to **string** | The time spent working on the issue as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). Required when creating a worklog if &#x60;timeSpentSeconds&#x60; isn&#39;t provided. Optional when updating a worklog. Cannot be provided if &#x60;timeSpentSecond&#x60; is provided. | [optional] 
**TimeSpentSeconds** | Pointer to **int64** | The time in seconds spent working on the issue. Required when creating a worklog if &#x60;timeSpent&#x60; isn&#39;t provided. Optional when updating a worklog. Cannot be provided if &#x60;timeSpent&#x60; is provided. | [optional] 
**UpdateAuthor** | Pointer to [**UserDetails**](UserDetails.md) | Details of the user who last updated the worklog. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The datetime on which the worklog was last updated. | [optional] [readonly] 
**Visibility** | Pointer to [**Visibility**](Visibility.md) | Details about any restrictions in the visibility of the worklog. Optional when creating or updating a worklog. | [optional] 

## Methods

### NewWorklog

`func NewWorklog() *Worklog`

NewWorklog instantiates a new Worklog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorklogWithDefaults

`func NewWorklogWithDefaults() *Worklog`

NewWorklogWithDefaults instantiates a new Worklog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Worklog) GetAuthor() UserDetails`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Worklog) GetAuthorOk() (*UserDetails, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Worklog) SetAuthor(v UserDetails)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Worklog) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComment

`func (o *Worklog) GetComment() interface{}`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Worklog) GetCommentOk() (*interface{}, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Worklog) SetComment(v interface{})`

SetComment sets Comment field to given value.

### HasComment

`func (o *Worklog) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Worklog) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Worklog) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCreated

`func (o *Worklog) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Worklog) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Worklog) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Worklog) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Worklog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Worklog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Worklog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Worklog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueId

`func (o *Worklog) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Worklog) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Worklog) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *Worklog) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetProperties

`func (o *Worklog) GetProperties() []EntityProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Worklog) GetPropertiesOk() (*[]EntityProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Worklog) SetProperties(v []EntityProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Worklog) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSelf

`func (o *Worklog) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Worklog) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Worklog) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Worklog) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStarted

`func (o *Worklog) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *Worklog) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *Worklog) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *Worklog) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTimeSpent

`func (o *Worklog) GetTimeSpent() string`

GetTimeSpent returns the TimeSpent field if non-nil, zero value otherwise.

### GetTimeSpentOk

`func (o *Worklog) GetTimeSpentOk() (*string, bool)`

GetTimeSpentOk returns a tuple with the TimeSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpent

`func (o *Worklog) SetTimeSpent(v string)`

SetTimeSpent sets TimeSpent field to given value.

### HasTimeSpent

`func (o *Worklog) HasTimeSpent() bool`

HasTimeSpent returns a boolean if a field has been set.

### GetTimeSpentSeconds

`func (o *Worklog) GetTimeSpentSeconds() int64`

GetTimeSpentSeconds returns the TimeSpentSeconds field if non-nil, zero value otherwise.

### GetTimeSpentSecondsOk

`func (o *Worklog) GetTimeSpentSecondsOk() (*int64, bool)`

GetTimeSpentSecondsOk returns a tuple with the TimeSpentSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpentSeconds

`func (o *Worklog) SetTimeSpentSeconds(v int64)`

SetTimeSpentSeconds sets TimeSpentSeconds field to given value.

### HasTimeSpentSeconds

`func (o *Worklog) HasTimeSpentSeconds() bool`

HasTimeSpentSeconds returns a boolean if a field has been set.

### GetUpdateAuthor

`func (o *Worklog) GetUpdateAuthor() UserDetails`

GetUpdateAuthor returns the UpdateAuthor field if non-nil, zero value otherwise.

### GetUpdateAuthorOk

`func (o *Worklog) GetUpdateAuthorOk() (*UserDetails, bool)`

GetUpdateAuthorOk returns a tuple with the UpdateAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthor

`func (o *Worklog) SetUpdateAuthor(v UserDetails)`

SetUpdateAuthor sets UpdateAuthor field to given value.

### HasUpdateAuthor

`func (o *Worklog) HasUpdateAuthor() bool`

HasUpdateAuthor returns a boolean if a field has been set.

### GetUpdated

`func (o *Worklog) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Worklog) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Worklog) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Worklog) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVisibility

`func (o *Worklog) GetVisibility() Visibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Worklog) GetVisibilityOk() (*Visibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Worklog) SetVisibility(v Visibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Worklog) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


