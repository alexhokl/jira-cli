# IssueBeanTransitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional transition details in the response. | [optional] [readonly] 
**Fields** | Pointer to [**map[string]IssueBeanEditmetaAllOfFieldsValue**](IssueBeanEditmetaAllOfFieldsValue.md) | Details of the fields associated with the issue transition screen. Use this information to populate &#x60;fields&#x60; and &#x60;update&#x60; in a transition request. | [optional] [readonly] 
**HasScreen** | Pointer to **bool** | Whether there is a screen associated with the issue transition. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the issue transition. Required when specifying a transition to undertake. | [optional] 
**IsAvailable** | Pointer to **bool** | Whether the transition is available to be performed. | [optional] [readonly] 
**IsConditional** | Pointer to **bool** | Whether the issue has to meet criteria before the issue transition is applied. | [optional] [readonly] 
**IsGlobal** | Pointer to **bool** | Whether the issue transition is global, that is, the transition is applied to issues regardless of their status. | [optional] [readonly] 
**IsInitial** | Pointer to **bool** | Whether this is the initial issue transition for the workflow. | [optional] [readonly] 
**Looped** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The name of the issue transition. | [optional] [readonly] 
**To** | Pointer to [**IssueBeanTransitionsInnerTo**](IssueBeanTransitionsInnerTo.md) |  | [optional] 

## Methods

### NewIssueBeanTransitionsInner

`func NewIssueBeanTransitionsInner() *IssueBeanTransitionsInner`

NewIssueBeanTransitionsInner instantiates a new IssueBeanTransitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBeanTransitionsInnerWithDefaults

`func NewIssueBeanTransitionsInnerWithDefaults() *IssueBeanTransitionsInner`

NewIssueBeanTransitionsInnerWithDefaults instantiates a new IssueBeanTransitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *IssueBeanTransitionsInner) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *IssueBeanTransitionsInner) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *IssueBeanTransitionsInner) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *IssueBeanTransitionsInner) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *IssueBeanTransitionsInner) GetFields() map[string]IssueBeanEditmetaAllOfFieldsValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IssueBeanTransitionsInner) GetFieldsOk() (*map[string]IssueBeanEditmetaAllOfFieldsValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IssueBeanTransitionsInner) SetFields(v map[string]IssueBeanEditmetaAllOfFieldsValue)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IssueBeanTransitionsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHasScreen

`func (o *IssueBeanTransitionsInner) GetHasScreen() bool`

GetHasScreen returns the HasScreen field if non-nil, zero value otherwise.

### GetHasScreenOk

`func (o *IssueBeanTransitionsInner) GetHasScreenOk() (*bool, bool)`

GetHasScreenOk returns a tuple with the HasScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScreen

`func (o *IssueBeanTransitionsInner) SetHasScreen(v bool)`

SetHasScreen sets HasScreen field to given value.

### HasHasScreen

`func (o *IssueBeanTransitionsInner) HasHasScreen() bool`

HasHasScreen returns a boolean if a field has been set.

### GetId

`func (o *IssueBeanTransitionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueBeanTransitionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueBeanTransitionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueBeanTransitionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAvailable

`func (o *IssueBeanTransitionsInner) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *IssueBeanTransitionsInner) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *IssueBeanTransitionsInner) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *IssueBeanTransitionsInner) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetIsConditional

`func (o *IssueBeanTransitionsInner) GetIsConditional() bool`

GetIsConditional returns the IsConditional field if non-nil, zero value otherwise.

### GetIsConditionalOk

`func (o *IssueBeanTransitionsInner) GetIsConditionalOk() (*bool, bool)`

GetIsConditionalOk returns a tuple with the IsConditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConditional

`func (o *IssueBeanTransitionsInner) SetIsConditional(v bool)`

SetIsConditional sets IsConditional field to given value.

### HasIsConditional

`func (o *IssueBeanTransitionsInner) HasIsConditional() bool`

HasIsConditional returns a boolean if a field has been set.

### GetIsGlobal

`func (o *IssueBeanTransitionsInner) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *IssueBeanTransitionsInner) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *IssueBeanTransitionsInner) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *IssueBeanTransitionsInner) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### GetIsInitial

`func (o *IssueBeanTransitionsInner) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *IssueBeanTransitionsInner) GetIsInitialOk() (*bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitial

`func (o *IssueBeanTransitionsInner) SetIsInitial(v bool)`

SetIsInitial sets IsInitial field to given value.

### HasIsInitial

`func (o *IssueBeanTransitionsInner) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### GetLooped

`func (o *IssueBeanTransitionsInner) GetLooped() bool`

GetLooped returns the Looped field if non-nil, zero value otherwise.

### GetLoopedOk

`func (o *IssueBeanTransitionsInner) GetLoopedOk() (*bool, bool)`

GetLoopedOk returns a tuple with the Looped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLooped

`func (o *IssueBeanTransitionsInner) SetLooped(v bool)`

SetLooped sets Looped field to given value.

### HasLooped

`func (o *IssueBeanTransitionsInner) HasLooped() bool`

HasLooped returns a boolean if a field has been set.

### GetName

`func (o *IssueBeanTransitionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueBeanTransitionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueBeanTransitionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueBeanTransitionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTo

`func (o *IssueBeanTransitionsInner) GetTo() IssueBeanTransitionsInnerTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IssueBeanTransitionsInner) GetToOk() (*IssueBeanTransitionsInnerTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IssueBeanTransitionsInner) SetTo(v IssueBeanTransitionsInnerTo)`

SetTo sets To field to given value.

### HasTo

`func (o *IssueBeanTransitionsInner) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


