# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**UserDetails**](UserDetails.md) | The ID of the user who created the comment. | [optional] [readonly] 
**Body** | Pointer to **interface{}** | The comment text in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/). | [optional] 
**Created** | Pointer to **time.Time** | The date and time at which the comment was created. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the comment. | [optional] [readonly] 
**JsdAuthorCanSeeRequest** | Pointer to **bool** | Whether the comment was added from an email sent by a person who is not part of the issue. See [Allow external emails to be added as comments on issues](https://support.atlassian.com/jira-service-management-cloud/docs/allow-external-emails-to-be-added-as-comments-on-issues/)for information on setting up this feature. | [optional] [readonly] 
**JsdPublic** | Pointer to **bool** | Whether the comment is visible in Jira Service Desk. Defaults to true when comments are created in the Jira Cloud Platform. This includes when the site doesn&#39;t use Jira Service Desk or the project isn&#39;t a Jira Service Desk project and, therefore, there is no Jira Service Desk for the issue to be visible on. To create a comment with its visibility in Jira Service Desk set to false, use the Jira Service Desk REST API [Create request comment](https://developer.atlassian.com/cloud/jira/service-desk/rest/#api-rest-servicedeskapi-request-issueIdOrKey-comment-post) operation. | [optional] [readonly] 
**Properties** | Pointer to [**[]EntityProperty**](EntityProperty.md) | A list of comment properties. Optional on create and update. | [optional] 
**RenderedBody** | Pointer to **string** | The rendered version of the comment. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the comment. | [optional] [readonly] 
**UpdateAuthor** | Pointer to [**UserDetails**](UserDetails.md) | The ID of the user who updated the comment last. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The date and time at which the comment was updated last. | [optional] [readonly] 
**Visibility** | Pointer to [**Visibility**](Visibility.md) | The group or role to which this comment is visible. Optional on create and update. | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Comment) GetAuthor() UserDetails`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Comment) GetAuthorOk() (*UserDetails, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Comment) SetAuthor(v UserDetails)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Comment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBody

`func (o *Comment) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Comment) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Comment) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *Comment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Comment) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Comment) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetCreated

`func (o *Comment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Comment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Comment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Comment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJsdAuthorCanSeeRequest

`func (o *Comment) GetJsdAuthorCanSeeRequest() bool`

GetJsdAuthorCanSeeRequest returns the JsdAuthorCanSeeRequest field if non-nil, zero value otherwise.

### GetJsdAuthorCanSeeRequestOk

`func (o *Comment) GetJsdAuthorCanSeeRequestOk() (*bool, bool)`

GetJsdAuthorCanSeeRequestOk returns a tuple with the JsdAuthorCanSeeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsdAuthorCanSeeRequest

`func (o *Comment) SetJsdAuthorCanSeeRequest(v bool)`

SetJsdAuthorCanSeeRequest sets JsdAuthorCanSeeRequest field to given value.

### HasJsdAuthorCanSeeRequest

`func (o *Comment) HasJsdAuthorCanSeeRequest() bool`

HasJsdAuthorCanSeeRequest returns a boolean if a field has been set.

### GetJsdPublic

`func (o *Comment) GetJsdPublic() bool`

GetJsdPublic returns the JsdPublic field if non-nil, zero value otherwise.

### GetJsdPublicOk

`func (o *Comment) GetJsdPublicOk() (*bool, bool)`

GetJsdPublicOk returns a tuple with the JsdPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsdPublic

`func (o *Comment) SetJsdPublic(v bool)`

SetJsdPublic sets JsdPublic field to given value.

### HasJsdPublic

`func (o *Comment) HasJsdPublic() bool`

HasJsdPublic returns a boolean if a field has been set.

### GetProperties

`func (o *Comment) GetProperties() []EntityProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Comment) GetPropertiesOk() (*[]EntityProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Comment) SetProperties(v []EntityProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Comment) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRenderedBody

`func (o *Comment) GetRenderedBody() string`

GetRenderedBody returns the RenderedBody field if non-nil, zero value otherwise.

### GetRenderedBodyOk

`func (o *Comment) GetRenderedBodyOk() (*string, bool)`

GetRenderedBodyOk returns a tuple with the RenderedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedBody

`func (o *Comment) SetRenderedBody(v string)`

SetRenderedBody sets RenderedBody field to given value.

### HasRenderedBody

`func (o *Comment) HasRenderedBody() bool`

HasRenderedBody returns a boolean if a field has been set.

### GetSelf

`func (o *Comment) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Comment) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Comment) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Comment) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetUpdateAuthor

`func (o *Comment) GetUpdateAuthor() UserDetails`

GetUpdateAuthor returns the UpdateAuthor field if non-nil, zero value otherwise.

### GetUpdateAuthorOk

`func (o *Comment) GetUpdateAuthorOk() (*UserDetails, bool)`

GetUpdateAuthorOk returns a tuple with the UpdateAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthor

`func (o *Comment) SetUpdateAuthor(v UserDetails)`

SetUpdateAuthor sets UpdateAuthor field to given value.

### HasUpdateAuthor

`func (o *Comment) HasUpdateAuthor() bool`

HasUpdateAuthor returns a boolean if a field has been set.

### GetUpdated

`func (o *Comment) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Comment) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Comment) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Comment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVisibility

`func (o *Comment) GetVisibility() Visibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Comment) GetVisibilityOk() (*Visibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Comment) SetVisibility(v Visibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Comment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


