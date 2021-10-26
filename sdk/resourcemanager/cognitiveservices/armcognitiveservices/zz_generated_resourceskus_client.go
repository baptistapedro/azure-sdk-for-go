//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
func NewResourceSKUsClient(con *arm.Connection, subscriptionID string) *ResourceSKUsClient {
	return &ResourceSKUsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Gets the list of Microsoft.CognitiveServices SKUs available for your Subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ResourceSKUsClient) List(options *ResourceSKUsListOptions) *ResourceSKUsListPager {
	return &ResourceSKUsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ResourceSKUsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceSKUListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ResourceSKUsClient) listCreateRequest(ctx context.Context, options *ResourceSKUsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceSKUsClient) listHandleResponse(resp *http.Response) (ResourceSKUsListResponse, error) {
	result := ResourceSKUsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUListResult); err != nil {
		return ResourceSKUsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ResourceSKUsClient) listHandleError(resp *http.Response) error {
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