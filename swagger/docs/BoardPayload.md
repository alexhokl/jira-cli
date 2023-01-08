# BoardPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardFilterJQL** | Pointer to **string** | Takes in a JQL string to create a new filter. If no value is provided, it&#39;ll default to a JQL filter for the project creating | [optional] 
**CardColorStrategy** | Pointer to **string** | Card color settings of the board | [optional] 
**CardLayout** | Pointer to [**CardLayout**](CardLayout.md) |  | [optional] 
**CardLayouts** | Pointer to [**[]CardLayoutField**](CardLayoutField.md) | Card layout settings of the board | [optional] 
**Columns** | Pointer to [**[]BoardColumnPayload**](BoardColumnPayload.md) | The columns of the board | [optional] 
**Features** | Pointer to [**[]BoardFeaturePayload**](BoardFeaturePayload.md) | Feature settings for the board | [optional] 
**Name** | Pointer to **string** | The name of the board | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**QuickFilters** | Pointer to [**[]QuickFilterPayload**](QuickFilterPayload.md) | The quick filters for the board. | [optional] 
**SupportsSprint** | Pointer to **bool** | Whether sprints are supported on the board | [optional] [default to true]
**Swimlanes** | Pointer to [**SwimlanesPayload**](SwimlanesPayload.md) |  | [optional] 
**WorkingDaysConfig** | Pointer to [**WorkingDaysConfig**](WorkingDaysConfig.md) |  | [optional] 

## Methods

### NewBoardPayload

`func NewBoardPayload() *BoardPayload`

NewBoardPayload instantiates a new BoardPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardPayloadWithDefaults

`func NewBoardPayloadWithDefaults() *BoardPayload`

NewBoardPayloadWithDefaults instantiates a new BoardPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardFilterJQL

`func (o *BoardPayload) GetBoardFilterJQL() string`

GetBoardFilterJQL returns the BoardFilterJQL field if non-nil, zero value otherwise.

### GetBoardFilterJQLOk

`func (o *BoardPayload) GetBoardFilterJQLOk() (*string, bool)`

GetBoardFilterJQLOk returns a tuple with the BoardFilterJQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardFilterJQL

`func (o *BoardPayload) SetBoardFilterJQL(v string)`

SetBoardFilterJQL sets BoardFilterJQL field to given value.

### HasBoardFilterJQL

`func (o *BoardPayload) HasBoardFilterJQL() bool`

HasBoardFilterJQL returns a boolean if a field has been set.

### GetCardColorStrategy

`func (o *BoardPayload) GetCardColorStrategy() string`

GetCardColorStrategy returns the CardColorStrategy field if non-nil, zero value otherwise.

### GetCardColorStrategyOk

`func (o *BoardPayload) GetCardColorStrategyOk() (*string, bool)`

GetCardColorStrategyOk returns a tuple with the CardColorStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardColorStrategy

`func (o *BoardPayload) SetCardColorStrategy(v string)`

SetCardColorStrategy sets CardColorStrategy field to given value.

### HasCardColorStrategy

`func (o *BoardPayload) HasCardColorStrategy() bool`

HasCardColorStrategy returns a boolean if a field has been set.

### GetCardLayout

`func (o *BoardPayload) GetCardLayout() CardLayout`

GetCardLayout returns the CardLayout field if non-nil, zero value otherwise.

### GetCardLayoutOk

`func (o *BoardPayload) GetCardLayoutOk() (*CardLayout, bool)`

GetCardLayoutOk returns a tuple with the CardLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLayout

`func (o *BoardPayload) SetCardLayout(v CardLayout)`

SetCardLayout sets CardLayout field to given value.

### HasCardLayout

`func (o *BoardPayload) HasCardLayout() bool`

HasCardLayout returns a boolean if a field has been set.

### GetCardLayouts

`func (o *BoardPayload) GetCardLayouts() []CardLayoutField`

GetCardLayouts returns the CardLayouts field if non-nil, zero value otherwise.

