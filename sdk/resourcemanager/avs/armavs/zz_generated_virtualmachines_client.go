//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachinesClient contains the methods for the VirtualMachines group.
// Don't use this type directly, use NewVirtualMachinesClient() instead.
type VirtualMachinesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient with the specified values.
func NewVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachinesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VirtualMachinesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get a virtual machine by id in a private cloud cluster
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, options *VirtualMachinesGetOptions) (VirtualMachinesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, options)
	if err != nil {
		return VirtualMachinesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachinesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachinesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, options *VirtualMachinesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines/{virtualMachineId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if virtualMachineID == "" {
		return nil, errors.New("parameter virtualMachineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineId}", url.PathEscape(virtualMachineID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachinesClient) getHandleResponse(resp *http.Response) (VirtualMachinesGetResponse, error) {
	result := VirtualMachinesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachine); err != nil {
		return VirtualMachinesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachinesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List of virtual machines in a private cloud cluster
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachinesClient) List(resourceGroupName string, privateCloudName string, clusterName string, options *VirtualMachinesListOptions) *VirtualMachinesListPager {
	return &VirtualMachinesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachinesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachinesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *VirtualMachinesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachinesClient) listHandleResponse(resp *http.Response) (VirtualMachinesListResponse, error) {
	result := VirtualMachinesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachinesList); err != nil {
		return VirtualMachinesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachinesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRestrictMovement - Enable or disable DRS-driven VM movement restriction
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachinesClient) BeginRestrictMovement(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesBeginRestrictMovementOptions) (VirtualMachinesRestrictMovementPollerResponse, error) {
	resp, err := client.restrictMovement(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, restrictMovement, options)
	if err != nil {
		return VirtualMachinesRestrictMovementPollerResponse{}, err
	}
	result := VirtualMachinesRestrictMovementPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachinesClient.RestrictMovement", "", resp, client.pl, client.restrictMovementHandleError)
	if err != nil {
		return VirtualMachinesRestrictMovementPollerResponse{}, err
	}
	result.Poller = &VirtualMachinesRestrictMovementPoller{
		pt: pt,
	}
	return result, nil
}

// RestrictMovement - Enable or disable DRS-driven VM movement restriction
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachinesClient) restrictMovement(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesBeginRestrictMovementOptions) (*http.Response, error) {
	req, err := client.restrictMovementCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, restrictMovement, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.restrictMovementHandleError(resp)
	}
	return resp, nil
}

// restrictMovementCreateRequest creates the RestrictMovement request.
func (client *VirtualMachinesClient) restrictMovementCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesBeginRestrictMovementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines/{virtualMachineId}/restrictMovement"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if virtualMachineID == "" {
		return nil, errors.New("parameter virtualMachineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineId}", url.PathEscape(virtualMachineID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, restrictMovement)
}

// restrictMovementHandleError handles the RestrictMovement error response.
func (client *VirtualMachinesClient) restrictMovementHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}