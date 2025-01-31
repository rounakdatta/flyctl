// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// CreateAppCreateAppCreateAppPayload includes the requested fields of the GraphQL type CreateAppPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of CreateApp
type CreateAppCreateAppCreateAppPayload struct {
	App CreateAppCreateAppCreateAppPayloadApp `json:"app"`
}

// GetApp returns CreateAppCreateAppCreateAppPayload.App, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayload) GetApp() CreateAppCreateAppCreateAppPayloadApp {
	return v.App
}

// CreateAppCreateAppCreateAppPayloadApp includes the requested fields of the GraphQL type App.
type CreateAppCreateAppCreateAppPayloadApp struct {
	// The unique application name
	Name string `json:"name"`
	// Application status
	Status string `json:"status"`
	// Organization that owns this app
	Organization CreateAppCreateAppCreateAppPayloadAppOrganization `json:"organization"`
}

// GetName returns CreateAppCreateAppCreateAppPayloadApp.Name, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayloadApp) GetName() string { return v.Name }

// GetStatus returns CreateAppCreateAppCreateAppPayloadApp.Status, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayloadApp) GetStatus() string { return v.Status }

// GetOrganization returns CreateAppCreateAppCreateAppPayloadApp.Organization, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayloadApp) GetOrganization() CreateAppCreateAppCreateAppPayloadAppOrganization {
	return v.Organization
}

// CreateAppCreateAppCreateAppPayloadAppOrganization includes the requested fields of the GraphQL type Organization.
type CreateAppCreateAppCreateAppPayloadAppOrganization struct {
	Id string `json:"id"`
	// Unique organization slug
	Slug string `json:"slug"`
}

// GetId returns CreateAppCreateAppCreateAppPayloadAppOrganization.Id, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayloadAppOrganization) GetId() string { return v.Id }

// GetSlug returns CreateAppCreateAppCreateAppPayloadAppOrganization.Slug, and is useful for accessing the field via an interface.
func (v *CreateAppCreateAppCreateAppPayloadAppOrganization) GetSlug() string { return v.Slug }

// CreateAppResponse is returned by CreateApp on success.
type CreateAppResponse struct {
	CreateApp CreateAppCreateAppCreateAppPayload `json:"createApp"`
}

// GetCreateApp returns CreateAppResponse.CreateApp, and is useful for accessing the field via an interface.
func (v *CreateAppResponse) GetCreateApp() CreateAppCreateAppCreateAppPayload { return v.CreateApp }

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload includes the requested fields of the GraphQL type EnsureMachineRemoteBuilderPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of EnsureMachineRemoteBuilder
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload struct {
	Machine EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine `json:"machine"`
	App     EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp     `json:"app"`
}

// GetMachine returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload.Machine, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload) GetMachine() EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine {
	return v.Machine
}

// GetApp returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload.App, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload) GetApp() EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp {
	return v.App
}

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp includes the requested fields of the GraphQL type App.
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp struct {
	// The unique application name
	Name string `json:"name"`
	// Organization that owns this app
	Organization EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization `json:"organization"`
}

// GetName returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp.Name, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp) GetName() string {
	return v.Name
}

// GetOrganization returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp.Organization, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadApp) GetOrganization() EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization {
	return v.Organization
}

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization includes the requested fields of the GraphQL type Organization.
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization struct {
	Id string `json:"id"`
	// Unique organization slug
	Slug string `json:"slug"`
}

// GetId returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization.Id, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization) GetId() string {
	return v.Id
}

// GetSlug returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization.Slug, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadAppOrganization) GetSlug() string {
	return v.Slug
}

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine includes the requested fields of the GraphQL type Machine.
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine struct {
	Id    string                                                                                                      `json:"id"`
	State string                                                                                                      `json:"state"`
	Ips   EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection `json:"ips"`
}

// GetId returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine.Id, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine) GetId() string {
	return v.Id
}

// GetState returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine.State, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine) GetState() string {
	return v.State
}

// GetIps returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine.Ips, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachine) GetIps() EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection {
	return v.Ips
}

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection includes the requested fields of the GraphQL type MachineIPConnection.
// The GraphQL type's documentation follows.
//
// The connection type for MachineIP.
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection struct {
	// A list of nodes.
	Nodes []EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP `json:"nodes"`
}

