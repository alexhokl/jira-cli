# JiraExpressionEvaluateContextBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Board** | Pointer to **int64** | The ID of the board that is available under the &#x60;board&#x60; variable when evaluating the expression. | [optional] 
**Custom** | Pointer to [**[]CustomContextVariable**](CustomContextVariable.md) | Custom context variables and their types. These variable types are available for use in a custom context:   *  &#x60;user&#x60;: A [user](https://developer.atlassian.com/cloud/jira/platform/jira-expressions-type-reference#user) specified as an Atlassian account ID.  *  &#x60;issue&#x60;: An [issue](https://developer.atlassian.com/cloud/jira/platform/jira-expressions-type-reference#issue) specified by ID or key. All the fields of the issue object are available in the Jira expression.  *  &#x60;json&#x60;: A JSON object containing custom content.  *  &#x60;list&#x60;: A JSON list of &#x60;user&#x60;, &#x60;issue&#x60;, or &#x60;json&#x60; variable types. | [optional] 
**CustomerRequest** | Pointer to **int64** | The ID of the customer request that is available under the &#x60;customerRequest&#x60; variable when evaluating the expression. This is the same as the ID of the underlying Jira issue, but the customer request context variable will have a different type. | [optional] 
**Issue** | Pointer to [**IdOrKeyBean**](IdOrKeyBean.md) | The issue that is available under the &#x60;issue&#x60; variable when evaluating the expression. | [optional] 
**Issues** | Pointer to [**JexpEvaluateCtxIssues**](JexpEvaluateCtxIssues.md) | The collection of issues that is available under the &#x60;issues&#x60; variable when evaluating the expression. | [optional] 
**Project** | Pointer to [**IdOrKeyBean**](IdOrKeyBean.md) | The project that is available under the &#x60;project&#x60; variable when evaluating the expression. | [optional] 
**ServiceDesk** | Pointer to **int64** | The ID of the service desk that is available under the &#x60;serviceDesk&#x60; variable when evaluating the expression. | [optional] 
**Sprint** | Pointer to **int64** | The ID of the sprint that is available under the &#x60;sprint&#x60; variable when evaluating the expression. | [optional] 

## Methods

### NewJiraExpressionEvaluateContextBean

`func NewJiraExpressionEvaluateContextBean() *JiraExpressionEvaluateContextBean`

NewJiraExpressionEvaluateContextBean instantiates a new JiraExpressionEvaluateContextBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionEvaluateContextBeanWithDefaults

`func NewJiraExpressionEvaluateContextBeanWithDefaults() *JiraExpressionEvaluateContextBean`

NewJiraExpressionEvaluateContextBeanWithDefaults instantiates a new JiraExpressionEvaluateContextBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoard

`func (o *JiraExpressionEvaluateContextBean) GetBoard() int64`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *JiraExpressionEvaluateContextBean) GetBoardOk() (*int64, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *JiraExpressionEvaluateContextBean) SetBoard(v int64)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *JiraExpressionEvaluateContextBean) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCustom

`func (o *JiraExpressionEvaluateContextBean) GetCustom() []CustomContextVariable`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *JiraExpressionEvaluateContextBean) GetCustomOk() (*[]CustomContextVariable, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *JiraExpressionEvaluateContextBean) SetCustom(v []CustomContextVariable)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *JiraExpressionEvaluateContextBean) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomerRequest

`func (o *JiraExpressionEvaluateContextBean) GetCustomerRequest() int64`

GetCustomerRequest returns the CustomerRequest field if non-nil, zero value otherwise.

### GetCustomerRequestOk

`func (o *JiraExpressionEvaluateContextBean) GetCustomerRequestOk() (*int64, bool)`

GetCustomerRequestOk returns a tuple with the CustomerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRequest

`func (o *JiraExpressionEvaluateContextBean) SetCustomerRequest(v int64)`

SetCustomerRequest sets CustomerRequest field to given value.

### HasCustomerRequest

`func (o *JiraExpressionEvaluateContextBean) HasCustomerRequest() bool`

HasCustomerRequest returns a boolean if a field has been set.

### GetIssue

`func (o *JiraExpressionEvaluateContextBean) GetIssue() IdOrKeyBean`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *JiraExpressionEvaluateContextBean) GetIssueOk() (*IdOrKeyBean, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *JiraExpressionEvaluateContextBean) SetIssue(v IdOrKeyBean)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *JiraExpressionEvaluateContextBean) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetIssues

`func (o *JiraExpressionEvaluateContextBean) GetIssues() JexpEvaluateCtxIssues`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *JiraExpressionEvaluateContextBean) GetIssuesOk() (*JexpEvaluateCtxIssues, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *JiraExpressionEvaluateContextBean) SetIssues(v JexpEvaluateCtxIssues)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *JiraExpressionEvaluateContextBean) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetProject

`func (o *JiraExpressionEvaluateContextBean) GetProject() IdOrKeyBean`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *JiraExpressionEvaluateContextBean) GetProjectOk() (*IdOrKeyBean, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *JiraExpressionEvaluateContextBean) SetProject(v IdOrKeyBean)`

SetProject sets Project field to given value.

### HasProject

`func (o *JiraExpressionEvaluateContextBean) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetServiceDesk

`func (o *JiraExpressionEvaluateContextBean) GetServiceDesk() int64`

GetServiceDesk returns the ServiceDesk field if non-nil, zero value otherwise.

### GetServiceDeskOk

`func (o *JiraExpressionEvaluateContextBean) GetServiceDeskOk() (*int64, bool)`

GetServiceDeskOk returns a tuple with the ServiceDesk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDesk

`func (o *JiraExpressionEvaluateContextBean) SetServiceDesk(v int64)`

SetServiceDesk sets ServiceDesk field to given value.

### HasServiceDesk

`func (o *JiraExpressionEvaluateContextBean) HasServiceDesk() bool`

HasServiceDesk returns a boolean if a field has been set.

### GetSprint

`func (o *JiraExpressionEvaluateContextBean) GetSprint() int64`

GetSprint returns the Sprint field if non-nil, zero value otherwise.

### GetSprintOk

`func (o *JiraExpressionEvaluateContextBean) GetSprintOk() (*int64, bool)`

GetSprintOk returns a tuple with the Sprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprint

`func (o *JiraExpressionEvaluateContextBean) SetSprint(v int64)`

SetSprint sets Sprint field to given value.

### HasSprint

`func (o *JiraExpressionEvaluateContextBean) HasSprint() bool`

HasSprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


