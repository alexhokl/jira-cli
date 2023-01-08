# FailedWebhooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | **int32** | The maximum number of items on the page. If the list of values is shorter than this number, then there are no more pages. | 
**Next** | Pointer to **string** | The URL to the next page of results. Present only if the request returned at least one result.The next page may be empty at the time of receiving the response, but new failed webhooks may appear in time. You can save the URL to the next page and query for new results periodically (for example, every hour). | [optional] 
**Values** | [**[]FailedWebhook**](FailedWebhook.md) | The list of webhooks. | 

## Methods

### NewFailedWebhooks

`func NewFailedWebhooks(maxResults int32, values []FailedWebhook, ) *FailedWebhooks`

NewFailedWebhooks instantiates a new FailedWebhooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedWebhooksWithDefaults

`func NewFailedWebhooksWithDefaults() *FailedWebhooks`

NewFailedWebhooksWithDefaults instantiates a new FailedWebhooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *FailedWebhooks) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *FailedWebhooks) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *FailedWebhooks) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.


### GetNext

`func (o *FailedWebhooks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FailedWebhooks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FailedWebhooks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FailedWebhooks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetValues

`func (o *FailedWebhooks) GetValues() []FailedWebhook`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FailedWebhooks) GetValuesOk() (*[]FailedWebhook, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FailedWebhooks) SetValues(v []FailedWebhook)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


