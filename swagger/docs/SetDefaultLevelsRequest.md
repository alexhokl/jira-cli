# SetDefaultLevelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | [**[]DefaultLevelValue**](DefaultLevelValue.md) | List of objects with issue security scheme ID and new default level ID. | 

## Methods

### NewSetDefaultLevelsRequest

`func NewSetDefaultLevelsRequest(defaultValues []DefaultLevelValue, ) *SetDefaultLevelsRequest`

NewSetDefaultLevelsRequest instantiates a new SetDefaultLevelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultLevelsRequestWithDefaults

`func NewSetDefaultLevelsRequestWithDefaults() *SetDefaultLevelsRequest`

NewSetDefaultLevelsRequestWithDefaults instantiates a new SetDefaultLevelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *SetDefaultLevelsRequest) GetDefaultValues() []DefaultLevelValue`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *SetDefaultLevelsRequest) GetDefaultValuesOk() (*[]DefaultLevelValue, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *SetDefaultLevelsRequest) SetDefaultValues(v []DefaultLevelValue)`

SetDefaultValues sets DefaultValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


