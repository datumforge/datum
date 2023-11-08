// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package datumclient

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/datumforge/datum/internal/ent/generated"
)

type DatumClient interface {
	GetOrganizationByID(ctx context.Context, organizationID string, interceptors ...clientv2.RequestInterceptor) (*GetOrganizationByID, error)
	GetAllOrganizations(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllOrganizations, error)
	CreateOrganization(ctx context.Context, input CreateOrganizationInput, interceptors ...clientv2.RequestInterceptor) (*CreateOrganization, error)
	UpdateOrganization(ctx context.Context, updateOrganizationID string, input UpdateOrganizationInput, interceptors ...clientv2.RequestInterceptor) (*UpdateOrganization, error)
	DeleteOrganization(ctx context.Context, deleteOrganizationID string, interceptors ...clientv2.RequestInterceptor) (*DeleteOrganization, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) DatumClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type Query struct {
	Node               generated.Noder         "json:\"node,omitempty\" graphql:\"node\""
	Nodes              []generated.Noder       "json:\"nodes\" graphql:\"nodes\""
	Groups             GroupConnection         "json:\"groups\" graphql:\"groups\""
	GroupSettingsSlice GroupSettingsConnection "json:\"groupSettingsSlice\" graphql:\"groupSettingsSlice\""
	Integrations       IntegrationConnection   "json:\"integrations\" graphql:\"integrations\""
	Organizations      OrganizationConnection  "json:\"organizations\" graphql:\"organizations\""
	Sessions           SessionConnection       "json:\"sessions\" graphql:\"sessions\""
	Users              UserConnection          "json:\"users\" graphql:\"users\""
	Group              Group                   "json:\"group\" graphql:\"group\""
	Integration        Integration             "json:\"integration\" graphql:\"integration\""
	Organization       Organization            "json:\"organization\" graphql:\"organization\""
	Session            Session                 "json:\"session\" graphql:\"session\""
	User               User                    "json:\"user\" graphql:\"user\""
	Service            Service                 "json:\"_service\" graphql:\"_service\""
}
type Mutation struct {
	CreateGroup        GroupCreatePayload        "json:\"createGroup\" graphql:\"createGroup\""
	UpdateGroup        GroupUpdatePayload        "json:\"updateGroup\" graphql:\"updateGroup\""
	DeleteGroup        GroupDeletePayload        "json:\"deleteGroup\" graphql:\"deleteGroup\""
	CreateIntegration  IntegrationCreatePayload  "json:\"createIntegration\" graphql:\"createIntegration\""
	UpdateIntegration  IntegrationUpdatePayload  "json:\"updateIntegration\" graphql:\"updateIntegration\""
	DeleteIntegration  IntegrationDeletePayload  "json:\"deleteIntegration\" graphql:\"deleteIntegration\""
	CreateOrganization OrganizationCreatePayload "json:\"createOrganization\" graphql:\"createOrganization\""
	UpdateOrganization OrganizationUpdatePayload "json:\"updateOrganization\" graphql:\"updateOrganization\""
	DeleteOrganization OrganizationDeletePayload "json:\"deleteOrganization\" graphql:\"deleteOrganization\""
	CreateSession      SessionCreatePayload      "json:\"createSession\" graphql:\"createSession\""
	UpdateSession      SessionUpdatePayload      "json:\"updateSession\" graphql:\"updateSession\""
	DeleteSession      SessionDeletePayload      "json:\"deleteSession\" graphql:\"deleteSession\""
	CreateUser         UserCreatePayload         "json:\"createUser\" graphql:\"createUser\""
	UpdateUser         UserUpdatePayload         "json:\"updateUser\" graphql:\"updateUser\""
	DeleteUser         UserDeletePayload         "json:\"deleteUser\" graphql:\"deleteUser\""
}
type GetOrganizationByID_Organization_Parent struct {
	ID   string "json:\"id\" graphql:\"id\""
	Name string "json:\"name\" graphql:\"name\""
}

func (t *GetOrganizationByID_Organization_Parent) GetID() string {
	if t == nil {
		t = &GetOrganizationByID_Organization_Parent{}
	}
	return t.ID
}
func (t *GetOrganizationByID_Organization_Parent) GetName() string {
	if t == nil {
		t = &GetOrganizationByID_Organization_Parent{}
	}
	return t.Name
}

type GetOrganizationByID_Organization struct {
	ID          string                                   "json:\"id\" graphql:\"id\""
	Name        string                                   "json:\"name\" graphql:\"name\""
	Parent      *GetOrganizationByID_Organization_Parent "json:\"parent,omitempty\" graphql:\"parent\""
	Description *string                                  "json:\"description,omitempty\" graphql:\"description\""
}

func (t *GetOrganizationByID_Organization) GetID() string {
	if t == nil {
		t = &GetOrganizationByID_Organization{}
	}
	return t.ID
}
func (t *GetOrganizationByID_Organization) GetName() string {
	if t == nil {
		t = &GetOrganizationByID_Organization{}
	}
	return t.Name
}
func (t *GetOrganizationByID_Organization) GetParent() *GetOrganizationByID_Organization_Parent {
	if t == nil {
		t = &GetOrganizationByID_Organization{}
	}
	return t.Parent
}
func (t *GetOrganizationByID_Organization) GetDescription() *string {
	if t == nil {
		t = &GetOrganizationByID_Organization{}
	}
	return t.Description
}

type GetAllOrganizations_Organizations_Edges_Node struct {
	ID          string  "json:\"id\" graphql:\"id\""
	Name        string  "json:\"name\" graphql:\"name\""
	Description *string "json:\"description,omitempty\" graphql:\"description\""
}

func (t *GetAllOrganizations_Organizations_Edges_Node) GetID() string {
	if t == nil {
		t = &GetAllOrganizations_Organizations_Edges_Node{}
	}
	return t.ID
}
func (t *GetAllOrganizations_Organizations_Edges_Node) GetName() string {
	if t == nil {
		t = &GetAllOrganizations_Organizations_Edges_Node{}
	}
	return t.Name
}
func (t *GetAllOrganizations_Organizations_Edges_Node) GetDescription() *string {
	if t == nil {
		t = &GetAllOrganizations_Organizations_Edges_Node{}
	}
	return t.Description
}

type GetAllOrganizations_Organizations_Edges struct {
	Node *GetAllOrganizations_Organizations_Edges_Node "json:\"node,omitempty\" graphql:\"node\""
}

func (t *GetAllOrganizations_Organizations_Edges) GetNode() *GetAllOrganizations_Organizations_Edges_Node {
	if t == nil {
		t = &GetAllOrganizations_Organizations_Edges{}
	}
	return t.Node
}

type GetAllOrganizations_Organizations struct {
	Edges []*GetAllOrganizations_Organizations_Edges "json:\"edges,omitempty\" graphql:\"edges\""
}

func (t *GetAllOrganizations_Organizations) GetEdges() []*GetAllOrganizations_Organizations_Edges {
	if t == nil {
		t = &GetAllOrganizations_Organizations{}
	}
	return t.Edges
}

type CreateOrganization_CreateOrganization_Organization_Parent struct {
	ID   string "json:\"id\" graphql:\"id\""
	Name string "json:\"name\" graphql:\"name\""
}

func (t *CreateOrganization_CreateOrganization_Organization_Parent) GetID() string {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization_Parent{}
	}
	return t.ID
}
func (t *CreateOrganization_CreateOrganization_Organization_Parent) GetName() string {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization_Parent{}
	}
	return t.Name
}

