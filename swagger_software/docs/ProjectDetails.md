# ProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrls** | Pointer to [**IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfAvatarUrls**](IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfAvatarUrls.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the project. | [optional] 
**Key** | Pointer to **string** | The key of the project. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the project. | [optional] [readonly] 
**ProjectCategory** | Pointer to [**IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfProjectCategory**](IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfProjectCategory.md) |  | [optional] 
**ProjectTypeKey** | Pointer to **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes) of the project. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the project details. | [optional] [readonly] 
**Simplified** | Pointer to **bool** | Whether or not the project is simplified. | [optional] [readonly] 

## Methods

### NewProjectDetails

`func NewProjectDetails() *ProjectDetails`

NewProjectDetails instantiates a new ProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsWithDefaults

`func NewProjectDetailsWithDefaults() *ProjectDetails`

NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrls

`func (o *ProjectDetails) GetAvatarUrls() IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfAvatarUrls`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *ProjectDetails) GetAvatarUrlsOk() (*IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfAvatarUrls, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *ProjectDetails) SetAvatarUrls(v IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfAvatarUrls)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *ProjectDetails) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetId

`func (o *ProjectDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ProjectDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ProjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectCategory

`func (o *ProjectDetails) GetProjectCategory() IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfProjectCategory`

GetProjectCategory returns the ProjectCategory field if non-nil, zero value otherwise.

### GetProjectCategoryOk

`func (o *ProjectDetails) GetProjectCategoryOk() (*IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfProjectCategory, bool)`

GetProjectCategoryOk returns a tuple with the ProjectCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCategory

`func (o *ProjectDetails) SetProjectCategory(v IssueBeanTransitionsInnerToAllOfScopeAllOfProjectAllOfProjectCategory)`

SetProjectCategory sets ProjectCategory field to given value.

### HasProjectCategory

`func (o *ProjectDetails) HasProjectCategory() bool`

HasProjectCategory returns a boolean if a field has been set.

### GetProjectTypeKey

`func (o *ProjectDetails) GetProjectTypeKey() string`

GetProjectTypeKey returns the ProjectTypeKey field if non-nil, zero value otherwise.

### GetProjectTypeKeyOk

`func (o *ProjectDetails) GetProjectTypeKeyOk() (*string, bool)`

GetProjectTypeKeyOk returns a tuple with the ProjectTypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeKey

`func (o *ProjectDetails) SetProjectTypeKey(v string)`

SetProjectTypeKey sets ProjectTypeKey field to given value.

### HasProjectTypeKey

`func (o *ProjectDetails) HasProjectTypeKey() bool`

HasProjectTypeKey returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ProjectDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSimplified

`func (o *ProjectDetails) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *ProjectDetails) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *ProjectDetails) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *ProjectDetails) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


