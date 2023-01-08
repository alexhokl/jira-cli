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
import (
	"time"
)

// Metadata for an issue attachment.
type AttachmentMetadata struct {
	// Details of the user who attached the file.
	Author *AllOfAttachmentMetadataAuthor `json:"author,omitempty"`
	// The URL of the attachment.
	Content string `json:"content,omitempty"`
	// The datetime the attachment was created.
	Created time.Time `json:"created,omitempty"`
	// The name of the attachment file.
	Filename string `json:"filename,omitempty"`
	// The ID of the attachment.
	Id int64 `json:"id,omitempty"`
	// The MIME type of the attachment.
	MimeType string `json:"mimeType,omitempty"`
	// Additional properties of the attachment.
	Properties map[string]Object `json:"properties,omitempty"`
	// The URL of the attachment metadata details.
	Self string `json:"self,omitempty"`
	// The size of the attachment.
	Size int64 `json:"size,omitempty"`
	// The URL of a thumbnail representing the attachment.
	Thumbnail string `json:"thumbnail,omitempty"`
}
