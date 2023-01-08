# IssueLinkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the issue link type and is used as follows:   *  In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is the type of issue link. Required on create when &#x60;name&#x60; isn&#39;t provided. Otherwise, read only.  *  In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is read only. | [optional] 
**Inward** | Pointer to **string** | The description of the issue link type inward link and is used as follows:   *  In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is read only.  *  In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only. | [optional] 
**Name** | Pointer to **string** | The name of the issue link type and is used as follows:   *  In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is the type of issue link. Required on create when &#x60;id&#x60; isn&#39;t provided. Otherwise, read only.  *  In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only. | [optional] 
**Outward** | Pointer to **string** | The description of the issue link type outward link and is used as follows:   *  In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is read only.  *  In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only. | [optional] 
**Self** | Pointer to **string** | The URL of the issue link type. Read only. | [optional] [readonly] 

## Methods

### NewIssueLinkType

`func NewIssueLinkType() *IssueLinkType`

NewIssueLinkType instantiates a new IssueLinkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLinkTypeWithDefaults

`func NewIssueLinkTypeWithDefaults() *IssueLinkType`

NewIssueLinkTypeWithDefaults instantiates a new IssueLinkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueLinkType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueLinkType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueLinkType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueLinkType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInward

`func (o *IssueLinkType) GetInward() string`

GetInward returns the Inward field if non-nil, zero value otherwise.

### GetInwardOk

`func (o *IssueLinkType) GetInwardOk() (*string, bool)`

GetInwardOk returns a tuple with the Inward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInward

`func (o *IssueLinkType) SetInward(v string)`

SetInward sets Inward field to given value.

### HasInward

`func (o *IssueLinkType) HasInward() bool`

HasInward returns a boolean if a field has been set.

### GetName

`func (o *IssueLinkType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueLinkType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueLinkType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueLinkType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutward

`func (o *IssueLinkType) GetOutward() string`

GetOutward returns the Outward field if non-nil, zero value otherwise.

### GetOutwardOk

`func (o *IssueLinkType) GetOutwardOk() (*string, bool)`

GetOutwardOk returns a tuple with the Outward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutward

`func (o *IssueLinkType) SetOutward(v string)`

SetOutward sets Outward field to given value.

### HasOutward

`func (o *IssueLinkType) HasOutward() bool`

HasOutward returns a boolean if a field has been set.

### GetSelf

`func (o *IssueLinkType) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueLinkType) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueLinkType) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IssueLinkType) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


