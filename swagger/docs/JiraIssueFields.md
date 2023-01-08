# JiraIssueFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CascadingSelectFields** | Pointer to [**[]JiraCascadingSelectField**](JiraCascadingSelectField.md) | Add or clear a cascading select field:   *  To add, specify &#x60;optionId&#x60; for both parent and child.  *  To clear the child, set its &#x60;optionId&#x60; to null.  *  To clear both, set the parent&#39;s &#x60;optionId&#x60; to null. | [optional] 
**ClearableNumberFields** | Pointer to [**[]JiraNumberField**](JiraNumberField.md) | Add or clear a number field:   *  To add, specify a numeric &#x60;value&#x60;.  *  To clear, set &#x60;value&#x60; to &#x60;null&#x60;. | [optional] 
**ColorFields** | Pointer to [**[]JiraColorField**](JiraColorField.md) | Add or clear a color field:   *  To add, specify the color &#x60;name&#x60;. Available colors are: &#x60;purple&#x60;, &#x60;blue&#x60;, &#x60;green&#x60;, &#x60;teal&#x60;, &#x60;yellow&#x60;, &#x60;orange&#x60;, &#x60;grey&#x60;, &#x60;dark purple&#x60;, &#x60;dark blue&#x60;, &#x60;dark green&#x60;, &#x60;dark teal&#x60;, &#x60;dark yellow&#x60;, &#x60;dark orange&#x60;, &#x60;dark grey&#x60;.  *  To clear, set the color &#x60;name&#x60; to an empty string. | [optional] 
**DatePickerFields** | Pointer to [**[]JiraDateField**](JiraDateField.md) | Add or clear a date picker field:   *  To add, specify the date in &#x60;d/mmm/yy&#x60; format or ISO format &#x60;dd-mm-yyyy&#x60;.  *  To clear, set &#x60;formattedDate&#x60; to an empty string. | [optional] 
**DateTimePickerFields** | Pointer to [**[]JiraDateTimeField**](JiraDateTimeField.md) | Add or clear the planned start date and time:   *  To add, specify the date and time in ISO format for &#x60;formattedDateTime&#x60;.  *  To clear, provide an empty string for &#x60;formattedDateTime&#x60;. | [optional] 
**IssueType** | Pointer to [**JiraIssueTypeField**](JiraIssueTypeField.md) | Set the issue type field by providing an &#x60;issueTypeId&#x60;. | [optional] 
**LabelsFields** | Pointer to [**[]JiraLabelsField**](JiraLabelsField.md) | Edit a labels field:   *  Options include &#x60;ADD&#x60;, &#x60;REPLACE&#x60;, &#x60;REMOVE&#x60;, or &#x60;REMOVE_ALL&#x60; for bulk edits.  *  To clear labels, use the &#x60;REMOVE_ALL&#x60; option with an empty &#x60;labels&#x60; array. | [optional] 
**MultipleGroupPickerFields** | Pointer to [**[]JiraMultipleGroupPickerField**](JiraMultipleGroupPickerField.md) | Add or clear a multi-group picker field:   *  To add groups, provide an array of groups with &#x60;groupName&#x60;s.  *  To clear all groups, use an empty &#x60;groups&#x60; array. | [optional] 
**MultipleSelectClearableUserPickerFields** | Pointer to [**[]JiraMultipleSelectUserPickerField**](JiraMultipleSelectUserPickerField.md) | Assign or unassign multiple users to/from a field:   *  To assign, provide an array of user &#x60;accountId&#x60;s.  *  To clear, set &#x60;users&#x60; to &#x60;null&#x60;. | [optional] 
**MultipleSelectFields** | Pointer to [**[]JiraMultipleSelectField**](JiraMultipleSelectField.md) | Add or clear a multi-select field:   *  To add, provide an array of options with &#x60;optionId&#x60;s.  *  To clear, use an empty &#x60;options&#x60; array. | [optional] 
**MultipleVersionPickerFields** | Pointer to [**[]JiraMultipleVersionPickerField**](JiraMultipleVersionPickerField.md) | Edit a multi-version picker field like Fix Versions/Affects Versions:   *  Options include &#x60;ADD&#x60;, &#x60;REPLACE&#x60;, &#x60;REMOVE&#x60;, or &#x60;REMOVE_ALL&#x60; for bulk edits.  *  To clear the field, use the &#x60;REMOVE_ALL&#x60; option with an empty &#x60;versions&#x60; array. | [optional] 
**MultiselectComponents** | Pointer to [**JiraMultiSelectComponentField**](JiraMultiSelectComponentField.md) | Edit a multi select components field:   *  Options include &#x60;ADD&#x60;, &#x60;REPLACE&#x60;, &#x60;REMOVE&#x60;, or &#x60;REMOVE_ALL&#x60; for bulk edits.  *  To clear, use the &#x60;REMOVE_ALL&#x60; option with an empty &#x60;components&#x60; array. | [optional] 
**OriginalEstimateField** | Pointer to [**JiraDurationField**](JiraDurationField.md) | Edit the original estimate field. | [optional] 
**Priority** | Pointer to [**JiraPriorityField**](JiraPriorityField.md) | Set the priority of an issue by specifying a &#x60;priorityId&#x60;. | [optional] 
**RichTextFields** | Pointer to [**[]JiraRichTextField**](JiraRichTextField.md) | Add or clear a rich text field:   *  To add, provide &#x60;adfValue&#x60;. Note that rich text fields only support ADF values.  *  To clear, use an empty &#x60;richText&#x60; object.  For ADF format details, refer to: [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure). | [optional] 
**SingleGroupPickerFields** | Pointer to [**[]JiraSingleGroupPickerField**](JiraSingleGroupPickerField.md) | Add or clear a single group picker field:   *  To add, specify the group with &#x60;groupName&#x60;.  *  To clear, set &#x60;groupName&#x60; to an empty string. | [optional] 
**SingleLineTextFields** | Pointer to [**[]JiraSingleLineTextField**](JiraSingleLineTextField.md) | Add or clear a single line text field:   *  To add, provide the &#x60;text&#x60; value.  *  To clear, set &#x60;text&#x60; to an empty string. | [optional] 
**SingleSelectClearableUserPickerFields** | Pointer to [**[]JiraSingleSelectUserPickerField**](JiraSingleSelectUserPickerField.md) | Edit assignment for single select user picker fields like Assignee/Reporter:   *  To assign an issue, specify the user&#39;s &#x60;accountId&#x60;.  *  To unassign an issue, set &#x60;user&#x60; to &#x60;null&#x60;.  *  For automatic assignment, set &#x60;accountId&#x60; to &#x60;-1&#x60;. | [optional] 
**SingleSelectFields** | Pointer to [**[]JiraSingleSelectField**](JiraSingleSelectField.md) | Add or clear a single select field:   *  To add, specify the option with an &#x60;optionId&#x60;.  *  To clear, pass an option with &#x60;optionId&#x60; as &#x60;-1&#x60;. | [optional] 
**SingleVersionPickerFields** | Pointer to [**[]JiraSingleVersionPickerField**](JiraSingleVersionPickerField.md) | Add or clear a single version picker field:   *  To add, specify the version with a &#x60;versionId&#x60;.  *  To clear, set &#x60;versionId&#x60; to &#x60;-1&#x60;. | [optional] 
**Status** | Pointer to [**JiraStatusInput**](JiraStatusInput.md) |  | [optional] 
**TimeTrackingField** | Pointer to [**JiraTimeTrackingField**](JiraTimeTrackingField.md) | Edit the time tracking field. | [optional] 
**UrlFields** | Pointer to [**[]JiraUrlField**](JiraUrlField.md) | Add or clear a URL field:   *  To add, provide the &#x60;url&#x60; with the desired URL value.  *  To clear, set &#x60;url&#x60; to an empty string. | [optional] 