type CreateOrganization_CreateOrganization_Organization struct {
	ID          string                                                     "json:\"id\" graphql:\"id\""
	Name        string                                                     "json:\"name\" graphql:\"name\""
	Description *string                                                    "json:\"description,omitempty\" graphql:\"description\""
	Parent      *CreateOrganization_CreateOrganization_Organization_Parent "json:\"parent,omitempty\" graphql:\"parent\""
}

func (t *CreateOrganization_CreateOrganization_Organization) GetID() string {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization{}
	}
	return t.ID
}
func (t *CreateOrganization_CreateOrganization_Organization) GetName() string {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization{}
	}
	return t.Name
}
func (t *CreateOrganization_CreateOrganization_Organization) GetDescription() *string {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization{}
	}
	return t.Description
}
func (t *CreateOrganization_CreateOrganization_Organization) GetParent() *CreateOrganization_CreateOrganization_Organization_Parent {
	if t == nil {
		t = &CreateOrganization_CreateOrganization_Organization{}
	}
	return t.Parent
}

type CreateOrganization_CreateOrganization struct {
	Organization CreateOrganization_CreateOrganization_Organization "json:\"organization\" graphql:\"organization\""
}

func (t *CreateOrganization_CreateOrganization) GetOrganization() *CreateOrganization_CreateOrganization_Organization {
	if t == nil {
		t = &CreateOrganization_CreateOrganization{}
	}
	return &t.Organization
}

