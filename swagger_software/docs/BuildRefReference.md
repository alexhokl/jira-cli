# BuildRefReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the ref the build ran on  | 
**Uri** | **string** | An identifer for the ref.  In most cases this should be the URL of the tag/branch etc. in the SCM provider.  For cases where the build was executed against a local repository etc. this should be something that uniquely identifies the ref.  | 

## Methods

### NewBuildRefReference

`func NewBuildRefReference(name string, uri string, ) *BuildRefReference`

NewBuildRefReference instantiates a new BuildRefReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRefReferenceWithDefaults

`func NewBuildRefReferenceWithDefaults() *BuildRefReference`

NewBuildRefReferenceWithDefaults instantiates a new BuildRefReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BuildRefReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildRefReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildRefReference) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *BuildRefReference) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BuildRefReference) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BuildRefReference) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


