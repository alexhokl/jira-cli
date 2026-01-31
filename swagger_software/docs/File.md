# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path of the file. Max length is 1024 characters. | 
**Url** | **string** | The URL of this file. Max length is 2000 characters. | 
**ChangeType** | **string** | The operation performed on this file | 
**LinesAdded** | **int32** | Number of lines added to the file | 
**LinesRemoved** | **int32** | Number of lines removed from the file | 

## Methods

### NewFile

`func NewFile(path string, url string, changeType string, linesAdded int32, linesRemoved int32, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *File) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *File) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *File) SetPath(v string)`

SetPath sets Path field to given value.


### GetUrl

`func (o *File) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *File) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *File) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetChangeType

`func (o *File) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *File) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *File) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.


### GetLinesAdded

`func (o *File) GetLinesAdded() int32`

GetLinesAdded returns the LinesAdded field if non-nil, zero value otherwise.

### GetLinesAddedOk

`func (o *File) GetLinesAddedOk() (*int32, bool)`

GetLinesAddedOk returns a tuple with the LinesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesAdded

`func (o *File) SetLinesAdded(v int32)`

SetLinesAdded sets LinesAdded field to given value.


### GetLinesRemoved

`func (o *File) GetLinesRemoved() int32`

GetLinesRemoved returns the LinesRemoved field if non-nil, zero value otherwise.

### GetLinesRemovedOk

`func (o *File) GetLinesRemovedOk() (*int32, bool)`

GetLinesRemovedOk returns a tuple with the LinesRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesRemoved

`func (o *File) SetLinesRemoved(v int32)`

SetLinesRemoved sets LinesRemoved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