## Methods

### NewJiraIssueFields

`func NewJiraIssueFields() *JiraIssueFields`

NewJiraIssueFields instantiates a new JiraIssueFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraIssueFieldsWithDefaults

`func NewJiraIssueFieldsWithDefaults() *JiraIssueFields`

NewJiraIssueFieldsWithDefaults instantiates a new JiraIssueFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadingSelectFields

`func (o *JiraIssueFields) GetCascadingSelectFields() []JiraCascadingSelectField`

GetCascadingSelectFields returns the CascadingSelectFields field if non-nil, zero value otherwise.

### GetCascadingSelectFieldsOk

`func (o *JiraIssueFields) GetCascadingSelectFieldsOk() (*[]JiraCascadingSelectField, bool)`

GetCascadingSelectFieldsOk returns a tuple with the CascadingSelectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadingSelectFields

`func (o *JiraIssueFields) SetCascadingSelectFields(v []JiraCascadingSelectField)`

SetCascadingSelectFields sets CascadingSelectFields field to given value.

### HasCascadingSelectFields

`func (o *JiraIssueFields) HasCascadingSelectFields() bool`

HasCascadingSelectFields returns a boolean if a field has been set.

### GetClearableNumberFields

`func (o *JiraIssueFields) GetClearableNumberFields() []JiraNumberField`

