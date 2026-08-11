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

// ProjectService contains methods and other services that help with interacting
// with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.options = opts
	return
}

// Create a new project in the given organization.
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update an existing project.
func (r *ProjectService) Update(ctx context.Context, projectID string, params ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List projects in an organization. Requires `organization_id` or a project-scoped
// API key.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[Project], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/projects"
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

// List projects in an organization. Requires `organization_id` or a project-scoped
// API key.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[Project] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a project by ID.
func (r *ProjectService) Delete(ctx context.Context, projectID string, body ProjectDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v2/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Get a project by ID.
func (r *ProjectService) Get(ctx context.Context, projectID string, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get usage for a project
func (r *ProjectService) GetUsage(ctx context.Context, projectID string, query ProjectGetUsageParams, opts ...option.RequestOption) (res *ProjectGetUsageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/projects/%s/usage", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// API response schema for a project.
type Project struct {
	// The project's unique identifier.
	ID string `json:"id" api:"required"`
	// The project's display name.
	Name string `json:"name" api:"required"`
	// The organization the project belongs to.
	OrganizationID string `json:"organization_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Whether this project is the default project for its organization.
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
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponse struct {
	Plan ProjectGetUsageResponsePlan `json:"plan" api:"required"`
	// Account usage totals shown alongside the plan.
	Usage ProjectGetUsageResponseUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plan        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectGetUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponsePlan struct {
	Limits ProjectGetUsageResponsePlanLimits `json:"limits" api:"required"`
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
	CurrentBillingPeriod ProjectGetUsageResponsePlanCurrentBillingPeriod `json:"current_billing_period" api:"nullable"`
	// The date the plan ends on
	EndingBefore time.Time `json:"ending_before" api:"nullable" format:"date-time"`
	// The number of payment failures for this organization
	FailureCount int64 `json:"failure_count"`
	// Whether the organization has a failed payment that requires support contact
	IsPaymentFailed bool `json:"is_payment_failed"`
	// The ID of the customer in Metronome
	MetronomeCustomerID string                                       `json:"metronome_customer_id" api:"nullable"`
	RecurringCredits    []ProjectGetUsageResponsePlanRecurringCredit `json:"recurring_credits" api:"nullable"`
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
func (r ProjectGetUsageResponsePlan) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponsePlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponsePlanLimits struct {
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
func (r ProjectGetUsageResponsePlanLimits) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponsePlanLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current billing period
type ProjectGetUsageResponsePlanCurrentBillingPeriod struct {
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
func (r ProjectGetUsageResponsePlanCurrentBillingPeriod) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponsePlanCurrentBillingPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponsePlanRecurringCredit struct {
	CreditAmount int64                                                `json:"credit_amount" api:"required"`
	CreditType   ProjectGetUsageResponsePlanRecurringCreditCreditType `json:"credit_type" api:"required"`
	Name         string                                               `json:"name" api:"required"`
	Priority     float64                                              `json:"priority" api:"required"`
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
func (r ProjectGetUsageResponsePlanRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponsePlanRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponsePlanRecurringCreditCreditType struct {
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
func (r ProjectGetUsageResponsePlanRecurringCreditCreditType) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponsePlanRecurringCreditCreditType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account usage totals shown alongside the plan.
type ProjectGetUsageResponseUsage struct {
	// Any of "configured_spend_limit_exceeded", "free_credits_exhausted",
	// "has_spending_alert", "internal_spending_alert", "plan_spend_limit_exceeded",
	// "plan_spend_limit_soft_alert".
	ActiveAlerts                []string                                             `json:"active_alerts"`
	ActiveFreeCreditsUsage      []ProjectGetUsageResponseUsageActiveFreeCreditsUsage `json:"active_free_credits_usage"`
	CurrentInvoiceTotalUsdCents int64                                                `json:"current_invoice_total_usd_cents" api:"nullable"`
	TotalExtractionAgents       int64                                                `json:"total_extraction_agents"`
	TotalIndexedPages           int64                                                `json:"total_indexed_pages"`
	TotalIndexes                int64                                                `json:"total_indexes"`
	TotalUsers                  int64                                                `json:"total_users"`
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
func (r ProjectGetUsageResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetUsageResponseUsageActiveFreeCreditsUsage struct {
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
func (r ProjectGetUsageResponseUsageActiveFreeCreditsUsage) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetUsageResponseUsageActiveFreeCreditsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	OrganizationID string `query:"organization_id" api:"required" format:"uuid" json:"-"`
	// The project's display name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ProjectNewParams]'s query parameters as `url.Values`.
func (r ProjectNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectUpdateParams struct {
	// The project's new display name.
	Name           string            `json:"name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ProjectUpdateParams]'s query parameters as `url.Values`.
func (r ProjectUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectListParams struct {
	Name           param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectDeleteParams]'s query parameters as `url.Values`.
func (r ProjectDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectGetParams]'s query parameters as `url.Values`.
func (r ProjectGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectGetUsageParams struct {
	OrganizationID         param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	GetCurrentInvoiceTotal param.Opt[bool]   `query:"get_current_invoice_total,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectGetUsageParams]'s query parameters as `url.Values`.
func (r ProjectGetUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
