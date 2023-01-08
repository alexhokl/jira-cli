# AttachmentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**User**](User.md) | Details of the user who attached the file. | [optional] [readonly] 
**Content** | Pointer to **string** | The URL of the attachment. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The datetime the attachment was created. | [optional] [readonly] 
**Filename** | Pointer to **string** | The name of the attachment file. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the attachment. | [optional] [readonly] 
**MimeType** | Pointer to **string** | The MIME type of the attachment. | [optional] [readonly] 
**Properties** | Pointer to **map[string]interface{}** | Additional properties of the attachment. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the attachment metadata details. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the attachment. | [optional] [readonly] 
**Thumbnail** | Pointer to **string** | The URL of a thumbnail representing the attachment. | [optional] [readonly] 

## Methods

### NewAttachmentMetadata

`func NewAttachmentMetadata() *AttachmentMetadata`

NewAttachmentMetadata instantiates a new AttachmentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentMetadataWithDefaults

`func NewAttachmentMetadataWithDefaults() *AttachmentMetadata`

NewAttachmentMetadataWithDefaults instantiates a new AttachmentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AttachmentMetadata) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AttachmentMetadata) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AttachmentMetadata) SetAuthor(v User)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AttachmentMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetContent

`func (o *AttachmentMetadata) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AttachmentMetadata) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AttachmentMetadata) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AttachmentMetadata) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *AttachmentMetadata) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AttachmentMetadata) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AttachmentMetadata) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AttachmentMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFilename

`func (o *AttachmentMetadata) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AttachmentMetadata) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AttachmentMetadata) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AttachmentMetadata) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *AttachmentMetadata) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentMetadata) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentMetadata) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimeType

`func (o *AttachmentMetadata) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AttachmentMetadata) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AttachmentMetadata) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AttachmentMetadata) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetProperties

`func (o *AttachmentMetadata) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AttachmentMetadata) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AttachmentMetadata) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AttachmentMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSelf

`func (o *AttachmentMetadata) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AttachmentMetadata) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AttachmentMetadata) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AttachmentMetadata) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSize

`func (o *AttachmentMetadata) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AttachmentMetadata) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AttachmentMetadata) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AttachmentMetadata) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThumbnail

`func (o *AttachmentMetadata) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *AttachmentMetadata) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *AttachmentMetadata) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *AttachmentMetadata) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