// GetNodes returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection.Nodes, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnection) GetNodes() []EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP {
	return v.Nodes
}

// EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP includes the requested fields of the GraphQL type MachineIP.
type EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP struct {
	Family string `json:"family"`
	Kind   string `json:"kind"`
	Ip     string `json:"ip"`
}

// GetFamily returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP.Family, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP) GetFamily() string {
	return v.Family
}

// GetKind returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP.Kind, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP) GetKind() string {
	return v.Kind
}

// GetIp returns EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP.Ip, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayloadMachineIpsMachineIPConnectionNodesMachineIP) GetIp() string {
	return v.Ip
}

// EnsureRemoteBuilderResponse is returned by EnsureRemoteBuilder on success.
type EnsureRemoteBuilderResponse struct {
	EnsureMachineRemoteBuilder EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload `json:"ensureMachineRemoteBuilder"`
}

// GetEnsureMachineRemoteBuilder returns EnsureRemoteBuilderResponse.EnsureMachineRemoteBuilder, and is useful for accessing the field via an interface.
func (v *EnsureRemoteBuilderResponse) GetEnsureMachineRemoteBuilder() EnsureRemoteBuilderEnsureMachineRemoteBuilderEnsureMachineRemoteBuilderPayload {
	return v.EnsureMachineRemoteBuilder
}

// GetAppApp includes the requested fields of the GraphQL type App.
type GetAppApp struct {
	// The unique application name
	Name string `json:"name"`
	// Fly platform version
	PlatformVersion PlatformVersionEnum `json:"platformVersion"`
	// Autogenerated hostname for this application
	Hostname string `json:"hostname"`
	// Organization that owns this app
	Organization GetAppAppOrganization `json:"organization"`
}

// GetName returns GetAppApp.Name, and is useful for accessing the field via an interface.
func (v *GetAppApp) GetName() string { return v.Name }

// GetPlatformVersion returns GetAppApp.PlatformVersion, and is useful for accessing the field via an interface.
func (v *GetAppApp) GetPlatformVersion() PlatformVersionEnum { return v.PlatformVersion }

// GetHostname returns GetAppApp.Hostname, and is useful for accessing the field via an interface.
func (v *GetAppApp) GetHostname() string { return v.Hostname }

// GetOrganization returns GetAppApp.Organization, and is useful for accessing the field via an interface.
func (v *GetAppApp) GetOrganization() GetAppAppOrganization { return v.Organization }

// GetAppAppOrganization includes the requested fields of the GraphQL type Organization.
type GetAppAppOrganization struct {
	Id string `json:"id"`
	// Unique organization slug
	Slug string `json:"slug"`
}

// GetId returns GetAppAppOrganization.Id, and is useful for accessing the field via an interface.
func (v *GetAppAppOrganization) GetId() string { return v.Id }

// GetSlug returns GetAppAppOrganization.Slug, and is useful for accessing the field via an interface.
func (v *GetAppAppOrganization) GetSlug() string { return v.Slug }

// GetAppResponse is returned by GetApp on success.
type GetAppResponse struct {
	// Find an app by name
	App GetAppApp `json:"app"`
}

// GetApp returns GetAppResponse.App, and is useful for accessing the field via an interface.
func (v *GetAppResponse) GetApp() GetAppApp { return v.App }

// GetNearestRegionNearestRegion includes the requested fields of the GraphQL type Region.
type GetNearestRegionNearestRegion struct {
	// The IATA airport code for this region
	Code string `json:"code"`
	// The name of this region
	Name             string `json:"name"`
	GatewayAvailable bool   `json:"gatewayAvailable"`
}

// GetCode returns GetNearestRegionNearestRegion.Code, and is useful for accessing the field via an interface.
func (v *GetNearestRegionNearestRegion) GetCode() string { return v.Code }

// GetName returns GetNearestRegionNearestRegion.Name, and is useful for accessing the field via an interface.
func (v *GetNearestRegionNearestRegion) GetName() string { return v.Name }

