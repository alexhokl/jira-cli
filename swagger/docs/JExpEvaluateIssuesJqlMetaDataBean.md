# JExpEvaluateIssuesJqlMetaDataBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Indicates whether this is the last page of the paginated response. | [optional] [readonly] 
**NextPageToken** | **string** | Next Page token for the next page of issues. | 

## Methods

### NewJExpEvaluateIssuesJqlMetaDataBean

`func NewJExpEvaluateIssuesJqlMetaDataBean(nextPageToken string, ) *JExpEvaluateIssuesJqlMetaDataBean`

NewJExpEvaluateIssuesJqlMetaDataBean instantiates a new JExpEvaluateIssuesJqlMetaDataBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJExpEvaluateIssuesJqlMetaDataBeanWithDefaults

`func NewJExpEvaluateIssuesJqlMetaDataBeanWithDefaults() *JExpEvaluateIssuesJqlMetaDataBean`

NewJExpEvaluateIssuesJqlMetaDataBeanWithDefaults instantiates a new JExpEvaluateIssuesJqlMetaDataBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *JExpEvaluateIssuesJqlMetaDataBean) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *JExpEvaluateIssuesJqlMetaDataBean) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *JExpEvaluateIssuesJqlMetaDataBean) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *JExpEvaluateIssuesJqlMetaDataBean) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetNextPageToken

`func (o *JExpEvaluateIssuesJqlMetaDataBean) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *JExpEvaluateIssuesJqlMetaDataBean) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *JExpEvaluateIssuesJqlMetaDataBean) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


