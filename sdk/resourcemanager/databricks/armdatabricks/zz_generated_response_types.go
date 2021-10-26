//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// OutboundNetworkDependenciesEndpointsListResponse contains the response from method OutboundNetworkDependenciesEndpoints.List.
type OutboundNetworkDependenciesEndpointsListResponse struct {
	OutboundNetworkDependenciesEndpointsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutboundNetworkDependenciesEndpointsListResult contains the result from method OutboundNetworkDependenciesEndpoints.List.
type OutboundNetworkDependenciesEndpointsListResult struct {
	// Collection of outbound network dependency endpoints
	OutboundEnvironmentEndpointArray []*OutboundEnvironmentEndpoint
}

// PrivateEndpointConnectionsCreatePollerResponse contains the response from method PrivateEndpointConnections.Create.
type PrivateEndpointConnectionsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointConnectionsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsCreateResponse, error) {
	respType := PrivateEndpointConnectionsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsCreatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsCreatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsCreateResponse contains the response from method PrivateEndpointConnections.Create.
type PrivateEndpointConnectionsCreateResponse struct {
	PrivateEndpointConnectionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateResult contains the result from method PrivateEndpointConnections.Create.
type PrivateEndpointConnectionsCreateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListResponse contains the response from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResponse struct {
	PrivateEndpointConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListResult contains the result from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResult struct {
	PrivateEndpointConnectionsList
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	GroupIDInformation
}

// PrivateLinkResourcesListResponse contains the response from method PrivateLinkResources.List.
type PrivateLinkResourcesListResponse struct {
	PrivateLinkResourcesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListResult contains the result from method PrivateLinkResources.List.
type PrivateLinkResourcesListResult struct {
	PrivateLinkResourcesList
}

// VNetPeeringCreateOrUpdatePollerResponse contains the response from method VNetPeering.CreateOrUpdate.
type VNetPeeringCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VNetPeeringCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l VNetPeeringCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VNetPeeringCreateOrUpdateResponse, error) {
	respType := VNetPeeringCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.VirtualNetworkPeering)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VNetPeeringCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *VNetPeeringCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *VNetPeeringClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VNetPeeringClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &VNetPeeringCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VNetPeeringCreateOrUpdateResponse contains the response from method VNetPeering.CreateOrUpdate.
type VNetPeeringCreateOrUpdateResponse struct {
	VNetPeeringCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VNetPeeringCreateOrUpdateResult contains the result from method VNetPeering.CreateOrUpdate.
type VNetPeeringCreateOrUpdateResult struct {
	VirtualNetworkPeering
}

// VNetPeeringDeletePollerResponse contains the response from method VNetPeering.Delete.
type VNetPeeringDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VNetPeeringDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l VNetPeeringDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VNetPeeringDeleteResponse, error) {
	respType := VNetPeeringDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VNetPeeringDeletePollerResponse from the provided client and resume token.
func (l *VNetPeeringDeletePollerResponse) Resume(ctx context.Context, client *VNetPeeringClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VNetPeeringClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &VNetPeeringDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VNetPeeringDeleteResponse contains the response from method VNetPeering.Delete.
type VNetPeeringDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VNetPeeringGetResponse contains the response from method VNetPeering.Get.
type VNetPeeringGetResponse struct {
	VNetPeeringGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VNetPeeringGetResult contains the result from method VNetPeering.Get.
type VNetPeeringGetResult struct {
	VirtualNetworkPeering
}

// VNetPeeringListByWorkspaceResponse contains the response from method VNetPeering.ListByWorkspace.
type VNetPeeringListByWorkspaceResponse struct {
	VNetPeeringListByWorkspaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VNetPeeringListByWorkspaceResult contains the result from method VNetPeering.ListByWorkspace.
type VNetPeeringListByWorkspaceResult struct {
	VirtualNetworkPeeringList
}

// WorkspacesCreateOrUpdatePollerResponse contains the response from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l WorkspacesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesCreateOrUpdateResponse, error) {
	respType := WorkspacesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Workspace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *WorkspacesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesCreateOrUpdateResponse contains the response from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdateResponse struct {
	WorkspacesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesCreateOrUpdateResult contains the result from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdateResult struct {
	Workspace
}

// WorkspacesDeletePollerResponse contains the response from method Workspaces.Delete.
type WorkspacesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l WorkspacesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesDeleteResponse, error) {
	respType := WorkspacesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesDeletePollerResponse from the provided client and resume token.
func (l *WorkspacesDeletePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesDeleteResponse contains the response from method Workspaces.Delete.
type WorkspacesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesGetResponse contains the response from method Workspaces.Get.
type WorkspacesGetResponse struct {
	WorkspacesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesGetResult contains the result from method Workspaces.Get.
type WorkspacesGetResult struct {
	Workspace
}

// WorkspacesListByResourceGroupResponse contains the response from method Workspaces.ListByResourceGroup.
type WorkspacesListByResourceGroupResponse struct {
	WorkspacesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListByResourceGroupResult contains the result from method Workspaces.ListByResourceGroup.
type WorkspacesListByResourceGroupResult struct {
	WorkspaceListResult
}

// WorkspacesListBySubscriptionResponse contains the response from method Workspaces.ListBySubscription.
type WorkspacesListBySubscriptionResponse struct {
	WorkspacesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListBySubscriptionResult contains the result from method Workspaces.ListBySubscription.
type WorkspacesListBySubscriptionResult struct {
	WorkspaceListResult
}

// WorkspacesUpdatePollerResponse contains the response from method Workspaces.Update.
type WorkspacesUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l WorkspacesUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesUpdateResponse, error) {
	respType := WorkspacesUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Workspace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesUpdatePollerResponse from the provided client and resume token.
func (l *WorkspacesUpdatePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesUpdateResponse contains the response from method Workspaces.Update.
type WorkspacesUpdateResponse struct {
	WorkspacesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesUpdateResult contains the result from method Workspaces.Update.
type WorkspacesUpdateResult struct {
	Workspace
}