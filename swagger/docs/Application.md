# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the application. Used in conjunction with the (remote) object icon title to display a tooltip for the link&#39;s icon. The tooltip takes the format \&quot;\\[application name\\] icon title\&quot;. Blank items are excluded from the tooltip title. If both items are blank, the icon tooltop displays as \&quot;Web Link\&quot;. Grouping and sorting of links may place links without an application name last. | [optional] 
**Type** | Pointer to **string** | The name-spaced type of the application, used by registered rendering apps. | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Application) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Application) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Application) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Application) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Application) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


