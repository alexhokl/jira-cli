# JiraExpressionsComplexityBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beans** | [**JiraExpressionsComplexityValueBean**](JiraExpressionsComplexityValueBean.md) | The number of Jira REST API beans returned in the response. | 
**ExpensiveOperations** | [**JiraExpressionsComplexityValueBean**](JiraExpressionsComplexityValueBean.md) | The number of expensive operations executed while evaluating the expression. Expensive operations are those that load additional data, such as entity properties, comments, or custom fields. | 
**PrimitiveValues** | [**JiraExpressionsComplexityValueBean**](JiraExpressionsComplexityValueBean.md) | The number of primitive values returned in the response. | 
**Steps** | [**JiraExpressionsComplexityValueBean**](JiraExpressionsComplexityValueBean.md) | The number of steps it took to evaluate the expression, where a step is a high-level operation performed by the expression. A step is an operation such as arithmetic, accessing a property, accessing a context variable, or calling a function. | 

## Methods

### NewJiraExpressionsComplexityBean

`func NewJiraExpressionsComplexityBean(beans JiraExpressionsComplexityValueBean, expensiveOperations JiraExpressionsComplexityValueBean, primitiveValues JiraExpressionsComplexityValueBean, steps JiraExpressionsComplexityValueBean, ) *JiraExpressionsComplexityBean`

NewJiraExpressionsComplexityBean instantiates a new JiraExpressionsComplexityBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionsComplexityBeanWithDefaults

`func NewJiraExpressionsComplexityBeanWithDefaults() *JiraExpressionsComplexityBean`

NewJiraExpressionsComplexityBeanWithDefaults instantiates a new JiraExpressionsComplexityBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeans

`func (o *JiraExpressionsComplexityBean) GetBeans() JiraExpressionsComplexityValueBean`

GetBeans returns the Beans field if non-nil, zero value otherwise.

### GetBeansOk

`func (o *JiraExpressionsComplexityBean) GetBeansOk() (*JiraExpressionsComplexityValueBean, bool)`

GetBeansOk returns a tuple with the Beans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeans

`func (o *JiraExpressionsComplexityBean) SetBeans(v JiraExpressionsComplexityValueBean)`

SetBeans sets Beans field to given value.


### GetExpensiveOperations

`func (o *JiraExpressionsComplexityBean) GetExpensiveOperations() JiraExpressionsComplexityValueBean`

GetExpensiveOperations returns the ExpensiveOperations field if non-nil, zero value otherwise.

### GetExpensiveOperationsOk

`func (o *JiraExpressionsComplexityBean) GetExpensiveOperationsOk() (*JiraExpressionsComplexityValueBean, bool)`

GetExpensiveOperationsOk returns a tuple with the ExpensiveOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensiveOperations

`func (o *JiraExpressionsComplexityBean) SetExpensiveOperations(v JiraExpressionsComplexityValueBean)`

SetExpensiveOperations sets ExpensiveOperations field to given value.


### GetPrimitiveValues

`func (o *JiraExpressionsComplexityBean) GetPrimitiveValues() JiraExpressionsComplexityValueBean`

GetPrimitiveValues returns the PrimitiveValues field if non-nil, zero value otherwise.

### GetPrimitiveValuesOk

`func (o *JiraExpressionsComplexityBean) GetPrimitiveValuesOk() (*JiraExpressionsComplexityValueBean, bool)`

GetPrimitiveValuesOk returns a tuple with the PrimitiveValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimitiveValues

`func (o *JiraExpressionsComplexityBean) SetPrimitiveValues(v JiraExpressionsComplexityValueBean)`

SetPrimitiveValues sets PrimitiveValues field to given value.


### GetSteps

`func (o *JiraExpressionsComplexityBean) GetSteps() JiraExpressionsComplexityValueBean`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *JiraExpressionsComplexityBean) GetStepsOk() (*JiraExpressionsComplexityValueBean, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *JiraExpressionsComplexityBean) SetSteps(v JiraExpressionsComplexityValueBean)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


