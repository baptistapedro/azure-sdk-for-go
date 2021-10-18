//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// WorkflowRunActionsClient contains the methods for the WorkflowRunActions group.
// Don't use this type directly, use NewWorkflowRunActionsClient() instead.
type WorkflowRunActionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkflowRunActionsClient creates a new instance of WorkflowRunActionsClient with the specified values.
func NewWorkflowRunActionsClient(con *arm.Connection, subscriptionID string) *WorkflowRunActionsClient {
	return &WorkflowRunActionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a workflow run action.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunActionsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsGetOptions) (WorkflowRunActionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
	if err != nil {
		return WorkflowRunActionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunActionsClient) getHandleResponse(resp *http.Response) (WorkflowRunActionsGetResponse, error) {
	result := WorkflowRunActionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunAction); err != nil {
		return WorkflowRunActionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkflowRunActionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of workflow run actions.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunActionsClient) List(resourceGroupName string, workflowName string, runName string, options *WorkflowRunActionsListOptions) *WorkflowRunActionsListPager {
	return &WorkflowRunActionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowRunActionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkflowRunActionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunActionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunActionsClient) listHandleResponse(resp *http.Response) (WorkflowRunActionsListResponse, error) {
	result := WorkflowRunActionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionListResult); err != nil {
		return WorkflowRunActionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkflowRunActionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListExpressionTraces - Lists a workflow run expression trace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunActionsClient) ListExpressionTraces(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsListExpressionTracesOptions) (WorkflowRunActionsListExpressionTracesResponse, error) {
	req, err := client.listExpressionTracesCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
	if err != nil {
		return WorkflowRunActionsListExpressionTracesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionsListExpressionTracesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionsListExpressionTracesResponse{}, client.listExpressionTracesHandleError(resp)
	}
	return client.listExpressionTracesHandleResponse(resp)
}

// listExpressionTracesCreateRequest creates the ListExpressionTraces request.
func (client *WorkflowRunActionsClient) listExpressionTracesCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsListExpressionTracesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/listExpressionTraces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listExpressionTracesHandleResponse handles the ListExpressionTraces response.
func (client *WorkflowRunActionsClient) listExpressionTracesHandleResponse(resp *http.Response) (WorkflowRunActionsListExpressionTracesResponse, error) {
	result := WorkflowRunActionsListExpressionTracesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressionTraces); err != nil {
		return WorkflowRunActionsListExpressionTracesResponse{}, err
	}
	return result, nil
}

// listExpressionTracesHandleError handles the ListExpressionTraces error response.
func (client *WorkflowRunActionsClient) listExpressionTracesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}