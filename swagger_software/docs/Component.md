# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | **string** | The DevOpsComponentData schema version used for this devops component data.  Placeholder to support potential schema changes in the future.  | [default to "1.0"]
**Id** | **string** | The identifier for the DevOps Component. Must be unique for a given Provider.  | 
**UpdateSequenceNumber** | **int64** | An ID used to apply an ordering to updates for this DevOps Component in the case of out-of-order receipt of update requests.  This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the Provider system, but other alternatives are valid (e.g. a Provider could store a counter against each DevOps Component and increment that on each update to Jira).  Updates for a DevOps Component that are received with an updateSqeuenceId lower than what is currently stored will be ignored.  | 
**Name** | **string** | The human-readable name for the DevOps Component. Will be shown in the UI.  | 
**ProviderName** | Pointer to **string** | The human-readable name for the Provider that owns this DevOps Component. Will be shown in the UI.  | [optional] 
**Description** | **string** | A description of the DevOps Component in Markdown format. Will be shown in the UI.  | 
**Url** | **string** | A URL users can use to link to a summary view of this devops component, if appropriate.  This could be any location that makes sense in the Provider system (e.g. if the summary information comes from a specific project, it might make sense to link the user to the component in that project).  | 
**AvatarUrl** | **string** | A URL to display a logo representing this devops component, if available.  | 
**Tier** | **string** | The tier of the component. Will be shown in the UI.  | 
**ComponentType** | **string** | The type of the component. Will be shown in the UI.  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user the last time the DevOps Component was updated.  Expected format is an RFC3339 formatted string.  | 

## Methods

### NewComponent

`func NewComponent(schemaVersion string, id string, updateSequenceNumber int64, name string, description string, url string, avatarUrl string, tier string, componentType string, lastUpdated time.Time, ) *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *Component) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *Component) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *Component) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetId

`func (o *Component) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Component) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Component) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateSequenceNumber

`func (o *Component) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *Component) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *Component) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetName

`func (o *Component) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Component) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Component) SetName(v string)`

SetName sets Name field to given value.


### GetProviderName

`func (o *Component) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *Component) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *Component) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *Component) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetDescription

`func (o *Component) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Component) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Component) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *Component) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Component) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Component) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *Component) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Component) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Component) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetTier

`func (o *Component) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Component) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Component) SetTier(v string)`

SetTier sets Tier field to given value.


### GetComponentType

`func (o *Component) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *Component) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *Component) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.


### GetLastUpdated

`func (o *Component) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Component) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Component) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


