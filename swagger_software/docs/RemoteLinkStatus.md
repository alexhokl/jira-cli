# RemoteLinkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | **string** | Appearance is a fixed set of appearance types affecting the colour of the status lozenge in the UI. The colours they correspond to are equivalent to atlaskit&#39;s [Lozenge](https://atlaskit.atlassian.com/packages/core/lozenge) component.  | 
**Label** | **string** | The human-readable description for the Remote Link status.  Will be shown in the UI.  | 

## Methods

### NewRemoteLinkStatus

`func NewRemoteLinkStatus(appearance string, label string, ) *RemoteLinkStatus`

NewRemoteLinkStatus instantiates a new RemoteLinkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLinkStatusWithDefaults

`func NewRemoteLinkStatusWithDefaults() *RemoteLinkStatus`

NewRemoteLinkStatusWithDefaults instantiates a new RemoteLinkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *RemoteLinkStatus) GetAppearance() string`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *RemoteLinkStatus) GetAppearanceOk() (*string, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *RemoteLinkStatus) SetAppearance(v string)`

SetAppearance sets Appearance field to given value.


### GetLabel

`func (o *RemoteLinkStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RemoteLinkStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RemoteLinkStatus) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


