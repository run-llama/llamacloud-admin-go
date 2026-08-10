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

// InviteService contains methods and other services that help with interacting
// with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInviteService] method instead.
type InviteService struct {
	options []option.RequestOption
}

// NewInviteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInviteService(opts ...option.RequestOption) (r InviteService) {
	r = InviteService{}
	r.options = opts
	return
}

// List the current user's pending invitations, cursor-paginated.
func (r *InviteService) List(ctx context.Context, query InviteListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[Invite], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/invites"
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

// List the current user's pending invitations, cursor-paginated.
func (r *InviteService) ListAutoPaging(ctx context.Context, query InviteListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[Invite] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Decline a pending invitation.
func (r *InviteService) Delete(ctx context.Context, inviteID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v2/invites/%s", url.PathEscape(inviteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Accept a pending invitation. Returns the joined organization id.
func (r *InviteService) Accept(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *InviteAcceptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/invites/%s/accept", url.PathEscape(inviteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A pending invitation visible to the invitee.
type Invite struct {
	// The invite's unique identifier.
	ID string `json:"id" api:"required"`
	// The organization the user is invited to.
	OrganizationID string `json:"organization_id" api:"required"`
	// The organization's display name.
	OrganizationName string `json:"organization_name" api:"required"`
	// The role being granted (e.g. admin, viewer).
	Role string `json:"role" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		OrganizationID   respjson.Field
		OrganizationName respjson.Field
		Role             respjson.Field
		CreatedAt        respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invite) RawJSON() string { return r.JSON.raw }
func (r *Invite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for accepting an invitation.
type InviteAcceptResponse struct {
	// The organization the user just joined.
	OrganizationID string `json:"organization_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InviteAcceptResponse) RawJSON() string { return r.JSON.raw }
func (r *InviteAcceptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InviteListParams struct {
	PageSize  param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InviteListParams]'s query parameters as `url.Values`.
func (r InviteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
