//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

const (
	module  = "armdeploymentmanager"
	version = "v0.1.0"
)

// DeploymentMode - Describes the type of ARM deployment to be performed on the resource.
type DeploymentMode string

const (
	DeploymentModeIncremental DeploymentMode = "Incremental"
	DeploymentModeComplete    DeploymentMode = "Complete"
)

// PossibleDeploymentModeValues returns the possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{
		DeploymentModeIncremental,
		DeploymentModeComplete,
	}
}

// ToPtr returns a *DeploymentMode pointing to the current value.
func (c DeploymentMode) ToPtr() *DeploymentMode {
	return &c
}

// RestAuthLocation - The location of the authentication key/value pair in the request.
type RestAuthLocation string

const (
	RestAuthLocationQuery  RestAuthLocation = "Query"
	RestAuthLocationHeader RestAuthLocation = "Header"
)

// PossibleRestAuthLocationValues returns the possible values for the RestAuthLocation const type.
func PossibleRestAuthLocationValues() []RestAuthLocation {
	return []RestAuthLocation{
		RestAuthLocationQuery,
		RestAuthLocationHeader,
	}
}

// ToPtr returns a *RestAuthLocation pointing to the current value.
func (c RestAuthLocation) ToPtr() *RestAuthLocation {
	return &c
}

// RestAuthType - The authentication type.
type RestAuthType string

const (
	RestAuthTypeAPIKey          RestAuthType = "ApiKey"
	RestAuthTypeRolloutIdentity RestAuthType = "RolloutIdentity"
)

// PossibleRestAuthTypeValues returns the possible values for the RestAuthType const type.
func PossibleRestAuthTypeValues() []RestAuthType {
	return []RestAuthType{
		RestAuthTypeAPIKey,
		RestAuthTypeRolloutIdentity,
	}
}

// ToPtr returns a *RestAuthType pointing to the current value.
func (c RestAuthType) ToPtr() *RestAuthType {
	return &c
}

// RestMatchQuantifier - Indicates whether any or all of the expressions should match with the response content.
type RestMatchQuantifier string

const (
	RestMatchQuantifierAll RestMatchQuantifier = "All"
	RestMatchQuantifierAny RestMatchQuantifier = "Any"
)

// PossibleRestMatchQuantifierValues returns the possible values for the RestMatchQuantifier const type.
func PossibleRestMatchQuantifierValues() []RestMatchQuantifier {
	return []RestMatchQuantifier{
		RestMatchQuantifierAll,
		RestMatchQuantifierAny,
	}
}

// ToPtr returns a *RestMatchQuantifier pointing to the current value.
func (c RestMatchQuantifier) ToPtr() *RestMatchQuantifier {
	return &c
}

// RestRequestMethod - The HTTP method to use for the request.
type RestRequestMethod string

const (
	RestRequestMethodGET  RestRequestMethod = "GET"
	RestRequestMethodPOST RestRequestMethod = "POST"
)

// PossibleRestRequestMethodValues returns the possible values for the RestRequestMethod const type.
func PossibleRestRequestMethodValues() []RestRequestMethod {
	return []RestRequestMethod{
		RestRequestMethodGET,
		RestRequestMethodPOST,
	}
}

// ToPtr returns a *RestRequestMethod pointing to the current value.
func (c RestRequestMethod) ToPtr() *RestRequestMethod {
	return &c
}

// StepType - The type of step.
type StepType string

const (
	StepTypeWait        StepType = "Wait"
	StepTypeHealthCheck StepType = "HealthCheck"
)

// PossibleStepTypeValues returns the possible values for the StepType const type.
func PossibleStepTypeValues() []StepType {
	return []StepType{
		StepTypeWait,
		StepTypeHealthCheck,
	}
}

// ToPtr returns a *StepType pointing to the current value.
func (c StepType) ToPtr() *StepType {
	return &c
}