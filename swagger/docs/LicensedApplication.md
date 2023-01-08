# LicensedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the application. | [readonly] 
**Plan** | **string** | The licensing plan. | [readonly] 

## Methods

### NewLicensedApplication

`func NewLicensedApplication(id string, plan string, ) *LicensedApplication`

NewLicensedApplication instantiates a new LicensedApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensedApplicationWithDefaults

`func NewLicensedApplicationWithDefaults() *LicensedApplication`

NewLicensedApplicationWithDefaults instantiates a new LicensedApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicensedApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensedApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensedApplication) SetId(v string)`

SetId sets Id field to given value.


### GetPlan

`func (o *LicensedApplication) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LicensedApplication) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LicensedApplication) SetPlan(v string)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


