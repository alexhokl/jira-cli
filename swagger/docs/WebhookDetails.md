# WebhookDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** | The Jira events that trigger the webhook. | 
**FieldIdsFilter** | Pointer to **[]string** | A list of field IDs. When the issue changelog contains any of the fields, the webhook &#x60;jira:issue_updated&#x60; is sent. If this parameter is not present, the app is notified about all field updates. | [optional] 
**IssuePropertyKeysFilter** | Pointer to **[]string** | A list of issue property keys. A change of those issue properties triggers the &#x60;issue_property_set&#x60; or &#x60;issue_property_deleted&#x60; webhooks. If this parameter is not present, the app is notified about all issue property updates. | [optional] 
**JqlFilter** | **string** | The JQL filter that specifies which issues the webhook is sent for. Only a subset of JQL can be used. The supported elements are:   *  Fields: &#x60;issueKey&#x60;, &#x60;project&#x60;, &#x60;issuetype&#x60;, &#x60;status&#x60;, &#x60;assignee&#x60;, &#x60;reporter&#x60;, &#x60;issue.property&#x60;, and &#x60;cf[id]&#x60;. For custom fields (&#x60;cf[id]&#x60;), only the epic label custom field is supported.\&quot;.  *  Operators: &#x60;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;IN&#x60;, and &#x60;NOT IN&#x60;. | 

## Methods

### NewWebhookDetails

`func NewWebhookDetails(events []string, jqlFilter string, ) *WebhookDetails`

NewWebhookDetails instantiates a new WebhookDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDetailsWithDefaults

`func NewWebhookDetailsWithDefaults() *WebhookDetails`

NewWebhookDetailsWithDefaults instantiates a new WebhookDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhookDetails) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookDetails) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookDetails) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetFieldIdsFilter

`func (o *WebhookDetails) GetFieldIdsFilter() []string`

GetFieldIdsFilter returns the FieldIdsFilter field if non-nil, zero value otherwise.

### GetFieldIdsFilterOk

`func (o *WebhookDetails) GetFieldIdsFilterOk() (*[]string, bool)`

GetFieldIdsFilterOk returns a tuple with the FieldIdsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdsFilter

`func (o *WebhookDetails) SetFieldIdsFilter(v []string)`

SetFieldIdsFilter sets FieldIdsFilter field to given value.

### HasFieldIdsFilter

`func (o *WebhookDetails) HasFieldIdsFilter() bool`

HasFieldIdsFilter returns a boolean if a field has been set.

### GetIssuePropertyKeysFilter

`func (o *WebhookDetails) GetIssuePropertyKeysFilter() []string`

GetIssuePropertyKeysFilter returns the IssuePropertyKeysFilter field if non-nil, zero value otherwise.

### GetIssuePropertyKeysFilterOk

`func (o *WebhookDetails) GetIssuePropertyKeysFilterOk() (*[]string, bool)`

GetIssuePropertyKeysFilterOk returns a tuple with the IssuePropertyKeysFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuePropertyKeysFilter

`func (o *WebhookDetails) SetIssuePropertyKeysFilter(v []string)`

SetIssuePropertyKeysFilter sets IssuePropertyKeysFilter field to given value.

### HasIssuePropertyKeysFilter

`func (o *WebhookDetails) HasIssuePropertyKeysFilter() bool`

HasIssuePropertyKeysFilter returns a boolean if a field has been set.

### GetJqlFilter

`func (o *WebhookDetails) GetJqlFilter() string`

GetJqlFilter returns the JqlFilter field if non-nil, zero value otherwise.

### GetJqlFilterOk

`func (o *WebhookDetails) GetJqlFilterOk() (*string, bool)`

GetJqlFilterOk returns a tuple with the JqlFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqlFilter

`func (o *WebhookDetails) SetJqlFilter(v string)`

SetJqlFilter sets JqlFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


