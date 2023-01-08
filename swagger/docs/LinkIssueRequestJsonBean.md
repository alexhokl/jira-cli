# LinkIssueRequestJsonBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**Comment**](Comment.md) |  | [optional] 
**InwardIssue** | [**LinkedIssue**](LinkedIssue.md) |  | 
**OutwardIssue** | [**LinkedIssue**](LinkedIssue.md) |  | 
**Type** | [**IssueLinkType**](IssueLinkType.md) |  | 

## Methods

### NewLinkIssueRequestJsonBean

`func NewLinkIssueRequestJsonBean(inwardIssue LinkedIssue, outwardIssue LinkedIssue, type_ IssueLinkType, ) *LinkIssueRequestJsonBean`

NewLinkIssueRequestJsonBean instantiates a new LinkIssueRequestJsonBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkIssueRequestJsonBeanWithDefaults

`func NewLinkIssueRequestJsonBeanWithDefaults() *LinkIssueRequestJsonBean`

NewLinkIssueRequestJsonBeanWithDefaults instantiates a new LinkIssueRequestJsonBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *LinkIssueRequestJsonBean) GetComment() Comment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LinkIssueRequestJsonBean) GetCommentOk() (*Comment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LinkIssueRequestJsonBean) SetComment(v Comment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LinkIssueRequestJsonBean) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInwardIssue

`func (o *LinkIssueRequestJsonBean) GetInwardIssue() LinkedIssue`

GetInwardIssue returns the InwardIssue field if non-nil, zero value otherwise.

### GetInwardIssueOk

`func (o *LinkIssueRequestJsonBean) GetInwardIssueOk() (*LinkedIssue, bool)`

GetInwardIssueOk returns a tuple with the InwardIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInwardIssue

`func (o *LinkIssueRequestJsonBean) SetInwardIssue(v LinkedIssue)`

SetInwardIssue sets InwardIssue field to given value.


### GetOutwardIssue

`func (o *LinkIssueRequestJsonBean) GetOutwardIssue() LinkedIssue`

GetOutwardIssue returns the OutwardIssue field if non-nil, zero value otherwise.

### GetOutwardIssueOk

`func (o *LinkIssueRequestJsonBean) GetOutwardIssueOk() (*LinkedIssue, bool)`

GetOutwardIssueOk returns a tuple with the OutwardIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutwardIssue

`func (o *LinkIssueRequestJsonBean) SetOutwardIssue(v LinkedIssue)`

SetOutwardIssue sets OutwardIssue field to given value.


### GetType

`func (o *LinkIssueRequestJsonBean) GetType() IssueLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkIssueRequestJsonBean) GetTypeOk() (*IssueLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkIssueRequestJsonBean) SetType(v IssueLinkType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


