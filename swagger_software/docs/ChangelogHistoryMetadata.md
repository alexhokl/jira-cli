# ChangelogHistoryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityDescription** | Pointer to **string** | The activity described in the history record. | [optional] 
**ActivityDescriptionKey** | Pointer to **string** | The key of the activity described in the history record. | [optional] 
**Actor** | Pointer to [**ChangelogHistoryMetadataAllOfActor**](ChangelogHistoryMetadataAllOfActor.md) |  | [optional] 
**Cause** | Pointer to [**ChangelogHistoryMetadataAllOfCause**](ChangelogHistoryMetadataAllOfCause.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the history record. | [optional] 
**DescriptionKey** | Pointer to **string** | The description key of the history record. | [optional] 
**EmailDescription** | Pointer to **string** | The description of the email address associated the history record. | [optional] 
**EmailDescriptionKey** | Pointer to **string** | The description key of the email address associated the history record. | [optional] 
**ExtraData** | Pointer to **map[string]string** | Additional arbitrary information about the history record. | [optional] 
**Generator** | Pointer to [**ChangelogHistoryMetadataAllOfGenerator**](ChangelogHistoryMetadataAllOfGenerator.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the history record. | [optional] 

## Methods

### NewChangelogHistoryMetadata

`func NewChangelogHistoryMetadata() *ChangelogHistoryMetadata`

NewChangelogHistoryMetadata instantiates a new ChangelogHistoryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogHistoryMetadataWithDefaults

`func NewChangelogHistoryMetadataWithDefaults() *ChangelogHistoryMetadata`

NewChangelogHistoryMetadataWithDefaults instantiates a new ChangelogHistoryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityDescription

`func (o *ChangelogHistoryMetadata) GetActivityDescription() string`

GetActivityDescription returns the ActivityDescription field if non-nil, zero value otherwise.

### GetActivityDescriptionOk

`func (o *ChangelogHistoryMetadata) GetActivityDescriptionOk() (*string, bool)`

GetActivityDescriptionOk returns a tuple with the ActivityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDescription

`func (o *ChangelogHistoryMetadata) SetActivityDescription(v string)`

SetActivityDescription sets ActivityDescription field to given value.

### HasActivityDescription

`func (o *ChangelogHistoryMetadata) HasActivityDescription() bool`

HasActivityDescription returns a boolean if a field has been set.

### GetActivityDescriptionKey

`func (o *ChangelogHistoryMetadata) GetActivityDescriptionKey() string`

GetActivityDescriptionKey returns the ActivityDescriptionKey field if non-nil, zero value otherwise.

### GetActivityDescriptionKeyOk

`func (o *ChangelogHistoryMetadata) GetActivityDescriptionKeyOk() (*string, bool)`

GetActivityDescriptionKeyOk returns a tuple with the ActivityDescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDescriptionKey

`func (o *ChangelogHistoryMetadata) SetActivityDescriptionKey(v string)`

SetActivityDescriptionKey sets ActivityDescriptionKey field to given value.

### HasActivityDescriptionKey

`func (o *ChangelogHistoryMetadata) HasActivityDescriptionKey() bool`

HasActivityDescriptionKey returns a boolean if a field has been set.

### GetActor

`func (o *ChangelogHistoryMetadata) GetActor() ChangelogHistoryMetadataAllOfActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ChangelogHistoryMetadata) GetActorOk() (*ChangelogHistoryMetadataAllOfActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ChangelogHistoryMetadata) SetActor(v ChangelogHistoryMetadataAllOfActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *ChangelogHistoryMetadata) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCause

`func (o *ChangelogHistoryMetadata) GetCause() ChangelogHistoryMetadataAllOfCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ChangelogHistoryMetadata) GetCauseOk() (*ChangelogHistoryMetadataAllOfCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ChangelogHistoryMetadata) SetCause(v ChangelogHistoryMetadataAllOfCause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ChangelogHistoryMetadata) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetDescription

`func (o *ChangelogHistoryMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangelogHistoryMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangelogHistoryMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangelogHistoryMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionKey

`func (o *ChangelogHistoryMetadata) GetDescriptionKey() string`

GetDescriptionKey returns the DescriptionKey field if non-nil, zero value otherwise.

### GetDescriptionKeyOk

`func (o *ChangelogHistoryMetadata) GetDescriptionKeyOk() (*string, bool)`

GetDescriptionKeyOk returns a tuple with the DescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionKey

`func (o *ChangelogHistoryMetadata) SetDescriptionKey(v string)`

SetDescriptionKey sets DescriptionKey field to given value.

### HasDescriptionKey

`func (o *ChangelogHistoryMetadata) HasDescriptionKey() bool`

HasDescriptionKey returns a boolean if a field has been set.

### GetEmailDescription

`func (o *ChangelogHistoryMetadata) GetEmailDescription() string`

GetEmailDescription returns the EmailDescription field if non-nil, zero value otherwise.

### GetEmailDescriptionOk

`func (o *ChangelogHistoryMetadata) GetEmailDescriptionOk() (*string, bool)`

GetEmailDescriptionOk returns a tuple with the EmailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDescription

`func (o *ChangelogHistoryMetadata) SetEmailDescription(v string)`

SetEmailDescription sets EmailDescription field to given value.

### HasEmailDescription

`func (o *ChangelogHistoryMetadata) HasEmailDescription() bool`

HasEmailDescription returns a boolean if a field has been set.

### GetEmailDescriptionKey

`func (o *ChangelogHistoryMetadata) GetEmailDescriptionKey() string`

GetEmailDescriptionKey returns the EmailDescriptionKey field if non-nil, zero value otherwise.

### GetEmailDescriptionKeyOk

`func (o *ChangelogHistoryMetadata) GetEmailDescriptionKeyOk() (*string, bool)`

GetEmailDescriptionKeyOk returns a tuple with the EmailDescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDescriptionKey

`func (o *ChangelogHistoryMetadata) SetEmailDescriptionKey(v string)`

SetEmailDescriptionKey sets EmailDescriptionKey field to given value.

### HasEmailDescriptionKey

`func (o *ChangelogHistoryMetadata) HasEmailDescriptionKey() bool`

HasEmailDescriptionKey returns a boolean if a field has been set.

### GetExtraData

`func (o *ChangelogHistoryMetadata) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChangelogHistoryMetadata) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChangelogHistoryMetadata) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChangelogHistoryMetadata) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetGenerator

`func (o *ChangelogHistoryMetadata) GetGenerator() ChangelogHistoryMetadataAllOfGenerator`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *ChangelogHistoryMetadata) GetGeneratorOk() (*ChangelogHistoryMetadataAllOfGenerator, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *ChangelogHistoryMetadata) SetGenerator(v ChangelogHistoryMetadataAllOfGenerator)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *ChangelogHistoryMetadata) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetType

`func (o *ChangelogHistoryMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChangelogHistoryMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChangelogHistoryMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChangelogHistoryMetadata) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


