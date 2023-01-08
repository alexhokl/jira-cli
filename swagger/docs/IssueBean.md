# IssueBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changelog** | Pointer to [**PageOfChangelogs**](PageOfChangelogs.md) | Details of changelogs associated with the issue. | [optional] [readonly] 
**Editmeta** | Pointer to [**IssueUpdateMetadata**](IssueUpdateMetadata.md) | The metadata for the fields on the issue that can be amended. | [optional] [readonly] 
**Expand** | Pointer to **string** | Expand options that include additional issue details in the response. | [optional] [readonly] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**FieldsToInclude** | Pointer to [**IncludedFields**](IncludedFields.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the issue. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the issue. | [optional] [readonly] 
**Names** | Pointer to **map[string]string** | The ID and name of each field present on the issue. | [optional] [readonly] 
**Operations** | Pointer to [**Operations**](Operations.md) | The operations that can be performed on the issue. | [optional] [readonly] 
**Properties** | Pointer to **map[string]interface{}** | Details of the issue properties identified in the request. | [optional] [readonly] 
**RenderedFields** | Pointer to **map[string]interface{}** | The rendered value of each field present on the issue. | [optional] [readonly] 
**Schema** | Pointer to [**map[string]JsonTypeBean**](JsonTypeBean.md) | The schema describing each field present on the issue. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the issue details. | [optional] [readonly] 
**Transitions** | Pointer to [**[]IssueTransition**](IssueTransition.md) | The transitions that can be performed on the issue. | [optional] [readonly] 
**VersionedRepresentations** | Pointer to **map[string]map[string]interface{}** | The versions of each field on the issue. | [optional] [readonly] 

## Methods

### NewIssueBean

`func NewIssueBean() *IssueBean`

NewIssueBean instantiates a new IssueBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBeanWithDefaults

`func NewIssueBeanWithDefaults() *IssueBean`

NewIssueBeanWithDefaults instantiates a new IssueBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangelog

`func (o *IssueBean) GetChangelog() PageOfChangelogs`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *IssueBean) GetChangelogOk() (*PageOfChangelogs, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *IssueBean) SetChangelog(v PageOfChangelogs)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *IssueBean) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetEditmeta

`func (o *IssueBean) GetEditmeta() IssueUpdateMetadata`

GetEditmeta returns the Editmeta field if non-nil, zero value otherwise.

### GetEditmetaOk

`func (o *IssueBean) GetEditmetaOk() (*IssueUpdateMetadata, bool)`

GetEditmetaOk returns a tuple with the Editmeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditmeta

`func (o *IssueBean) SetEditmeta(v IssueUpdateMetadata)`

SetEditmeta sets Editmeta field to given value.

### HasEditmeta

`func (o *IssueBean) HasEditmeta() bool`

HasEditmeta returns a boolean if a field has been set.

### GetExpand

`func (o *IssueBean) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *IssueBean) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *IssueBean) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *IssueBean) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *IssueBean) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IssueBean) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IssueBean) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *IssueBean) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFieldsToInclude

`func (o *IssueBean) GetFieldsToInclude() IncludedFields`

GetFieldsToInclude returns the FieldsToInclude field if non-nil, zero value otherwise.

### GetFieldsToIncludeOk

`func (o *IssueBean) GetFieldsToIncludeOk() (*IncludedFields, bool)`

GetFieldsToIncludeOk returns a tuple with the FieldsToInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsToInclude

`func (o *IssueBean) SetFieldsToInclude(v IncludedFields)`

SetFieldsToInclude sets FieldsToInclude field to given value.

### HasFieldsToInclude

`func (o *IssueBean) HasFieldsToInclude() bool`

HasFieldsToInclude returns a boolean if a field has been set.

### GetId

`func (o *IssueBean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueBean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueBean) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *IssueBean) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueBean) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueBean) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueBean) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNames

`func (o *IssueBean) GetNames() map[string]string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *IssueBean) GetNamesOk() (*map[string]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *IssueBean) SetNames(v map[string]string)`

SetNames sets Names field to given value.

### HasNames

`func (o *IssueBean) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetOperations

`func (o *IssueBean) GetOperations() Operations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *IssueBean) GetOperationsOk() (*Operations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *IssueBean) SetOperations(v Operations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *IssueBean) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetProperties

`func (o *IssueBean) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueBean) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueBean) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueBean) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRenderedFields

`func (o *IssueBean) GetRenderedFields() map[string]interface{}`

GetRenderedFields returns the RenderedFields field if non-nil, zero value otherwise.

### GetRenderedFieldsOk

`func (o *IssueBean) GetRenderedFieldsOk() (*map[string]interface{}, bool)`

GetRenderedFieldsOk returns a tuple with the RenderedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedFields

`func (o *IssueBean) SetRenderedFields(v map[string]interface{})`

SetRenderedFields sets RenderedFields field to given value.

### HasRenderedFields

`func (o *IssueBean) HasRenderedFields() bool`

HasRenderedFields returns a boolean if a field has been set.

### GetSchema

`func (o *IssueBean) GetSchema() map[string]JsonTypeBean`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IssueBean) GetSchemaOk() (*map[string]JsonTypeBean, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IssueBean) SetSchema(v map[string]JsonTypeBean)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IssueBean) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSelf

`func (o *IssueBean) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueBean) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueBean) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IssueBean) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTransitions

`func (o *IssueBean) GetTransitions() []IssueTransition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *IssueBean) GetTransitionsOk() (*[]IssueTransition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *IssueBean) SetTransitions(v []IssueTransition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *IssueBean) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetVersionedRepresentations

`func (o *IssueBean) GetVersionedRepresentations() map[string]map[string]interface{}`

GetVersionedRepresentations returns the VersionedRepresentations field if non-nil, zero value otherwise.

### GetVersionedRepresentationsOk

`func (o *IssueBean) GetVersionedRepresentationsOk() (*map[string]map[string]interface{}, bool)`

GetVersionedRepresentationsOk returns a tuple with the VersionedRepresentations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedRepresentations

`func (o *IssueBean) SetVersionedRepresentations(v map[string]map[string]interface{})`

SetVersionedRepresentations sets VersionedRepresentations field to given value.

### HasVersionedRepresentations

`func (o *IssueBean) HasVersionedRepresentations() bool`

HasVersionedRepresentations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


