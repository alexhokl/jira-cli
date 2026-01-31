# IssueIdOrKeysAssociation1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationType** | **string** | Defines the association type.  | 
**Values** | **[]string** | The Jira issue id or keys to associate the Security information with.  The number of values counted across all associationTypes (issueIdOrKeys) must not exceed a limit of 500.  | 

## Methods

### NewIssueIdOrKeysAssociation1

`func NewIssueIdOrKeysAssociation1(associationType string, values []string, ) *IssueIdOrKeysAssociation1`

NewIssueIdOrKeysAssociation1 instantiates a new IssueIdOrKeysAssociation1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueIdOrKeysAssociation1WithDefaults

`func NewIssueIdOrKeysAssociation1WithDefaults() *IssueIdOrKeysAssociation1`

NewIssueIdOrKeysAssociation1WithDefaults instantiates a new IssueIdOrKeysAssociation1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationType

`func (o *IssueIdOrKeysAssociation1) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *IssueIdOrKeysAssociation1) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *IssueIdOrKeysAssociation1) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.


### GetValues

`func (o *IssueIdOrKeysAssociation1) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IssueIdOrKeysAssociation1) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IssueIdOrKeysAssociation1) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


