# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**File** | Pointer to ***os.File** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**InputStream** | Pointer to **map[string]interface{}** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**Readable** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFile

`func (o *Resource) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Resource) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Resource) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *Resource) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFilename

`func (o *Resource) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Resource) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Resource) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Resource) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetInputStream

`func (o *Resource) GetInputStream() map[string]interface{}`

GetInputStream returns the InputStream field if non-nil, zero value otherwise.

### GetInputStreamOk

`func (o *Resource) GetInputStreamOk() (*map[string]interface{}, bool)`

GetInputStreamOk returns a tuple with the InputStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputStream

`func (o *Resource) SetInputStream(v map[string]interface{})`

SetInputStream sets InputStream field to given value.

### HasInputStream

`func (o *Resource) HasInputStream() bool`

HasInputStream returns a boolean if a field has been set.

### GetOpen

`func (o *Resource) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Resource) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Resource) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Resource) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetReadable

`func (o *Resource) GetReadable() bool`

GetReadable returns the Readable field if non-nil, zero value otherwise.

### GetReadableOk

`func (o *Resource) GetReadableOk() (*bool, bool)`

GetReadableOk returns a tuple with the Readable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadable

`func (o *Resource) SetReadable(v bool)`

SetReadable sets Readable field to given value.

### HasReadable

`func (o *Resource) HasReadable() bool`

HasReadable returns a boolean if a field has been set.

### GetUri

`func (o *Resource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Resource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Resource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Resource) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUrl

`func (o *Resource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Resource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Resource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Resource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


