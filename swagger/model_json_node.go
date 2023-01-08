/*
 * The Jira Cloud platform REST API
 *
 * Jira Cloud platform REST API documentation
 *
 * API version: 1001.0.0-SNAPSHOT
 * Contact: ecosystem@atlassian.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type JsonNode struct {
	Array bool `json:"array,omitempty"`
	BigDecimal bool `json:"bigDecimal,omitempty"`
	BigInteger bool `json:"bigInteger,omitempty"`
	BigIntegerValue int32 `json:"bigIntegerValue,omitempty"`
	Binary bool `json:"binary,omitempty"`
	BinaryValue []string `json:"binaryValue,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	BooleanValue bool `json:"booleanValue,omitempty"`
	ContainerNode bool `json:"containerNode,omitempty"`
	DecimalValue float64 `json:"decimalValue,omitempty"`
	Double bool `json:"double,omitempty"`
	DoubleValue float64 `json:"doubleValue,omitempty"`
	Elements *interface{} `json:"elements,omitempty"`
	FieldNames *interface{} `json:"fieldNames,omitempty"`
	Fields *interface{} `json:"fields,omitempty"`
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`
	Int_ bool `json:"int,omitempty"`
	IntValue int32 `json:"intValue,omitempty"`
	IntegralNumber bool `json:"integralNumber,omitempty"`
	Long bool `json:"long,omitempty"`
	LongValue int64 `json:"longValue,omitempty"`
	MissingNode bool `json:"missingNode,omitempty"`
	Null bool `json:"null,omitempty"`
	Number bool `json:"number,omitempty"`
	NumberType string `json:"numberType,omitempty"`
	NumberValue float64 `json:"numberValue,omitempty"`
	Object bool `json:"object,omitempty"`
	Pojo bool `json:"pojo,omitempty"`
	TextValue string `json:"textValue,omitempty"`
	Textual bool `json:"textual,omitempty"`
	ValueAsBoolean bool `json:"valueAsBoolean,omitempty"`
	ValueAsDouble float64 `json:"valueAsDouble,omitempty"`
	ValueAsInt int32 `json:"valueAsInt,omitempty"`
	ValueAsLong int64 `json:"valueAsLong,omitempty"`
	ValueAsText string `json:"valueAsText,omitempty"`
	ValueNode bool `json:"valueNode,omitempty"`
}
