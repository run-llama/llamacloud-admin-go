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
	"github.com/run-llama/llamacloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llamacloud-admin-go/option"
	"github.com/run-llama/llamacloud-admin-go/packages/pagination"
	"github.com/run-llama/llamacloud-admin-go/packages/param"
	"github.com/run-llama/llamacloud-admin-go/packages/respjson"
)

// OrganizationService contains methods and other services that help with
// interacting with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	options []option.RequestOption
	Users   OrganizationUserService
	Roles   OrganizationRoleService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.options = opts
	r.Users = NewOrganizationUserService(opts...)
	r.Roles = NewOrganizationRoleService(opts...)
	return
}

// Create a new organization.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing organization.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List organizations the current user can access.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[Organization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/organizations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List organizations the current user can access.
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[Organization] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an organization by ID.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v2/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get an organization by ID.
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// API response schema for an organization.
type Organization struct {
	// The organization's unique identifier.
	ID string `json:"id" api:"required"`
	// The organization's display name.
	Name string `json:"name" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Additional organization metadata.
	Metadata map[string]any `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user's membership in an organization, including roles.
type OrganizationMember struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The organization's ID.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// The roles of the user in the organization.
	Roles []UserOrganizationRole `json:"roles" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The user's email address.
	Email string `json:"email" api:"nullable" format:"email"`
	// The email address of the user who added the user to the organization.
	//
	// Deprecated: deprecated
	InvitedByUserEmail string `json:"invited_by_user_email" api:"nullable" format:"email"`
	// The user ID of the user who added the user to the organization.
	InvitedByUserID string `json:"invited_by_user_id" api:"nullable"`
	// Whether the user's membership is pending account signup.
	Pending bool `json:"pending"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// The user's ID.
	UserID string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		OrganizationID     respjson.Field
		Roles              respjson.Field
		CreatedAt          respjson.Field
		Email              respjson.Field
		InvitedByUserEmail respjson.Field
		InvitedByUserID    respjson.Field
		Pending            respjson.Field
		UpdatedAt          respjson.Field
		UserID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationMember) RawJSON() string { return r.JSON.raw }
func (r *OrganizationMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for a role.
type Role struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// A name for the role.
	Name string `json:"name" api:"required"`
	// The actual permissions of the role.
	Permissions []RolePermission `json:"permissions" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Permissions respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for a permission.
type RolePermission struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the permission is granted or not.
	Access bool `json:"access" api:"required"`
	// A description for the permission.
	Description string `json:"description" api:"required"`
	// A name for the permission.
	Name string `json:"name" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Access      respjson.Field
		Description respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolePermission) RawJSON() string { return r.JSON.raw }
func (r *RolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for a user's role in an organization.
type UserOrganizationRole struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The organization's ID.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// The role.
	Role Role `json:"role" api:"required"`
	// The user's ID.
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The project ID scope.
	ProjectIDs []string `json:"project_ids" api:"nullable" format:"uuid"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		OrganizationID respjson.Field
		Role           respjson.Field
		UserID         respjson.Field
		CreatedAt      respjson.Field
		ProjectIDs     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserOrganizationRole) RawJSON() string { return r.JSON.raw }
func (r *UserOrganizationRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationNewParams struct {
	// The organization's display name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateParams struct {
	// The organization's new display name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListParams struct {
	Name      param.Opt[string] `query:"name,omitzero" json:"-"`
	PageSize  param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
