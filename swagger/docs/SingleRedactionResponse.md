# SingleRedactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | An unique id for the redaction request | 
**Successful** | **bool** | Indicates if redaction was success/failure | 

## Methods

### NewSingleRedactionResponse

`func NewSingleRedactionResponse(externalId string, successful bool, ) *SingleRedactionResponse`

NewSingleRedactionResponse instantiates a new SingleRedactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRedactionResponseWithDefaults

`func NewSingleRedactionResponseWithDefaults() *SingleRedactionResponse`

NewSingleRedactionResponseWithDefaults instantiates a new SingleRedactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *SingleRedactionResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SingleRedactionResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SingleRedactionResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetSuccessful

`func (o *SingleRedactionResponse) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *SingleRedactionResponse) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *SingleRedactionResponse) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


