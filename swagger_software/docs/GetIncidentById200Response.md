# GetIncidentById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | **string** | The IncidentData schema version used for this incident data.  Placeholder to support potential schema changes in the future.  | [default to "1.0"]
**Id** | **string** | The identifier for the Incident. Must be unique for a given Provider.  | 
**UpdateSequenceNumber** | **int64** | An ID used to apply an ordering to updates for this Incident in the case of out-of-order receipt of update requests.  This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the Provider system, but other alternatives are valid (e.g. a Provider could store a counter against each Incident and increment that on each update to Jira).  Updates for a Incident that are received with an updateSqeuenceId lower than what is currently stored will be ignored.  | 
**AffectedComponents** | **[]string** | The IDs of the Components impacted by this Incident. Must be unique for a given Provider.  | 
**Summary** | **string** | The human-readable summary for the Incident. Will be shown in the UI.  If not provided, will use the ID for display.  | 
**Description** | **string** | A description of the issue in Markdown format. Will be shown in the UI and used when creating Jira Issues.  | 
**Url** | **string** | A URL users can use to link to a summary view of this incident, if appropriate.  This could be any location that makes sense in the Provider system (e.g. if the summary information comes from a specific project, it might make sense to link the user to the incident in that project).  | 
**CreatedDate** | **time.Time** | The timestamp to present to the user that shows when the Incident was raised.  Expected format is an RFC3339 formatted string.  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user the last time the Incident was updated.  Expected format is an RFC3339 formatted string.  | 
**Severity** | Pointer to [**IncidentSeverity**](IncidentSeverity.md) |  | [optional] 
**Status** | **string** | The current status of the Incident.  | 
**Associations** | Pointer to [**[]Associations**](Associations.md) | The IDs of the Jira issues related to this Incident. Must be unique for a given Provider.  | [optional] 

## Methods

### NewGetIncidentById200Response

`func NewGetIncidentById200Response(schemaVersion string, id string, updateSequenceNumber int64, affectedComponents []string, summary string, description string, url string, createdDate time.Time, lastUpdated time.Time, status string, ) *GetIncidentById200Response`

NewGetIncidentById200Response instantiates a new GetIncidentById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncidentById200ResponseWithDefaults

`func NewGetIncidentById200ResponseWithDefaults() *GetIncidentById200Response`

NewGetIncidentById200ResponseWithDefaults instantiates a new GetIncidentById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetIncidentById200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetIncidentById200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetIncidentById200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetId

`func (o *GetIncidentById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIncidentById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIncidentById200Response) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateSequenceNumber

`func (o *GetIncidentById200Response) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *GetIncidentById200Response) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *GetIncidentById200Response) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetAffectedComponents

`func (o *GetIncidentById200Response) GetAffectedComponents() []string`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *GetIncidentById200Response) GetAffectedComponentsOk() (*[]string, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *GetIncidentById200Response) SetAffectedComponents(v []string)`

SetAffectedComponents sets AffectedComponents field to given value.


### GetSummary

`func (o *GetIncidentById200Response) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetIncidentById200Response) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetIncidentById200Response) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *GetIncidentById200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIncidentById200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIncidentById200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *GetIncidentById200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIncidentById200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIncidentById200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedDate

`func (o *GetIncidentById200Response) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *GetIncidentById200Response) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *GetIncidentById200Response) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetLastUpdated

`func (o *GetIncidentById200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetIncidentById200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetIncidentById200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetSeverity

`func (o *GetIncidentById200Response) GetSeverity() IncidentSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIncidentById200Response) GetSeverityOk() (*IncidentSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIncidentById200Response) SetSeverity(v IncidentSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIncidentById200Response) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *GetIncidentById200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIncidentById200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIncidentById200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAssociations

`func (o *GetIncidentById200Response) GetAssociations() []Associations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *GetIncidentById200Response) GetAssociationsOk() (*[]Associations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *GetIncidentById200Response) SetAssociations(v []Associations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *GetIncidentById200Response) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


