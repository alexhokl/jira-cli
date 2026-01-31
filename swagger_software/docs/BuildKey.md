# BuildKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineId** | **string** | An ID that relates a sequence of builds. Depending on your system this might be a project ID, pipeline ID, plan key etc. - whatever logical unit you use to group a sequence of builds.  The combination of &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60; must uniquely identify the build.  | 
**BuildNumber** | **int64** | Identifies a build within the sequence of builds identified by the build &#x60;pipelineId&#x60;.  Used to identify the &#39;most recent&#39; build in that sequence of builds.  The combination of &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60; must uniquely identify the build.  | 

## Methods

### NewBuildKey

`func NewBuildKey(pipelineId string, buildNumber int64, ) *BuildKey`

NewBuildKey instantiates a new BuildKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildKeyWithDefaults

`func NewBuildKeyWithDefaults() *BuildKey`

NewBuildKeyWithDefaults instantiates a new BuildKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineId

`func (o *BuildKey) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *BuildKey) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *BuildKey) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.


### GetBuildNumber

`func (o *BuildKey) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *BuildKey) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *BuildKey) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


