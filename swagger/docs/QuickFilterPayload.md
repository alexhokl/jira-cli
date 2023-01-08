# QuickFilterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the quick filter | [optional] 
**JqlQuery** | Pointer to **string** | The jql query for the quick filter | [optional] 
**Name** | Pointer to **string** | The name of the quick filter | [optional] 

## Methods

### NewQuickFilterPayload

`func NewQuickFilterPayload() *QuickFilterPayload`

NewQuickFilterPayload instantiates a new QuickFilterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickFilterPayloadWithDefaults

`func NewQuickFilterPayloadWithDefaults() *QuickFilterPayload`

NewQuickFilterPayloadWithDefaults instantiates a new QuickFilterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QuickFilterPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickFilterPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickFilterPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickFilterPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJqlQuery

`func (o *QuickFilterPayload) GetJqlQuery() string`

GetJqlQuery returns the JqlQuery field if non-nil, zero value otherwise.

### GetJqlQueryOk

`func (o *QuickFilterPayload) GetJqlQueryOk() (*string, bool)`

GetJqlQueryOk returns a tuple with the JqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqlQuery

`func (o *QuickFilterPayload) SetJqlQuery(v string)`

SetJqlQuery sets JqlQuery field to given value.

### HasJqlQuery

`func (o *QuickFilterPayload) HasJqlQuery() bool`

HasJqlQuery returns a boolean if a field has been set.

### GetName

`func (o *QuickFilterPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickFilterPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickFilterPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickFilterPayload) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


