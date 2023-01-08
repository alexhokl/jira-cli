# ReorderIssueResolutionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The ID of the resolution. Required if &#x60;position&#x60; isn&#39;t provided. | [optional] 
**Ids** | **[]string** | The list of resolution IDs to be reordered. Cannot contain duplicates nor after ID. | 
**Position** | Pointer to **string** | The position for issue resolutions to be moved to. Required if &#x60;after&#x60; isn&#39;t provided. | [optional] 

## Methods

### NewReorderIssueResolutionsRequest

`func NewReorderIssueResolutionsRequest(ids []string, ) *ReorderIssueResolutionsRequest`

NewReorderIssueResolutionsRequest instantiates a new ReorderIssueResolutionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderIssueResolutionsRequestWithDefaults

`func NewReorderIssueResolutionsRequestWithDefaults() *ReorderIssueResolutionsRequest`

NewReorderIssueResolutionsRequestWithDefaults instantiates a new ReorderIssueResolutionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ReorderIssueResolutionsRequest) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ReorderIssueResolutionsRequest) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ReorderIssueResolutionsRequest) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ReorderIssueResolutionsRequest) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetIds

`func (o *ReorderIssueResolutionsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ReorderIssueResolutionsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ReorderIssueResolutionsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetPosition

`func (o *ReorderIssueResolutionsRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReorderIssueResolutionsRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReorderIssueResolutionsRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ReorderIssueResolutionsRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