GetClearableNumberFields returns the ClearableNumberFields field if non-nil, zero value otherwise.

### GetClearableNumberFieldsOk

`func (o *JiraIssueFields) GetClearableNumberFieldsOk() (*[]JiraNumberField, bool)`

GetClearableNumberFieldsOk returns a tuple with the ClearableNumberFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearableNumberFields

`func (o *JiraIssueFields) SetClearableNumberFields(v []JiraNumberField)`

SetClearableNumberFields sets ClearableNumberFields field to given value.

### HasClearableNumberFields

`func (o *JiraIssueFields) HasClearableNumberFields() bool`

HasClearableNumberFields returns a boolean if a field has been set.

### GetColorFields

`func (o *JiraIssueFields) GetColorFields() []JiraColorField`

GetColorFields returns the ColorFields field if non-nil, zero value otherwise.

### GetColorFieldsOk

`func (o *JiraIssueFields) GetColorFieldsOk() (*[]JiraColorField, bool)`

GetColorFieldsOk returns a tuple with the ColorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFields

`func (o *JiraIssueFields) SetColorFields(v []JiraColorField)`

SetColorFields sets ColorFields field to given value.

### HasColorFields

`func (o *JiraIssueFields) HasColorFields() bool`

HasColorFields returns a boolean if a field has been set.

### GetDatePickerFields

`func (o *JiraIssueFields) GetDatePickerFields() []JiraDateField`

GetDatePickerFields returns the DatePickerFields field if non-nil, zero value otherwise.

### GetDatePickerFieldsOk

`func (o *JiraIssueFields) GetDatePickerFieldsOk() (*[]JiraDateField, bool)`

GetDatePickerFieldsOk returns a tuple with the DatePickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePickerFields

`func (o *JiraIssueFields) SetDatePickerFields(v []JiraDateField)`

SetDatePickerFields sets DatePickerFields field to given value.

### HasDatePickerFields

`func (o *JiraIssueFields) HasDatePickerFields() bool`

HasDatePickerFields returns a boolean if a field has been set.

### GetDateTimePickerFields

`func (o *JiraIssueFields) GetDateTimePickerFields() []JiraDateTimeField`

GetDateTimePickerFields returns the DateTimePickerFields field if non-nil, zero value otherwise.

### GetDateTimePickerFieldsOk

`func (o *JiraIssueFields) GetDateTimePickerFieldsOk() (*[]JiraDateTimeField, bool)`

GetDateTimePickerFieldsOk returns a tuple with the DateTimePickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimePickerFields

`func (o *JiraIssueFields) SetDateTimePickerFields(v []JiraDateTimeField)`

SetDateTimePickerFields sets DateTimePickerFields field to given value.

### HasDateTimePickerFields

`func (o *JiraIssueFields) HasDateTimePickerFields() bool`

HasDateTimePickerFields returns a boolean if a field has been set.

### GetIssueType

`func (o *JiraIssueFields) GetIssueType() JiraIssueTypeField`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *JiraIssueFields) GetIssueTypeOk() (*JiraIssueTypeField, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *JiraIssueFields) SetIssueType(v JiraIssueTypeField)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *JiraIssueFields) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLabelsFields

`func (o *JiraIssueFields) GetLabelsFields() []JiraLabelsField`

GetLabelsFields returns the LabelsFields field if non-nil, zero value otherwise.

### GetLabelsFieldsOk

`func (o *JiraIssueFields) GetLabelsFieldsOk() (*[]JiraLabelsField, bool)`

GetLabelsFieldsOk returns a tuple with the LabelsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsFields

`func (o *JiraIssueFields) SetLabelsFields(v []JiraLabelsField)`

SetLabelsFields sets LabelsFields field to given value.

### HasLabelsFields

