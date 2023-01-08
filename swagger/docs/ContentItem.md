# ContentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the content entity.   *  For redacting an issue field, this will be the field ID (e.g., summary, customfield\\_10000).  *  For redacting a comment, this will be the comment ID.  *  For redacting a worklog, this will be the worklog ID. | 
**EntityType** | **string** | The type of the entity to redact; It will be one of the following:   *  **issuefieldvalue** \\- To redact in issue fields  *  **issue-comment** \\- To redact in issue comments.  *  **issue-worklog** \\- To redact in issue worklogs | 
**Id** | **string** | This would be the issue ID | 

## Methods

### NewContentItem

`func NewContentItem(entityId string, entityType string, id string, ) *ContentItem`

NewContentItem instantiates a new ContentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentItemWithDefaults

`func NewContentItemWithDefaults() *ContentItem`

NewContentItemWithDefaults instantiates a new ContentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ContentItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ContentItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ContentItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *ContentItem) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ContentItem) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ContentItem) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *ContentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentItem) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


