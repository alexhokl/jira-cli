# ConnectModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modules** | **[]map[string]interface{}** | A list of app modules in the same format as the &#x60;modules&#x60; property in the [app descriptor](https://developer.atlassian.com/cloud/jira/platform/app-descriptor/). | 

## Methods

### NewConnectModules

`func NewConnectModules(modules []map[string]interface{}, ) *ConnectModules`

NewConnectModules instantiates a new ConnectModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectModulesWithDefaults

`func NewConnectModulesWithDefaults() *ConnectModules`

NewConnectModulesWithDefaults instantiates a new ConnectModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModules

`func (o *ConnectModules) GetModules() []map[string]interface{}`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ConnectModules) GetModulesOk() (*[]map[string]interface{}, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ConnectModules) SetModules(v []map[string]interface{})`

SetModules sets Modules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


