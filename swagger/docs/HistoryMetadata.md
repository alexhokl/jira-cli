# HistoryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityDescription** | Pointer to **string** | The activity described in the history record. | [optional] 
**ActivityDescriptionKey** | Pointer to **string** | The key of the activity described in the history record. | [optional] 
**Actor** | Pointer to [**HistoryMetadataParticipant**](HistoryMetadataParticipant.md) | Details of the user whose action created the history record. | [optional] 
**Cause** | Pointer to [**HistoryMetadataParticipant**](HistoryMetadataParticipant.md) | Details of the cause that triggered the creation the history record. | [optional] 
**Description** | Pointer to **string** | The description of the history record. | [optional] 
**DescriptionKey** | Pointer to **string** | The description key of the history record. | [optional] 
**EmailDescription** | Pointer to **string** | The description of the email address associated the history record. | [optional] 
**EmailDescriptionKey** | Pointer to **string** | The description key of the email address associated the history record. | [optional] 
**ExtraData** | Pointer to **map[string]string** | Additional arbitrary information about the history record. | [optional] 
**Generator** | Pointer to [**HistoryMetadataParticipant**](HistoryMetadataParticipant.md) | Details of the system that generated the history record. | [optional] 
**Type** | Pointer to **string** | The type of the history record. | [optional] 

## Methods

### NewHistoryMetadata

`func NewHistoryMetadata() *HistoryMetadata`

NewHistoryMetadata instantiates a new HistoryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryMetadataWithDefaults

`func NewHistoryMetadataWithDefaults() *HistoryMetadata`

NewHistoryMetadataWithDefaults instantiates a new HistoryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityDescription

`func (o *HistoryMetadata) GetActivityDescription() string`

GetActivityDescription returns the ActivityDescription field if non-nil, zero value otherwise.

### GetActivityDescriptionOk

`func (o *HistoryMetadata) GetActivityDescriptionOk() (*string, bool)`

GetActivityDescriptionOk returns a tuple with the ActivityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDescription

`func (o *HistoryMetadata) SetActivityDescription(v string)`

SetActivityDescription sets ActivityDescription field to given value.

### HasActivityDescription

`func (o *HistoryMetadata) HasActivityDescription() bool`

HasActivityDescription returns a boolean if a field has been set.

### GetActivityDescriptionKey

`func (o *HistoryMetadata) GetActivityDescriptionKey() string`

GetActivityDescriptionKey returns the ActivityDescriptionKey field if non-nil, zero value otherwise.

### GetActivityDescriptionKeyOk

`func (o *HistoryMetadata) GetActivityDescriptionKeyOk() (*string, bool)`

GetActivityDescriptionKeyOk returns a tuple with the ActivityDescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDescriptionKey

`func (o *HistoryMetadata) SetActivityDescriptionKey(v string)`

SetActivityDescriptionKey sets ActivityDescriptionKey field to given value.

### HasActivityDescriptionKey

`func (o *HistoryMetadata) HasActivityDescriptionKey() bool`

HasActivityDescriptionKey returns a boolean if a field has been set.

### GetActor

`func (o *HistoryMetadata) GetActor() HistoryMetadataParticipant`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *HistoryMetadata) GetActorOk() (*HistoryMetadataParticipant, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *HistoryMetadata) SetActor(v HistoryMetadataParticipant)`

SetActor sets Actor field to given value.

### HasActor

`func (o *HistoryMetadata) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCause

`func (o *HistoryMetadata) GetCause() HistoryMetadataParticipant`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HistoryMetadata) GetCauseOk() (*HistoryMetadataParticipant, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HistoryMetadata) SetCause(v HistoryMetadataParticipant)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HistoryMetadata) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetDescription

`func (o *HistoryMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HistoryMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionKey

`func (o *HistoryMetadata) GetDescriptionKey() string`

GetDescriptionKey returns the DescriptionKey field if non-nil, zero value otherwise.

### GetDescriptionKeyOk

`func (o *HistoryMetadata) GetDescriptionKeyOk() (*string, bool)`

GetDescriptionKeyOk returns a tuple with the DescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionKey

`func (o *HistoryMetadata) SetDescriptionKey(v string)`

SetDescriptionKey sets DescriptionKey field to given value.

### HasDescriptionKey

`func (o *HistoryMetadata) HasDescriptionKey() bool`

HasDescriptionKey returns a boolean if a field has been set.

### GetEmailDescription

`func (o *HistoryMetadata) GetEmailDescription() string`

GetEmailDescription returns the EmailDescription field if non-nil, zero value otherwise.

### GetEmailDescriptionOk

`func (o *HistoryMetadata) GetEmailDescriptionOk() (*string, bool)`

GetEmailDescriptionOk returns a tuple with the EmailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDescription

`func (o *HistoryMetadata) SetEmailDescription(v string)`

SetEmailDescription sets EmailDescription field to given value.

### HasEmailDescription

`func (o *HistoryMetadata) HasEmailDescription() bool`

HasEmailDescription returns a boolean if a field has been set.

### GetEmailDescriptionKey

`func (o *HistoryMetadata) GetEmailDescriptionKey() string`

GetEmailDescriptionKey returns the EmailDescriptionKey field if non-nil, zero value otherwise.

### GetEmailDescriptionKeyOk

`func (o *HistoryMetadata) GetEmailDescriptionKeyOk() (*string, bool)`

GetEmailDescriptionKeyOk returns a tuple with the EmailDescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDescriptionKey

`func (o *HistoryMetadata) SetEmailDescriptionKey(v string)`

SetEmailDescriptionKey sets EmailDescriptionKey field to given value.

### HasEmailDescriptionKey

`func (o *HistoryMetadata) HasEmailDescriptionKey() bool`

HasEmailDescriptionKey returns a boolean if a field has been set.

### GetExtraData

`func (o *HistoryMetadata) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *HistoryMetadata) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *HistoryMetadata) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *HistoryMetadata) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetGenerator

`func (o *HistoryMetadata) GetGenerator() HistoryMetadataParticipant`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *HistoryMetadata) GetGeneratorOk() (*HistoryMetadataParticipant, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *HistoryMetadata) SetGenerator(v HistoryMetadataParticipant)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *HistoryMetadata) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetType

`func (o *HistoryMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoryMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoryMetadata) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


