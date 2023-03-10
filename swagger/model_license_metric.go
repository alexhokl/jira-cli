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

// A license metric
type LicenseMetric struct {
	// The key of the license metric.
	Key string `json:"key,omitempty"`
	// The value for the license metric.
	Value string `json:"value,omitempty"`
}
