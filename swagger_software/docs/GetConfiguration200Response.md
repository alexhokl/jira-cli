# GetConfiguration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnConfig** | Pointer to [**GetConfiguration200ResponseColumnConfig**](GetConfiguration200ResponseColumnConfig.md) |  | [optional] 
**Estimation** | Pointer to [**GetConfiguration200ResponseEstimation**](GetConfiguration200ResponseEstimation.md) |  | [optional] 
**Filter** | Pointer to [**GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner**](GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**CreateBoardRequestLocation**](CreateBoardRequestLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ranking** | Pointer to [**GetConfiguration200ResponseRanking**](GetConfiguration200ResponseRanking.md) |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**SubQuery** | Pointer to [**GetConfiguration200ResponseSubQuery**](GetConfiguration200ResponseSubQuery.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetConfiguration200Response

`func NewGetConfiguration200Response() *GetConfiguration200Response`

NewGetConfiguration200Response instantiates a new GetConfiguration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfiguration200ResponseWithDefaults

`func NewGetConfiguration200ResponseWithDefaults() *GetConfiguration200Response`

NewGetConfiguration200ResponseWithDefaults instantiates a new GetConfiguration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnConfig

`func (o *GetConfiguration200Response) GetColumnConfig() GetConfiguration200ResponseColumnConfig`

GetColumnConfig returns the ColumnConfig field if non-nil, zero value otherwise.

### GetColumnConfigOk

`func (o *GetConfiguration200Response) GetColumnConfigOk() (*GetConfiguration200ResponseColumnConfig, bool)`

GetColumnConfigOk returns a tuple with the ColumnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnConfig

`func (o *GetConfiguration200Response) SetColumnConfig(v GetConfiguration200ResponseColumnConfig)`

SetColumnConfig sets ColumnConfig field to given value.

### HasColumnConfig

`func (o *GetConfiguration200Response) HasColumnConfig() bool`

HasColumnConfig returns a boolean if a field has been set.

### GetEstimation

`func (o *GetConfiguration200Response) GetEstimation() GetConfiguration200ResponseEstimation`

GetEstimation returns the Estimation field if non-nil, zero value otherwise.

### GetEstimationOk

`func (o *GetConfiguration200Response) GetEstimationOk() (*GetConfiguration200ResponseEstimation, bool)`

GetEstimationOk returns a tuple with the Estimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimation

`func (o *GetConfiguration200Response) SetEstimation(v GetConfiguration200ResponseEstimation)`

SetEstimation sets Estimation field to given value.

### HasEstimation

`func (o *GetConfiguration200Response) HasEstimation() bool`

HasEstimation returns a boolean if a field has been set.

### GetFilter

`func (o *GetConfiguration200Response) GetFilter() GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetConfiguration200Response) GetFilterOk() (*GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetConfiguration200Response) SetFilter(v GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetConfiguration200Response) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *GetConfiguration200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConfiguration200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConfiguration200Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetConfiguration200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *GetConfiguration200Response) GetLocation() CreateBoardRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetConfiguration200Response) GetLocationOk() (*CreateBoardRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetConfiguration200Response) SetLocation(v CreateBoardRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetConfiguration200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *GetConfiguration200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetConfiguration200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetConfiguration200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetConfiguration200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRanking

`func (o *GetConfiguration200Response) GetRanking() GetConfiguration200ResponseRanking`

GetRanking returns the Ranking field if non-nil, zero value otherwise.

### GetRankingOk

`func (o *GetConfiguration200Response) GetRankingOk() (*GetConfiguration200ResponseRanking, bool)`

GetRankingOk returns a tuple with the Ranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanking

`func (o *GetConfiguration200Response) SetRanking(v GetConfiguration200ResponseRanking)`

SetRanking sets Ranking field to given value.

### HasRanking

`func (o *GetConfiguration200Response) HasRanking() bool`

HasRanking returns a boolean if a field has been set.

### GetSelf

`func (o *GetConfiguration200Response) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GetConfiguration200Response) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GetConfiguration200Response) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GetConfiguration200Response) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubQuery

`func (o *GetConfiguration200Response) GetSubQuery() GetConfiguration200ResponseSubQuery`

GetSubQuery returns the SubQuery field if non-nil, zero value otherwise.

### GetSubQueryOk

`func (o *GetConfiguration200Response) GetSubQueryOk() (*GetConfiguration200ResponseSubQuery, bool)`

GetSubQueryOk returns a tuple with the SubQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubQuery

`func (o *GetConfiguration200Response) SetSubQuery(v GetConfiguration200ResponseSubQuery)`

SetSubQuery sets SubQuery field to given value.

### HasSubQuery

`func (o *GetConfiguration200Response) HasSubQuery() bool`

HasSubQuery returns a boolean if a field has been set.

### GetType

`func (o *GetConfiguration200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetConfiguration200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetConfiguration200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetConfiguration200Response) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