type UpdateOrganization_UpdateOrganization_Organization_Parent struct {
	ID   string "json:\"id\" graphql:\"id\""
	Name string "json:\"name\" graphql:\"name\""
}

func (t *UpdateOrganization_UpdateOrganization_Organization_Parent) GetID() string {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization_Parent{}
	}
	return t.ID
}
func (t *UpdateOrganization_UpdateOrganization_Organization_Parent) GetName() string {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization_Parent{}
	}
	return t.Name
}

type UpdateOrganization_UpdateOrganization_Organization struct {
	ID          string                                                     "json:\"id\" graphql:\"id\""
	Name        string                                                     "json:\"name\" graphql:\"name\""
	Parent      *UpdateOrganization_UpdateOrganization_Organization_Parent "json:\"parent,omitempty\" graphql:\"parent\""
	Description *string                                                    "json:\"description,omitempty\" graphql:\"description\""
}

func (t *UpdateOrganization_UpdateOrganization_Organization) GetID() string {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization{}
	}
	return t.ID
}
func (t *UpdateOrganization_UpdateOrganization_Organization) GetName() string {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization{}
	}
	return t.Name
}
func (t *UpdateOrganization_UpdateOrganization_Organization) GetParent() *UpdateOrganization_UpdateOrganization_Organization_Parent {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization{}
	}
	return t.Parent
}
func (t *UpdateOrganization_UpdateOrganization_Organization) GetDescription() *string {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization_Organization{}
	}
	return t.Description
}

type UpdateOrganization_UpdateOrganization struct {
	Organization UpdateOrganization_UpdateOrganization_Organization "json:\"organization\" graphql:\"organization\""
}

func (t *UpdateOrganization_UpdateOrganization) GetOrganization() *UpdateOrganization_UpdateOrganization_Organization {
	if t == nil {
		t = &UpdateOrganization_UpdateOrganization{}
	}
	return &t.Organization
}

type DeleteOrganization_DeleteOrganization struct {
	DeletedID string "json:\"deletedID\" graphql:\"deletedID\""
}

func (t *DeleteOrganization_DeleteOrganization) GetDeletedID() string {
	if t == nil {
		t = &DeleteOrganization_DeleteOrganization{}
	}
	return t.DeletedID
}

type GetOrganizationByID struct {
	Organization GetOrganizationByID_Organization "json:\"organization\" graphql:\"organization\""
}

func (t *GetOrganizationByID) GetOrganization() *GetOrganizationByID_Organization {
	if t == nil {
		t = &GetOrganizationByID{}
	}
	return &t.Organization
}

type GetAllOrganizations struct {
	Organizations GetAllOrganizations_Organizations "json:\"organizations\" graphql:\"organizations\""
}

func (t *GetAllOrganizations) GetOrganizations() *GetAllOrganizations_Organizations {
	if t == nil {
		t = &GetAllOrganizations{}
	}
	return &t.Organizations
}

