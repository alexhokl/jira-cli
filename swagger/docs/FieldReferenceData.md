# FieldReferenceData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | **string** | Whether the field provide auto-complete suggestions. | [optional] [default to null]
**Cfid** | **string** | If the item is a custom field, the ID of the custom field. | [optional] [default to null]
**Deprecated** | **string** | Whether this field has been deprecated. | [optional] [default to null]
**DeprecatedSearcherKey** | **string** | The searcher key of the field, only passed when the field is deprecated. | [optional] [default to null]
**DisplayName** | **string** | The display name contains the following:   *  for system fields, the field name. For example, &#x60;Summary&#x60;.  *  for collapsed custom fields, the field name followed by a hyphen and then the field name and field type. For example, &#x60;Component - Component[Dropdown]&#x60;.  *  for other custom fields, the field name followed by a hyphen and then the custom field ID. For example, &#x60;Component - cf[10061]&#x60;. | [optional] [default to null]
**Operators** | **[]string** | The valid search operators for the field. | [optional] [default to null]
**Orderable** | **string** | Whether the field can be used in a query&#x27;s &#x60;ORDER BY&#x60; clause. | [optional] [default to null]
**Searchable** | **string** | Whether the content of this field can be searched. | [optional] [default to null]
**Types** | **[]string** | The data types of items in the field. | [optional] [default to null]
**Value** | **string** | The field identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

