# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]LicensedApplication**](LicensedApplication.md) | The applications under this license. | [readonly] 

## Methods

### NewLicense

`func NewLicense(applications []LicensedApplication, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *License) GetApplications() []LicensedApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *License) GetApplicationsOk() (*[]LicensedApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *License) SetApplications(v []LicensedApplication)`

SetApplications sets Applications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


