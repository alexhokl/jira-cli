# SetDefaultPriorityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the new default issue priority. Must be an existing ID or null. Setting this to null erases the default priority setting. | 

## Methods

### NewSetDefaultPriorityRequest

`func NewSetDefaultPriorityRequest(id string, ) *SetDefaultPriorityRequest`

NewSetDefaultPriorityRequest instantiates a new SetDefaultPriorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultPriorityRequestWithDefaults

`func NewSetDefaultPriorityRequestWithDefaults() *SetDefaultPriorityRequest`

NewSetDefaultPriorityRequestWithDefaults instantiates a new SetDefaultPriorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetDefaultPriorityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetDefaultPriorityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetDefaultPriorityRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


