# CreateIssueSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The issue source type. This must be \&quot;Board\&quot;, \&quot;Project\&quot; or \&quot;Filter\&quot;. | 
**Value** | **int64** | The issue source value. This must be a board ID if the type is \&quot;Board\&quot;, a project ID if the type is \&quot;Project\&quot; or a filter ID if the type is \&quot;Filter\&quot;. | 

## Methods

### NewCreateIssueSourceRequest

`func NewCreateIssueSourceRequest(type_ string, value int64, ) *CreateIssueSourceRequest`

NewCreateIssueSourceRequest instantiates a new CreateIssueSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueSourceRequestWithDefaults

`func NewCreateIssueSourceRequestWithDefaults() *CreateIssueSourceRequest`

NewCreateIssueSourceRequestWithDefaults instantiates a new CreateIssueSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateIssueSourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIssueSourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIssueSourceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CreateIssueSourceRequest) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateIssueSourceRequest) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateIssueSourceRequest) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


