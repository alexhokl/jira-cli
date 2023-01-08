# TabPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | The list of resource identifier of the field associated to the tab. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/\\#api-rest-api-3-screens-screenid-tabs-tabid-fields-post | [optional] 
**Name** | Pointer to **string** | The name of the tab | [optional] 

## Methods

### NewTabPayload

`func NewTabPayload() *TabPayload`

NewTabPayload instantiates a new TabPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTabPayloadWithDefaults

`func NewTabPayloadWithDefaults() *TabPayload`

NewTabPayloadWithDefaults instantiates a new TabPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *TabPayload) GetFields() []ProjectCreateResourceIdentifier`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TabPayload) GetFieldsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TabPayload) SetFields(v []ProjectCreateResourceIdentifier)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TabPayload) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *TabPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TabPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TabPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TabPayload) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


