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

// APIKeyService contains methods and other services that help with interacting
// with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.options = opts
	return
}

// Create an API key.
//
// Scope it to a project with `project_id`, which requires read access to that
// project; omit it for a key that reaches every project you can read. A
// project-scoped key cannot escape its own project: it confines an omitted
// `project_id` to that project and refuses any other. The response carries the
// secret in `redacted_api_key`, and only this once.
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys.
//
// Name a `project_id` to list every key on that project, which its members share;
// naming one you cannot read is a 404. Omit it to list your own. A project-scoped
// key sees only its own project either way.
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[APIKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/api-keys"
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

// List API keys.
//
// Name a `project_id` to list every key on that project, which its members share;
// naming one you cannot read is a 404. Omit it to list your own. A project-scoped
// key sees only its own project either way.
func (r *APIKeyService) ListAutoPaging(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[APIKey] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Revoke an API key.
//
// Revoking a project key takes access away from everyone using it, so it needs
// key-management permission on that project. Your own unscoped keys need only that
// you own them. A project-scoped key revokes only within its own project, unscoped
// keys included.
func (r *APIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *APIKeyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/api-keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Schema for an API Key.
type APIKey struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The key with its middle masked, except on the create response, which returns the
	// full secret once and never again.
	RedactedAPIKey string `json:"redacted_api_key" api:"required"`
	UserID         string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// When the API key expires. Null if the key never expires.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Any of "agent", "user".
	KeyType   APIKeyKeyType  `json:"key_type"`
	Metadata  map[string]any `json:"metadata" api:"nullable"`
	Name      string         `json:"name" api:"nullable"`
	ProjectID string         `json:"project_id" api:"nullable" format:"uuid"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		RedactedAPIKey respjson.Field
		UserID         respjson.Field
		CreatedAt      respjson.Field
		ExpiresAt      respjson.Field
		KeyType        respjson.Field
		Metadata       respjson.Field
		Name           respjson.Field
		ProjectID      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyKeyType string

const (
	APIKeyKeyTypeAgent APIKeyKeyType = "agent"
	APIKeyKeyTypeUser  APIKeyKeyType = "user"
)

// Confirmation that a resource was deleted.
type APIKeyDeleteResponse struct {
	// Maximum seconds until cached information expires
	CacheTtlSeconds int64 `json:"cache_ttl_seconds" api:"required"`
	// Whether the resource was deleted
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheTtlSeconds respjson.Field
		Success         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	// When the API key should expire. If not set, the key never expires.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	Name      param.Opt[string]    `json:"name,omitzero"`
	// The project ID to associate with the API key.
	ProjectID param.Opt[string] `json:"project_id,omitzero" format:"uuid"`
	// Any of "agent", "user".
	KeyType APIKeyNewParamsKeyType `json:"key_type,omitzero"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParamsKeyType string

const (
	APIKeyNewParamsKeyTypeAgent APIKeyNewParamsKeyType = "agent"
	APIKeyNewParamsKeyTypeUser  APIKeyNewParamsKeyType = "user"
)

type APIKeyListParams struct {
	Name      param.Opt[string] `query:"name,omitzero" json:"-"`
	PageSize  param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Any of "agent", "user".
	KeyType APIKeyListParamsKeyType `query:"key_type,omitzero" json:"-"`
	Expand  []string                `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIKeyListParamsKeyType string

const (
	APIKeyListParamsKeyTypeAgent APIKeyListParamsKeyType = "agent"
	APIKeyListParamsKeyTypeUser  APIKeyListParamsKeyType = "user"
)
