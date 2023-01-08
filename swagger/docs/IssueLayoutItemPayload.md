# IssueLayoutItemPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Additional properties for this item. This field is only used when the type is FIELD. | [optional] 
**SectionType** | Pointer to **string** | The item section type | [optional] 
**Type** | Pointer to **string** | The item type. Currently only support FIELD | [optional] 

## Methods

### NewIssueLayoutItemPayload

`func NewIssueLayoutItemPayload() *IssueLayoutItemPayload`

NewIssueLayoutItemPayload instantiates a new IssueLayoutItemPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLayoutItemPayloadWithDefaults

`func NewIssueLayoutItemPayloadWithDefaults() *IssueLayoutItemPayload`

NewIssueLayoutItemPayloadWithDefaults instantiates a new IssueLayoutItemPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *IssueLayoutItemPayload) GetItemKey() ProjectCreateResourceIdentifier`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *IssueLayoutItemPayload) GetItemKeyOk() (*ProjectCreateResourceIdentifier, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *IssueLayoutItemPayload) SetItemKey(v ProjectCreateResourceIdentifier)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *IssueLayoutItemPayload) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetProperties

`func (o *IssueLayoutItemPayload) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueLayoutItemPayload) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueLayoutItemPayload) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueLayoutItemPayload) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSectionType

`func (o *IssueLayoutItemPayload) GetSectionType() string`

GetSectionType returns the SectionType field if non-nil, zero value otherwise.

### GetSectionTypeOk

`func (o *IssueLayoutItemPayload) GetSectionTypeOk() (*string, bool)`

GetSectionTypeOk returns a tuple with the SectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionType

`func (o *IssueLayoutItemPayload) SetSectionType(v string)`

SetSectionType sets SectionType field to given value.

### HasSectionType

`func (o *IssueLayoutItemPayload) HasSectionType() bool`

HasSectionType returns a boolean if a field has been set.

### GetType

`func (o *IssueLayoutItemPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueLayoutItemPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueLayoutItemPayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueLayoutItemPayload) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


