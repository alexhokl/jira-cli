# BulkFetchIssueRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **[]string** | Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where &#x60;expand&#x60; is specified, &#x60;expand&#x60; is defined as a list of values. The expand options are:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;operations&#x60; Returns all possible operations for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent. This returns a maximum of 40 changelogs. If you require more, please refer to [Bulk fetch changelogs](#api-rest-api-3-changelog-bulkfetch-post).  *  &#x60;versionedRepresentations&#x60; Instead of &#x60;fields&#x60;, returns &#x60;versionedRepresentations&#x60; a JSON array containing each version of a field&#39;s value, with the highest numbered item representing the most recent version. | [optional] 
**Fields** | Pointer to **[]string** | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  The default is &#x60;*navigable&#x60;.  Examples:   *  &#x60;summary,comment&#x60; Returns the summary and comments fields only.  *  &#x60;-description&#x60; Returns all navigable (default) fields except description.  *  &#x60;*all,-comment&#x60; Returns all fields except comments.  Multiple &#x60;fields&#x60; parameters can be included in a request.  Note: All navigable fields are returned by default. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. | [optional] 
**FieldsByKeys** | Pointer to **bool** | Reference fields by their key (rather than ID). The default is &#x60;false&#x60;. | [optional] 
**IssueIdsOrKeys** | **[]string** | An array of issue IDs or issue keys to fetch. You can mix issue IDs and keys in the same query. | 
**Properties** | Pointer to **[]string** | A list of issue property keys of issue properties to be included in the results. A maximum of 5 issue property keys can be specified. | [optional] 

## Methods

### NewBulkFetchIssueRequestBean

`func NewBulkFetchIssueRequestBean(issueIdsOrKeys []string, ) *BulkFetchIssueRequestBean`

NewBulkFetchIssueRequestBean instantiates a new BulkFetchIssueRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkFetchIssueRequestBeanWithDefaults

`func NewBulkFetchIssueRequestBeanWithDefaults() *BulkFetchIssueRequestBean`

NewBulkFetchIssueRequestBeanWithDefaults instantiates a new BulkFetchIssueRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *BulkFetchIssueRequestBean) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *BulkFetchIssueRequestBean) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *BulkFetchIssueRequestBean) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *BulkFetchIssueRequestBean) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *BulkFetchIssueRequestBean) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BulkFetchIssueRequestBean) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BulkFetchIssueRequestBean) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BulkFetchIssueRequestBean) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFieldsByKeys

`func (o *BulkFetchIssueRequestBean) GetFieldsByKeys() bool`

GetFieldsByKeys returns the FieldsByKeys field if non-nil, zero value otherwise.

### GetFieldsByKeysOk

`func (o *BulkFetchIssueRequestBean) GetFieldsByKeysOk() (*bool, bool)`

GetFieldsByKeysOk returns a tuple with the FieldsByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsByKeys

`func (o *BulkFetchIssueRequestBean) SetFieldsByKeys(v bool)`

SetFieldsByKeys sets FieldsByKeys field to given value.

### HasFieldsByKeys

`func (o *BulkFetchIssueRequestBean) HasFieldsByKeys() bool`

HasFieldsByKeys returns a boolean if a field has been set.

### GetIssueIdsOrKeys

`func (o *BulkFetchIssueRequestBean) GetIssueIdsOrKeys() []string`

GetIssueIdsOrKeys returns the IssueIdsOrKeys field if non-nil, zero value otherwise.

### GetIssueIdsOrKeysOk

`func (o *BulkFetchIssueRequestBean) GetIssueIdsOrKeysOk() (*[]string, bool)`

GetIssueIdsOrKeysOk returns a tuple with the IssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdsOrKeys

`func (o *BulkFetchIssueRequestBean) SetIssueIdsOrKeys(v []string)`

SetIssueIdsOrKeys sets IssueIdsOrKeys field to given value.


### GetProperties

`func (o *BulkFetchIssueRequestBean) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BulkFetchIssueRequestBean) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BulkFetchIssueRequestBean) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BulkFetchIssueRequestBean) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


