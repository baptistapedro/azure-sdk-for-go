//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExportConfigurationsClient contains the methods for the ExportConfigurations group.
// Don't use this type directly, use NewExportConfigurationsClient() instead.
type ExportConfigurationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExportConfigurationsClient creates a new instance of ExportConfigurationsClient with the specified values.
func NewExportConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExportConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ExportConfigurationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create a Continuous Export configuration of an Application Insights component.
// If the operation fails it returns a generic error.
func (client *ExportConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest, options *ExportConfigurationsCreateOptions) (ExportConfigurationsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, exportProperties, options)
	if err != nil {
		return ExportConfigurationsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ExportConfigurationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest, options *ExportConfigurationsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, exportProperties)
}

// createHandleResponse handles the Create response.
func (client *ExportConfigurationsClient) createHandleResponse(resp *http.Response) (ExportConfigurationsCreateResponse, error) {
	result := ExportConfigurationsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentExportConfigurationArray); err != nil {
		return ExportConfigurationsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *ExportConfigurationsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Delete a Continuous Export configuration of an Application Insights component.
// If the operation fails it returns a generic error.
func (client *ExportConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsDeleteOptions) (ExportConfigurationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, exportID, options)
	if err != nil {
		return ExportConfigurationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ExportConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ExportConfigurationsClient) deleteHandleResponse(resp *http.Response) (ExportConfigurationsDeleteResponse, error) {
	result := ExportConfigurationsDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentExportConfiguration); err != nil {
		return ExportConfigurationsDeleteResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExportConfigurationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get the Continuous Export configuration for this export id.
// If the operation fails it returns a generic error.
func (client *ExportConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsGetOptions) (ExportConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, exportID, options)
	if err != nil {
		return ExportConfigurationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExportConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExportConfigurationsClient) getHandleResponse(resp *http.Response) (ExportConfigurationsGetResponse, error) {
	result := ExportConfigurationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentExportConfiguration); err != nil {
		return ExportConfigurationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExportConfigurationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets a list of Continuous Export configuration of an Application Insights component.
// If the operation fails it returns a generic error.
func (client *ExportConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *ExportConfigurationsListOptions) (ExportConfigurationsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ExportConfigurationsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ExportConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ExportConfigurationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExportConfigurationsClient) listHandleResponse(resp *http.Response) (ExportConfigurationsListResponse, error) {
	result := ExportConfigurationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentExportConfigurationArray); err != nil {
		return ExportConfigurationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExportConfigurationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Update the Continuous Export configuration for this export id.
// If the operation fails it returns a generic error.
func (client *ExportConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest, options *ExportConfigurationsUpdateOptions) (ExportConfigurationsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, exportID, exportProperties, options)
	if err != nil {
		return ExportConfigurationsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ExportConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest, options *ExportConfigurationsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, exportProperties)
}

// updateHandleResponse handles the Update response.
func (client *ExportConfigurationsClient) updateHandleResponse(resp *http.Response) (ExportConfigurationsUpdateResponse, error) {
	result := ExportConfigurationsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentExportConfiguration); err != nil {
		return ExportConfigurationsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ExportConfigurationsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}