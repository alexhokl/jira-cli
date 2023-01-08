# RemoteObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to [**Icon**](Icon.md) | Details of the icon for the item. If no icon is defined, the default link icon is used in Jira. | [optional] 
**Status** | Pointer to [**Status**](Status.md) | The status of the item. | [optional] 
**Summary** | Pointer to **string** | The summary details of the item. | [optional] 
**Title** | **string** | The title of the item. | 
**Url** | **string** | The URL of the item. | 

## Methods

### NewRemoteObject

`func NewRemoteObject(title string, url string, ) *RemoteObject`

NewRemoteObject instantiates a new RemoteObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteObjectWithDefaults

`func NewRemoteObjectWithDefaults() *RemoteObject`

NewRemoteObjectWithDefaults instantiates a new RemoteObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *RemoteObject) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RemoteObject) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RemoteObject) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RemoteObject) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetStatus

`func (o *RemoteObject) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteObject) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteObject) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *RemoteObject) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RemoteObject) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RemoteObject) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RemoteObject) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *RemoteObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RemoteObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RemoteObject) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *RemoteObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteObject) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


