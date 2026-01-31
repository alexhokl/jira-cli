# RemoteLinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | The schema version used for this data.  Placeholder to support potential schema changes in the future.  | [optional] [default to "1.0"]
**Id** | **string** | The identifier for the Remote Link. Must be unique for a given Provider.  | 
**UpdateSequenceNumber** | **int64** | An ID used to apply an ordering to updates for this Remote Link in the case of out-of-order receipt of update requests.  It must be a monotonically increasing number. For example, epoch time could be one way to generate the &#x60;updateSequenceNumber&#x60;.  Updates for a Remote Link that is received with an &#x60;updateSqeuenceNumber&#x60; less than or equal to what is currently stored will be ignored.  | 
**DisplayName** | **string** | The human-readable name for the Remote Link.  Will be shown in the UI.  | 
**Url** | **string** | The URL to this Remote Link in your system.  | 
**Type** | **string** | The type of the Remote Link. The current supported types are &#39;document&#39;, &#39;alert&#39;, &#39;test&#39;, &#39;security&#39;, &#39;logFile&#39;, &#39;prototype&#39;, &#39;coverage&#39;, &#39;bugReport&#39; and &#39;other&#39;  | 
**Description** | Pointer to **string** | An optional description to attach to this Remote Link.  This may be anything that makes sense in your system.  | [optional] 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user as a summary of when Remote Link was last updated.  | 
**Associations** | Pointer to [**[]RemoteLinkDataAssociationsInner**](RemoteLinkDataAssociationsInner.md) | The entities to associate the Remote Link information with.  | [optional] 
**Status** | Pointer to [**RemoteLinkStatus**](RemoteLinkStatus.md) |  | [optional] 
**ActionIds** | Pointer to **[]string** | Optional list of actionIds. They are associated with the actions the provider is able to provide when they registered. Indicates which actions this Remote Link has.  If any actions have a templateUrl that requires string substitution, then &#x60;attributeMap&#x60; must be passed in.  | [optional] 
**AttributeMap** | Pointer to **map[string]string** | Map of key/values (string to string mapping). This is used to build the urls for actions from the templateUrl the provider registered their available actions with.  | [optional] 

## Methods

### NewRemoteLinkData

`func NewRemoteLinkData(id string, updateSequenceNumber int64, displayName string, url string, type_ string, lastUpdated time.Time, ) *RemoteLinkData`

NewRemoteLinkData instantiates a new RemoteLinkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLinkDataWithDefaults

`func NewRemoteLinkDataWithDefaults() *RemoteLinkData`

NewRemoteLinkDataWithDefaults instantiates a new RemoteLinkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RemoteLinkData) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RemoteLinkData) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RemoteLinkData) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RemoteLinkData) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetId

`func (o *RemoteLinkData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteLinkData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteLinkData) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateSequenceNumber

`func (o *RemoteLinkData) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *RemoteLinkData) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *RemoteLinkData) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetDisplayName

`func (o *RemoteLinkData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteLinkData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoteLinkData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUrl

`func (o *RemoteLinkData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteLinkData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteLinkData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *RemoteLinkData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemoteLinkData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemoteLinkData) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *RemoteLinkData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteLinkData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteLinkData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RemoteLinkData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RemoteLinkData) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RemoteLinkData) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RemoteLinkData) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetAssociations

`func (o *RemoteLinkData) GetAssociations() []RemoteLinkDataAssociationsInner`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *RemoteLinkData) GetAssociationsOk() (*[]RemoteLinkDataAssociationsInner, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *RemoteLinkData) SetAssociations(v []RemoteLinkDataAssociationsInner)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *RemoteLinkData) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetStatus

`func (o *RemoteLinkData) GetStatus() RemoteLinkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteLinkData) GetStatusOk() (*RemoteLinkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteLinkData) SetStatus(v RemoteLinkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteLinkData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActionIds

`func (o *RemoteLinkData) GetActionIds() []string`

GetActionIds returns the ActionIds field if non-nil, zero value otherwise.

### GetActionIdsOk

`func (o *RemoteLinkData) GetActionIdsOk() (*[]string, bool)`

GetActionIdsOk returns a tuple with the ActionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionIds

`func (o *RemoteLinkData) SetActionIds(v []string)`

SetActionIds sets ActionIds field to given value.

### HasActionIds

`func (o *RemoteLinkData) HasActionIds() bool`

HasActionIds returns a boolean if a field has been set.

### GetAttributeMap

`func (o *RemoteLinkData) GetAttributeMap() map[string]string`

GetAttributeMap returns the AttributeMap field if non-nil, zero value otherwise.

### GetAttributeMapOk

`func (o *RemoteLinkData) GetAttributeMapOk() (*map[string]string, bool)`

GetAttributeMapOk returns a tuple with the AttributeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMap

`func (o *RemoteLinkData) SetAttributeMap(v map[string]string)`

SetAttributeMap sets AttributeMap field to given value.

### HasAttributeMap

`func (o *RemoteLinkData) HasAttributeMap() bool`

HasAttributeMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


