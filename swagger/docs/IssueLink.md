# IssueLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the issue link. | [optional] [readonly] 
**InwardIssue** | [**LinkedIssue**](LinkedIssue.md) | Provides details about the linked issue. If presenting this link in a user interface, use the &#x60;inward&#x60; field of the issue link type to label the link. | 
**OutwardIssue** | [**LinkedIssue**](LinkedIssue.md) | Provides details about the linked issue. If presenting this link in a user interface, use the &#x60;outward&#x60; field of the issue link type to label the link. | 
**Self** | Pointer to **string** | The URL of the issue link. | [optional] [readonly] 
**Type** | [**IssueLinkType**](IssueLinkType.md) | The type of link between the issues. | 

## Methods

### NewIssueLink

`func NewIssueLink(inwardIssue LinkedIssue, outwardIssue LinkedIssue, type_ IssueLinkType, ) *IssueLink`

NewIssueLink instantiates a new IssueLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLinkWithDefaults

`func NewIssueLinkWithDefaults() *IssueLink`

NewIssueLinkWithDefaults instantiates a new IssueLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInwardIssue

`func (o *IssueLink) GetInwardIssue() LinkedIssue`

GetInwardIssue returns the InwardIssue field if non-nil, zero value otherwise.

### GetInwardIssueOk

`func (o *IssueLink) GetInwardIssueOk() (*LinkedIssue, bool)`

GetInwardIssueOk returns a tuple with the InwardIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInwardIssue

`func (o *IssueLink) SetInwardIssue(v LinkedIssue)`

SetInwardIssue sets InwardIssue field to given value.


### GetOutwardIssue

`func (o *IssueLink) GetOutwardIssue() LinkedIssue`

GetOutwardIssue returns the OutwardIssue field if non-nil, zero value otherwise.

### GetOutwardIssueOk

`func (o *IssueLink) GetOutwardIssueOk() (*LinkedIssue, bool)`

GetOutwardIssueOk returns a tuple with the OutwardIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutwardIssue

`func (o *IssueLink) SetOutwardIssue(v LinkedIssue)`

SetOutwardIssue sets OutwardIssue field to given value.


### GetSelf

`func (o *IssueLink) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueLink) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueLink) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IssueLink) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetType

`func (o *IssueLink) GetType() IssueLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueLink) GetTypeOk() (*IssueLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueLink) SetType(v IssueLinkType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


