# JsonNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to **bool** |  | [optional] 
**BigDecimal** | Pointer to **bool** |  | [optional] 
**BigInteger** | Pointer to **bool** |  | [optional] 
**BigIntegerValue** | Pointer to **int32** |  | [optional] 
**Binary** | Pointer to **bool** |  | [optional] 
**BinaryValue** | Pointer to **[]string** |  | [optional] 
**Boolean** | Pointer to **bool** |  | [optional] 
**BooleanValue** | Pointer to **bool** |  | [optional] 
**ContainerNode** | Pointer to **bool** |  | [optional] 
**DecimalValue** | Pointer to **float32** |  | [optional] 
**Double** | Pointer to **bool** |  | [optional] 
**DoubleValue** | Pointer to **float64** |  | [optional] 
**Elements** | Pointer to **map[string]interface{}** |  | [optional] 
**FieldNames** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**FloatingPointNumber** | Pointer to **bool** |  | [optional] 
**Int** | Pointer to **bool** |  | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**IntegralNumber** | Pointer to **bool** |  | [optional] 
**Long** | Pointer to **bool** |  | [optional] 
**LongValue** | Pointer to **int64** |  | [optional] 
**MissingNode** | Pointer to **bool** |  | [optional] 
**Null** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **bool** |  | [optional] 
**NumberType** | Pointer to **string** |  | [optional] 
**NumberValue** | Pointer to **float32** |  | [optional] 
**Object** | Pointer to **bool** |  | [optional] 
**Pojo** | Pointer to **bool** |  | [optional] 
**TextValue** | Pointer to **string** |  | [optional] 
**Textual** | Pointer to **bool** |  | [optional] 
**ValueAsBoolean** | Pointer to **bool** |  | [optional] 
**ValueAsDouble** | Pointer to **float64** |  | [optional] 
**ValueAsInt** | Pointer to **int32** |  | [optional] 
**ValueAsLong** | Pointer to **int64** |  | [optional] 
**ValueAsText** | Pointer to **string** |  | [optional] 
**ValueNode** | Pointer to **bool** |  | [optional] 

## Methods

### NewJsonNode

`func NewJsonNode() *JsonNode`

NewJsonNode instantiates a new JsonNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonNodeWithDefaults

`func NewJsonNodeWithDefaults() *JsonNode`

NewJsonNodeWithDefaults instantiates a new JsonNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *JsonNode) GetArray() bool`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *JsonNode) GetArrayOk() (*bool, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *JsonNode) SetArray(v bool)`

SetArray sets Array field to given value.

### HasArray

`func (o *JsonNode) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBigDecimal

`func (o *JsonNode) GetBigDecimal() bool`

GetBigDecimal returns the BigDecimal field if non-nil, zero value otherwise.

### GetBigDecimalOk

`func (o *JsonNode) GetBigDecimalOk() (*bool, bool)`

GetBigDecimalOk returns a tuple with the BigDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigDecimal

`func (o *JsonNode) SetBigDecimal(v bool)`

SetBigDecimal sets BigDecimal field to given value.

### HasBigDecimal

`func (o *JsonNode) HasBigDecimal() bool`

HasBigDecimal returns a boolean if a field has been set.

### GetBigInteger

`func (o *JsonNode) GetBigInteger() bool`

GetBigInteger returns the BigInteger field if non-nil, zero value otherwise.

### GetBigIntegerOk

`func (o *JsonNode) GetBigIntegerOk() (*bool, bool)`

GetBigIntegerOk returns a tuple with the BigInteger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigInteger

`func (o *JsonNode) SetBigInteger(v bool)`

SetBigInteger sets BigInteger field to given value.

### HasBigInteger

`func (o *JsonNode) HasBigInteger() bool`

HasBigInteger returns a boolean if a field has been set.

### GetBigIntegerValue

`func (o *JsonNode) GetBigIntegerValue() int32`

GetBigIntegerValue returns the BigIntegerValue field if non-nil, zero value otherwise.

### GetBigIntegerValueOk

`func (o *JsonNode) GetBigIntegerValueOk() (*int32, bool)`

GetBigIntegerValueOk returns a tuple with the BigIntegerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIntegerValue

`func (o *JsonNode) SetBigIntegerValue(v int32)`

SetBigIntegerValue sets BigIntegerValue field to given value.

### HasBigIntegerValue

`func (o *JsonNode) HasBigIntegerValue() bool`

HasBigIntegerValue returns a boolean if a field has been set.

### GetBinary

`func (o *JsonNode) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *JsonNode) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *JsonNode) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *JsonNode) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetBinaryValue

