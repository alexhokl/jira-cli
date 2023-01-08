# \ProjectTemplatesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectWithCustomTemplate**](ProjectTemplatesAPI.md#CreateProjectWithCustomTemplate) | **Post** /rest/api/3/project-template | Create custom project
[**EditTemplate**](ProjectTemplatesAPI.md#EditTemplate) | **Put** /rest/api/3/project-template/edit-template | Edit a custom project template
[**LiveTemplate**](ProjectTemplatesAPI.md#LiveTemplate) | **Get** /rest/api/3/project-template/live-template | Gets a custom project template
[**RemoveTemplate**](ProjectTemplatesAPI.md#RemoveTemplate) | **Delete** /rest/api/3/project-template/remove-template | Deletes a custom project template
[**SaveTemplate**](ProjectTemplatesAPI.md#SaveTemplate) | **Post** /rest/api/3/project-template/save-template | Save a custom project template



## CreateProjectWithCustomTemplate

> CreateProjectWithCustomTemplate(ctx).ProjectCustomTemplateCreateRequestDTO(projectCustomTemplateCreateRequestDTO).Execute()

Create custom project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectCustomTemplateCreateRequestDTO := *openapiclient.NewProjectCustomTemplateCreateRequestDTO() // ProjectCustomTemplateCreateRequestDTO | The JSON payload containing the project details and capabilities

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTemplatesAPI.CreateProjectWithCustomTemplate(context.Background()).ProjectCustomTemplateCreateRequestDTO(projectCustomTemplateCreateRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.CreateProjectWithCustomTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectWithCustomTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCustomTemplateCreateRequestDTO** | [**ProjectCustomTemplateCreateRequestDTO**](ProjectCustomTemplateCreateRequestDTO.md) | The JSON payload containing the project details and capabilities | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTemplate

> interface{} EditTemplate(ctx).EditTemplateRequest(editTemplateRequest).Execute()

Edit a custom project template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	editTemplateRequest := *openapiclient.NewEditTemplateRequest() // EditTemplateRequest | The object containing the updated template details: name, description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.EditTemplate(context.Background()).EditTemplateRequest(editTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.EditTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTemplate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.EditTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editTemplateRequest** | [**EditTemplateRequest**](EditTemplateRequest.md) | The object containing the updated template details: name, description | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveTemplate

> ProjectTemplateModel LiveTemplate(ctx).ProjectId(projectId).TemplateKey(templateKey).Execute()

Gets a custom project template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectId := "projectId_example" // string | optional - The \\{@link String\\} containing the project key linked to the custom template to retrieve (optional)
	templateKey := "templateKey_example" // string | optional - The \\{@link String\\} containing the key of the custom template to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.LiveTemplate(context.Background()).ProjectId(projectId).TemplateKey(templateKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.LiveTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveTemplate`: ProjectTemplateModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.LiveTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | optional - The \\{@link String\\} containing the project key linked to the custom template to retrieve | 
 **templateKey** | **string** | optional - The \\{@link String\\} containing the key of the custom template to retrieve | 

### Return type

[**ProjectTemplateModel**](ProjectTemplateModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTemplate

> interface{} RemoveTemplate(ctx).TemplateKey(templateKey).Execute()

Deletes a custom project template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	templateKey := "templateKey_example" // string | The \\{@link String\\} containing the key of the custom template to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.RemoveTemplate(context.Background()).TemplateKey(templateKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.RemoveTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTemplate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.RemoveTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateKey** | **string** | The \\{@link String\\} containing the key of the custom template to remove | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveTemplate

> SaveTemplateResponse SaveTemplate(ctx).SaveTemplateRequest(saveTemplateRequest).Execute()

Save a custom project template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	saveTemplateRequest := *openapiclient.NewSaveTemplateRequest() // SaveTemplateRequest | The object containing the template basic details: name, description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.SaveTemplate(context.Background()).SaveTemplateRequest(saveTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.SaveTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveTemplate`: SaveTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.SaveTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveTemplateRequest** | [**SaveTemplateRequest**](SaveTemplateRequest.md) | The object containing the template basic details: name, description | 

### Return type

[**SaveTemplateResponse**](SaveTemplateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

