# SubmitDeploymentsResponse1DetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the gating status details.  | 
**IssueKey** | **string** | An issue key that references an issue in Jira.  | 
**IssueLink** | **string** | A full HTTPS link to the Jira issue for the change request gating this Deployment. This field is provided if the details type is issue.  | 

## Methods

### NewSubmitDeploymentsResponse1DetailsInner

`func NewSubmitDeploymentsResponse1DetailsInner(type_ string, issueKey string, issueLink string, ) *SubmitDeploymentsResponse1DetailsInner`

NewSubmitDeploymentsResponse1DetailsInner instantiates a new SubmitDeploymentsResponse1DetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDeploymentsResponse1DetailsInnerWithDefaults

`func NewSubmitDeploymentsResponse1DetailsInnerWithDefaults() *SubmitDeploymentsResponse1DetailsInner`

NewSubmitDeploymentsResponse1DetailsInnerWithDefaults instantiates a new SubmitDeploymentsResponse1DetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubmitDeploymentsResponse1DetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmitDeploymentsResponse1DetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmitDeploymentsResponse1DetailsInner) SetType(v string)`

SetType sets Type field to given value.


### GetIssueKey

`func (o *SubmitDeploymentsResponse1DetailsInner) GetIssueKey() string`

GetIssueKey returns the IssueKey field if non-nil, zero value otherwise.

### GetIssueKeyOk

`func (o *SubmitDeploymentsResponse1DetailsInner) GetIssueKeyOk() (*string, bool)`

GetIssueKeyOk returns a tuple with the IssueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKey

`func (o *SubmitDeploymentsResponse1DetailsInner) SetIssueKey(v string)`

SetIssueKey sets IssueKey field to given value.


### GetIssueLink

`func (o *SubmitDeploymentsResponse1DetailsInner) GetIssueLink() string`

GetIssueLink returns the IssueLink field if non-nil, zero value otherwise.

### GetIssueLinkOk

`func (o *SubmitDeploymentsResponse1DetailsInner) GetIssueLinkOk() (*string, bool)`

GetIssueLinkOk returns a tuple with the IssueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLink

`func (o *SubmitDeploymentsResponse1DetailsInner) SetIssueLink(v string)`

SetIssueLink sets IssueLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


