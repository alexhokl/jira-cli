# DocumentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The version UUID. | [optional] 
**VersionNumber** | Pointer to **int64** | The version number. | [optional] 

## Methods

### NewDocumentVersion

`func NewDocumentVersion() *DocumentVersion`

NewDocumentVersion instantiates a new DocumentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVersionWithDefaults

`func NewDocumentVersionWithDefaults() *DocumentVersion`

NewDocumentVersionWithDefaults instantiates a new DocumentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionNumber

`func (o *DocumentVersion) GetVersionNumber() int64`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *DocumentVersion) GetVersionNumberOk() (*int64, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *DocumentVersion) SetVersionNumber(v int64)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *DocumentVersion) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