`func (o *JiraIssueFields) HasLabelsFields() bool`

HasLabelsFields returns a boolean if a field has been set.

### GetMultipleGroupPickerFields

`func (o *JiraIssueFields) GetMultipleGroupPickerFields() []JiraMultipleGroupPickerField`

GetMultipleGroupPickerFields returns the MultipleGroupPickerFields field if non-nil, zero value otherwise.

### GetMultipleGroupPickerFieldsOk

`func (o *JiraIssueFields) GetMultipleGroupPickerFieldsOk() (*[]JiraMultipleGroupPickerField, bool)`

GetMultipleGroupPickerFieldsOk returns a tuple with the MultipleGroupPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleGroupPickerFields

`func (o *JiraIssueFields) SetMultipleGroupPickerFields(v []JiraMultipleGroupPickerField)`

SetMultipleGroupPickerFields sets MultipleGroupPickerFields field to given value.

### HasMultipleGroupPickerFields

`func (o *JiraIssueFields) HasMultipleGroupPickerFields() bool`

HasMultipleGroupPickerFields returns a boolean if a field has been set.

### GetMultipleSelectClearableUserPickerFields

`func (o *JiraIssueFields) GetMultipleSelectClearableUserPickerFields() []JiraMultipleSelectUserPickerField`

GetMultipleSelectClearableUserPickerFields returns the MultipleSelectClearableUserPickerFields field if non-nil, zero value otherwise.

### GetMultipleSelectClearableUserPickerFieldsOk

`func (o *JiraIssueFields) GetMultipleSelectClearableUserPickerFieldsOk() (*[]JiraMultipleSelectUserPickerField, bool)`

GetMultipleSelectClearableUserPickerFieldsOk returns a tuple with the MultipleSelectClearableUserPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSelectClearableUserPickerFields

`func (o *JiraIssueFields) SetMultipleSelectClearableUserPickerFields(v []JiraMultipleSelectUserPickerField)`

SetMultipleSelectClearableUserPickerFields sets MultipleSelectClearableUserPickerFields field to given value.

### HasMultipleSelectClearableUserPickerFields

`func (o *JiraIssueFields) HasMultipleSelectClearableUserPickerFields() bool`

HasMultipleSelectClearableUserPickerFields returns a boolean if a field has been set.

### GetMultipleSelectFields

`func (o *JiraIssueFields) GetMultipleSelectFields() []JiraMultipleSelectField`

GetMultipleSelectFields returns the MultipleSelectFields field if non-nil, zero value otherwise.

### GetMultipleSelectFieldsOk

`func (o *JiraIssueFields) GetMultipleSelectFieldsOk() (*[]JiraMultipleSelectField, bool)`

GetMultipleSelectFieldsOk returns a tuple with the MultipleSelectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSelectFields

`func (o *JiraIssueFields) SetMultipleSelectFields(v []JiraMultipleSelectField)`

SetMultipleSelectFields sets MultipleSelectFields field to given value.

### HasMultipleSelectFields

`func (o *JiraIssueFields) HasMultipleSelectFields() bool`

HasMultipleSelectFields returns a boolean if a field has been set.

### GetMultipleVersionPickerFields

`func (o *JiraIssueFields) GetMultipleVersionPickerFields() []JiraMultipleVersionPickerField`

GetMultipleVersionPickerFields returns the MultipleVersionPickerFields field if non-nil, zero value otherwise.

### GetMultipleVersionPickerFieldsOk

`func (o *JiraIssueFields) GetMultipleVersionPickerFieldsOk() (*[]JiraMultipleVersionPickerField, bool)`

GetMultipleVersionPickerFieldsOk returns a tuple with the MultipleVersionPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVersionPickerFields

`func (o *JiraIssueFields) SetMultipleVersionPickerFields(v []JiraMultipleVersionPickerField)`

SetMultipleVersionPickerFields sets MultipleVersionPickerFields field to given value.

### HasMultipleVersionPickerFields

`func (o *JiraIssueFields) HasMultipleVersionPickerFields() bool`

HasMultipleVersionPickerFields returns a boolean if a field has been set.

### GetMultiselectComponents

`func (o *JiraIssueFields) GetMultiselectComponents() JiraMultiSelectComponentField`

GetMultiselectComponents returns the MultiselectComponents field if non-nil, zero value otherwise.

### GetMultiselectComponentsOk

