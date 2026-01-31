# Review

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | **string** | The PostIncidentReviewData schema version used for this post-incident review data.  Placeholder to support potential schema changes in the future.  | [default to "1.0"]
**Id** | **string** | The identifier for the Review. Must be unique for a given Provider.  | 
**UpdateSequenceNumber** | **int64** | An ID used to apply an ordering to updates for this Review in the case of out-of-order receipt of update requests.  This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the Provider system, but other alternatives are valid (e.g. a Provider could store a counter against each Review and increment that on each update to Jira).  Updates for a Review that are received with an updateSqeuenceId lower than what is currently stored will be ignored.  | 
**Reviews** | **[]string** | The IDs of the Incidents covered by this Review. Must be unique for a given Provider.  | 
**Summary** | **string** | The human-readable summary for the Post-Incident Review. Will be shown in the UI.  If not provided, will use the ID for display.  | 
**Description** | **string** | A description of the review in Markdown format. Will be shown in the UI and used when creating Jira Issues.  | 
**Url** | **string** | A URL users can use to link to a summary view of this review, if appropriate.  This could be any location that makes sense in the Provider system (e.g. if the summary information comes from a specific project, it might make sense to link the user to the review in that project).  | 
**CreatedDate** | **time.Time** | The timestamp to present to the user that shows when the Review was raised.  Expected format is an RFC3339 formatted string.  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user the last time the Review was updated.  Expected format is an RFC3339 formatted string.  | 
**Status** | **string** | The current status of the Post-Incident Review.  | 
**Associations** | Pointer to [**[]Associations**](Associations.md) | The IDs of the Jira issues related to this Incident. Must be unique for a given Provider.  | [optional] 

## Methods

### NewReview

`func NewReview(schemaVersion string, id string, updateSequenceNumber int64, reviews []string, summary string, description string, url string, createdDate time.Time, lastUpdated time.Time, status string, ) *Review`

NewReview instantiates a new Review object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWithDefaults

`func NewReviewWithDefaults() *Review`

NewReviewWithDefaults instantiates a new Review object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *Review) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *Review) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *Review) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetId

`func (o *Review) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Review) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Review) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateSequenceNumber

`func (o *Review) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *Review) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *Review) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetReviews

`func (o *Review) GetReviews() []string`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *Review) GetReviewsOk() (*[]string, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *Review) SetReviews(v []string)`

SetReviews sets Reviews field to given value.


### GetSummary

`func (o *Review) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Review) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Review) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *Review) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Review) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Review) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *Review) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Review) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Review) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedDate

`func (o *Review) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Review) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Review) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetLastUpdated

`func (o *Review) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Review) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Review) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetStatus

`func (o *Review) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Review) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Review) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAssociations

`func (o *Review) GetAssociations() []Associations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Review) GetAssociationsOk() (*[]Associations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Review) SetAssociations(v []Associations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Review) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


