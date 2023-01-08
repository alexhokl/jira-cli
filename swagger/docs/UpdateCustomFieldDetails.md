# UpdateCustomFieldDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the custom field. The maximum length is 40000 characters. | [optional] 
**Name** | Pointer to **string** | The name of the custom field. It doesn&#39;t have to be unique. The maximum length is 255 characters. | [optional] 
**SearcherKey** | Pointer to **string** | The searcher that defines the way the field is searched in Jira. It can be set to &#x60;null&#x60;, otherwise you must specify the valid searcher for the field type, as listed below (abbreviated values shown):   *  &#x60;cascadingselect&#x60;: &#x60;cascadingselectsearcher&#x60;  *  &#x60;datepicker&#x60;: &#x60;daterange&#x60;  *  &#x60;datetime&#x60;: &#x60;datetimerange&#x60;  *  &#x60;float&#x60;: &#x60;exactnumber&#x60; or &#x60;numberrange&#x60;  *  &#x60;grouppicker&#x60;: &#x60;grouppickersearcher&#x60;  *  &#x60;importid&#x60;: &#x60;exactnumber&#x60; or &#x60;numberrange&#x60;  *  &#x60;labels&#x60;: &#x60;labelsearcher&#x60;  *  &#x60;multicheckboxes&#x60;: &#x60;multiselectsearcher&#x60;  *  &#x60;multigrouppicker&#x60;: &#x60;multiselectsearcher&#x60;  *  &#x60;multiselect&#x60;: &#x60;multiselectsearcher&#x60;  *  &#x60;multiuserpicker&#x60;: &#x60;userpickergroupsearcher&#x60;  *  &#x60;multiversion&#x60;: &#x60;versionsearcher&#x60;  *  &#x60;project&#x60;: &#x60;projectsearcher&#x60;  *  &#x60;radiobuttons&#x60;: &#x60;multiselectsearcher&#x60;  *  &#x60;readonlyfield&#x60;: &#x60;textsearcher&#x60;  *  &#x60;select&#x60;: &#x60;multiselectsearcher&#x60;  *  &#x60;textarea&#x60;: &#x60;textsearcher&#x60;  *  &#x60;textfield&#x60;: &#x60;textsearcher&#x60;  *  &#x60;url&#x60;: &#x60;exacttextsearcher&#x60;  *  &#x60;userpicker&#x60;: &#x60;userpickergroupsearcher&#x60;  *  &#x60;version&#x60;: &#x60;versionsearcher&#x60; | [optional] 

## Methods

### NewUpdateCustomFieldDetails

`func NewUpdateCustomFieldDetails() *UpdateCustomFieldDetails`

NewUpdateCustomFieldDetails instantiates a new UpdateCustomFieldDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomFieldDetailsWithDefaults

`func NewUpdateCustomFieldDetailsWithDefaults() *UpdateCustomFieldDetails`

NewUpdateCustomFieldDetailsWithDefaults instantiates a new UpdateCustomFieldDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCustomFieldDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomFieldDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCustomFieldDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCustomFieldDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomFieldDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomFieldDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomFieldDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomFieldDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSearcherKey

`func (o *UpdateCustomFieldDetails) GetSearcherKey() string`

GetSearcherKey returns the SearcherKey field if non-nil, zero value otherwise.

### GetSearcherKeyOk

`func (o *UpdateCustomFieldDetails) GetSearcherKeyOk() (*string, bool)`

GetSearcherKeyOk returns a tuple with the SearcherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearcherKey

`func (o *UpdateCustomFieldDetails) SetSearcherKey(v string)`

SetSearcherKey sets SearcherKey field to given value.

### HasSearcherKey

`func (o *UpdateCustomFieldDetails) HasSearcherKey() bool`

HasSearcherKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


