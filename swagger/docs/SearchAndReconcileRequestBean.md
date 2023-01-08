# SearchAndReconcileRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where &#x60;expand&#x60; is specified, &#x60;expand&#x60; is defined as a comma-delimited string of values. The expand options are:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;operations&#x60; Returns all possible operations for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  &#x60;versionedRepresentations&#x60; Instead of &#x60;fields&#x60;, returns &#x60;versionedRepresentations&#x60; a JSON array containing each version of a field&#39;s value, with the highest numbered item representing the most recent version.  Examples: &#x60;\&quot;names,changelog\&quot;&#x60; Returns the display name of each field as well as a list of recent updates to an issue. | [optional] 
**Fields** | Pointer to **[]string** | A list of fields to return for each issue. Use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  &#x60;id&#x60; Returns only issue IDs.  *  Any issue field, prefixed with a dash to exclude.  The default is &#x60;id&#x60;.  Examples:   *  &#x60;summary,comment&#x60; Returns the summary and comments fields only.  *  &#x60;*all,-comment&#x60; Returns all fields except comments.  Multiple &#x60;fields&#x60; parameters can be included in a request.  Note: By default, this resource returns IDs only. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. | [optional] 
**FieldsByKeys** | Pointer to **bool** | Reference fields by their key (rather than ID). The default is &#x60;false&#x60;. | [optional] 
**Jql** | Pointer to **string** | A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. For performance reasons, this parameter requires a bounded query. A bounded query is a query with a search restriction.   *  Example of an unbounded query: &#x60;order by key desc&#x60;.  *  Example of a bounded query: &#x60;assignee &#x3D; currentUser() order by key&#x60;.  Additionally, &#x60;orderBy&#x60; clause can contain a maximum of 7 fields. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of items to return per page. To manage page size, API may return fewer items per page where a large number of fields are requested. The greatest number of items returned per page is achieved when requesting &#x60;id&#x60; or &#x60;key&#x60; only. It returns max 5000 issues. | [optional] [default to 50]
**NextPageToken** | Pointer to **string** | The token for a page to fetch that is not the first page. The first page has a &#x60;nextPageToken&#x60; of &#x60;null&#x60;. Use the &#x60;nextPageToken&#x60; to fetch the next page of issues. | [optional] 
**Properties** | Pointer to **[]string** | A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list. | [optional] 
**ReconcileIssues** | Pointer to **[]int64** | Strong consistency issue ids to be reconciled with search results. Accepts max 50 ids. This list of ids should be consistent with each paginated request across different pages. | [optional] 

## Methods

### NewSearchAndReconcileRequestBean

`func NewSearchAndReconcileRequestBean() *SearchAndReconcileRequestBean`

NewSearchAndReconcileRequestBean instantiates a new SearchAndReconcileRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAndReconcileRequestBeanWithDefaults

`func NewSearchAndReconcileRequestBeanWithDefaults() *SearchAndReconcileRequestBean`

NewSearchAndReconcileRequestBeanWithDefaults instantiates a new SearchAndReconcileRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *SearchAndReconcileRequestBean) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *SearchAndReconcileRequestBean) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *SearchAndReconcileRequestBean) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *SearchAndReconcileRequestBean) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *SearchAndReconcileRequestBean) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchAndReconcileRequestBean) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchAndReconcileRequestBean) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchAndReconcileRequestBean) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFieldsByKeys

`func (o *SearchAndReconcileRequestBean) GetFieldsByKeys() bool`

GetFieldsByKeys returns the FieldsByKeys field if non-nil, zero value otherwise.

### GetFieldsByKeysOk

`func (o *SearchAndReconcileRequestBean) GetFieldsByKeysOk() (*bool, bool)`

GetFieldsByKeysOk returns a tuple with the FieldsByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsByKeys

`func (o *SearchAndReconcileRequestBean) SetFieldsByKeys(v bool)`

SetFieldsByKeys sets FieldsByKeys field to given value.

### HasFieldsByKeys

`func (o *SearchAndReconcileRequestBean) HasFieldsByKeys() bool`

HasFieldsByKeys returns a boolean if a field has been set.

### GetJql

`func (o *SearchAndReconcileRequestBean) GetJql() string`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *SearchAndReconcileRequestBean) GetJqlOk() (*string, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *SearchAndReconcileRequestBean) SetJql(v string)`

SetJql sets Jql field to given value.

### HasJql

`func (o *SearchAndReconcileRequestBean) HasJql() bool`

HasJql returns a boolean if a field has been set.

### GetMaxResults

`func (o *SearchAndReconcileRequestBean) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchAndReconcileRequestBean) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchAndReconcileRequestBean) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchAndReconcileRequestBean) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPageToken

`func (o *SearchAndReconcileRequestBean) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SearchAndReconcileRequestBean) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SearchAndReconcileRequestBean) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SearchAndReconcileRequestBean) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetProperties

`func (o *SearchAndReconcileRequestBean) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SearchAndReconcileRequestBean) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SearchAndReconcileRequestBean) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SearchAndReconcileRequestBean) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetReconcileIssues

`func (o *SearchAndReconcileRequestBean) GetReconcileIssues() []int64`

GetReconcileIssues returns the ReconcileIssues field if non-nil, zero value otherwise.

### GetReconcileIssuesOk

`func (o *SearchAndReconcileRequestBean) GetReconcileIssuesOk() (*[]int64, bool)`

GetReconcileIssuesOk returns a tuple with the ReconcileIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconcileIssues

`func (o *SearchAndReconcileRequestBean) SetReconcileIssues(v []int64)`

SetReconcileIssues sets ReconcileIssues field to given value.

### HasReconcileIssues

`func (o *SearchAndReconcileRequestBean) HasReconcileIssues() bool`

HasReconcileIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


