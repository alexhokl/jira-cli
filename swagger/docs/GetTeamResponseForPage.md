# GetTeamResponseForPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The team ID. | 
**Name** | Pointer to **string** | The team name. This is returned if the type is \&quot;PlanOnly\&quot;. | [optional] 
**Type** | **string** | The team type. This is \&quot;PlanOnly\&quot; or \&quot;Atlassian\&quot;. | 

## Methods

### NewGetTeamResponseForPage

`func NewGetTeamResponseForPage(id string, type_ string, ) *GetTeamResponseForPage`

NewGetTeamResponseForPage instantiates a new GetTeamResponseForPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamResponseForPageWithDefaults

`func NewGetTeamResponseForPageWithDefaults() *GetTeamResponseForPage`

NewGetTeamResponseForPageWithDefaults instantiates a new GetTeamResponseForPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTeamResponseForPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamResponseForPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamResponseForPage) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetTeamResponseForPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTeamResponseForPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTeamResponseForPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTeamResponseForPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetTeamResponseForPage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTeamResponseForPage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTeamResponseForPage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


