# IssueCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional project details in the response. | [optional] [readonly] 
**Projects** | Pointer to [**[]ProjectIssueCreateMetadata**](ProjectIssueCreateMetadata.md) | List of projects and their issue creation metadata. | [optional] [readonly] 

## Methods

### NewIssueCreateMetadata

`func NewIssueCreateMetadata() *IssueCreateMetadata`

NewIssueCreateMetadata instantiates a new IssueCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCreateMetadataWithDefaults

`func NewIssueCreateMetadataWithDefaults() *IssueCreateMetadata`

NewIssueCreateMetadataWithDefaults instantiates a new IssueCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *IssueCreateMetadata) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *IssueCreateMetadata) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *IssueCreateMetadata) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *IssueCreateMetadata) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetProjects

`func (o *IssueCreateMetadata) GetProjects() []ProjectIssueCreateMetadata`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *IssueCreateMetadata) GetProjectsOk() (*[]ProjectIssueCreateMetadata, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *IssueCreateMetadata) SetProjects(v []ProjectIssueCreateMetadata)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *IssueCreateMetadata) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


