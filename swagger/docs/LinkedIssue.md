# LinkedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**Fields**](Fields.md) | The fields associated with the issue. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of an issue. Required if &#x60;key&#x60; isn&#39;t provided. | [optional] 
**Key** | Pointer to **string** | The key of an issue. Required if &#x60;id&#x60; isn&#39;t provided. | [optional] 
**Self** | Pointer to **string** | The URL of the issue. | [optional] [readonly] 

## Methods

### NewLinkedIssue

`func NewLinkedIssue() *LinkedIssue`

NewLinkedIssue instantiates a new LinkedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedIssueWithDefaults

`func NewLinkedIssueWithDefaults() *LinkedIssue`

NewLinkedIssueWithDefaults instantiates a new LinkedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *LinkedIssue) GetFields() Fields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LinkedIssue) GetFieldsOk() (*Fields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LinkedIssue) SetFields(v Fields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LinkedIssue) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *LinkedIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedIssue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *LinkedIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LinkedIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LinkedIssue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LinkedIssue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSelf

`func (o *LinkedIssue) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LinkedIssue) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LinkedIssue) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LinkedIssue) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


