# IncidentSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The severity level of the Incident with P1 being the highest and P5 being the lowest | 

## Methods

### NewIncidentSeverity

`func NewIncidentSeverity(level string, ) *IncidentSeverity`

NewIncidentSeverity instantiates a new IncidentSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSeverityWithDefaults

`func NewIncidentSeverityWithDefaults() *IncidentSeverity`

NewIncidentSeverityWithDefaults instantiates a new IncidentSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *IncidentSeverity) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IncidentSeverity) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IncidentSeverity) SetLevel(v string)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


