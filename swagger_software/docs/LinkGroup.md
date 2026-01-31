# LinkGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]LinkGroup**](LinkGroup.md) |  | [optional] 
**Header** | Pointer to [**LinkGroupHeader**](LinkGroupHeader.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]LinkGroupHeader**](LinkGroupHeader.md) |  | [optional] 
**StyleClass** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewLinkGroup

`func NewLinkGroup() *LinkGroup`

NewLinkGroup instantiates a new LinkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkGroupWithDefaults

`func NewLinkGroupWithDefaults() *LinkGroup`

NewLinkGroupWithDefaults instantiates a new LinkGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *LinkGroup) GetGroups() []LinkGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LinkGroup) GetGroupsOk() (*[]LinkGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LinkGroup) SetGroups(v []LinkGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *LinkGroup) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHeader

`func (o *LinkGroup) GetHeader() LinkGroupHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *LinkGroup) GetHeaderOk() (*LinkGroupHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *LinkGroup) SetHeader(v LinkGroupHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *LinkGroup) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetId

`func (o *LinkGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *LinkGroup) GetLinks() []LinkGroupHeader`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkGroup) GetLinksOk() (*[]LinkGroupHeader, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkGroup) SetLinks(v []LinkGroupHeader)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStyleClass

`func (o *LinkGroup) GetStyleClass() string`

GetStyleClass returns the StyleClass field if non-nil, zero value otherwise.

### GetStyleClassOk

`func (o *LinkGroup) GetStyleClassOk() (*string, bool)`

GetStyleClassOk returns a tuple with the StyleClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleClass

`func (o *LinkGroup) SetStyleClass(v string)`

SetStyleClass sets StyleClass field to given value.

### HasStyleClass

`func (o *LinkGroup) HasStyleClass() bool`

HasStyleClass returns a boolean if a field has been set.

### GetWeight

`func (o *LinkGroup) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *LinkGroup) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *LinkGroup) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *LinkGroup) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


