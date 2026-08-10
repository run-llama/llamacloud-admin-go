// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/run-llama/llama-cloud-admin-go/internal/apijson"
	"github.com/run-llama/llama-cloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llama-cloud-admin-go/option"
	"github.com/run-llama/llama-cloud-admin-go/packages/param"
	"github.com/run-llama/llama-cloud-admin-go/packages/respjson"
)

// AdminUserService contains methods and other services that help with interacting
// with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUserService] method instead.
type AdminUserService struct {
	options []option.RequestOption
}

// NewAdminUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminUserService(opts ...option.RequestOption) (r AdminUserService) {
	r = AdminUserService{}
	r.options = opts
	return
}

// Get a user's resolved custom claims.
//
// Claims that have not been explicitly set fall back to their system default.
// Returns 404 if the user does not exist.
//
// Global admin only.
func (r *AdminUserService) GetClaims(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserClaims, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/admin/users/%s/claims", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Additively update a user's custom claims.
//
// Claims in `set_claims` are added or overwritten; claims named in `remove_claims`
// are reset to their system default. Claims not referenced by either field are
// left unchanged, so a single claim can be changed without resending the full set.
// Returns the user's resolved claims after the update.
//
// Returns 404 if the user does not exist.
//
// Global admin only.
func (r *AdminUserService) UpdateClaims(ctx context.Context, userID string, body AdminUserUpdateClaimsParams, opts ...option.RequestOption) (res *UserClaims, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/admin/users/%s/claims", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Custom claims that dictate various limits or allowed behaviors. Currently these
// claims reside at a per user level. Claims may expand to a per organization level
// or project in the future.
type CustomClaims struct {
	// Whether the user is allowed to delete organizations.
	AllowOrgDeletion bool `json:"allow_org_deletion"`
	// Whether the user is allowed to create organizations.
	AllowedOrgCreation bool `json:"allowed_org_creation"`
	// Whether the user is allowed to access API data sources.
	APIDatasourceAccess bool `json:"api_datasource_access"`
	// Cap on how many organizations this user may create. None means unlimited. Only
	// enforced when allowed_org_creation is True.
	MaximumOrgCreation int64 `json:"maximum_org_creation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowOrgDeletion    respjson.Field
		AllowedOrgCreation  respjson.Field
		APIDatasourceAccess respjson.Field
		MaximumOrgCreation  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomClaims) RawJSON() string { return r.JSON.raw }
func (r *CustomClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user's fully resolved custom claims after applying system defaults.
type UserClaims struct {
	// The user's resolved custom claims.
	Claims CustomClaims `json:"claims" api:"required"`
	// The user ID the claims belong to.
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Claims      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserClaims) RawJSON() string { return r.JSON.raw }
func (r *UserClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminUserUpdateClaimsParams struct {
	// Names of claims to reset to their system default.
	//
	// Any of "allow_org_deletion", "allowed_org_creation", "api_datasource_access",
	// "maximum_org_creation".
	RemoveClaims []string `json:"remove_claims,omitzero"`
	// A partial set of custom claims for additive updates.
	//
	// Every field is optional. Only the claims explicitly provided in a request are
	// added or overwritten; claims left unset are not touched, so callers can change a
	// single claim without resending the full claim set.
	SetClaims AdminUserUpdateClaimsParamsSetClaims `json:"set_claims,omitzero"`
	paramObj
}

func (r AdminUserUpdateClaimsParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminUserUpdateClaimsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminUserUpdateClaimsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A partial set of custom claims for additive updates.
//
// Every field is optional. Only the claims explicitly provided in a request are
// added or overwritten; claims left unset are not touched, so callers can change a
// single claim without resending the full claim set.
type AdminUserUpdateClaimsParamsSetClaims struct {
	// Whether the user is allowed to delete organizations.
	AllowOrgDeletion param.Opt[bool] `json:"allow_org_deletion,omitzero"`
	// Whether the user is allowed to create organizations.
	AllowedOrgCreation param.Opt[bool] `json:"allowed_org_creation,omitzero"`
	// Whether the user is allowed to access API data sources.
	APIDatasourceAccess param.Opt[bool] `json:"api_datasource_access,omitzero"`
	// Cap on how many organizations this user may create. None means unlimited. Only
	// enforced when allowed_org_creation is True.
	MaximumOrgCreation param.Opt[int64] `json:"maximum_org_creation,omitzero"`
	paramObj
}

func (r AdminUserUpdateClaimsParamsSetClaims) MarshalJSON() (data []byte, err error) {
	type shadow AdminUserUpdateClaimsParamsSetClaims
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminUserUpdateClaimsParamsSetClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
