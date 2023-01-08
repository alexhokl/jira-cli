# StatusMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeId** | **string** | The ID of the issue type. | 
**NewStatusId** | **string** | The ID of the new status. | 
**StatusId** | **string** | The ID of the status. | 

## Methods

### NewStatusMapping

`func NewStatusMapping(issueTypeId string, newStatusId string, statusId string, ) *StatusMapping`

NewStatusMapping instantiates a new StatusMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMappingWithDefaults

`func NewStatusMappingWithDefaults() *StatusMapping`

NewStatusMappingWithDefaults instantiates a new StatusMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeId

`func (o *StatusMapping) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *StatusMapping) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *StatusMapping) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.


### GetNewStatusId

`func (o *StatusMapping) GetNewStatusId() string`

GetNewStatusId returns the NewStatusId field if non-nil, zero value otherwise.

### GetNewStatusIdOk

`func (o *StatusMapping) GetNewStatusIdOk() (*string, bool)`

GetNewStatusIdOk returns a tuple with the NewStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatusId

`func (o *StatusMapping) SetNewStatusId(v string)`

SetNewStatusId sets NewStatusId field to given value.


### GetStatusId

`func (o *StatusMapping) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *StatusMapping) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *StatusMapping) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


