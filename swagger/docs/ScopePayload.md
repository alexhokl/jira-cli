# ScopePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the scope. Use &#x60;GLOBAL&#x60; or empty for company-managed project, and &#x60;PROJECT&#x60; for team-managed project | [optional] 

## Methods

### NewScopePayload

`func NewScopePayload() *ScopePayload`

NewScopePayload instantiates a new ScopePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopePayloadWithDefaults

`func NewScopePayloadWithDefaults() *ScopePayload`

NewScopePayloadWithDefaults instantiates a new ScopePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScopePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScopePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScopePayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScopePayload) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


