# DevInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | [**[]Repository**](Repository.md) | List of repositories containing development information. Must not contain duplicates. Maximum number of entities across all repositories is 1000. | 
**PreventTransitions** | Pointer to **bool** | Flag to prevent automatic issue transitions and smart commits being fired, default is false. | [optional] 
**OperationType** | Pointer to **string** | Indicates the operation being performed by the provider system when sending this data. \&quot;NORMAL\&quot; - Data received during normal operation (e.g. a user pushing a branch). \&quot;BACKFILL\&quot; - Data received while backfilling existing data (e.g. indexing a newly connected account). Default is \&quot;NORMAL\&quot;. Please note that \&quot;BACKFILL\&quot; operations have a much higher rate-limiting threshold but are also processed slower in comparison to \&quot;NORMAL\&quot; operations. | [optional] 
**Properties** | Pointer to **map[string]string** | Arbitrary properties to tag the submitted repositories with. These properties can be used for delete operations to e.g. clean up all development information associated with an account in the event that the account is removed from the provider system. Note that these properties will never be returned with repository or entity data. They are not intended for use as metadata to associate with a repository. Maximum length of each key or value is 255 characters. Maximum allowed number of properties key/value pairs is 5. Properties keys cannot start with &#39;_&#39; character. Properties keys cannot contain &#39;:&#39; character.  | [optional] 
**ProviderMetadata** | Pointer to [**ProviderMetadata**](ProviderMetadata.md) |  | [optional] 

## Methods

### NewDevInformation

`func NewDevInformation(repositories []Repository, ) *DevInformation`

NewDevInformation instantiates a new DevInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevInformationWithDefaults

`func NewDevInformationWithDefaults() *DevInformation`

NewDevInformationWithDefaults instantiates a new DevInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *DevInformation) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *DevInformation) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *DevInformation) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.


### GetPreventTransitions

`func (o *DevInformation) GetPreventTransitions() bool`

GetPreventTransitions returns the PreventTransitions field if non-nil, zero value otherwise.

### GetPreventTransitionsOk

`func (o *DevInformation) GetPreventTransitionsOk() (*bool, bool)`

GetPreventTransitionsOk returns a tuple with the PreventTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTransitions

`func (o *DevInformation) SetPreventTransitions(v bool)`

SetPreventTransitions sets PreventTransitions field to given value.

### HasPreventTransitions

`func (o *DevInformation) HasPreventTransitions() bool`

HasPreventTransitions returns a boolean if a field has been set.

### GetOperationType

`func (o *DevInformation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *DevInformation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *DevInformation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *DevInformation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetProperties

`func (o *DevInformation) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DevInformation) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DevInformation) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DevInformation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *DevInformation) GetProviderMetadata() ProviderMetadata`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *DevInformation) GetProviderMetadataOk() (*ProviderMetadata, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *DevInformation) SetProviderMetadata(v ProviderMetadata)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *DevInformation) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


