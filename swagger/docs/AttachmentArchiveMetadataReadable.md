# AttachmentArchiveMetadataReadable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AttachmentArchiveItemReadable**](AttachmentArchiveItemReadable.md) | The list of the items included in the archive. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the attachment. | [optional] [readonly] 
**MediaType** | Pointer to **string** | The MIME type of the attachment. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the archive file. | [optional] [readonly] 
**TotalEntryCount** | Pointer to **int64** | The number of items included in the archive. | [optional] [readonly] 

## Methods

### NewAttachmentArchiveMetadataReadable

`func NewAttachmentArchiveMetadataReadable() *AttachmentArchiveMetadataReadable`

NewAttachmentArchiveMetadataReadable instantiates a new AttachmentArchiveMetadataReadable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentArchiveMetadataReadableWithDefaults

`func NewAttachmentArchiveMetadataReadableWithDefaults() *AttachmentArchiveMetadataReadable`

NewAttachmentArchiveMetadataReadableWithDefaults instantiates a new AttachmentArchiveMetadataReadable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AttachmentArchiveMetadataReadable) GetEntries() []AttachmentArchiveItemReadable`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AttachmentArchiveMetadataReadable) GetEntriesOk() (*[]AttachmentArchiveItemReadable, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AttachmentArchiveMetadataReadable) SetEntries(v []AttachmentArchiveItemReadable)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AttachmentArchiveMetadataReadable) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetId

`func (o *AttachmentArchiveMetadataReadable) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentArchiveMetadataReadable) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentArchiveMetadataReadable) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentArchiveMetadataReadable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaType

`func (o *AttachmentArchiveMetadataReadable) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AttachmentArchiveMetadataReadable) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AttachmentArchiveMetadataReadable) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AttachmentArchiveMetadataReadable) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetName

`func (o *AttachmentArchiveMetadataReadable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentArchiveMetadataReadable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentArchiveMetadataReadable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentArchiveMetadataReadable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalEntryCount

`func (o *AttachmentArchiveMetadataReadable) GetTotalEntryCount() int64`

GetTotalEntryCount returns the TotalEntryCount field if non-nil, zero value otherwise.

### GetTotalEntryCountOk

`func (o *AttachmentArchiveMetadataReadable) GetTotalEntryCountOk() (*int64, bool)`

GetTotalEntryCountOk returns a tuple with the TotalEntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntryCount

`func (o *AttachmentArchiveMetadataReadable) SetTotalEntryCount(v int64)`

SetTotalEntryCount sets TotalEntryCount field to given value.

### HasTotalEntryCount

`func (o *AttachmentArchiveMetadataReadable) HasTotalEntryCount() bool`

HasTotalEntryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


