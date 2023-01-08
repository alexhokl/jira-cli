# CreateDateFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCustomFieldId** | Pointer to **int64** | A date custom field ID. This is required if the type is \&quot;DateCustomField\&quot;. | [optional] 
**Type** | **string** | The date field type. This must be \&quot;DueDate\&quot;, \&quot;TargetStartDate\&quot;, \&quot;TargetEndDate\&quot; or \&quot;DateCustomField\&quot;. | 

## Methods

### NewCreateDateFieldRequest

`func NewCreateDateFieldRequest(type_ string, ) *CreateDateFieldRequest`

NewCreateDateFieldRequest instantiates a new CreateDateFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDateFieldRequestWithDefaults

`func NewCreateDateFieldRequestWithDefaults() *CreateDateFieldRequest`

NewCreateDateFieldRequestWithDefaults instantiates a new CreateDateFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCustomFieldId

`func (o *CreateDateFieldRequest) GetDateCustomFieldId() int64`

GetDateCustomFieldId returns the DateCustomFieldId field if non-nil, zero value otherwise.

### GetDateCustomFieldIdOk

`func (o *CreateDateFieldRequest) GetDateCustomFieldIdOk() (*int64, bool)`

GetDateCustomFieldIdOk returns a tuple with the DateCustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomFieldId

`func (o *CreateDateFieldRequest) SetDateCustomFieldId(v int64)`

SetDateCustomFieldId sets DateCustomFieldId field to given value.

### HasDateCustomFieldId

`func (o *CreateDateFieldRequest) HasDateCustomFieldId() bool`

HasDateCustomFieldId returns a boolean if a field has been set.

### GetType

`func (o *CreateDateFieldRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDateFieldRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDateFieldRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


