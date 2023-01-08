# VersionRelatedWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the related work | 
**IssueId** | Pointer to **int64** | The ID of the issue associated with the related work (if there is one). Cannot be updated via the Rest API. | [optional] [readonly] 
**RelatedWorkId** | Pointer to **string** | The id of the related work. For the native release note related work item, this will be null, and Rest API does not support updating it. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the related work | [optional] 
**Url** | Pointer to **string** | The URL of the related work. Will be null for the native release note related work item, but is otherwise required. | [optional] 

## Methods

### NewVersionRelatedWork

`func NewVersionRelatedWork(category string, ) *VersionRelatedWork`

NewVersionRelatedWork instantiates a new VersionRelatedWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionRelatedWorkWithDefaults

`func NewVersionRelatedWorkWithDefaults() *VersionRelatedWork`

NewVersionRelatedWorkWithDefaults instantiates a new VersionRelatedWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *VersionRelatedWork) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VersionRelatedWork) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VersionRelatedWork) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetIssueId

`func (o *VersionRelatedWork) GetIssueId() int64`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *VersionRelatedWork) GetIssueIdOk() (*int64, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *VersionRelatedWork) SetIssueId(v int64)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *VersionRelatedWork) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetRelatedWorkId

`func (o *VersionRelatedWork) GetRelatedWorkId() string`

GetRelatedWorkId returns the RelatedWorkId field if non-nil, zero value otherwise.

### GetRelatedWorkIdOk

`func (o *VersionRelatedWork) GetRelatedWorkIdOk() (*string, bool)`

GetRelatedWorkIdOk returns a tuple with the RelatedWorkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedWorkId

`func (o *VersionRelatedWork) SetRelatedWorkId(v string)`

SetRelatedWorkId sets RelatedWorkId field to given value.

### HasRelatedWorkId

`func (o *VersionRelatedWork) HasRelatedWorkId() bool`

HasRelatedWorkId returns a boolean if a field has been set.

### GetTitle

`func (o *VersionRelatedWork) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VersionRelatedWork) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VersionRelatedWork) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VersionRelatedWork) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *VersionRelatedWork) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VersionRelatedWork) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VersionRelatedWork) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VersionRelatedWork) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


