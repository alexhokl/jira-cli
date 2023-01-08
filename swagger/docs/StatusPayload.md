# StatusPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status | [optional] 
**Name** | Pointer to **string** | The name of the status | [optional] 
**OnConflict** | Pointer to **string** | The conflict strategy for the status already exists. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters; NEW - Create a new entity | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**StatusCategory** | Pointer to **string** | The status category of the status. The value is case-sensitive. | [optional] 

## Methods

### NewStatusPayload

`func NewStatusPayload() *StatusPayload`

NewStatusPayload instantiates a new StatusPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusPayloadWithDefaults

`func NewStatusPayloadWithDefaults() *StatusPayload`

NewStatusPayloadWithDefaults instantiates a new StatusPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StatusPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StatusPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatusPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *StatusPayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *StatusPayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *StatusPayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *StatusPayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *StatusPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *StatusPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *StatusPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *StatusPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetStatusCategory

`func (o *StatusPayload) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *StatusPayload) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *StatusPayload) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *StatusPayload) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


