# ProjectIssueCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrls** | Pointer to [**AvatarUrlsBean**](AvatarUrlsBean.md) | List of the project&#39;s avatars, returning the avatar size and associated URL. | [optional] [readonly] 
**Expand** | Pointer to **string** | Expand options that include additional project issue create metadata details in the response. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the project. | [optional] [readonly] 
**Issuetypes** | Pointer to [**[]IssueTypeIssueCreateMetadata**](IssueTypeIssueCreateMetadata.md) | List of the issue types supported by the project. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the project. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the project. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the project. | [optional] [readonly] 

## Methods

### NewProjectIssueCreateMetadata

`func NewProjectIssueCreateMetadata() *ProjectIssueCreateMetadata`

NewProjectIssueCreateMetadata instantiates a new ProjectIssueCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueCreateMetadataWithDefaults

`func NewProjectIssueCreateMetadataWithDefaults() *ProjectIssueCreateMetadata`

NewProjectIssueCreateMetadataWithDefaults instantiates a new ProjectIssueCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrls

`func (o *ProjectIssueCreateMetadata) GetAvatarUrls() AvatarUrlsBean`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *ProjectIssueCreateMetadata) GetAvatarUrlsOk() (*AvatarUrlsBean, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *ProjectIssueCreateMetadata) SetAvatarUrls(v AvatarUrlsBean)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *ProjectIssueCreateMetadata) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetExpand

`func (o *ProjectIssueCreateMetadata) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *ProjectIssueCreateMetadata) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *ProjectIssueCreateMetadata) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *ProjectIssueCreateMetadata) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetId

`func (o *ProjectIssueCreateMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectIssueCreateMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectIssueCreateMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectIssueCreateMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuetypes

`func (o *ProjectIssueCreateMetadata) GetIssuetypes() []IssueTypeIssueCreateMetadata`

GetIssuetypes returns the Issuetypes field if non-nil, zero value otherwise.

### GetIssuetypesOk

`func (o *ProjectIssueCreateMetadata) GetIssuetypesOk() (*[]IssueTypeIssueCreateMetadata, bool)`

GetIssuetypesOk returns a tuple with the Issuetypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuetypes

`func (o *ProjectIssueCreateMetadata) SetIssuetypes(v []IssueTypeIssueCreateMetadata)`

SetIssuetypes sets Issuetypes field to given value.

### HasIssuetypes

`func (o *ProjectIssueCreateMetadata) HasIssuetypes() bool`

HasIssuetypes returns a boolean if a field has been set.

### GetKey

`func (o *ProjectIssueCreateMetadata) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectIssueCreateMetadata) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectIssueCreateMetadata) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectIssueCreateMetadata) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ProjectIssueCreateMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectIssueCreateMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectIssueCreateMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectIssueCreateMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectIssueCreateMetadata) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectIssueCreateMetadata) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectIssueCreateMetadata) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ProjectIssueCreateMetadata) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


