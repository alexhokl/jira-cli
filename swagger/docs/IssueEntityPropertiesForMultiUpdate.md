# IssueEntityPropertiesForMultiUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueID** | Pointer to **int64** | The ID of the issue. | [optional] 
**Properties** | Pointer to [**map[string]JsonNode**](JsonNode.md) | Entity properties to set on the issue. The maximum length of an issue property value is 32768 characters. | [optional] 

## Methods

### NewIssueEntityPropertiesForMultiUpdate

`func NewIssueEntityPropertiesForMultiUpdate() *IssueEntityPropertiesForMultiUpdate`

NewIssueEntityPropertiesForMultiUpdate instantiates a new IssueEntityPropertiesForMultiUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEntityPropertiesForMultiUpdateWithDefaults

`func NewIssueEntityPropertiesForMultiUpdateWithDefaults() *IssueEntityPropertiesForMultiUpdate`

NewIssueEntityPropertiesForMultiUpdateWithDefaults instantiates a new IssueEntityPropertiesForMultiUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueID

`func (o *IssueEntityPropertiesForMultiUpdate) GetIssueID() int64`

GetIssueID returns the IssueID field if non-nil, zero value otherwise.

### GetIssueIDOk

`func (o *IssueEntityPropertiesForMultiUpdate) GetIssueIDOk() (*int64, bool)`

GetIssueIDOk returns a tuple with the IssueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueID

`func (o *IssueEntityPropertiesForMultiUpdate) SetIssueID(v int64)`

SetIssueID sets IssueID field to given value.

### HasIssueID

`func (o *IssueEntityPropertiesForMultiUpdate) HasIssueID() bool`

HasIssueID returns a boolean if a field has been set.

### GetProperties

`func (o *IssueEntityPropertiesForMultiUpdate) GetProperties() map[string]JsonNode`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueEntityPropertiesForMultiUpdate) GetPropertiesOk() (*map[string]JsonNode, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueEntityPropertiesForMultiUpdate) SetProperties(v map[string]JsonNode)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueEntityPropertiesForMultiUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


