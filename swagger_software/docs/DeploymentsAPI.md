# \DeploymentsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeploymentByKey**](DeploymentsAPI.md#DeleteDeploymentByKey) | **Delete** /rest/deployments/0.1/pipelines/{pipelineId}/environments/{environmentId}/deployments/{deploymentSequenceNumber} | Delete a deployment by key
[**DeleteDeploymentsByProperty**](DeploymentsAPI.md#DeleteDeploymentsByProperty) | **Delete** /rest/deployments/0.1/bulkByProperties | Delete deployments by Property
[**GetDeploymentByKey**](DeploymentsAPI.md#GetDeploymentByKey) | **Get** /rest/deployments/0.1/pipelines/{pipelineId}/environments/{environmentId}/deployments/{deploymentSequenceNumber} | Get a deployment by key
[**GetDeploymentGatingStatusByKey**](DeploymentsAPI.md#GetDeploymentGatingStatusByKey) | **Get** /rest/deployments/0.1/pipelines/{pipelineId}/environments/{environmentId}/deployments/{deploymentSequenceNumber}/gating-status | Get deployment gating status by key
[**SubmitDeployments**](DeploymentsAPI.md#SubmitDeployments) | **Post** /rest/deployments/0.1/bulk | Submit deployment data



## DeleteDeploymentByKey

> DeleteDeploymentByKey(ctx, pipelineId, environmentId, deploymentSequenceNumber).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()

Delete a deployment by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraDeploymentInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	pipelineId := "pipelineId_example" // string | The ID of the deployment's pipeline. 
	environmentId := "environmentId_example" // string | The ID of the deployment's environment. 
	deploymentSequenceNumber := int64(789) // int64 | The deployment's deploymentSequenceNumber. 
	updateSequenceNumber := int64(789) // int64 | This parameter usage is no longer supported.  An optional `_updateSequenceNumber` to use to control deletion.  Only stored data with an `updateSequenceNumber` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentsAPI.DeleteDeploymentByKey(context.Background(), pipelineId, environmentId, deploymentSequenceNumber).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.DeleteDeploymentByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The ID of the deployment&#39;s pipeline.  | 
**environmentId** | **string** | The ID of the deployment&#39;s environment.  | 
**deploymentSequenceNumber** | **int64** | The deployment&#39;s deploymentSequenceNumber.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraDeploymentInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 



 **updateSequenceNumber** | **int64** | This parameter usage is no longer supported.  An optional &#x60;_updateSequenceNumber&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceNumber&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentsByProperty

> DeleteDeploymentsByProperty(ctx).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()

Delete deployments by Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraDeploymentInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	updateSequenceNumber := int64(789) // int64 | This parameter usage is no longer supported.  An optional `updateSequenceNumber` to use to control deletion.  Only stored data with an `updateSequenceNumber` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentsAPI.DeleteDeploymentsByProperty(context.Background()).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.DeleteDeploymentsByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentsByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraDeploymentInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 
 **updateSequenceNumber** | **int64** | This parameter usage is no longer supported.  An optional &#x60;updateSequenceNumber&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceNumber&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentByKey

> DeploymentData1 GetDeploymentByKey(ctx, pipelineId, environmentId, deploymentSequenceNumber).Authorization(authorization).Execute()

Get a deployment by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraDeploymentInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	pipelineId := "pipelineId_example" // string | The ID of the deployment's pipeline. 
	environmentId := "environmentId_example" // string | The ID of the deployment's environment. 
	deploymentSequenceNumber := int64(789) // int64 | The deployment's deploymentSequenceNumber. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.GetDeploymentByKey(context.Background(), pipelineId, environmentId, deploymentSequenceNumber).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.GetDeploymentByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentByKey`: DeploymentData1
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.GetDeploymentByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The ID of the deployment&#39;s pipeline.  | 
**environmentId** | **string** | The ID of the deployment&#39;s environment.  | 
**deploymentSequenceNumber** | **int64** | The deployment&#39;s deploymentSequenceNumber.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraDeploymentInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 




### Return type

[**DeploymentData1**](DeploymentData1.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentGatingStatusByKey

> SubmitDeploymentsResponse1 GetDeploymentGatingStatusByKey(ctx, pipelineId, environmentId, deploymentSequenceNumber).Execute()

Get deployment gating status by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	pipelineId := "pipelineId_example" // string | The ID of the Deployment's pipeline. 
	environmentId := "environmentId_example" // string | The ID of the Deployment's environment. 
	deploymentSequenceNumber := int64(789) // int64 | The Deployment's deploymentSequenceNumber. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.GetDeploymentGatingStatusByKey(context.Background(), pipelineId, environmentId, deploymentSequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.GetDeploymentGatingStatusByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentGatingStatusByKey`: SubmitDeploymentsResponse1
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.GetDeploymentGatingStatusByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The ID of the Deployment&#39;s pipeline.  | 
**environmentId** | **string** | The ID of the Deployment&#39;s environment.  | 
**deploymentSequenceNumber** | **int64** | The Deployment&#39;s deploymentSequenceNumber.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentGatingStatusByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SubmitDeploymentsResponse1**](SubmitDeploymentsResponse1.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitDeployments

> SubmitDeploymentsResponse SubmitDeployments(ctx).Authorization(authorization).SubmitDeploymentRequest(submitDeploymentRequest).Execute()

Submit deployment data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraDeploymentInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	submitDeploymentRequest := *openapiclient.NewSubmitDeploymentRequest([]openapiclient.DeploymentData{*openapiclient.NewDeploymentData(int64(100), int64(1), "Deployment number 16 of Data Depot", "http://mydeployer.com/project1/1111-222-333/prod-east", "The bits are being transferred", time.Now(), "in_progress", *openapiclient.NewPipeline("e9c906a7-451f-4fa6-ae1a-c389e2e2d87c", "Data Depot Deployment", "http://mydeployer.com/project1"), *openapiclient.NewEnvironment("8ec94d72-a4fc-4ac0-b31d-c5a595f373ba", "US East", "production"))}) // SubmitDeploymentRequest | Deployment data to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.SubmitDeployments(context.Background()).Authorization(authorization).SubmitDeploymentRequest(submitDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.SubmitDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitDeployments`: SubmitDeploymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.SubmitDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraDeploymentInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 
 **submitDeploymentRequest** | [**SubmitDeploymentRequest**](SubmitDeploymentRequest.md) | Deployment data to submit.  | 

### Return type

[**SubmitDeploymentsResponse**](SubmitDeploymentsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