type CreateOrganization struct {
	CreateOrganization CreateOrganization_CreateOrganization "json:\"createOrganization\" graphql:\"createOrganization\""
}

func (t *CreateOrganization) GetCreateOrganization() *CreateOrganization_CreateOrganization {
	if t == nil {
		t = &CreateOrganization{}
	}
	return &t.CreateOrganization
}

type UpdateOrganization struct {
	UpdateOrganization UpdateOrganization_UpdateOrganization "json:\"updateOrganization\" graphql:\"updateOrganization\""
}

func (t *UpdateOrganization) GetUpdateOrganization() *UpdateOrganization_UpdateOrganization {
	if t == nil {
		t = &UpdateOrganization{}
	}
	return &t.UpdateOrganization
}

type DeleteOrganization struct {
	DeleteOrganization DeleteOrganization_DeleteOrganization "json:\"deleteOrganization\" graphql:\"deleteOrganization\""
}

func (t *DeleteOrganization) GetDeleteOrganization() *DeleteOrganization_DeleteOrganization {
	if t == nil {
		t = &DeleteOrganization{}
	}
	return &t.DeleteOrganization
}

const GetOrganizationByIDDocument = `query GetOrganizationByID ($organizationId: ID!) {
	organization(id: $organizationId) {
		id
		name
		parent {
			id
			name
		}
		description
	}
}
`

func (c *Client) GetOrganizationByID(ctx context.Context, organizationID string, interceptors ...clientv2.RequestInterceptor) (*GetOrganizationByID, error) {
	vars := map[string]interface{}{
		"organizationId": organizationID,
	}

	var res GetOrganizationByID
	if err := c.Client.Post(ctx, "GetOrganizationByID", GetOrganizationByIDDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const GetAllOrganizationsDocument = `query GetAllOrganizations {
	organizations {
		edges {
			node {
				id
				name
				description
			}
		}
	}
}
`

func (c *Client) GetAllOrganizations(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllOrganizations, error) {
	vars := map[string]interface{}{}

	var res GetAllOrganizations
	if err := c.Client.Post(ctx, "GetAllOrganizations", GetAllOrganizationsDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const CreateOrganizationDocument = `mutation CreateOrganization ($input: CreateOrganizationInput!) {
	createOrganization(input: $input) {
		organization {
			id
			name
			description
			parent {
				id
				name
			}
		}
	}
}
`

func (c *Client) CreateOrganization(ctx context.Context, input CreateOrganizationInput, interceptors ...clientv2.RequestInterceptor) (*CreateOrganization, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateOrganization
	if err := c.Client.Post(ctx, "CreateOrganization", CreateOrganizationDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const UpdateOrganizationDocument = `mutation UpdateOrganization ($updateOrganizationId: ID!, $input: UpdateOrganizationInput!) {
	updateOrganization(id: $updateOrganizationId, input: $input) {
		organization {
			id
			name
			parent {
				id
				name
			}
			description
		}
	}
}
`

func (c *Client) UpdateOrganization(ctx context.Context, updateOrganizationID string, input UpdateOrganizationInput, interceptors ...clientv2.RequestInterceptor) (*UpdateOrganization, error) {
	vars := map[string]interface{}{
		"updateOrganizationId": updateOrganizationID,
		"input":                input,
	}

	var res UpdateOrganization
	if err := c.Client.Post(ctx, "UpdateOrganization", UpdateOrganizationDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const DeleteOrganizationDocument = `mutation DeleteOrganization ($deleteOrganizationId: ID!) {
	deleteOrganization(id: $deleteOrganizationId) {
		deletedID
	}
}
`

func (c *Client) DeleteOrganization(ctx context.Context, deleteOrganizationID string, interceptors ...clientv2.RequestInterceptor) (*DeleteOrganization, error) {
	vars := map[string]interface{}{
		"deleteOrganizationId": deleteOrganizationID,
	}

	var res DeleteOrganization
	if err := c.Client.Post(ctx, "DeleteOrganization", DeleteOrganizationDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}
