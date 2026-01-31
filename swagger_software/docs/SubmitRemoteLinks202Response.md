# SubmitRemoteLinks202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedRemoteLinks** | Pointer to **[]string** | The IDs of Remote Links that have been accepted for submission.  A Remote Link may be rejected if it was only associated with unknown issue keys, unknown service IDs, or if the submitted data for that Remote Link does not match the required schema.  Note that a Remote Link that isn&#39;t updated due to it&#39;s &#x60;updateSequenceNumber&#x60; being out of order is not considered a failed submission.  | [optional] 
**RejectedRemoteLinks** | Pointer to [**map[string][]SubmitRemoteLinks202ResponseRejectedRemoteLinksValueInner**](array.md) | Details of Remote Links that have not been accepted for submission, usually due to a problem with the request data.  A Remote Link may be rejected if it was only associated with unknown issue keys, unknown service IDs, or if the submitted data for the Remote Link does not match the required schema.  The object (if present) will be keyed by Remote Link ID and include any errors associated with that Remote Link that have prevented it being submitted.  | [optional] 
**UnknownAssociations** | Pointer to [**[]RemoteLinkDataAssociationsInner**](RemoteLinkDataAssociationsInner.md) | Issue keys or services IDs or keys that are not known on this Jira instance (if any).  | [optional] 

## Methods

### NewSubmitRemoteLinks202Response

`func NewSubmitRemoteLinks202Response() *SubmitRemoteLinks202Response`

NewSubmitRemoteLinks202Response instantiates a new SubmitRemoteLinks202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRemoteLinks202ResponseWithDefaults

`func NewSubmitRemoteLinks202ResponseWithDefaults() *SubmitRemoteLinks202Response`

NewSubmitRemoteLinks202ResponseWithDefaults instantiates a new SubmitRemoteLinks202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedRemoteLinks

`func (o *SubmitRemoteLinks202Response) GetAcceptedRemoteLinks() []string`

GetAcceptedRemoteLinks returns the AcceptedRemoteLinks field if non-nil, zero value otherwise.

### GetAcceptedRemoteLinksOk

`func (o *SubmitRemoteLinks202Response) GetAcceptedRemoteLinksOk() (*[]string, bool)`

GetAcceptedRemoteLinksOk returns a tuple with the AcceptedRemoteLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRemoteLinks

`func (o *SubmitRemoteLinks202Response) SetAcceptedRemoteLinks(v []string)`

SetAcceptedRemoteLinks sets AcceptedRemoteLinks field to given value.

### HasAcceptedRemoteLinks

`func (o *SubmitRemoteLinks202Response) HasAcceptedRemoteLinks() bool`

HasAcceptedRemoteLinks returns a boolean if a field has been set.

### GetRejectedRemoteLinks

`func (o *SubmitRemoteLinks202Response) GetRejectedRemoteLinks() map[string][]SubmitRemoteLinks202ResponseRejectedRemoteLinksValueInner`

GetRejectedRemoteLinks returns the RejectedRemoteLinks field if non-nil, zero value otherwise.

### GetRejectedRemoteLinksOk

`func (o *SubmitRemoteLinks202Response) GetRejectedRemoteLinksOk() (*map[string][]SubmitRemoteLinks202ResponseRejectedRemoteLinksValueInner, bool)`

GetRejectedRemoteLinksOk returns a tuple with the RejectedRemoteLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedRemoteLinks

`func (o *SubmitRemoteLinks202Response) SetRejectedRemoteLinks(v map[string][]SubmitRemoteLinks202ResponseRejectedRemoteLinksValueInner)`

SetRejectedRemoteLinks sets RejectedRemoteLinks field to given value.

### HasRejectedRemoteLinks

`func (o *SubmitRemoteLinks202Response) HasRejectedRemoteLinks() bool`

HasRejectedRemoteLinks returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *SubmitRemoteLinks202Response) GetUnknownAssociations() []RemoteLinkDataAssociationsInner`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *SubmitRemoteLinks202Response) GetUnknownAssociationsOk() (*[]RemoteLinkDataAssociationsInner, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *SubmitRemoteLinks202Response) SetUnknownAssociations(v []RemoteLinkDataAssociationsInner)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *SubmitRemoteLinks202Response) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


