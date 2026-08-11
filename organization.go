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

// Get usage for a specific organization.
func (r *OrganizationService) GetUsage(ctx context.Context, organizationID string, query OrganizationGetUsageParams, opts ...option.RequestOption) (res *UsageAndPlan, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/usage", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type UsageAndPlan struct {
	Plan UsageAndPlanPlan `json:"plan" api:"required"`
	// Account usage totals shown alongside the plan.
	Usage UsageAndPlanUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plan        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlan) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAndPlanPlan struct {
	Limits UsageAndPlanPlanLimits `json:"limits" api:"required"`
	// Any of "contract", "plan".
	MetronomePlanType      string `json:"metronome_plan_type" api:"required"`
	MetronomeRateCardAlias string `json:"metronome_rate_card_alias" api:"required"`
	// Any of "enterprise", "enterprise_contract", "enterprise_poc", "free",
	// "free_contract", "free_v1", "free_v2", "llama_parse", "pro", "pro_v1", "pro_v2",
	// "starter_v1", "starter_v2", "unknown", "yc_deal_v1".
	Name string `json:"name" api:"required"`
	// Any of "ANNUAL", "MONTHLY", "QUARTERLY".
	PlanFrequency string `json:"plan_frequency" api:"required"`
	// The ID of the plan in Metronome
	ID string `json:"id" api:"nullable"`
	// The current billing period
	CurrentBillingPeriod UsageAndPlanPlanCurrentBillingPeriod `json:"current_billing_period" api:"nullable"`
	// The date the plan ends on
	EndingBefore time.Time `json:"ending_before" api:"nullable" format:"date-time"`
	// The number of payment failures for this organization
	FailureCount int64 `json:"failure_count"`
	// Whether the organization has a failed payment that requires support contact
	IsPaymentFailed bool `json:"is_payment_failed"`
	// The ID of the customer in Metronome
	MetronomeCustomerID string                            `json:"metronome_customer_id" api:"nullable"`
	RecurringCredits    []UsageAndPlanPlanRecurringCredit `json:"recurring_credits" api:"nullable"`
	// The date the plan starts on
	StartingOn time.Time `json:"starting_on" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limits                 respjson.Field
		MetronomePlanType      respjson.Field
		MetronomeRateCardAlias respjson.Field
		Name                   respjson.Field
		PlanFrequency          respjson.Field
		ID                     respjson.Field
		CurrentBillingPeriod   respjson.Field
		EndingBefore           respjson.Field
		FailureCount           respjson.Field
		IsPaymentFailed        respjson.Field
		MetronomeCustomerID    respjson.Field
		RecurringCredits       respjson.Field
		StartingOn             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanPlan) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAndPlanPlanLimits struct {
	// Whether usage is allowed after credit grants are exhausted
	AllowPayAsYouGo               bool  `json:"allow_pay_as_you_go" api:"required"`
	MaxConcurrentIndexJobs        int64 `json:"max_concurrent_index_jobs" api:"required"`
	MaxConcurrentParseJobsOther   int64 `json:"max_concurrent_parse_jobs_other" api:"required"`
	MaxConcurrentParseJobsPremium int64 `json:"max_concurrent_parse_jobs_premium" api:"required"`
	MaxDataSinks                  int64 `json:"max_data_sinks" api:"required"`
	MaxDataSources                int64 `json:"max_data_sources" api:"required"`
	MaxEmbeddingModels            int64 `json:"max_embedding_models" api:"required"`
	MaxExtractionAgents           int64 `json:"max_extraction_agents" api:"required"`
	MaxExtractionJobs             int64 `json:"max_extraction_jobs" api:"required"`
	MaxExtractionRuns             int64 `json:"max_extraction_runs" api:"required"`
	MaxFilesPerIndex              int64 `json:"max_files_per_index" api:"required"`
	MaxIndexes                    int64 `json:"max_indexes" api:"required"`
	MaxMonthlyInvoiceTotalUsd     int64 `json:"max_monthly_invoice_total_usd" api:"required"`
	MaxOrganizations              int64 `json:"max_organizations" api:"required"`
	MaxPagesPerIndex              int64 `json:"max_pages_per_index" api:"required"`
	MaxProjects                   int64 `json:"max_projects" api:"required"`
	MaxPublishedAgents            int64 `json:"max_published_agents" api:"required"`
	MaxReportAgentSessions        int64 `json:"max_report_agent_sessions" api:"required"`
	MaxUsers                      int64 `json:"max_users" api:"required"`
	MfaEnabled                    bool  `json:"mfa_enabled" api:"required"`
	SSOEnabled                    bool  `json:"sso_enabled" api:"required"`
	SubscriptionCostUsd           int64 `json:"subscription_cost_usd" api:"required"`
	MaxDirectories                int64 `json:"max_directories" api:"nullable"`
	MaxDirectoryFilesPerDirectory int64 `json:"max_directory_files_per_directory" api:"nullable"`
	MaxDirectoryIngestFiles       int64 `json:"max_directory_ingest_files" api:"nullable"`
	MaxDirectorySyncPlanActions   int64 `json:"max_directory_sync_plan_actions" api:"nullable"`
	// The amount of USD cents at which a soft alert should be triggered
	SpendingSoftAlertsUsdCents []int64 `json:"spending_soft_alerts_usd_cents" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowPayAsYouGo               respjson.Field
		MaxConcurrentIndexJobs        respjson.Field
		MaxConcurrentParseJobsOther   respjson.Field
		MaxConcurrentParseJobsPremium respjson.Field
		MaxDataSinks                  respjson.Field
		MaxDataSources                respjson.Field
		MaxEmbeddingModels            respjson.Field
		MaxExtractionAgents           respjson.Field
		MaxExtractionJobs             respjson.Field
		MaxExtractionRuns             respjson.Field
		MaxFilesPerIndex              respjson.Field
		MaxIndexes                    respjson.Field
		MaxMonthlyInvoiceTotalUsd     respjson.Field
		MaxOrganizations              respjson.Field
		MaxPagesPerIndex              respjson.Field
		MaxProjects                   respjson.Field
		MaxPublishedAgents            respjson.Field
		MaxReportAgentSessions        respjson.Field
		MaxUsers                      respjson.Field
		MfaEnabled                    respjson.Field
		SSOEnabled                    respjson.Field
		SubscriptionCostUsd           respjson.Field
		MaxDirectories                respjson.Field
		MaxDirectoryFilesPerDirectory respjson.Field
		MaxDirectoryIngestFiles       respjson.Field
		MaxDirectorySyncPlanActions   respjson.Field
		SpendingSoftAlertsUsdCents    respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanPlanLimits) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanPlanLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current billing period
