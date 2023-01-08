# VersionIssueCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldUsage** | Pointer to [**[]VersionUsageInCustomField**](VersionUsageInCustomField.md) | List of custom fields using the version. | [optional] [readonly] 
**IssueCountWithCustomFieldsShowingVersion** | Pointer to **int64** | Count of issues where a version custom field is set to the version. | [optional] [readonly] 
**IssuesAffectedCount** | Pointer to **int64** | Count of issues where the &#x60;affectedVersion&#x60; is set to the version. | [optional] [readonly] 
**IssuesFixedCount** | Pointer to **int64** | Count of issues where the &#x60;fixVersion&#x60; is set to the version. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these count details. | [optional] [readonly] 

## Methods

### NewVersionIssueCounts

`func NewVersionIssueCounts() *VersionIssueCounts`

NewVersionIssueCounts instantiates a new VersionIssueCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionIssueCountsWithDefaults

`func NewVersionIssueCountsWithDefaults() *VersionIssueCounts`

NewVersionIssueCountsWithDefaults instantiates a new VersionIssueCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldUsage

`func (o *VersionIssueCounts) GetCustomFieldUsage() []VersionUsageInCustomField`

GetCustomFieldUsage returns the CustomFieldUsage field if non-nil, zero value otherwise.

### GetCustomFieldUsageOk

`func (o *VersionIssueCounts) GetCustomFieldUsageOk() (*[]VersionUsageInCustomField, bool)`

GetCustomFieldUsageOk returns a tuple with the CustomFieldUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldUsage

`func (o *VersionIssueCounts) SetCustomFieldUsage(v []VersionUsageInCustomField)`

SetCustomFieldUsage sets CustomFieldUsage field to given value.

### HasCustomFieldUsage

`func (o *VersionIssueCounts) HasCustomFieldUsage() bool`

HasCustomFieldUsage returns a boolean if a field has been set.

### GetIssueCountWithCustomFieldsShowingVersion

`func (o *VersionIssueCounts) GetIssueCountWithCustomFieldsShowingVersion() int64`

GetIssueCountWithCustomFieldsShowingVersion returns the IssueCountWithCustomFieldsShowingVersion field if non-nil, zero value otherwise.

### GetIssueCountWithCustomFieldsShowingVersionOk

`func (o *VersionIssueCounts) GetIssueCountWithCustomFieldsShowingVersionOk() (*int64, bool)`

GetIssueCountWithCustomFieldsShowingVersionOk returns a tuple with the IssueCountWithCustomFieldsShowingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCountWithCustomFieldsShowingVersion

`func (o *VersionIssueCounts) SetIssueCountWithCustomFieldsShowingVersion(v int64)`

SetIssueCountWithCustomFieldsShowingVersion sets IssueCountWithCustomFieldsShowingVersion field to given value.

### HasIssueCountWithCustomFieldsShowingVersion

`func (o *VersionIssueCounts) HasIssueCountWithCustomFieldsShowingVersion() bool`

HasIssueCountWithCustomFieldsShowingVersion returns a boolean if a field has been set.

### GetIssuesAffectedCount

`func (o *VersionIssueCounts) GetIssuesAffectedCount() int64`

GetIssuesAffectedCount returns the IssuesAffectedCount field if non-nil, zero value otherwise.

### GetIssuesAffectedCountOk

`func (o *VersionIssueCounts) GetIssuesAffectedCountOk() (*int64, bool)`

GetIssuesAffectedCountOk returns a tuple with the IssuesAffectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesAffectedCount

`func (o *VersionIssueCounts) SetIssuesAffectedCount(v int64)`

SetIssuesAffectedCount sets IssuesAffectedCount field to given value.

### HasIssuesAffectedCount

`func (o *VersionIssueCounts) HasIssuesAffectedCount() bool`

HasIssuesAffectedCount returns a boolean if a field has been set.

### GetIssuesFixedCount

`func (o *VersionIssueCounts) GetIssuesFixedCount() int64`

GetIssuesFixedCount returns the IssuesFixedCount field if non-nil, zero value otherwise.

### GetIssuesFixedCountOk

`func (o *VersionIssueCounts) GetIssuesFixedCountOk() (*int64, bool)`

GetIssuesFixedCountOk returns a tuple with the IssuesFixedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesFixedCount

`func (o *VersionIssueCounts) SetIssuesFixedCount(v int64)`

SetIssuesFixedCount sets IssuesFixedCount field to given value.

### HasIssuesFixedCount

`func (o *VersionIssueCounts) HasIssuesFixedCount() bool`

HasIssuesFixedCount returns a boolean if a field has been set.

### GetSelf

`func (o *VersionIssueCounts) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *VersionIssueCounts) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *VersionIssueCounts) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *VersionIssueCounts) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


