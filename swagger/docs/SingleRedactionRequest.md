# SingleRedactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentItem** | [**ContentItem**](ContentItem.md) |  | 
**ExternalId** | **string** | Unique id for the redaction request; ID format should be of UUID | 
**Reason** | **string** | The reason why the content is being redacted | 
**RedactionPosition** | [**RedactionPosition**](RedactionPosition.md) |  | 

## Methods

### NewSingleRedactionRequest

`func NewSingleRedactionRequest(contentItem ContentItem, externalId string, reason string, redactionPosition RedactionPosition, ) *SingleRedactionRequest`

NewSingleRedactionRequest instantiates a new SingleRedactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRedactionRequestWithDefaults

`func NewSingleRedactionRequestWithDefaults() *SingleRedactionRequest`

NewSingleRedactionRequestWithDefaults instantiates a new SingleRedactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentItem

`func (o *SingleRedactionRequest) GetContentItem() ContentItem`

GetContentItem returns the ContentItem field if non-nil, zero value otherwise.

### GetContentItemOk

`func (o *SingleRedactionRequest) GetContentItemOk() (*ContentItem, bool)`

GetContentItemOk returns a tuple with the ContentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentItem

`func (o *SingleRedactionRequest) SetContentItem(v ContentItem)`

SetContentItem sets ContentItem field to given value.


### GetExternalId

`func (o *SingleRedactionRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SingleRedactionRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SingleRedactionRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetReason

`func (o *SingleRedactionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SingleRedactionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SingleRedactionRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRedactionPosition

`func (o *SingleRedactionRequest) GetRedactionPosition() RedactionPosition`

GetRedactionPosition returns the RedactionPosition field if non-nil, zero value otherwise.

### GetRedactionPositionOk

`func (o *SingleRedactionRequest) GetRedactionPositionOk() (*RedactionPosition, bool)`

GetRedactionPositionOk returns a tuple with the RedactionPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactionPosition

`func (o *SingleRedactionRequest) SetRedactionPosition(v RedactionPosition)`

SetRedactionPosition sets RedactionPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


