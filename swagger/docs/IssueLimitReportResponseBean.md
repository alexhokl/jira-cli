# IssueLimitReportResponseBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuesApproachingLimit** | Pointer to **map[string]map[string]int64** | A list of ids of issues approaching the limit and their field count | [optional] 
**IssuesBreachingLimit** | Pointer to **map[string]map[string]int64** | A list of ids of issues breaching the limit and their field count | [optional] 
**Limits** | Pointer to **map[string]int32** | The fields and their defined limits | [optional] 

## Methods

### NewIssueLimitReportResponseBean

`func NewIssueLimitReportResponseBean() *IssueLimitReportResponseBean`

NewIssueLimitReportResponseBean instantiates a new IssueLimitReportResponseBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLimitReportResponseBeanWithDefaults

`func NewIssueLimitReportResponseBeanWithDefaults() *IssueLimitReportResponseBean`

NewIssueLimitReportResponseBeanWithDefaults instantiates a new IssueLimitReportResponseBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuesApproachingLimit

`func (o *IssueLimitReportResponseBean) GetIssuesApproachingLimit() map[string]map[string]int64`

GetIssuesApproachingLimit returns the IssuesApproachingLimit field if non-nil, zero value otherwise.

### GetIssuesApproachingLimitOk

`func (o *IssueLimitReportResponseBean) GetIssuesApproachingLimitOk() (*map[string]map[string]int64, bool)`

GetIssuesApproachingLimitOk returns a tuple with the IssuesApproachingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesApproachingLimit

`func (o *IssueLimitReportResponseBean) SetIssuesApproachingLimit(v map[string]map[string]int64)`

SetIssuesApproachingLimit sets IssuesApproachingLimit field to given value.

### HasIssuesApproachingLimit

`func (o *IssueLimitReportResponseBean) HasIssuesApproachingLimit() bool`

HasIssuesApproachingLimit returns a boolean if a field has been set.

### GetIssuesBreachingLimit

`func (o *IssueLimitReportResponseBean) GetIssuesBreachingLimit() map[string]map[string]int64`

GetIssuesBreachingLimit returns the IssuesBreachingLimit field if non-nil, zero value otherwise.

### GetIssuesBreachingLimitOk

`func (o *IssueLimitReportResponseBean) GetIssuesBreachingLimitOk() (*map[string]map[string]int64, bool)`

GetIssuesBreachingLimitOk returns a tuple with the IssuesBreachingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesBreachingLimit

`func (o *IssueLimitReportResponseBean) SetIssuesBreachingLimit(v map[string]map[string]int64)`

SetIssuesBreachingLimit sets IssuesBreachingLimit field to given value.

### HasIssuesBreachingLimit

`func (o *IssueLimitReportResponseBean) HasIssuesBreachingLimit() bool`

HasIssuesBreachingLimit returns a boolean if a field has been set.

### GetLimits

`func (o *IssueLimitReportResponseBean) GetLimits() map[string]int32`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IssueLimitReportResponseBean) GetLimitsOk() (*map[string]int32, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IssueLimitReportResponseBean) SetLimits(v map[string]int32)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IssueLimitReportResponseBean) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