`func (o *JiraIssueFields) GetMultiselectComponentsOk() (*JiraMultiSelectComponentField, bool)`

GetMultiselectComponentsOk returns a tuple with the MultiselectComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiselectComponents

`func (o *JiraIssueFields) SetMultiselectComponents(v JiraMultiSelectComponentField)`

SetMultiselectComponents sets MultiselectComponents field to given value.

### HasMultiselectComponents

`func (o *JiraIssueFields) HasMultiselectComponents() bool`

HasMultiselectComponents returns a boolean if a field has been set.

### GetOriginalEstimateField

`func (o *JiraIssueFields) GetOriginalEstimateField() JiraDurationField`

GetOriginalEstimateField returns the OriginalEstimateField field if non-nil, zero value otherwise.

### GetOriginalEstimateFieldOk

`func (o *JiraIssueFields) GetOriginalEstimateFieldOk() (*JiraDurationField, bool)`

GetOriginalEstimateFieldOk returns a tuple with the OriginalEstimateField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimateField

`func (o *JiraIssueFields) SetOriginalEstimateField(v JiraDurationField)`

SetOriginalEstimateField sets OriginalEstimateField field to given value.

### HasOriginalEstimateField

`func (o *JiraIssueFields) HasOriginalEstimateField() bool`

HasOriginalEstimateField returns a boolean if a field has been set.

### GetPriority

`func (o *JiraIssueFields) GetPriority() JiraPriorityField`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JiraIssueFields) GetPriorityOk() (*JiraPriorityField, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JiraIssueFields) SetPriority(v JiraPriorityField)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JiraIssueFields) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRichTextFields

`func (o *JiraIssueFields) GetRichTextFields() []JiraRichTextField`

GetRichTextFields returns the RichTextFields field if non-nil, zero value otherwise.

### GetRichTextFieldsOk

`func (o *JiraIssueFields) GetRichTextFieldsOk() (*[]JiraRichTextField, bool)`

GetRichTextFieldsOk returns a tuple with the RichTextFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichTextFields

`func (o *JiraIssueFields) SetRichTextFields(v []JiraRichTextField)`

SetRichTextFields sets RichTextFields field to given value.

### HasRichTextFields

`func (o *JiraIssueFields) HasRichTextFields() bool`

HasRichTextFields returns a boolean if a field has been set.

### GetSingleGroupPickerFields

`func (o *JiraIssueFields) GetSingleGroupPickerFields() []JiraSingleGroupPickerField`

GetSingleGroupPickerFields returns the SingleGroupPickerFields field if non-nil, zero value otherwise.

### GetSingleGroupPickerFieldsOk

`func (o *JiraIssueFields) GetSingleGroupPickerFieldsOk() (*[]JiraSingleGroupPickerField, bool)`

GetSingleGroupPickerFieldsOk returns a tuple with the SingleGroupPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleGroupPickerFields

`func (o *JiraIssueFields) SetSingleGroupPickerFields(v []JiraSingleGroupPickerField)`

SetSingleGroupPickerFields sets SingleGroupPickerFields field to given value.

### HasSingleGroupPickerFields

`func (o *JiraIssueFields) HasSingleGroupPickerFields() bool`

HasSingleGroupPickerFields returns a boolean if a field has been set.

### GetSingleLineTextFields

`func (o *JiraIssueFields) GetSingleLineTextFields() []JiraSingleLineTextField`

GetSingleLineTextFields returns the SingleLineTextFields field if non-nil, zero value otherwise.

### GetSingleLineTextFieldsOk

`func (o *JiraIssueFields) GetSingleLineTextFieldsOk() (*[]JiraSingleLineTextField, bool)`

GetSingleLineTextFieldsOk returns a tuple with the SingleLineTextFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleLineTextFields

`func (o *JiraIssueFields) SetSingleLineTextFields(v []JiraSingleLineTextField)`

SetSingleLineTextFields sets SingleLineTextFields field to given value.

### HasSingleLineTextFields

`func (o *JiraIssueFields) HasSingleLineTextFields() bool`

HasSingleLineTextFields returns a boolean if a field has been set.

### GetSingleSelectClearableUserPickerFields

`func (o *JiraIssueFields) GetSingleSelectClearableUserPickerFields() []JiraSingleSelectUserPickerField`

