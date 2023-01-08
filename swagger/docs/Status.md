# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to [**Icon**](Icon.md) | Details of the icon representing the status. If not provided, no status icon displays in Jira. | [optional] 
**Resolved** | Pointer to **bool** | Whether the item is resolved. If set to \&quot;true\&quot;, the link to the issue is displayed in a strikethrough font, otherwise the link displays in normal font. | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *Status) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Status) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Status) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Status) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetResolved

`func (o *Status) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *Status) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *Status) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *Status) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


