# JQLCountRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jql** | Pointer to **string** | A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. For performance reasons, this parameter requires a bounded query. A bounded query is a query with a search restriction. | [optional] 

## Methods

### NewJQLCountRequestBean

`func NewJQLCountRequestBean() *JQLCountRequestBean`

NewJQLCountRequestBean instantiates a new JQLCountRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJQLCountRequestBeanWithDefaults

`func NewJQLCountRequestBeanWithDefaults() *JQLCountRequestBean`

NewJQLCountRequestBeanWithDefaults instantiates a new JQLCountRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJql

`func (o *JQLCountRequestBean) GetJql() string`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *JQLCountRequestBean) GetJqlOk() (*string, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *JQLCountRequestBean) SetJql(v string)`

SetJql sets Jql field to given value.

### HasJql

`func (o *JQLCountRequestBean) HasJql() bool`

HasJql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


