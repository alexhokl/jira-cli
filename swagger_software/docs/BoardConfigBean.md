# BoardConfigBean

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

### NewBoardConfigBean

`func NewBoardConfigBean() *BoardConfigBean`

NewBoardConfigBean instantiates a new BoardConfigBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardConfigBeanWithDefaults

`func NewBoardConfigBeanWithDefaults() *BoardConfigBean`

NewBoardConfigBeanWithDefaults instantiates a new BoardConfigBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnConfig

`func (o *BoardConfigBean) GetColumnConfig() GetConfiguration200ResponseColumnConfig`

GetColumnConfig returns the ColumnConfig field if non-nil, zero value otherwise.

### GetColumnConfigOk

`func (o *BoardConfigBean) GetColumnConfigOk() (*GetConfiguration200ResponseColumnConfig, bool)`

GetColumnConfigOk returns a tuple with the ColumnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnConfig

`func (o *BoardConfigBean) SetColumnConfig(v GetConfiguration200ResponseColumnConfig)`

SetColumnConfig sets ColumnConfig field to given value.

### HasColumnConfig

`func (o *BoardConfigBean) HasColumnConfig() bool`

HasColumnConfig returns a boolean if a field has been set.

### GetEstimation

`func (o *BoardConfigBean) GetEstimation() GetConfiguration200ResponseEstimation`

GetEstimation returns the Estimation field if non-nil, zero value otherwise.

### GetEstimationOk

`func (o *BoardConfigBean) GetEstimationOk() (*GetConfiguration200ResponseEstimation, bool)`

GetEstimationOk returns a tuple with the Estimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimation

`func (o *BoardConfigBean) SetEstimation(v GetConfiguration200ResponseEstimation)`

SetEstimation sets Estimation field to given value.

### HasEstimation

`func (o *BoardConfigBean) HasEstimation() bool`

HasEstimation returns a boolean if a field has been set.

### GetFilter

`func (o *BoardConfigBean) GetFilter() GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BoardConfigBean) GetFilterOk() (*GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BoardConfigBean) SetFilter(v GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BoardConfigBean) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *BoardConfigBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardConfigBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardConfigBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BoardConfigBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *BoardConfigBean) GetLocation() CreateBoardRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BoardConfigBean) GetLocationOk() (*CreateBoardRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BoardConfigBean) SetLocation(v CreateBoardRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BoardConfigBean) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *BoardConfigBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardConfigBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardConfigBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardConfigBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRanking

`func (o *BoardConfigBean) GetRanking() GetConfiguration200ResponseRanking`

GetRanking returns the Ranking field if non-nil, zero value otherwise.

### GetRankingOk

`func (o *BoardConfigBean) GetRankingOk() (*GetConfiguration200ResponseRanking, bool)`

GetRankingOk returns a tuple with the Ranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanking

`func (o *BoardConfigBean) SetRanking(v GetConfiguration200ResponseRanking)`

SetRanking sets Ranking field to given value.

### HasRanking

`func (o *BoardConfigBean) HasRanking() bool`

HasRanking returns a boolean if a field has been set.

### GetSelf

`func (o *BoardConfigBean) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *BoardConfigBean) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *BoardConfigBean) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *BoardConfigBean) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubQuery

`func (o *BoardConfigBean) GetSubQuery() GetConfiguration200ResponseSubQuery`

GetSubQuery returns the SubQuery field if non-nil, zero value otherwise.

### GetSubQueryOk

`func (o *BoardConfigBean) GetSubQueryOk() (*GetConfiguration200ResponseSubQuery, bool)`

GetSubQueryOk returns a tuple with the SubQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubQuery

`func (o *BoardConfigBean) SetSubQuery(v GetConfiguration200ResponseSubQuery)`

SetSubQuery sets SubQuery field to given value.

### HasSubQuery

`func (o *BoardConfigBean) HasSubQuery() bool`

HasSubQuery returns a boolean if a field has been set.

### GetType

`func (o *BoardConfigBean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardConfigBean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardConfigBean) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BoardConfigBean) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


