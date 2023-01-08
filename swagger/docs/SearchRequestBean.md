# SearchRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **[]string** | Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where &#x60;expand&#x60; is specified, &#x60;expand&#x60; is defined as a list of values. The expand options are:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;operations&#x60; Returns all possible operations for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  &#x60;versionedRepresentations&#x60; Instead of &#x60;fields&#x60;, returns &#x60;versionedRepresentations&#x60; a JSON array containing each version of a field&#39;s value, with the highest numbered item representing the most recent version. | [optional] 
**Fields** | Pointer to **[]string** | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  The default is &#x60;*navigable&#x60;.  Examples:   *  &#x60;summary,comment&#x60; Returns the summary and comments fields only.  *  &#x60;-description&#x60; Returns all navigable (default) fields except description.  *  &#x60;*all,-comment&#x60; Returns all fields except comments.  Multiple &#x60;fields&#x60; parameters can be included in a request.  Note: All navigable fields are returned by default. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. | [optional] 
**FieldsByKeys** | Pointer to **bool** | Reference fields by their key (rather than ID). The default is &#x60;false&#x60;. | [optional] 
**Jql** | Pointer to **string** | A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of items to return per page. | [optional] [default to 50]
**Properties** | Pointer to **[]string** | A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list. | [optional] 
**StartAt** | Pointer to **int32** | The index of the first item to return in the page of results (page offset). The base index is &#x60;0&#x60;. | [optional] 
**ValidateQuery** | Pointer to **string** | Determines how to validate the JQL query and treat the validation results. Supported values:   *  &#x60;strict&#x60; Returns a 400 response code if any errors are found, along with a list of all errors (and warnings).  *  &#x60;warn&#x60; Returns all errors as warnings.  *  &#x60;none&#x60; No validation is performed.  *  &#x60;true&#x60; *Deprecated* A legacy synonym for &#x60;strict&#x60;.  *  &#x60;false&#x60; *Deprecated* A legacy synonym for &#x60;warn&#x60;.  The default is &#x60;strict&#x60;.  Note: If the JQL is not correctly formed a 400 response code is returned, regardless of the &#x60;validateQuery&#x60; value. | [optional] 

## Methods

### NewSearchRequestBean

`func NewSearchRequestBean() *SearchRequestBean`

NewSearchRequestBean instantiates a new SearchRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestBeanWithDefaults

`func NewSearchRequestBeanWithDefaults() *SearchRequestBean`

NewSearchRequestBeanWithDefaults instantiates a new SearchRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *SearchRequestBean) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *SearchRequestBean) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *SearchRequestBean) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *SearchRequestBean) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *SearchRequestBean) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchRequestBean) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchRequestBean) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchRequestBean) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFieldsByKeys

`func (o *SearchRequestBean) GetFieldsByKeys() bool`

GetFieldsByKeys returns the FieldsByKeys field if non-nil, zero value otherwise.

### GetFieldsByKeysOk

`func (o *SearchRequestBean) GetFieldsByKeysOk() (*bool, bool)`

GetFieldsByKeysOk returns a tuple with the FieldsByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsByKeys

`func (o *SearchRequestBean) SetFieldsByKeys(v bool)`

SetFieldsByKeys sets FieldsByKeys field to given value.

### HasFieldsByKeys

`func (o *SearchRequestBean) HasFieldsByKeys() bool`

HasFieldsByKeys returns a boolean if a field has been set.

### GetJql

`func (o *SearchRequestBean) GetJql() string`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *SearchRequestBean) GetJqlOk() (*string, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *SearchRequestBean) SetJql(v string)`

SetJql sets Jql field to given value.

### HasJql

`func (o *SearchRequestBean) HasJql() bool`

HasJql returns a boolean if a field has been set.

### GetMaxResults

`func (o *SearchRequestBean) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchRequestBean) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchRequestBean) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchRequestBean) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetProperties

`func (o *SearchRequestBean) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SearchRequestBean) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SearchRequestBean) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SearchRequestBean) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStartAt

`func (o *SearchRequestBean) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *SearchRequestBean) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *SearchRequestBean) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *SearchRequestBean) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetValidateQuery

`func (o *SearchRequestBean) GetValidateQuery() string`

GetValidateQuery returns the ValidateQuery field if non-nil, zero value otherwise.

### GetValidateQueryOk

`func (o *SearchRequestBean) GetValidateQueryOk() (*string, bool)`

GetValidateQueryOk returns a tuple with the ValidateQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateQuery

`func (o *SearchRequestBean) SetValidateQuery(v string)`

SetValidateQuery sets ValidateQuery field to given value.

### HasValidateQuery

`func (o *SearchRequestBean) HasValidateQuery() bool`

HasValidateQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


