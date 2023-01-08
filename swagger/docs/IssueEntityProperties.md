# IssueEntityProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitiesIds** | Pointer to **[]int64** | A list of entity property IDs. | [optional] 
**Properties** | Pointer to [**map[string]JsonNode**](JsonNode.md) | A list of entity property keys and values. | [optional] 

## Methods

### NewIssueEntityProperties

`func NewIssueEntityProperties() *IssueEntityProperties`

NewIssueEntityProperties instantiates a new IssueEntityProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEntityPropertiesWithDefaults

`func NewIssueEntityPropertiesWithDefaults() *IssueEntityProperties`

NewIssueEntityPropertiesWithDefaults instantiates a new IssueEntityProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitiesIds

`func (o *IssueEntityProperties) GetEntitiesIds() []int64`

GetEntitiesIds returns the EntitiesIds field if non-nil, zero value otherwise.

### GetEntitiesIdsOk

`func (o *IssueEntityProperties) GetEntitiesIdsOk() (*[]int64, bool)`

GetEntitiesIdsOk returns a tuple with the EntitiesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesIds

`func (o *IssueEntityProperties) SetEntitiesIds(v []int64)`

SetEntitiesIds sets EntitiesIds field to given value.

### HasEntitiesIds

`func (o *IssueEntityProperties) HasEntitiesIds() bool`

HasEntitiesIds returns a boolean if a field has been set.

### GetProperties

`func (o *IssueEntityProperties) GetProperties() map[string]JsonNode`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueEntityProperties) GetPropertiesOk() (*map[string]JsonNode, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueEntityProperties) SetProperties(v map[string]JsonNode)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueEntityProperties) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


