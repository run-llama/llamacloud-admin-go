// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/run-llama/llamacloud-admin-go/internal/apijson"
	"github.com/run-llama/llamacloud-admin-go/internal/apiquery"
	shimjson "github.com/run-llama/llamacloud-admin-go/internal/encoding/json"
	"github.com/run-llama/llamacloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llamacloud-admin-go/option"
	"github.com/run-llama/llamacloud-admin-go/packages/param"
	"github.com/run-llama/llamacloud-admin-go/packages/respjson"
)

// OrganizationUserService contains methods and other services that help with
// interacting with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserService] method instead.
type OrganizationUserService struct {
	options []option.RequestOption
}

// NewOrganizationUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserService(opts ...option.RequestOption) (r OrganizationUserService) {
	r = OrganizationUserService{}
	r.options = opts
	return
}

// Remove users from an organization.
func (r *OrganizationUserService) Delete(ctx context.Context, memberUserID string, params OrganizationUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	if memberUserID == "" {
		err = errors.New("missing required member_user_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users/%s", params.OrganizationID, url.PathEscape(memberUserID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Add a user to an organization.
func (r *OrganizationUserService) Add(ctx context.Context, organizationID string, body OrganizationUserAddParams, opts ...option.RequestOption) (res *[]OrganizationMember, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Add a user to a project.
func (r *OrganizationUserService) AddToProject(ctx context.Context, userID string, params OrganizationUserAddToProjectParams, opts ...option.RequestOption) (res *OrganizationUserAddToProjectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users/%s/projects", params.OrganizationID, url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Assign a role to a user in an organization.
func (r *OrganizationUserService) AssignRole(ctx context.Context, organizationID string, body OrganizationUserAssignRoleParams, opts ...option.RequestOption) (res *UserOrganizationRole, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users/roles", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Get all users in an organization.
func (r *OrganizationUserService) ListMembers(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *[]OrganizationMember, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all projects for a user in an organization.
func (r *OrganizationUserService) ListProjects(ctx context.Context, userID string, query OrganizationUserListProjectsParams, opts ...option.RequestOption) (res *[]OrganizationUserListProjectsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users/%s/projects", query.OrganizationID, url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove a user from a project.
func (r *OrganizationUserService) RemoveFromProject(ctx context.Context, projectID string, body OrganizationUserRemoveFromProjectParams, opts ...option.RequestOption) (res *OrganizationUserRemoveFromProjectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if body.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/users/%s/projects/%s", body.OrganizationID, url.PathEscape(body.UserID), projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type OrganizationUserAddToProjectResponse = any

// Schema for a project.
type OrganizationUserListProjectsResponse struct {
	// Unique identifier
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// The Organization ID the project is under.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Whether this project is the default project for the user.
	IsDefault bool `json:"is_default"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		CreatedAt      respjson.Field
		IsDefault      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserListProjectsResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListProjectsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserRemoveFromProjectResponse = any

type OrganizationUserDeleteParams struct {
	OrganizationID string `path:"organization_id" api:"required" format:"uuid" json:"-"`
	Body           []string
	paramObj
}

func (r OrganizationUserDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OrganizationUserDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserAddParams struct {
	Body []OrganizationUserAddParamsBody
	paramObj
}

func (r OrganizationUserAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OrganizationUserAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to add a user to an organization.
//
// The property ProjectIDs is required.
type OrganizationUserAddParamsBody struct {
	// The project IDs to add the user to.
	ProjectIDs []string `json:"project_ids,omitzero" api:"required" format:"uuid"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The role ID to assign to the user.
	RoleID param.Opt[string] `json:"role_id,omitzero" format:"uuid"`
	// The user's ID.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r OrganizationUserAddParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserAddParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserAddParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserAddToProjectParams struct {
	OrganizationID string            `path:"organization_id" api:"required" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserAddToProjectParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserAddToProjectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationUserAssignRoleParams struct {
	// The organization's ID.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// The role's ID.
	RoleID string `json:"role_id" api:"required" format:"uuid"`
	// The user's ID.
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r OrganizationUserAssignRoleParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserAssignRoleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserAssignRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserListProjectsParams struct {
	OrganizationID string `path:"organization_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type OrganizationUserRemoveFromProjectParams struct {
	OrganizationID string `path:"organization_id" api:"required" format:"uuid" json:"-"`
	UserID         string `path:"user_id" api:"required" json:"-"`
	paramObj
}
