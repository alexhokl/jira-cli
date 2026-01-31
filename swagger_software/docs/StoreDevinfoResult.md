# StoreDevinfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedDevinfoEntities** | Pointer to [**map[string]EntityIds**](EntityIds.md) | The IDs of devinfo entities that have been accepted for submission grouped by their repository IDs. Note that a devinfo entity that isn&#39;t updated due to it&#39;s updateSequenceId being out of order is not considered a failed submission. | [optional] 
**FailedDevinfoEntities** | Pointer to [**map[string]RepositoryErrors**](RepositoryErrors.md) | IDs of devinfo entities that have not been accepted for submission and caused error descriptions, usually due to a problem with the request data. The entities (if present) will be grouped by their repository id and type. Entity IDs are listed with errors associated with that devinfo entity that have prevented it being submitted.  | [optional] 
**UnknownIssueKeys** | Pointer to **[]string** | Issue keys that are not known on this Jira instance (if any). These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist. If a devinfo entity has been associated with issue keys other than those in this array it will still be stored against those valid keys.  | [optional] 
**UnknownAssociations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | Associations that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a development information entity has been associated with any other association other than those in this array it will still be stored against those valid associations. If a development information entity was only associated with the associations in this array, it is deemed to be invalid and it won&#39;t be persisted.  | [optional] 

## Methods

### NewStoreDevinfoResult

`func NewStoreDevinfoResult() *StoreDevinfoResult`

NewStoreDevinfoResult instantiates a new StoreDevinfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDevinfoResultWithDefaults

`func NewStoreDevinfoResultWithDefaults() *StoreDevinfoResult`

NewStoreDevinfoResultWithDefaults instantiates a new StoreDevinfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedDevinfoEntities

`func (o *StoreDevinfoResult) GetAcceptedDevinfoEntities() map[string]EntityIds`

GetAcceptedDevinfoEntities returns the AcceptedDevinfoEntities field if non-nil, zero value otherwise.

### GetAcceptedDevinfoEntitiesOk

`func (o *StoreDevinfoResult) GetAcceptedDevinfoEntitiesOk() (*map[string]EntityIds, bool)`

GetAcceptedDevinfoEntitiesOk returns a tuple with the AcceptedDevinfoEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDevinfoEntities

`func (o *StoreDevinfoResult) SetAcceptedDevinfoEntities(v map[string]EntityIds)`

SetAcceptedDevinfoEntities sets AcceptedDevinfoEntities field to given value.

### HasAcceptedDevinfoEntities

`func (o *StoreDevinfoResult) HasAcceptedDevinfoEntities() bool`

HasAcceptedDevinfoEntities returns a boolean if a field has been set.

### GetFailedDevinfoEntities

`func (o *StoreDevinfoResult) GetFailedDevinfoEntities() map[string]RepositoryErrors`

GetFailedDevinfoEntities returns the FailedDevinfoEntities field if non-nil, zero value otherwise.

### GetFailedDevinfoEntitiesOk

`func (o *StoreDevinfoResult) GetFailedDevinfoEntitiesOk() (*map[string]RepositoryErrors, bool)`

GetFailedDevinfoEntitiesOk returns a tuple with the FailedDevinfoEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDevinfoEntities

`func (o *StoreDevinfoResult) SetFailedDevinfoEntities(v map[string]RepositoryErrors)`

SetFailedDevinfoEntities sets FailedDevinfoEntities field to given value.

### HasFailedDevinfoEntities

`func (o *StoreDevinfoResult) HasFailedDevinfoEntities() bool`

HasFailedDevinfoEntities returns a boolean if a field has been set.

### GetUnknownIssueKeys

`func (o *StoreDevinfoResult) GetUnknownIssueKeys() []string`

GetUnknownIssueKeys returns the UnknownIssueKeys field if non-nil, zero value otherwise.

### GetUnknownIssueKeysOk

`func (o *StoreDevinfoResult) GetUnknownIssueKeysOk() (*[]string, bool)`

GetUnknownIssueKeysOk returns a tuple with the UnknownIssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownIssueKeys

`func (o *StoreDevinfoResult) SetUnknownIssueKeys(v []string)`

SetUnknownIssueKeys sets UnknownIssueKeys field to given value.

### HasUnknownIssueKeys

`func (o *StoreDevinfoResult) HasUnknownIssueKeys() bool`

HasUnknownIssueKeys returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *StoreDevinfoResult) GetUnknownAssociations() []IssueIdOrKeysAssociation`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *StoreDevinfoResult) GetUnknownAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *StoreDevinfoResult) SetUnknownAssociations(v []IssueIdOrKeysAssociation)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *StoreDevinfoResult) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