`func (o *JsonNode) GetBinaryValue() []string`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *JsonNode) GetBinaryValueOk() (*[]string, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *JsonNode) SetBinaryValue(v []string)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *JsonNode) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetBoolean

`func (o *JsonNode) GetBoolean() bool`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *JsonNode) GetBooleanOk() (*bool, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *JsonNode) SetBoolean(v bool)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *JsonNode) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### GetBooleanValue

`func (o *JsonNode) GetBooleanValue() bool`

GetBooleanValue returns the BooleanValue field if non-nil, zero value otherwise.

### GetBooleanValueOk

`func (o *JsonNode) GetBooleanValueOk() (*bool, bool)`

GetBooleanValueOk returns a tuple with the BooleanValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanValue

`func (o *JsonNode) SetBooleanValue(v bool)`

SetBooleanValue sets BooleanValue field to given value.

### HasBooleanValue

`func (o *JsonNode) HasBooleanValue() bool`

HasBooleanValue returns a boolean if a field has been set.

### GetContainerNode

`func (o *JsonNode) GetContainerNode() bool`

GetContainerNode returns the ContainerNode field if non-nil, zero value otherwise.

### GetContainerNodeOk

`func (o *JsonNode) GetContainerNodeOk() (*bool, bool)`

GetContainerNodeOk returns a tuple with the ContainerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNode

`func (o *JsonNode) SetContainerNode(v bool)`

SetContainerNode sets ContainerNode field to given value.

### HasContainerNode

`func (o *JsonNode) HasContainerNode() bool`

HasContainerNode returns a boolean if a field has been set.

### GetDecimalValue

`func (o *JsonNode) GetDecimalValue() float32`

GetDecimalValue returns the DecimalValue field if non-nil, zero value otherwise.

### GetDecimalValueOk

`func (o *JsonNode) GetDecimalValueOk() (*float32, bool)`

GetDecimalValueOk returns a tuple with the DecimalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalValue

`func (o *JsonNode) SetDecimalValue(v float32)`

SetDecimalValue sets DecimalValue field to given value.

### HasDecimalValue

`func (o *JsonNode) HasDecimalValue() bool`

HasDecimalValue returns a boolean if a field has been set.

### GetDouble

`func (o *JsonNode) GetDouble() bool`

GetDouble returns the Double field if non-nil, zero value otherwise.

### GetDoubleOk

`func (o *JsonNode) GetDoubleOk() (*bool, bool)`

GetDoubleOk returns a tuple with the Double field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDouble

`func (o *JsonNode) SetDouble(v bool)`

SetDouble sets Double field to given value.

### HasDouble

`func (o *JsonNode) HasDouble() bool`

HasDouble returns a boolean if a field has been set.

### GetDoubleValue

`func (o *JsonNode) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *JsonNode) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *JsonNode) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *JsonNode) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### GetElements

`func (o *JsonNode) GetElements() map[string]interface{}`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *JsonNode) GetElementsOk() (*map[string]interface{}, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *JsonNode) SetElements(v map[string]interface{})`

SetElements sets Elements field to given value.

### HasElements

`func (o *JsonNode) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetFieldNames

`func (o *JsonNode) GetFieldNames() map[string]interface{}`

GetFieldNames returns the FieldNames field if non-nil, zero value otherwise.

### GetFieldNamesOk

`func (o *JsonNode) GetFieldNamesOk() (*map[string]interface{}, bool)`

GetFieldNamesOk returns a tuple with the FieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNames

`func (o *JsonNode) SetFieldNames(v map[string]interface{})`

SetFieldNames sets FieldNames field to given value.

### HasFieldNames

`func (o *JsonNode) HasFieldNames() bool`

HasFieldNames returns a boolean if a field has been set.

### GetFields

`func (o *JsonNode) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *JsonNode) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *JsonNode) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *JsonNode) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFloatingPointNumber

`func (o *JsonNode) GetFloatingPointNumber() bool`

GetFloatingPointNumber returns the FloatingPointNumber field if non-nil, zero value otherwise.

### GetFloatingPointNumberOk

`func (o *JsonNode) GetFloatingPointNumberOk() (*bool, bool)`

GetFloatingPointNumberOk returns a tuple with the FloatingPointNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingPointNumber

`func (o *JsonNode) SetFloatingPointNumber(v bool)`

