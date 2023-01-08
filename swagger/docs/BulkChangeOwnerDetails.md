# BulkChangeOwnerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutofixName** | **bool** | Whether the name is fixed automatically if it&#39;s duplicated after changing owner. | 
**NewOwner** | **string** | The account id of the new owner. | 

## Methods

### NewBulkChangeOwnerDetails

`func NewBulkChangeOwnerDetails(autofixName bool, newOwner string, ) *BulkChangeOwnerDetails`

NewBulkChangeOwnerDetails instantiates a new BulkChangeOwnerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkChangeOwnerDetailsWithDefaults

`func NewBulkChangeOwnerDetailsWithDefaults() *BulkChangeOwnerDetails`

NewBulkChangeOwnerDetailsWithDefaults instantiates a new BulkChangeOwnerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutofixName

`func (o *BulkChangeOwnerDetails) GetAutofixName() bool`

GetAutofixName returns the AutofixName field if non-nil, zero value otherwise.

### GetAutofixNameOk

`func (o *BulkChangeOwnerDetails) GetAutofixNameOk() (*bool, bool)`

GetAutofixNameOk returns a tuple with the AutofixName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofixName

`func (o *BulkChangeOwnerDetails) SetAutofixName(v bool)`

SetAutofixName sets AutofixName field to given value.


### GetNewOwner

`func (o *BulkChangeOwnerDetails) GetNewOwner() string`

GetNewOwner returns the NewOwner field if non-nil, zero value otherwise.

### GetNewOwnerOk

`func (o *BulkChangeOwnerDetails) GetNewOwnerOk() (*string, bool)`

GetNewOwnerOk returns a tuple with the NewOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwner

`func (o *BulkChangeOwnerDetails) SetNewOwner(v string)`

SetNewOwner sets NewOwner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


