# FeatureFlagData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | The FeatureFlagData schema version used for this flag data.   Placeholder to support potential schema changes in the future.  | [optional] [default to "1.0"]
**Id** | **string** | The identifier for the Feature Flag. Must be unique for a given Provider.  | 
**Key** | **string** | The identifier that users would use to reference the Feature Flag in their source code etc.  Will be made available via the UI for users to copy into their source code etc.  | 
**UpdateSequenceId** | **int64** | An ID used to apply an ordering to updates for this Feature Flag in the case of out-of-order receipt of update requests.  This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the Provider system, but other alternatives are valid (e.g. a Provider could store a counter against each Feature Flag and increment that on each update to Jira).  Updates for a Feature Flag that are received with an updateSqeuenceId lower than what is currently stored will be ignored.  | 
**DisplayName** | Pointer to **string** | The human-readable name for the Feature Flag. Will be shown in the UI.  If not provided, will use the ID for display.  | [optional] 
**IssueKeys** | Pointer to **[]string** | The Jira issue keys to associate the Feature Flag information with.  | [optional] 
**Associations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | The Jira issue keys or IDs to associate the feature flag with. | [optional] 
**Summary** | [**FeatureFlagSummary**](FeatureFlagSummary.md) |  | 
**Details** | [**[]FeatureFlagDetails**](FeatureFlagDetails.md) | Detail information for this Feature Flag.  This may be information for each environment the Feature Flag is defined in or a selection of environments made by the user, as appropriate.  | 

## Methods

### NewFeatureFlagData

`func NewFeatureFlagData(id string, key string, updateSequenceId int64, summary FeatureFlagSummary, details []FeatureFlagDetails, ) *FeatureFlagData`

NewFeatureFlagData instantiates a new FeatureFlagData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagDataWithDefaults

`func NewFeatureFlagDataWithDefaults() *FeatureFlagData`

NewFeatureFlagDataWithDefaults instantiates a new FeatureFlagData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *FeatureFlagData) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *FeatureFlagData) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *FeatureFlagData) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *FeatureFlagData) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetId

`func (o *FeatureFlagData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureFlagData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureFlagData) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *FeatureFlagData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureFlagData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureFlagData) SetKey(v string)`

SetKey sets Key field to given value.


### GetUpdateSequenceId

`func (o *FeatureFlagData) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *FeatureFlagData) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *FeatureFlagData) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.


### GetDisplayName

`func (o *FeatureFlagData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FeatureFlagData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FeatureFlagData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FeatureFlagData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIssueKeys

`func (o *FeatureFlagData) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *FeatureFlagData) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *FeatureFlagData) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *FeatureFlagData) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *FeatureFlagData) GetAssociations() []IssueIdOrKeysAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *FeatureFlagData) GetAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *FeatureFlagData) SetAssociations(v []IssueIdOrKeysAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *FeatureFlagData) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetSummary

`func (o *FeatureFlagData) GetSummary() FeatureFlagSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FeatureFlagData) GetSummaryOk() (*FeatureFlagSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FeatureFlagData) SetSummary(v FeatureFlagSummary)`

SetSummary sets Summary field to given value.


### GetDetails

`func (o *FeatureFlagData) GetDetails() []FeatureFlagDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FeatureFlagData) GetDetailsOk() (*[]FeatureFlagDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FeatureFlagData) SetDetails(v []FeatureFlagDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