GetSingleSelectClearableUserPickerFields returns the SingleSelectClearableUserPickerFields field if non-nil, zero value otherwise.

### GetSingleSelectClearableUserPickerFieldsOk

`func (o *JiraIssueFields) GetSingleSelectClearableUserPickerFieldsOk() (*[]JiraSingleSelectUserPickerField, bool)`

GetSingleSelectClearableUserPickerFieldsOk returns a tuple with the SingleSelectClearableUserPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelectClearableUserPickerFields

`func (o *JiraIssueFields) SetSingleSelectClearableUserPickerFields(v []JiraSingleSelectUserPickerField)`

SetSingleSelectClearableUserPickerFields sets SingleSelectClearableUserPickerFields field to given value.

### HasSingleSelectClearableUserPickerFields

`func (o *JiraIssueFields) HasSingleSelectClearableUserPickerFields() bool`

HasSingleSelectClearableUserPickerFields returns a boolean if a field has been set.

### GetSingleSelectFields

`func (o *JiraIssueFields) GetSingleSelectFields() []JiraSingleSelectField`

GetSingleSelectFields returns the SingleSelectFields field if non-nil, zero value otherwise.

### GetSingleSelectFieldsOk

`func (o *JiraIssueFields) GetSingleSelectFieldsOk() (*[]JiraSingleSelectField, bool)`

GetSingleSelectFieldsOk returns a tuple with the SingleSelectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelectFields

`func (o *JiraIssueFields) SetSingleSelectFields(v []JiraSingleSelectField)`

SetSingleSelectFields sets SingleSelectFields field to given value.

### HasSingleSelectFields

`func (o *JiraIssueFields) HasSingleSelectFields() bool`

HasSingleSelectFields returns a boolean if a field has been set.

### GetSingleVersionPickerFields

`func (o *JiraIssueFields) GetSingleVersionPickerFields() []JiraSingleVersionPickerField`

GetSingleVersionPickerFields returns the SingleVersionPickerFields field if non-nil, zero value otherwise.

### GetSingleVersionPickerFieldsOk

`func (o *JiraIssueFields) GetSingleVersionPickerFieldsOk() (*[]JiraSingleVersionPickerField, bool)`

GetSingleVersionPickerFieldsOk returns a tuple with the SingleVersionPickerFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleVersionPickerFields

`func (o *JiraIssueFields) SetSingleVersionPickerFields(v []JiraSingleVersionPickerField)`

SetSingleVersionPickerFields sets SingleVersionPickerFields field to given value.

### HasSingleVersionPickerFields

`func (o *JiraIssueFields) HasSingleVersionPickerFields() bool`

HasSingleVersionPickerFields returns a boolean if a field has been set.

### GetStatus

`func (o *JiraIssueFields) GetStatus() JiraStatusInput`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JiraIssueFields) GetStatusOk() (*JiraStatusInput, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JiraIssueFields) SetStatus(v JiraStatusInput)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JiraIssueFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeTrackingField

`func (o *JiraIssueFields) GetTimeTrackingField() JiraTimeTrackingField`

GetTimeTrackingField returns the TimeTrackingField field if non-nil, zero value otherwise.

### GetTimeTrackingFieldOk

`func (o *JiraIssueFields) GetTimeTrackingFieldOk() (*JiraTimeTrackingField, bool)`

GetTimeTrackingFieldOk returns a tuple with the TimeTrackingField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTrackingField

`func (o *JiraIssueFields) SetTimeTrackingField(v JiraTimeTrackingField)`

SetTimeTrackingField sets TimeTrackingField field to given value.

### HasTimeTrackingField

`func (o *JiraIssueFields) HasTimeTrackingField() bool`

HasTimeTrackingField returns a boolean if a field has been set.

### GetUrlFields

`func (o *JiraIssueFields) GetUrlFields() []JiraUrlField`

GetUrlFields returns the UrlFields field if non-nil, zero value otherwise.

### GetUrlFieldsOk

`func (o *JiraIssueFields) GetUrlFieldsOk() (*[]JiraUrlField, bool)`

GetUrlFieldsOk returns a tuple with the UrlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFields

`func (o *JiraIssueFields) SetUrlFields(v []JiraUrlField)`

SetUrlFields sets UrlFields field to given value.

### HasUrlFields

`func (o *JiraIssueFields) HasUrlFields() bool`

HasUrlFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


