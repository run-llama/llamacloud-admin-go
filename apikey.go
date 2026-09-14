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

// Create a new API key.
//
// If project_id is specified, validates the user can read that project.
//
// Args: api_key_create: API key creation data user: Current user db: Database
// session
//
// Returns: The created API key with the secret key visible in redacted_api_key
// field
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys.
//
// If project_id is provided, validates user has access to that project. If
// project_id is not provided, scopes results to the current user.
//
// Args: user: Current user page_size: Number of items per page page_token: Token
// for pagination name: Filter by API key name project_id: Filter by project ID
// key_type: Filter by key type
//
// Returns: Paginated response with API keys
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
// If project_id is provided, validates user has access to that project. If
// project_id is not provided, scopes results to the current user.
//
// Args: user: Current user page_size: Number of items per page page_token: Token
// for pagination name: Filter by API key name project_id: Filter by project ID
// key_type: Filter by key type
//
// Returns: Paginated response with API keys
func (r *APIKeyService) ListAutoPaging(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[APIKey] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an API key.
//
// If the API key belongs to a project, validates user has admin permissions for
// that project. If the API key has no project, validates it belongs to the current
// user.
//
// Args: api_key_id: The ID of the API key to delete user: Current user
func (r *APIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/api-keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Schema for an API Key.
type APIKey struct {
	// Unique identifier
	ID             string `json:"id" api:"required" format:"uuid"`
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
