//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

const (
	module  = "armpowerbiprivatelinks"
	version = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// ActionsRequired - ActionsRequired
type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

// PossibleActionsRequiredValues returns the possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{
		ActionsRequiredNone,
		ActionsRequiredRecreate,
	}
}

// ToPtr returns a *ActionsRequired pointing to the current value.
func (c ActionsRequired) ToPtr() *ActionsRequired {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// PersistedConnectionStatus - Status of the connection.
type PersistedConnectionStatus string

const (
	PersistedConnectionStatusApproved     PersistedConnectionStatus = "Approved"
	PersistedConnectionStatusDisconnected PersistedConnectionStatus = "Disconnected"
	PersistedConnectionStatusPending      PersistedConnectionStatus = "Pending"
	PersistedConnectionStatusRejected     PersistedConnectionStatus = "Rejected"
)

// PossiblePersistedConnectionStatusValues returns the possible values for the PersistedConnectionStatus const type.
func PossiblePersistedConnectionStatusValues() []PersistedConnectionStatus {
	return []PersistedConnectionStatus{
		PersistedConnectionStatusApproved,
		PersistedConnectionStatusDisconnected,
		PersistedConnectionStatusPending,
		PersistedConnectionStatusRejected,
	}
}

// ToPtr returns a *PersistedConnectionStatus pointing to the current value.
func (c PersistedConnectionStatus) ToPtr() *PersistedConnectionStatus {
	return &c
}

// ResourceProvisioningState - Provisioning state of the Private Endpoint Connection.
type ResourceProvisioningState string

const (
	ResourceProvisioningStateCanceled  ResourceProvisioningState = "Canceled"
	ResourceProvisioningStateCreating  ResourceProvisioningState = "Creating"
	ResourceProvisioningStateDeleting  ResourceProvisioningState = "Deleting"
	ResourceProvisioningStateFailed    ResourceProvisioningState = "Failed"
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
	ResourceProvisioningStateUpdating  ResourceProvisioningState = "Updating"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateCreating,
		ResourceProvisioningStateDeleting,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
		ResourceProvisioningStateUpdating,
	}
}

// ToPtr returns a *ResourceProvisioningState pointing to the current value.
func (c ResourceProvisioningState) ToPtr() *ResourceProvisioningState {
	return &c
}