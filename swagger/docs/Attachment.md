# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**UserDetails**](UserDetails.md) | Details of the user who added the attachment. | [optional] [readonly] 
**Content** | Pointer to **string** | The content of the attachment. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The datetime the attachment was created. | [optional] [readonly] 
**Filename** | Pointer to **string** | The file name of the attachment. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the attachment. | [optional] [readonly] 
**MimeType** | Pointer to **string** | The MIME type of the attachment. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the attachment details response. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the attachment. | [optional] [readonly] 
**Thumbnail** | Pointer to **string** | The URL of a thumbnail representing the attachment. | [optional] [readonly] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Attachment) GetAuthor() UserDetails`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Attachment) GetAuthorOk() (*UserDetails, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Attachment) SetAuthor(v UserDetails)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Attachment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetContent

`func (o *Attachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Attachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Attachment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Attachment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *Attachment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Attachment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Attachment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Attachment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFilename

`func (o *Attachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Attachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Attachment) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Attachment) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *Attachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimeType

`func (o *Attachment) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Attachment) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Attachment) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Attachment) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSelf

`func (o *Attachment) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Attachment) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Attachment) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Attachment) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSize

`func (o *Attachment) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Attachment) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Attachment) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Attachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThumbnail

`func (o *Attachment) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Attachment) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Attachment) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Attachment) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


