//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StudentLabsClient contains the methods for the StudentLabs group.
// Don't use this type directly, use NewStudentLabsClient() instead.
type StudentLabsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewStudentLabsClient creates a new instance of StudentLabsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStudentLabsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*StudentLabsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &StudentLabsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get the details for a specified lab associated with the student lab.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// studentLabName - Student lab name.
// options - StudentLabsClientGetOptions contains the optional parameters for the StudentLabsClient.Get method.
func (client *StudentLabsClient) Get(ctx context.Context, studentLabName string, options *StudentLabsClientGetOptions) (StudentLabsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, studentLabName, options)
	if err != nil {
		return StudentLabsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StudentLabsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StudentLabsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StudentLabsClient) getCreateRequest(ctx context.Context, studentLabName string, options *StudentLabsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Education/studentLabs/{studentLabName}"
	if studentLabName == "" {
		return nil, errors.New("parameter studentLabName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{studentLabName}", url.PathEscape(studentLabName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StudentLabsClient) getHandleResponse(resp *http.Response) (StudentLabsClientGetResponse, error) {
	result := StudentLabsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StudentLabDetails); err != nil {
		return StudentLabsClientGetResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Get a list of all labs associated with the caller of the API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// options - StudentLabsClientListAllOptions contains the optional parameters for the StudentLabsClient.ListAll method.
func (client *StudentLabsClient) NewListAllPager(options *StudentLabsClientListAllOptions) *runtime.Pager[StudentLabsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[StudentLabsClientListAllResponse]{
		More: func(page StudentLabsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StudentLabsClientListAllResponse) (StudentLabsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StudentLabsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StudentLabsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StudentLabsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *StudentLabsClient) listAllCreateRequest(ctx context.Context, options *StudentLabsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Education/studentLabs"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *StudentLabsClient) listAllHandleResponse(resp *http.Response) (StudentLabsClientListAllResponse, error) {
	result := StudentLabsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StudentLabListResult); err != nil {
		return StudentLabsClientListAllResponse{}, err
	}
	return result, nil
}
