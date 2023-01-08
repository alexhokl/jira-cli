# TargetClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifications** | **map[string][]string** | An object with the key as the ID of the target classification and value with the list of the IDs of the current source classifications. | 
**IssueType** | Pointer to **string** | ID of the source issueType to which issues present in &#x60;issueIdOrKeys&#x60; belongs. | [optional] 
**ProjectKeyOrId** | Pointer to **string** | ID or key of the source project to which issues present in &#x60;issueIdOrKeys&#x60; belongs. | [optional] 

## Methods

### NewTargetClassification

`func NewTargetClassification(classifications map[string][]string, ) *TargetClassification`

NewTargetClassification instantiates a new TargetClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetClassificationWithDefaults

`func NewTargetClassificationWithDefaults() *TargetClassification`

NewTargetClassificationWithDefaults instantiates a new TargetClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifications

`func (o *TargetClassification) GetClassifications() map[string][]string`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *TargetClassification) GetClassificationsOk() (*map[string][]string, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *TargetClassification) SetClassifications(v map[string][]string)`

SetClassifications sets Classifications field to given value.


### GetIssueType

`func (o *TargetClassification) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *TargetClassification) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *TargetClassification) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *TargetClassification) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetProjectKeyOrId

`func (o *TargetClassification) GetProjectKeyOrId() string`

GetProjectKeyOrId returns the ProjectKeyOrId field if non-nil, zero value otherwise.

### GetProjectKeyOrIdOk

`func (o *TargetClassification) GetProjectKeyOrIdOk() (*string, bool)`

GetProjectKeyOrIdOk returns a tuple with the ProjectKeyOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeyOrId

`func (o *TargetClassification) SetProjectKeyOrId(v string)`

SetProjectKeyOrId sets ProjectKeyOrId field to given value.

### HasProjectKeyOrId

`func (o *TargetClassification) HasProjectKeyOrId() bool`

HasProjectKeyOrId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


