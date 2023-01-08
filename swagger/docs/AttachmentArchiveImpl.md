# AttachmentArchiveImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AttachmentArchiveEntry**](AttachmentArchiveEntry.md) | The list of the items included in the archive. | [optional] 
**TotalEntryCount** | Pointer to **int32** | The number of items in the archive. | [optional] 

## Methods

### NewAttachmentArchiveImpl

`func NewAttachmentArchiveImpl() *AttachmentArchiveImpl`

NewAttachmentArchiveImpl instantiates a new AttachmentArchiveImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentArchiveImplWithDefaults

`func NewAttachmentArchiveImplWithDefaults() *AttachmentArchiveImpl`

NewAttachmentArchiveImplWithDefaults instantiates a new AttachmentArchiveImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AttachmentArchiveImpl) GetEntries() []AttachmentArchiveEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AttachmentArchiveImpl) GetEntriesOk() (*[]AttachmentArchiveEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AttachmentArchiveImpl) SetEntries(v []AttachmentArchiveEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AttachmentArchiveImpl) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetTotalEntryCount

`func (o *AttachmentArchiveImpl) GetTotalEntryCount() int32`

GetTotalEntryCount returns the TotalEntryCount field if non-nil, zero value otherwise.

### GetTotalEntryCountOk

`func (o *AttachmentArchiveImpl) GetTotalEntryCountOk() (*int32, bool)`

GetTotalEntryCountOk returns a tuple with the TotalEntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntryCount

`func (o *AttachmentArchiveImpl) SetTotalEntryCount(v int32)`

SetTotalEntryCount sets TotalEntryCount field to given value.

### HasTotalEntryCount

`func (o *AttachmentArchiveImpl) HasTotalEntryCount() bool`

HasTotalEntryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


