# AuditRecordBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedItems** | Pointer to [**[]AssociatedItemBean**](AssociatedItemBean.md) | The list of items associated with the changed record. | [optional] [readonly] 
**AuthorKey** | Pointer to **string** | Deprecated, use &#x60;authorAccountId&#x60; instead. The key of the user who created the audit record. | [optional] [readonly] 
**Category** | Pointer to **string** | The category of the audit record. For a list of these categories, see the help article [Auditing in Jira applications](https://confluence.atlassian.com/x/noXKM). | [optional] [readonly] 
**ChangedValues** | Pointer to [**[]ChangedValueBean**](ChangedValueBean.md) | The list of values changed in the record event. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date and time on which the audit record was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the audit record. | [optional] [readonly] 
**EventSource** | Pointer to **string** | The event the audit record originated from. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the audit record. | [optional] [readonly] 
**ObjectItem** | Pointer to [**AssociatedItemBean**](AssociatedItemBean.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | The URL of the computer where the creation of the audit record was initiated. | [optional] [readonly] 
**Summary** | Pointer to **string** | The summary of the audit record. | [optional] [readonly] 

## Methods

### NewAuditRecordBean

`func NewAuditRecordBean() *AuditRecordBean`

NewAuditRecordBean instantiates a new AuditRecordBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditRecordBeanWithDefaults

`func NewAuditRecordBeanWithDefaults() *AuditRecordBean`

NewAuditRecordBeanWithDefaults instantiates a new AuditRecordBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedItems

`func (o *AuditRecordBean) GetAssociatedItems() []AssociatedItemBean`

GetAssociatedItems returns the AssociatedItems field if non-nil, zero value otherwise.

### GetAssociatedItemsOk

`func (o *AuditRecordBean) GetAssociatedItemsOk() (*[]AssociatedItemBean, bool)`

GetAssociatedItemsOk returns a tuple with the AssociatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedItems

`func (o *AuditRecordBean) SetAssociatedItems(v []AssociatedItemBean)`

SetAssociatedItems sets AssociatedItems field to given value.

### HasAssociatedItems

`func (o *AuditRecordBean) HasAssociatedItems() bool`

HasAssociatedItems returns a boolean if a field has been set.

### GetAuthorKey

`func (o *AuditRecordBean) GetAuthorKey() string`

GetAuthorKey returns the AuthorKey field if non-nil, zero value otherwise.

### GetAuthorKeyOk

`func (o *AuditRecordBean) GetAuthorKeyOk() (*string, bool)`

GetAuthorKeyOk returns a tuple with the AuthorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorKey

`func (o *AuditRecordBean) SetAuthorKey(v string)`

SetAuthorKey sets AuthorKey field to given value.

### HasAuthorKey

`func (o *AuditRecordBean) HasAuthorKey() bool`

HasAuthorKey returns a boolean if a field has been set.

### GetCategory

`func (o *AuditRecordBean) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditRecordBean) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditRecordBean) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AuditRecordBean) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChangedValues

`func (o *AuditRecordBean) GetChangedValues() []ChangedValueBean`

GetChangedValues returns the ChangedValues field if non-nil, zero value otherwise.

### GetChangedValuesOk

`func (o *AuditRecordBean) GetChangedValuesOk() (*[]ChangedValueBean, bool)`

GetChangedValuesOk returns a tuple with the ChangedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedValues

`func (o *AuditRecordBean) SetChangedValues(v []ChangedValueBean)`

SetChangedValues sets ChangedValues field to given value.

### HasChangedValues

`func (o *AuditRecordBean) HasChangedValues() bool`

HasChangedValues returns a boolean if a field has been set.

### GetCreated

`func (o *AuditRecordBean) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuditRecordBean) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuditRecordBean) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuditRecordBean) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *AuditRecordBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditRecordBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditRecordBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditRecordBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventSource

`func (o *AuditRecordBean) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *AuditRecordBean) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *AuditRecordBean) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *AuditRecordBean) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *AuditRecordBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditRecordBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditRecordBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuditRecordBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectItem

`func (o *AuditRecordBean) GetObjectItem() AssociatedItemBean`

GetObjectItem returns the ObjectItem field if non-nil, zero value otherwise.

### GetObjectItemOk

`func (o *AuditRecordBean) GetObjectItemOk() (*AssociatedItemBean, bool)`

GetObjectItemOk returns a tuple with the ObjectItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectItem

`func (o *AuditRecordBean) SetObjectItem(v AssociatedItemBean)`

SetObjectItem sets ObjectItem field to given value.

### HasObjectItem

`func (o *AuditRecordBean) HasObjectItem() bool`

HasObjectItem returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *AuditRecordBean) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *AuditRecordBean) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *AuditRecordBean) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *AuditRecordBean) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetSummary

`func (o *AuditRecordBean) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AuditRecordBean) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AuditRecordBean) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AuditRecordBean) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


