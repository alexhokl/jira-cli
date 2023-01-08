# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** | The Jira events that trigger the webhook. | 
**ExpirationDate** | Pointer to **int64** | The date after which the webhook is no longer sent. Use [Extend webhook life](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-webhooks/#api-rest-api-3-webhook-refresh-put) to extend the date. | [optional] [readonly] 
**FieldIdsFilter** | Pointer to **[]string** | A list of field IDs. When the issue changelog contains any of the fields, the webhook &#x60;jira:issue_updated&#x60; is sent. If this parameter is not present, the app is notified about all field updates. | [optional] 
**Id** | **int64** | The ID of the webhook. | 
**IssuePropertyKeysFilter** | Pointer to **[]string** | A list of issue property keys. A change of those issue properties triggers the &#x60;issue_property_set&#x60; or &#x60;issue_property_deleted&#x60; webhooks. If this parameter is not present, the app is notified about all issue property updates. | [optional] 
**JqlFilter** | **string** | The JQL filter that specifies which issues the webhook is sent for. | 
**Url** | **string** | The URL that specifies where the webhooks are sent. | 

## Methods

### NewWebhook

`func NewWebhook(events []string, id int64, jqlFilter string, url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Webhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetExpirationDate

`func (o *Webhook) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Webhook) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Webhook) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Webhook) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetFieldIdsFilter

`func (o *Webhook) GetFieldIdsFilter() []string`

GetFieldIdsFilter returns the FieldIdsFilter field if non-nil, zero value otherwise.

### GetFieldIdsFilterOk

`func (o *Webhook) GetFieldIdsFilterOk() (*[]string, bool)`

GetFieldIdsFilterOk returns a tuple with the FieldIdsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdsFilter

`func (o *Webhook) SetFieldIdsFilter(v []string)`

SetFieldIdsFilter sets FieldIdsFilter field to given value.

### HasFieldIdsFilter

`func (o *Webhook) HasFieldIdsFilter() bool`

HasFieldIdsFilter returns a boolean if a field has been set.

### GetId

`func (o *Webhook) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v int64)`

SetId sets Id field to given value.


### GetIssuePropertyKeysFilter

`func (o *Webhook) GetIssuePropertyKeysFilter() []string`

GetIssuePropertyKeysFilter returns the IssuePropertyKeysFilter field if non-nil, zero value otherwise.

### GetIssuePropertyKeysFilterOk

`func (o *Webhook) GetIssuePropertyKeysFilterOk() (*[]string, bool)`

GetIssuePropertyKeysFilterOk returns a tuple with the IssuePropertyKeysFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuePropertyKeysFilter

`func (o *Webhook) SetIssuePropertyKeysFilter(v []string)`

SetIssuePropertyKeysFilter sets IssuePropertyKeysFilter field to given value.

### HasIssuePropertyKeysFilter

`func (o *Webhook) HasIssuePropertyKeysFilter() bool`

HasIssuePropertyKeysFilter returns a boolean if a field has been set.

### GetJqlFilter

`func (o *Webhook) GetJqlFilter() string`

GetJqlFilter returns the JqlFilter field if non-nil, zero value otherwise.

### GetJqlFilterOk

`func (o *Webhook) GetJqlFilterOk() (*string, bool)`

GetJqlFilterOk returns a tuple with the JqlFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqlFilter

`func (o *Webhook) SetJqlFilter(v string)`

SetJqlFilter sets JqlFilter field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


