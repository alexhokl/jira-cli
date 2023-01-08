# AttachmentArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AttachmentArchiveEntry**](AttachmentArchiveEntry.md) |  | [optional] 
**MoreAvailable** | Pointer to **bool** |  | [optional] 
**TotalEntryCount** | Pointer to **int32** |  | [optional] 
**TotalNumberOfEntriesAvailable** | Pointer to **int32** |  | [optional] 

## Methods

### NewAttachmentArchive

`func NewAttachmentArchive() *AttachmentArchive`

NewAttachmentArchive instantiates a new AttachmentArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentArchiveWithDefaults

`func NewAttachmentArchiveWithDefaults() *AttachmentArchive`

NewAttachmentArchiveWithDefaults instantiates a new AttachmentArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AttachmentArchive) GetEntries() []AttachmentArchiveEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AttachmentArchive) GetEntriesOk() (*[]AttachmentArchiveEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AttachmentArchive) SetEntries(v []AttachmentArchiveEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AttachmentArchive) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetMoreAvailable

`func (o *AttachmentArchive) GetMoreAvailable() bool`

GetMoreAvailable returns the MoreAvailable field if non-nil, zero value otherwise.

### GetMoreAvailableOk

`func (o *AttachmentArchive) GetMoreAvailableOk() (*bool, bool)`

GetMoreAvailableOk returns a tuple with the MoreAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreAvailable

`func (o *AttachmentArchive) SetMoreAvailable(v bool)`

SetMoreAvailable sets MoreAvailable field to given value.

### HasMoreAvailable

`func (o *AttachmentArchive) HasMoreAvailable() bool`

HasMoreAvailable returns a boolean if a field has been set.

### GetTotalEntryCount

`func (o *AttachmentArchive) GetTotalEntryCount() int32`

GetTotalEntryCount returns the TotalEntryCount field if non-nil, zero value otherwise.

### GetTotalEntryCountOk

`func (o *AttachmentArchive) GetTotalEntryCountOk() (*int32, bool)`

GetTotalEntryCountOk returns a tuple with the TotalEntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntryCount

`func (o *AttachmentArchive) SetTotalEntryCount(v int32)`

SetTotalEntryCount sets TotalEntryCount field to given value.

### HasTotalEntryCount

`func (o *AttachmentArchive) HasTotalEntryCount() bool`

HasTotalEntryCount returns a boolean if a field has been set.

### GetTotalNumberOfEntriesAvailable

`func (o *AttachmentArchive) GetTotalNumberOfEntriesAvailable() int32`

GetTotalNumberOfEntriesAvailable returns the TotalNumberOfEntriesAvailable field if non-nil, zero value otherwise.

### GetTotalNumberOfEntriesAvailableOk

`func (o *AttachmentArchive) GetTotalNumberOfEntriesAvailableOk() (*int32, bool)`

GetTotalNumberOfEntriesAvailableOk returns a tuple with the TotalNumberOfEntriesAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfEntriesAvailable

`func (o *AttachmentArchive) SetTotalNumberOfEntriesAvailable(v int32)`

SetTotalNumberOfEntriesAvailable sets TotalNumberOfEntriesAvailable field to given value.

### HasTotalNumberOfEntriesAvailable

`func (o *AttachmentArchive) HasTotalNumberOfEntriesAvailable() bool`

HasTotalNumberOfEntriesAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


