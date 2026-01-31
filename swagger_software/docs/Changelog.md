# Changelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**ChangelogAuthor**](ChangelogAuthor.md) |  | [optional] 
**Created** | Pointer to **time.Time** | The date on which the change took place. | [optional] [readonly] 
**HistoryMetadata** | Pointer to [**ChangelogHistoryMetadata**](ChangelogHistoryMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the changelog. | [optional] [readonly] 
**Items** | Pointer to [**[]ChangelogItemsInner**](ChangelogItemsInner.md) | The list of items changed. | [optional] [readonly] 

## Methods

### NewChangelog

`func NewChangelog() *Changelog`

NewChangelog instantiates a new Changelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogWithDefaults

`func NewChangelogWithDefaults() *Changelog`

NewChangelogWithDefaults instantiates a new Changelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Changelog) GetAuthor() ChangelogAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Changelog) GetAuthorOk() (*ChangelogAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Changelog) SetAuthor(v ChangelogAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Changelog) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *Changelog) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Changelog) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Changelog) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Changelog) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHistoryMetadata

`func (o *Changelog) GetHistoryMetadata() ChangelogHistoryMetadata`

GetHistoryMetadata returns the HistoryMetadata field if non-nil, zero value otherwise.

### GetHistoryMetadataOk

`func (o *Changelog) GetHistoryMetadataOk() (*ChangelogHistoryMetadata, bool)`

GetHistoryMetadataOk returns a tuple with the HistoryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryMetadata

`func (o *Changelog) SetHistoryMetadata(v ChangelogHistoryMetadata)`

SetHistoryMetadata sets HistoryMetadata field to given value.

### HasHistoryMetadata

`func (o *Changelog) HasHistoryMetadata() bool`

HasHistoryMetadata returns a boolean if a field has been set.

### GetId

`func (o *Changelog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Changelog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Changelog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Changelog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *Changelog) GetItems() []ChangelogItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Changelog) GetItemsOk() (*[]ChangelogItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Changelog) SetItems(v []ChangelogItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *Changelog) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


