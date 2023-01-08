# FailedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The webhook body. | [optional] 
**FailureTime** | **int64** | The time the webhook was added to the list of failed webhooks (that is, the time of the last failed retry). | 
**Id** | **string** | The webhook ID, as sent in the &#x60;X-Atlassian-Webhook-Identifier&#x60; header with the webhook. | 
**Url** | **string** | The original webhook destination. | 

## Methods

### NewFailedWebhook

`func NewFailedWebhook(failureTime int64, id string, url string, ) *FailedWebhook`

NewFailedWebhook instantiates a new FailedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedWebhookWithDefaults

`func NewFailedWebhookWithDefaults() *FailedWebhook`

NewFailedWebhookWithDefaults instantiates a new FailedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *FailedWebhook) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FailedWebhook) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FailedWebhook) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *FailedWebhook) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFailureTime

`func (o *FailedWebhook) GetFailureTime() int64`

GetFailureTime returns the FailureTime field if non-nil, zero value otherwise.

### GetFailureTimeOk

`func (o *FailedWebhook) GetFailureTimeOk() (*int64, bool)`

GetFailureTimeOk returns a tuple with the FailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureTime

`func (o *FailedWebhook) SetFailureTime(v int64)`

SetFailureTime sets FailureTime field to given value.


### GetId

`func (o *FailedWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailedWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailedWebhook) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *FailedWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FailedWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FailedWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


