# LicenseMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of a specific license metric. | [optional] 
**Value** | Pointer to **string** | The calculated value of a licence metric linked to the key. An example licence metric is the approximate number of user accounts. | [optional] 

## Methods

### NewLicenseMetric

`func NewLicenseMetric() *LicenseMetric`

NewLicenseMetric instantiates a new LicenseMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseMetricWithDefaults

`func NewLicenseMetricWithDefaults() *LicenseMetric`

NewLicenseMetricWithDefaults instantiates a new LicenseMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LicenseMetric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LicenseMetric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LicenseMetric) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LicenseMetric) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *LicenseMetric) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LicenseMetric) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LicenseMetric) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LicenseMetric) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


