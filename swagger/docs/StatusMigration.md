# StatusMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewStatusReference** | **string** | The new status ID. | 
**OldStatusReference** | **string** | The old status ID. | 

## Methods

### NewStatusMigration

`func NewStatusMigration(newStatusReference string, oldStatusReference string, ) *StatusMigration`

NewStatusMigration instantiates a new StatusMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMigrationWithDefaults

`func NewStatusMigrationWithDefaults() *StatusMigration`

NewStatusMigrationWithDefaults instantiates a new StatusMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewStatusReference

`func (o *StatusMigration) GetNewStatusReference() string`

GetNewStatusReference returns the NewStatusReference field if non-nil, zero value otherwise.

### GetNewStatusReferenceOk

`func (o *StatusMigration) GetNewStatusReferenceOk() (*string, bool)`

GetNewStatusReferenceOk returns a tuple with the NewStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatusReference

`func (o *StatusMigration) SetNewStatusReference(v string)`

SetNewStatusReference sets NewStatusReference field to given value.


### GetOldStatusReference

`func (o *StatusMigration) GetOldStatusReference() string`

GetOldStatusReference returns the OldStatusReference field if non-nil, zero value otherwise.

### GetOldStatusReferenceOk

`func (o *StatusMigration) GetOldStatusReferenceOk() (*string, bool)`

GetOldStatusReferenceOk returns a tuple with the OldStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStatusReference

`func (o *StatusMigration) SetOldStatusReference(v string)`

SetOldStatusReference sets OldStatusReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


