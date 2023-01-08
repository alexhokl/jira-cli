# RegisteredWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedWebhookId** | Pointer to **int64** | The ID of the webhook. Returned if the webhook is created. | [optional] 
**Errors** | Pointer to **[]string** | Error messages specifying why the webhook creation failed. | [optional] 

## Methods

### NewRegisteredWebhook

`func NewRegisteredWebhook() *RegisteredWebhook`

NewRegisteredWebhook instantiates a new RegisteredWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredWebhookWithDefaults

`func NewRegisteredWebhookWithDefaults() *RegisteredWebhook`

NewRegisteredWebhookWithDefaults instantiates a new RegisteredWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedWebhookId

`func (o *RegisteredWebhook) GetCreatedWebhookId() int64`

GetCreatedWebhookId returns the CreatedWebhookId field if non-nil, zero value otherwise.

### GetCreatedWebhookIdOk

`func (o *RegisteredWebhook) GetCreatedWebhookIdOk() (*int64, bool)`

GetCreatedWebhookIdOk returns a tuple with the CreatedWebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWebhookId

`func (o *RegisteredWebhook) SetCreatedWebhookId(v int64)`

SetCreatedWebhookId sets CreatedWebhookId field to given value.

### HasCreatedWebhookId

`func (o *RegisteredWebhook) HasCreatedWebhookId() bool`

HasCreatedWebhookId returns a boolean if a field has been set.

### GetErrors

`func (o *RegisteredWebhook) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RegisteredWebhook) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RegisteredWebhook) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RegisteredWebhook) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