type UsageAndPlanPlanCurrentBillingPeriod struct {
	EndDate   time.Time `json:"end_date" api:"required" format:"date-time"`
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanPlanCurrentBillingPeriod) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanPlanCurrentBillingPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAndPlanPlanRecurringCredit struct {
	CreditAmount int64                                     `json:"credit_amount" api:"required"`
	CreditType   UsageAndPlanPlanRecurringCreditCreditType `json:"credit_type" api:"required"`
	Name         string                                    `json:"name" api:"required"`
	Priority     float64                                   `json:"priority" api:"required"`
	// The ID of the product in Metronome used to represent the credit grant
	ProductID string `json:"product_id" api:"required"`
	// The fraction of the credit that will roll over to the next period, between 0 and
	// 1
	RolloverFraction float64 `json:"rollover_fraction" api:"required"`
	// How many billing periods the credit grant will last for
	PeriodsDuration float64 `json:"periods_duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditAmount     respjson.Field
		CreditType       respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		ProductID        respjson.Field
		RolloverFraction respjson.Field
		PeriodsDuration  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanPlanRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanPlanRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAndPlanPlanRecurringCreditCreditType struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanPlanRecurringCreditCreditType) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanPlanRecurringCreditCreditType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account usage totals shown alongside the plan.
type UsageAndPlanUsage struct {
	// Any of "configured_spend_limit_exceeded", "free_credits_exhausted",
	// "has_spending_alert", "internal_spending_alert", "plan_spend_limit_exceeded",
	// "plan_spend_limit_soft_alert".
	ActiveAlerts                []string                                  `json:"active_alerts"`
	ActiveFreeCreditsUsage      []UsageAndPlanUsageActiveFreeCreditsUsage `json:"active_free_credits_usage"`
	CurrentInvoiceTotalUsdCents int64                                     `json:"current_invoice_total_usd_cents" api:"nullable"`
	TotalExtractionAgents       int64                                     `json:"total_extraction_agents"`
	TotalIndexedPages           int64                                     `json:"total_indexed_pages"`
	TotalIndexes                int64                                     `json:"total_indexes"`
	TotalUsers                  int64                                     `json:"total_users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveAlerts                respjson.Field
		ActiveFreeCreditsUsage      respjson.Field
		CurrentInvoiceTotalUsdCents respjson.Field
		TotalExtractionAgents       respjson.Field
		TotalIndexedPages           respjson.Field
		TotalIndexes                respjson.Field
		TotalUsers                  respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanUsage) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAndPlanUsageActiveFreeCreditsUsage struct {
	ExpiresAt        time.Time `json:"expires_at" api:"required" format:"date-time"`
	GrantName        string    `json:"grant_name" api:"required"`
	RemainingBalance int64     `json:"remaining_balance" api:"required"`
	StartingBalance  int64     `json:"starting_balance" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt        respjson.Field
		GrantName        respjson.Field
		RemainingBalance respjson.Field
		StartingBalance  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAndPlanUsageActiveFreeCreditsUsage) RawJSON() string { return r.JSON.raw }
func (r *UsageAndPlanUsageActiveFreeCreditsUsage) UnmarshalJSON(data []byte) error {
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

type OrganizationGetUsageParams struct {
	GetCurrentInvoiceTotal param.Opt[bool] `query:"get_current_invoice_total,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationGetUsageParams]'s query parameters as
// `url.Values`.
func (r OrganizationGetUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