SetFloatingPointNumber sets FloatingPointNumber field to given value.

### HasFloatingPointNumber

`func (o *JsonNode) HasFloatingPointNumber() bool`

HasFloatingPointNumber returns a boolean if a field has been set.

### GetInt

`func (o *JsonNode) GetInt() bool`

GetInt returns the Int field if non-nil, zero value otherwise.

### GetIntOk

`func (o *JsonNode) GetIntOk() (*bool, bool)`

GetIntOk returns a tuple with the Int field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt

`func (o *JsonNode) SetInt(v bool)`

SetInt sets Int field to given value.

### HasInt

`func (o *JsonNode) HasInt() bool`

HasInt returns a boolean if a field has been set.

### GetIntValue

`func (o *JsonNode) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *JsonNode) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *JsonNode) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *JsonNode) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetIntegralNumber

`func (o *JsonNode) GetIntegralNumber() bool`

GetIntegralNumber returns the IntegralNumber field if non-nil, zero value otherwise.

### GetIntegralNumberOk

`func (o *JsonNode) GetIntegralNumberOk() (*bool, bool)`

GetIntegralNumberOk returns a tuple with the IntegralNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegralNumber

`func (o *JsonNode) SetIntegralNumber(v bool)`

SetIntegralNumber sets IntegralNumber field to given value.

### HasIntegralNumber

`func (o *JsonNode) HasIntegralNumber() bool`

HasIntegralNumber returns a boolean if a field has been set.

### GetLong

`func (o *JsonNode) GetLong() bool`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *JsonNode) GetLongOk() (*bool, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *JsonNode) SetLong(v bool)`

SetLong sets Long field to given value.

### HasLong

`func (o *JsonNode) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetLongValue

`func (o *JsonNode) GetLongValue() int64`

GetLongValue returns the LongValue field if non-nil, zero value otherwise.

### GetLongValueOk

`func (o *JsonNode) GetLongValueOk() (*int64, bool)`

GetLongValueOk returns a tuple with the LongValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongValue

`func (o *JsonNode) SetLongValue(v int64)`

SetLongValue sets LongValue field to given value.

### HasLongValue

`func (o *JsonNode) HasLongValue() bool`

HasLongValue returns a boolean if a field has been set.

### GetMissingNode

`func (o *JsonNode) GetMissingNode() bool`

GetMissingNode returns the MissingNode field if non-nil, zero value otherwise.

### GetMissingNodeOk

`func (o *JsonNode) GetMissingNodeOk() (*bool, bool)`

GetMissingNodeOk returns a tuple with the MissingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingNode

`func (o *JsonNode) SetMissingNode(v bool)`

SetMissingNode sets MissingNode field to given value.

### HasMissingNode

`func (o *JsonNode) HasMissingNode() bool`

HasMissingNode returns a boolean if a field has been set.

### GetNull

`func (o *JsonNode) GetNull() bool`

GetNull returns the Null field if non-nil, zero value otherwise.

### GetNullOk

`func (o *JsonNode) GetNullOk() (*bool, bool)`

GetNullOk returns a tuple with the Null field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNull

`func (o *JsonNode) SetNull(v bool)`

SetNull sets Null field to given value.

### HasNull

`func (o *JsonNode) HasNull() bool`

HasNull returns a boolean if a field has been set.

### GetNumber

`func (o *JsonNode) GetNumber() bool`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *JsonNode) GetNumberOk() (*bool, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *JsonNode) SetNumber(v bool)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *JsonNode) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberType

`func (o *JsonNode) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *JsonNode) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *JsonNode) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *JsonNode) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### GetNumberValue

`func (o *JsonNode) GetNumberValue() float32`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *JsonNode) GetNumberValueOk() (*float32, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *JsonNode) SetNumberValue(v float32)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *JsonNode) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### GetObject

`func (o *JsonNode) GetObject() bool`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *JsonNode) GetObjectOk() (*bool, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *JsonNode) SetObject(v bool)`

SetObject sets Object field to given value.

### HasObject

`func (o *JsonNode) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPojo

`func (o *JsonNode) GetPojo() bool`

GetPojo returns the Pojo field if non-nil, zero value otherwise.

### GetPojoOk

`func (o *JsonNode) GetPojoOk() (*bool, bool)`

GetPojoOk returns a tuple with the Pojo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPojo

`func (o *JsonNode) SetPojo(v bool)`

SetPojo sets Pojo field to given value.

