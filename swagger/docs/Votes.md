# Votes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasVoted** | Pointer to **bool** | Whether the user making this request has voted on the issue. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these issue vote details. | [optional] [readonly] 
**Voters** | Pointer to [**[]User**](User.md) | List of the users who have voted on this issue. An empty list is returned when the calling user doesn&#39;t have the *View voters and watchers* project permission. | [optional] [readonly] 
**Votes** | Pointer to **int64** | The number of votes on the issue. | [optional] [readonly] 

## Methods

### NewVotes

`func NewVotes() *Votes`

NewVotes instantiates a new Votes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVotesWithDefaults

`func NewVotesWithDefaults() *Votes`

NewVotesWithDefaults instantiates a new Votes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasVoted

`func (o *Votes) GetHasVoted() bool`

GetHasVoted returns the HasVoted field if non-nil, zero value otherwise.

### GetHasVotedOk

`func (o *Votes) GetHasVotedOk() (*bool, bool)`

GetHasVotedOk returns a tuple with the HasVoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVoted

`func (o *Votes) SetHasVoted(v bool)`

SetHasVoted sets HasVoted field to given value.

### HasHasVoted

`func (o *Votes) HasHasVoted() bool`

HasHasVoted returns a boolean if a field has been set.

### GetSelf

`func (o *Votes) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Votes) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Votes) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Votes) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetVoters

`func (o *Votes) GetVoters() []User`

GetVoters returns the Voters field if non-nil, zero value otherwise.

### GetVotersOk

`func (o *Votes) GetVotersOk() (*[]User, bool)`

GetVotersOk returns a tuple with the Voters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoters

`func (o *Votes) SetVoters(v []User)`

SetVoters sets Voters field to given value.

### HasVoters

`func (o *Votes) HasVoters() bool`

HasVoters returns a boolean if a field has been set.

### GetVotes

`func (o *Votes) GetVotes() int64`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *Votes) GetVotesOk() (*int64, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *Votes) SetVotes(v int64)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *Votes) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