// GetGatewayAvailable returns GetNearestRegionNearestRegion.GatewayAvailable, and is useful for accessing the field via an interface.
func (v *GetNearestRegionNearestRegion) GetGatewayAvailable() bool { return v.GatewayAvailable }

// GetNearestRegionResponse is returned by GetNearestRegion on success.
type GetNearestRegionResponse struct {
	NearestRegion GetNearestRegionNearestRegion `json:"nearestRegion"`
}

// GetNearestRegion returns GetNearestRegionResponse.NearestRegion, and is useful for accessing the field via an interface.
func (v *GetNearestRegionResponse) GetNearestRegion() GetNearestRegionNearestRegion {
	return v.NearestRegion
}

type PlatformVersionEnum string

const (
	// App with only machines
	PlatformVersionEnumMachines PlatformVersionEnum = "machines"
	// Nomad managed application
	PlatformVersionEnumNomad PlatformVersionEnum = "nomad"
)

// __CreateAppInput is used internally by genqlient
type __CreateAppInput struct {
	Name           string `json:"name"`
	OrganizationId string `json:"organizationId"`
}

// GetName returns __CreateAppInput.Name, and is useful for accessing the field via an interface.
func (v *__CreateAppInput) GetName() string { return v.Name }

// GetOrganizationId returns __CreateAppInput.OrganizationId, and is useful for accessing the field via an interface.
func (v *__CreateAppInput) GetOrganizationId() string { return v.OrganizationId }

// __EnsureRemoteBuilderInput is used internally by genqlient
type __EnsureRemoteBuilderInput struct {
	OrganizationId string `json:"organizationId"`
}

// GetOrganizationId returns __EnsureRemoteBuilderInput.OrganizationId, and is useful for accessing the field via an interface.
func (v *__EnsureRemoteBuilderInput) GetOrganizationId() string { return v.OrganizationId }

// __GetAppInput is used internally by genqlient
type __GetAppInput struct {
	AppName string `json:"appName"`
}

// GetAppName returns __GetAppInput.AppName, and is useful for accessing the field via an interface.
func (v *__GetAppInput) GetAppName() string { return v.AppName }

func CreateApp(
	ctx context.Context,
	client graphql.Client,
	name string,
	organizationId string,
) (*CreateAppResponse, error) {
	req := &graphql.Request{
		OpName: "CreateApp",
		Query: `
mutation CreateApp ($name: String, $organizationId: ID!) {
	createApp(input: {name:$name,organizationId:$organizationId}) {
		app {
			name
			status
			organization {
				id
				slug
			}
		}
	}
}
`,
		Variables: &__CreateAppInput{
			Name:           name,
			OrganizationId: organizationId,
		},
	}
	var err error

	var data CreateAppResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func EnsureRemoteBuilder(
	ctx context.Context,
	client graphql.Client,
	organizationId string,
) (*EnsureRemoteBuilderResponse, error) {
	req := &graphql.Request{
		OpName: "EnsureRemoteBuilder",
		Query: `
mutation EnsureRemoteBuilder ($organizationId: ID!) {
	ensureMachineRemoteBuilder(input: {organizationId:$organizationId}) {
		machine {
			id
			state
			ips {
				nodes {
					family
					kind
					ip
				}
			}
		}
		app {
			name
			organization {
				id
				slug
			}
		}
	}
}
`,
		Variables: &__EnsureRemoteBuilderInput{
			OrganizationId: organizationId,
		},
	}
	var err error

	var data EnsureRemoteBuilderResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetApp(
	ctx context.Context,
	client graphql.Client,
	appName string,
) (*GetAppResponse, error) {
	req := &graphql.Request{
		OpName: "GetApp",
		Query: `
query GetApp ($appName: String!) {
	app(name: $appName) {
		name
		platformVersion
		hostname
		organization {
			id
			slug
		}
	}
}
`,
		Variables: &__GetAppInput{
			AppName: appName,
		},
	}
	var err error

	var data GetAppResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetNearestRegion(
	ctx context.Context,
	client graphql.Client,
) (*GetNearestRegionResponse, error) {
	req := &graphql.Request{
		OpName: "GetNearestRegion",
		Query: `
query GetNearestRegion {
	nearestRegion {
		code
		name
		gatewayAvailable
	}
}
`,
	}
	var err error

	var data GetNearestRegionResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