### HasPojo

`func (o *JsonNode) HasPojo() bool`

HasPojo returns a boolean if a field has been set.

### GetTextValue

`func (o *JsonNode) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *JsonNode) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *JsonNode) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *JsonNode) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.

### GetTextual

`func (o *JsonNode) GetTextual() bool`

GetTextual returns the Textual field if non-nil, zero value otherwise.

### GetTextualOk

`func (o *JsonNode) GetTextualOk() (*bool, bool)`

GetTextualOk returns a tuple with the Textual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextual

`func (o *JsonNode) SetTextual(v bool)`

SetTextual sets Textual field to given value.

### HasTextual

`func (o *JsonNode) HasTextual() bool`

HasTextual returns a boolean if a field has been set.

### GetValueAsBoolean

`func (o *JsonNode) GetValueAsBoolean() bool`

GetValueAsBoolean returns the ValueAsBoolean field if non-nil, zero value otherwise.

### GetValueAsBooleanOk

`func (o *JsonNode) GetValueAsBooleanOk() (*bool, bool)`

GetValueAsBooleanOk returns a tuple with the ValueAsBoolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsBoolean

`func (o *JsonNode) SetValueAsBoolean(v bool)`

SetValueAsBoolean sets ValueAsBoolean field to given value.

### HasValueAsBoolean

`func (o *JsonNode) HasValueAsBoolean() bool`

HasValueAsBoolean returns a boolean if a field has been set.

### GetValueAsDouble

`func (o *JsonNode) GetValueAsDouble() float64`

GetValueAsDouble returns the ValueAsDouble field if non-nil, zero value otherwise.

### GetValueAsDoubleOk

`func (o *JsonNode) GetValueAsDoubleOk() (*float64, bool)`

GetValueAsDoubleOk returns a tuple with the ValueAsDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsDouble

`func (o *JsonNode) SetValueAsDouble(v float64)`

SetValueAsDouble sets ValueAsDouble field to given value.

### HasValueAsDouble

`func (o *JsonNode) HasValueAsDouble() bool`

HasValueAsDouble returns a boolean if a field has been set.

### GetValueAsInt

`func (o *JsonNode) GetValueAsInt() int32`

GetValueAsInt returns the ValueAsInt field if non-nil, zero value otherwise.

### GetValueAsIntOk

`func (o *JsonNode) GetValueAsIntOk() (*int32, bool)`

GetValueAsIntOk returns a tuple with the ValueAsInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsInt

`func (o *JsonNode) SetValueAsInt(v int32)`

SetValueAsInt sets ValueAsInt field to given value.

### HasValueAsInt

`func (o *JsonNode) HasValueAsInt() bool`

HasValueAsInt returns a boolean if a field has been set.

### GetValueAsLong

`func (o *JsonNode) GetValueAsLong() int64`

GetValueAsLong returns the ValueAsLong field if non-nil, zero value otherwise.

### GetValueAsLongOk

`func (o *JsonNode) GetValueAsLongOk() (*int64, bool)`

GetValueAsLongOk returns a tuple with the ValueAsLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsLong

`func (o *JsonNode) SetValueAsLong(v int64)`

SetValueAsLong sets ValueAsLong field to given value.

### HasValueAsLong

`func (o *JsonNode) HasValueAsLong() bool`

HasValueAsLong returns a boolean if a field has been set.

### GetValueAsText

`func (o *JsonNode) GetValueAsText() string`

GetValueAsText returns the ValueAsText field if non-nil, zero value otherwise.

### GetValueAsTextOk

`func (o *JsonNode) GetValueAsTextOk() (*string, bool)`

GetValueAsTextOk returns a tuple with the ValueAsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsText

`func (o *JsonNode) SetValueAsText(v string)`

SetValueAsText sets ValueAsText field to given value.

### HasValueAsText

`func (o *JsonNode) HasValueAsText() bool`

HasValueAsText returns a boolean if a field has been set.

### GetValueNode

`func (o *JsonNode) GetValueNode() bool`

GetValueNode returns the ValueNode field if non-nil, zero value otherwise.

### GetValueNodeOk

`func (o *JsonNode) GetValueNodeOk() (*bool, bool)`

GetValueNodeOk returns a tuple with the ValueNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNode

`func (o *JsonNode) SetValueNode(v bool)`

SetValueNode sets ValueNode field to given value.

### HasValueNode

`func (o *JsonNode) HasValueNode() bool`

HasValueNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