### GetCardLayoutsOk

`func (o *BoardPayload) GetCardLayoutsOk() (*[]CardLayoutField, bool)`

GetCardLayoutsOk returns a tuple with the CardLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLayouts

`func (o *BoardPayload) SetCardLayouts(v []CardLayoutField)`

SetCardLayouts sets CardLayouts field to given value.

### HasCardLayouts

`func (o *BoardPayload) HasCardLayouts() bool`

HasCardLayouts returns a boolean if a field has been set.

### GetColumns

`func (o *BoardPayload) GetColumns() []BoardColumnPayload`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *BoardPayload) GetColumnsOk() (*[]BoardColumnPayload, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *BoardPayload) SetColumns(v []BoardColumnPayload)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *BoardPayload) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFeatures

`func (o *BoardPayload) GetFeatures() []BoardFeaturePayload`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BoardPayload) GetFeaturesOk() (*[]BoardFeaturePayload, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BoardPayload) SetFeatures(v []BoardFeaturePayload)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BoardPayload) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetName

`func (o *BoardPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *BoardPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *BoardPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *BoardPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *BoardPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetQuickFilters

`func (o *BoardPayload) GetQuickFilters() []QuickFilterPayload`

GetQuickFilters returns the QuickFilters field if non-nil, zero value otherwise.

### GetQuickFiltersOk

`func (o *BoardPayload) GetQuickFiltersOk() (*[]QuickFilterPayload, bool)`

GetQuickFiltersOk returns a tuple with the QuickFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickFilters

`func (o *BoardPayload) SetQuickFilters(v []QuickFilterPayload)`

SetQuickFilters sets QuickFilters field to given value.

### HasQuickFilters

`func (o *BoardPayload) HasQuickFilters() bool`

HasQuickFilters returns a boolean if a field has been set.

### GetSupportsSprint

`func (o *BoardPayload) GetSupportsSprint() bool`

GetSupportsSprint returns the SupportsSprint field if non-nil, zero value otherwise.

### GetSupportsSprintOk

`func (o *BoardPayload) GetSupportsSprintOk() (*bool, bool)`

GetSupportsSprintOk returns a tuple with the SupportsSprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSprint

`func (o *BoardPayload) SetSupportsSprint(v bool)`

SetSupportsSprint sets SupportsSprint field to given value.

### HasSupportsSprint

`func (o *BoardPayload) HasSupportsSprint() bool`

HasSupportsSprint returns a boolean if a field has been set.

### GetSwimlanes

`func (o *BoardPayload) GetSwimlanes() SwimlanesPayload`

GetSwimlanes returns the Swimlanes field if non-nil, zero value otherwise.

### GetSwimlanesOk

`func (o *BoardPayload) GetSwimlanesOk() (*SwimlanesPayload, bool)`

GetSwimlanesOk returns a tuple with the Swimlanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwimlanes

`func (o *BoardPayload) SetSwimlanes(v SwimlanesPayload)`

SetSwimlanes sets Swimlanes field to given value.

### HasSwimlanes

`func (o *BoardPayload) HasSwimlanes() bool`

HasSwimlanes returns a boolean if a field has been set.

### GetWorkingDaysConfig

`func (o *BoardPayload) GetWorkingDaysConfig() WorkingDaysConfig`

GetWorkingDaysConfig returns the WorkingDaysConfig field if non-nil, zero value otherwise.

### GetWorkingDaysConfigOk

`func (o *BoardPayload) GetWorkingDaysConfigOk() (*WorkingDaysConfig, bool)`

GetWorkingDaysConfigOk returns a tuple with the WorkingDaysConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDaysConfig

`func (o *BoardPayload) SetWorkingDaysConfig(v WorkingDaysConfig)`

SetWorkingDaysConfig sets WorkingDaysConfig field to given value.

### HasWorkingDaysConfig

`func (o *BoardPayload) HasWorkingDaysConfig() bool`

HasWorkingDaysConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


