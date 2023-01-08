# ScreenSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultScreen** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the screen scheme | [optional] 
**Name** | Pointer to **string** | The name of the screen scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Screens** | Pointer to [**map[string]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | Similar to the field layout scheme those mappings allow users to set different screens for different operations: default - always there, applied to all operations that don&#39;t have an explicit mapping &#x60;create&#x60;, &#x60;view&#x60;, &#x60;edit&#x60; - specific operations that are available and users can assign a different screen for each one of them https://support.atlassian.com/jira-cloud-administration/docs/manage-screen-schemes/\\#Associating-a-screen-with-an-issue-operation | [optional] 

## Methods

### NewScreenSchemePayload

`func NewScreenSchemePayload() *ScreenSchemePayload`

NewScreenSchemePayload instantiates a new ScreenSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenSchemePayloadWithDefaults

`func NewScreenSchemePayloadWithDefaults() *ScreenSchemePayload`

NewScreenSchemePayloadWithDefaults instantiates a new ScreenSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultScreen

`func (o *ScreenSchemePayload) GetDefaultScreen() ProjectCreateResourceIdentifier`

GetDefaultScreen returns the DefaultScreen field if non-nil, zero value otherwise.

### GetDefaultScreenOk

`func (o *ScreenSchemePayload) GetDefaultScreenOk() (*ProjectCreateResourceIdentifier, bool)`

GetDefaultScreenOk returns a tuple with the DefaultScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScreen

`func (o *ScreenSchemePayload) SetDefaultScreen(v ProjectCreateResourceIdentifier)`

SetDefaultScreen sets DefaultScreen field to given value.

### HasDefaultScreen

`func (o *ScreenSchemePayload) HasDefaultScreen() bool`

HasDefaultScreen returns a boolean if a field has been set.

### GetDescription

`func (o *ScreenSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScreenSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScreenSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *ScreenSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *ScreenSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *ScreenSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *ScreenSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetScreens

`func (o *ScreenSchemePayload) GetScreens() map[string]ProjectCreateResourceIdentifier`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *ScreenSchemePayload) GetScreensOk() (*map[string]ProjectCreateResourceIdentifier, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *ScreenSchemePayload) SetScreens(v map[string]ProjectCreateResourceIdentifier)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *ScreenSchemePayload) HasScreens() bool`

HasScreens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


