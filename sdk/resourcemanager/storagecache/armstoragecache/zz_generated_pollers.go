//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// CachesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type CachesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesCreateOrUpdateResponse will be returned.
func (p *CachesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (CachesCreateOrUpdateResponse, error) {
	respType := CachesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Cache)
	if err != nil {
		return CachesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesDebugInfoPoller provides polling facilities until the operation reaches a terminal state.
type CachesDebugInfoPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesDebugInfoPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesDebugInfoPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesDebugInfoResponse will be returned.
func (p *CachesDebugInfoPoller) FinalResponse(ctx context.Context) (CachesDebugInfoResponse, error) {
	respType := CachesDebugInfoResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesDebugInfoResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesDebugInfoPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesDeletePoller provides polling facilities until the operation reaches a terminal state.
type CachesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesDeleteResponse will be returned.
func (p *CachesDeletePoller) FinalResponse(ctx context.Context) (CachesDeleteResponse, error) {
	respType := CachesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesFlushPoller provides polling facilities until the operation reaches a terminal state.
type CachesFlushPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesFlushPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesFlushPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesFlushResponse will be returned.
func (p *CachesFlushPoller) FinalResponse(ctx context.Context) (CachesFlushResponse, error) {
	respType := CachesFlushResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesFlushResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesFlushPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesStartPoller provides polling facilities until the operation reaches a terminal state.
type CachesStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesStartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesStartResponse will be returned.
func (p *CachesStartPoller) FinalResponse(ctx context.Context) (CachesStartResponse, error) {
	respType := CachesStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesStopPoller provides polling facilities until the operation reaches a terminal state.
type CachesStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesStopPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesStopResponse will be returned.
func (p *CachesStopPoller) FinalResponse(ctx context.Context) (CachesStopResponse, error) {
	respType := CachesStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// CachesUpgradeFirmwarePoller provides polling facilities until the operation reaches a terminal state.
type CachesUpgradeFirmwarePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CachesUpgradeFirmwarePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CachesUpgradeFirmwarePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CachesUpgradeFirmwareResponse will be returned.
func (p *CachesUpgradeFirmwarePoller) FinalResponse(ctx context.Context) (CachesUpgradeFirmwareResponse, error) {
	respType := CachesUpgradeFirmwareResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return CachesUpgradeFirmwareResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CachesUpgradeFirmwarePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetFlushPoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetFlushPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetFlushPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetFlushPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetFlushResponse will be returned.
func (p *StorageTargetFlushPoller) FinalResponse(ctx context.Context) (StorageTargetFlushResponse, error) {
	respType := StorageTargetFlushResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StorageTargetFlushResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetFlushPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetResumePoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetResumePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetResumePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetResumePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetResumeResponse will be returned.
func (p *StorageTargetResumePoller) FinalResponse(ctx context.Context) (StorageTargetResumeResponse, error) {
	respType := StorageTargetResumeResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StorageTargetResumeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetResumePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetSuspendPoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetSuspendPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetSuspendPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetSuspendPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetSuspendResponse will be returned.
func (p *StorageTargetSuspendPoller) FinalResponse(ctx context.Context) (StorageTargetSuspendResponse, error) {
	respType := StorageTargetSuspendResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StorageTargetSuspendResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetSuspendPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetsCreateOrUpdateResponse will be returned.
func (p *StorageTargetsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (StorageTargetsCreateOrUpdateResponse, error) {
	respType := StorageTargetsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.StorageTarget)
	if err != nil {
		return StorageTargetsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetsDNSRefreshPoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetsDNSRefreshPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetsDNSRefreshPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetsDNSRefreshPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetsDNSRefreshResponse will be returned.
func (p *StorageTargetsDNSRefreshPoller) FinalResponse(ctx context.Context) (StorageTargetsDNSRefreshResponse, error) {
	respType := StorageTargetsDNSRefreshResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StorageTargetsDNSRefreshResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetsDNSRefreshPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StorageTargetsDeletePoller provides polling facilities until the operation reaches a terminal state.
type StorageTargetsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StorageTargetsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StorageTargetsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StorageTargetsDeleteResponse will be returned.
func (p *StorageTargetsDeletePoller) FinalResponse(ctx context.Context) (StorageTargetsDeleteResponse, error) {
	respType := StorageTargetsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StorageTargetsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StorageTargetsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}