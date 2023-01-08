# WebhookRegistrationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL that specifies where to send the webhooks. This URL must use the same base URL as the Connect app. Only a single URL per app is allowed to be registered. | 
**Webhooks** | [**[]WebhookDetails**](WebhookDetails.md) | A list of webhooks. | 

## Methods

### NewWebhookRegistrationDetails

`func NewWebhookRegistrationDetails(url string, webhooks []WebhookDetails, ) *WebhookRegistrationDetails`

NewWebhookRegistrationDetails instantiates a new WebhookRegistrationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRegistrationDetailsWithDefaults

`func NewWebhookRegistrationDetailsWithDefaults() *WebhookRegistrationDetails`

NewWebhookRegistrationDetailsWithDefaults instantiates a new WebhookRegistrationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookRegistrationDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRegistrationDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRegistrationDetails) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebhooks

`func (o *WebhookRegistrationDetails) GetWebhooks() []WebhookDetails`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookRegistrationDetails) GetWebhooksOk() (*[]WebhookDetails, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookRegistrationDetails) SetWebhooks(v []WebhookDetails)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


