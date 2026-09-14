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
	"github.com/run-llama/llamacloud-admin-go/packages/param"
	"github.com/run-llama/llamacloud-admin-go/packages/respjson"
)

// QuotaManagementService contains methods and other services that help with
// interacting with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaManagementService] method instead.
type QuotaManagementService struct {
	options []option.RequestOption
}

// NewQuotaManagementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuotaManagementService(opts ...option.RequestOption) (r QuotaManagementService) {
	r = QuotaManagementService{}
	r.options = opts
	return
}

// Create a quota configuration for your organization, or for a single project
// within it.
func (r *QuotaManagementService) New(ctx context.Context, params QuotaManagementNewParams, opts ...option.RequestOption) (res *QuotaConfiguration, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/quota-management"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a paginated list of quota configurations with optional filtering. When
// expand=true, returns resolved quotas (effective values after fallback chain) and
// pagination parameters are ignored.
func (r *QuotaManagementService) List(ctx context.Context, query QuotaManagementListParams, opts ...option.RequestOption) (res *QuotaManagementListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/quota-management"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a quota configuration by removing the override.
func (r *QuotaManagementService) Delete(ctx context.Context, quotaID string, body QuotaManagementDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if quotaID == "" {
		err = errors.New("missing required quota_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/quota-management/%s", quotaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Full quota configuration model.
type QuotaConfiguration struct {
	// The configuration metadata
	ConfigurationMetadata map[string]any `json:"configuration_metadata" api:"required"`
	// The quota configuration type
	//
	// Any of "allow_pay_as_you_go", "limit_agent_coder_daily_usage_usd",
	// "limit_agent_deployments", "limit_batch_files", "limit_classify_input_tokens",
	// "limit_daily_usage_credits", "limit_directories",
	// "limit_directory_files_per_directory",
	// "limit_directory_ingest_download_size_bytes", "limit_directory_ingest_files",
	// "limit_directory_sync_plan_actions", "limit_embedding_character",
	// "limit_files_per_index", "limit_max_monthly_invoice_total_usd_cents",
	// "limit_monthly_usage_credits", "limit_projects", "limit_split_categories",
	// "limit_total_file_count", "limit_total_file_storage_bytes", "limit_users",
	// "rate_limit_batch_api_creation", "rate_limit_chat_api_message",
	// "rate_limit_classify_api_creation", "rate_limit_classify_api_list",
	// "rate_limit_classify_api_query",
	// "rate_limit_concurrent_jobs_in_execution_default",
	// "rate_limit_concurrent_jobs_in_execution_doc_ingest",
	// "rate_limit_concurrent_jobs_in_execution_metadata_update",
	// "rate_limit_default_api_read", "rate_limit_default_api_write",
	// "rate_limit_directory_file_api_read", "rate_limit_directory_file_api_write",
	// "rate_limit_directory_ingest_project_job_creation",
	// "rate_limit_extract_agent_creation", "rate_limit_extract_api_creation",
	// "rate_limit_extract_api_list", "rate_limit_extract_api_query",
	// "rate_limit_extract_concurrent_default", "rate_limit_file_api_read",
	// "rate_limit_file_api_write", "rate_limit_index_v1_pipeline_concurrent_jobs",
	// "rate_limit_parse_api_creation", "rate_limit_parse_api_list",
	// "rate_limit_parse_api_query", "rate_limit_parse_concurrent_default",
	// "rate_limit_parse_concurrent_pages_agentic",
	// "rate_limit_parse_concurrent_pages_agentic_plus",
	// "rate_limit_parse_concurrent_pages_cost_effective",
	// "rate_limit_parse_concurrent_premium", "rate_limit_parse_token_bucket_agentic",
	// "rate_limit_parse_token_bucket_agentic_plus",
	// "rate_limit_parse_token_bucket_cost_effective",
	// "rate_limit_parse_token_bucket_unknown_tier",
	// "rate_limit_project_concurrent_jobs",
	// "rate_limit_project_concurrent_turbo_jobs", "rate_limit_split_api_creation",
	// "rate_limit_split_api_query", "rate_limit_spreadsheet_api_list",
	// "rate_limit_spreadsheet_api_query", "rate_limit_spreadsheet_creation",
	// "rate_limit_usage_api_query", "rate_limit_verify_api_creation",
	// "rate_limit_verify_api_list", "rate_limit_verify_api_query".
	ConfigurationType QuotaConfigurationConfigurationType `json:"configuration_type" api:"required"`
	// The quota configuration value
	ConfigurationValue QuotaConfigurationConfigurationValue `json:"configuration_value" api:"required"`
	// The source ID, e.g. the organization ID
	SourceID string `json:"source_id" api:"required"`
	// The source type, e.g. 'organization'
	//
	// Any of "GLOBAL", "organization", "plan_tier", "project".
	SourceType QuotaConfigurationSourceType `json:"source_type" api:"required"`
	// The status of the quota, i.e. 'ACTIVE' or 'INACTIVE'
	//
	// Any of "ACTIVE", "INACTIVE".
	Status QuotaConfigurationStatus `json:"status" api:"required"`
	// The system-generated UUID for the quota
	ID string `json:"id" api:"nullable" format:"uuid"`
	// The creation date of the quota configuration in the database
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The end date of the quota
	EndedAt time.Time `json:"ended_at" api:"nullable" format:"date-time"`
	// The idempotency key
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// The start date of the quota
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// The last updated date of the quota configuration in the database
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigurationMetadata respjson.Field
		ConfigurationType     respjson.Field
		ConfigurationValue    respjson.Field
		SourceID              respjson.Field
		SourceType            respjson.Field
		Status                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		EndedAt               respjson.Field
		IdempotencyKey        respjson.Field
		StartedAt             respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaConfiguration) RawJSON() string { return r.JSON.raw }
func (r *QuotaConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The quota configuration type
type QuotaConfigurationConfigurationType string

const (
	QuotaConfigurationConfigurationTypeAllowPayAsYouGo                                  QuotaConfigurationConfigurationType = "allow_pay_as_you_go"
	QuotaConfigurationConfigurationTypeLimitAgentCoderDailyUsageUsd                     QuotaConfigurationConfigurationType = "limit_agent_coder_daily_usage_usd"
	QuotaConfigurationConfigurationTypeLimitAgentDeployments                            QuotaConfigurationConfigurationType = "limit_agent_deployments"
	QuotaConfigurationConfigurationTypeLimitBatchFiles                                  QuotaConfigurationConfigurationType = "limit_batch_files"
	QuotaConfigurationConfigurationTypeLimitClassifyInputTokens                         QuotaConfigurationConfigurationType = "limit_classify_input_tokens"
	QuotaConfigurationConfigurationTypeLimitDailyUsageCredits                           QuotaConfigurationConfigurationType = "limit_daily_usage_credits"
	QuotaConfigurationConfigurationTypeLimitDirectories                                 QuotaConfigurationConfigurationType = "limit_directories"
	QuotaConfigurationConfigurationTypeLimitDirectoryFilesPerDirectory                  QuotaConfigurationConfigurationType = "limit_directory_files_per_directory"
	QuotaConfigurationConfigurationTypeLimitDirectoryIngestDownloadSizeBytes            QuotaConfigurationConfigurationType = "limit_directory_ingest_download_size_bytes"
	QuotaConfigurationConfigurationTypeLimitDirectoryIngestFiles                        QuotaConfigurationConfigurationType = "limit_directory_ingest_files"
	QuotaConfigurationConfigurationTypeLimitDirectorySyncPlanActions                    QuotaConfigurationConfigurationType = "limit_directory_sync_plan_actions"
	QuotaConfigurationConfigurationTypeLimitEmbeddingCharacter                          QuotaConfigurationConfigurationType = "limit_embedding_character"
	QuotaConfigurationConfigurationTypeLimitFilesPerIndex                               QuotaConfigurationConfigurationType = "limit_files_per_index"
	QuotaConfigurationConfigurationTypeLimitMaxMonthlyInvoiceTotalUsdCents              QuotaConfigurationConfigurationType = "limit_max_monthly_invoice_total_usd_cents"
	QuotaConfigurationConfigurationTypeLimitMonthlyUsageCredits                         QuotaConfigurationConfigurationType = "limit_monthly_usage_credits"
	QuotaConfigurationConfigurationTypeLimitProjects                                    QuotaConfigurationConfigurationType = "limit_projects"
	QuotaConfigurationConfigurationTypeLimitSplitCategories                             QuotaConfigurationConfigurationType = "limit_split_categories"
	QuotaConfigurationConfigurationTypeLimitTotalFileCount                              QuotaConfigurationConfigurationType = "limit_total_file_count"
	QuotaConfigurationConfigurationTypeLimitTotalFileStorageBytes                       QuotaConfigurationConfigurationType = "limit_total_file_storage_bytes"
	QuotaConfigurationConfigurationTypeLimitUsers                                       QuotaConfigurationConfigurationType = "limit_users"
	QuotaConfigurationConfigurationTypeRateLimitBatchAPICreation                        QuotaConfigurationConfigurationType = "rate_limit_batch_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitChatAPIMessage                          QuotaConfigurationConfigurationType = "rate_limit_chat_api_message"
	QuotaConfigurationConfigurationTypeRateLimitClassifyAPICreation                     QuotaConfigurationConfigurationType = "rate_limit_classify_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitClassifyAPIList                         QuotaConfigurationConfigurationType = "rate_limit_classify_api_list"
	QuotaConfigurationConfigurationTypeRateLimitClassifyAPIQuery                        QuotaConfigurationConfigurationType = "rate_limit_classify_api_query"
	QuotaConfigurationConfigurationTypeRateLimitConcurrentJobsInExecutionDefault        QuotaConfigurationConfigurationType = "rate_limit_concurrent_jobs_in_execution_default"
	QuotaConfigurationConfigurationTypeRateLimitConcurrentJobsInExecutionDocIngest      QuotaConfigurationConfigurationType = "rate_limit_concurrent_jobs_in_execution_doc_ingest"
	QuotaConfigurationConfigurationTypeRateLimitConcurrentJobsInExecutionMetadataUpdate QuotaConfigurationConfigurationType = "rate_limit_concurrent_jobs_in_execution_metadata_update"
	QuotaConfigurationConfigurationTypeRateLimitDefaultAPIRead                          QuotaConfigurationConfigurationType = "rate_limit_default_api_read"
	QuotaConfigurationConfigurationTypeRateLimitDefaultAPIWrite                         QuotaConfigurationConfigurationType = "rate_limit_default_api_write"
	QuotaConfigurationConfigurationTypeRateLimitDirectoryFileAPIRead                    QuotaConfigurationConfigurationType = "rate_limit_directory_file_api_read"
	QuotaConfigurationConfigurationTypeRateLimitDirectoryFileAPIWrite                   QuotaConfigurationConfigurationType = "rate_limit_directory_file_api_write"
	QuotaConfigurationConfigurationTypeRateLimitDirectoryIngestProjectJobCreation       QuotaConfigurationConfigurationType = "rate_limit_directory_ingest_project_job_creation"
	QuotaConfigurationConfigurationTypeRateLimitExtractAgentCreation                    QuotaConfigurationConfigurationType = "rate_limit_extract_agent_creation"
	QuotaConfigurationConfigurationTypeRateLimitExtractAPICreation                      QuotaConfigurationConfigurationType = "rate_limit_extract_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitExtractAPIList                          QuotaConfigurationConfigurationType = "rate_limit_extract_api_list"
	QuotaConfigurationConfigurationTypeRateLimitExtractAPIQuery                         QuotaConfigurationConfigurationType = "rate_limit_extract_api_query"
	QuotaConfigurationConfigurationTypeRateLimitExtractConcurrentDefault                QuotaConfigurationConfigurationType = "rate_limit_extract_concurrent_default"
	QuotaConfigurationConfigurationTypeRateLimitFileAPIRead                             QuotaConfigurationConfigurationType = "rate_limit_file_api_read"
	QuotaConfigurationConfigurationTypeRateLimitFileAPIWrite                            QuotaConfigurationConfigurationType = "rate_limit_file_api_write"
	QuotaConfigurationConfigurationTypeRateLimitIndexV1PipelineConcurrentJobs           QuotaConfigurationConfigurationType = "rate_limit_index_v1_pipeline_concurrent_jobs"
	QuotaConfigurationConfigurationTypeRateLimitParseAPICreation                        QuotaConfigurationConfigurationType = "rate_limit_parse_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitParseAPIList                            QuotaConfigurationConfigurationType = "rate_limit_parse_api_list"
	QuotaConfigurationConfigurationTypeRateLimitParseAPIQuery                           QuotaConfigurationConfigurationType = "rate_limit_parse_api_query"
	QuotaConfigurationConfigurationTypeRateLimitParseConcurrentDefault                  QuotaConfigurationConfigurationType = "rate_limit_parse_concurrent_default"
	QuotaConfigurationConfigurationTypeRateLimitParseConcurrentPagesAgentic             QuotaConfigurationConfigurationType = "rate_limit_parse_concurrent_pages_agentic"
	QuotaConfigurationConfigurationTypeRateLimitParseConcurrentPagesAgenticPlus         QuotaConfigurationConfigurationType = "rate_limit_parse_concurrent_pages_agentic_plus"
	QuotaConfigurationConfigurationTypeRateLimitParseConcurrentPagesCostEffective       QuotaConfigurationConfigurationType = "rate_limit_parse_concurrent_pages_cost_effective"
	QuotaConfigurationConfigurationTypeRateLimitParseConcurrentPremium                  QuotaConfigurationConfigurationType = "rate_limit_parse_concurrent_premium"
	QuotaConfigurationConfigurationTypeRateLimitParseTokenBucketAgentic                 QuotaConfigurationConfigurationType = "rate_limit_parse_token_bucket_agentic"
	QuotaConfigurationConfigurationTypeRateLimitParseTokenBucketAgenticPlus             QuotaConfigurationConfigurationType = "rate_limit_parse_token_bucket_agentic_plus"
	QuotaConfigurationConfigurationTypeRateLimitParseTokenBucketCostEffective           QuotaConfigurationConfigurationType = "rate_limit_parse_token_bucket_cost_effective"
	QuotaConfigurationConfigurationTypeRateLimitParseTokenBucketUnknownTier             QuotaConfigurationConfigurationType = "rate_limit_parse_token_bucket_unknown_tier"
	QuotaConfigurationConfigurationTypeRateLimitProjectConcurrentJobs                   QuotaConfigurationConfigurationType = "rate_limit_project_concurrent_jobs"
	QuotaConfigurationConfigurationTypeRateLimitProjectConcurrentTurboJobs              QuotaConfigurationConfigurationType = "rate_limit_project_concurrent_turbo_jobs"
	QuotaConfigurationConfigurationTypeRateLimitSplitAPICreation                        QuotaConfigurationConfigurationType = "rate_limit_split_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitSplitAPIQuery                           QuotaConfigurationConfigurationType = "rate_limit_split_api_query"
	QuotaConfigurationConfigurationTypeRateLimitSpreadsheetAPIList                      QuotaConfigurationConfigurationType = "rate_limit_spreadsheet_api_list"
	QuotaConfigurationConfigurationTypeRateLimitSpreadsheetAPIQuery                     QuotaConfigurationConfigurationType = "rate_limit_spreadsheet_api_query"
	QuotaConfigurationConfigurationTypeRateLimitSpreadsheetCreation                     QuotaConfigurationConfigurationType = "rate_limit_spreadsheet_creation"
	QuotaConfigurationConfigurationTypeRateLimitUsageAPIQuery                           QuotaConfigurationConfigurationType = "rate_limit_usage_api_query"
	QuotaConfigurationConfigurationTypeRateLimitVerifyAPICreation                       QuotaConfigurationConfigurationType = "rate_limit_verify_api_creation"
	QuotaConfigurationConfigurationTypeRateLimitVerifyAPIList                           QuotaConfigurationConfigurationType = "rate_limit_verify_api_list"
	QuotaConfigurationConfigurationTypeRateLimitVerifyAPIQuery                          QuotaConfigurationConfigurationType = "rate_limit_verify_api_query"
)

// The quota configuration value
type QuotaConfigurationConfigurationValue struct {
	// The rate numerator
	Numerator int64 `json:"numerator" api:"required"`
	// The rate limit denominator
	Denominator int64 `json:"denominator" api:"nullable"`
	// The default rate limit denominator units
	//
	// Any of "day", "hour", "minute", "second".
	DenominatorUnits string `json:"denominator_units" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Numerator        respjson.Field
		Denominator      respjson.Field
		DenominatorUnits respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaConfigurationConfigurationValue) RawJSON() string { return r.JSON.raw }
func (r *QuotaConfigurationConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source type, e.g. 'organization'
type QuotaConfigurationSourceType string

const (
	QuotaConfigurationSourceTypeGlobal       QuotaConfigurationSourceType = "GLOBAL"
	QuotaConfigurationSourceTypeOrganization QuotaConfigurationSourceType = "organization"
	QuotaConfigurationSourceTypePlanTier     QuotaConfigurationSourceType = "plan_tier"
	QuotaConfigurationSourceTypeProject      QuotaConfigurationSourceType = "project"
)

// The status of the quota, i.e. 'ACTIVE' or 'INACTIVE'
type QuotaConfigurationStatus string

const (
	QuotaConfigurationStatusActive   QuotaConfigurationStatus = "ACTIVE"
	QuotaConfigurationStatusInactive QuotaConfigurationStatus = "INACTIVE"
)

// Paginated list of quota configurations.
type QuotaManagementListResponse struct {
	Items []QuotaConfiguration `json:"items" api:"required"`
	Page  int64                `json:"page" api:"required"`
	Pages int64                `json:"pages" api:"required"`
	Size  int64                `json:"size" api:"required"`
	Total int64                `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Size        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaManagementListResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaManagementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaManagementNewParams struct {
	OrganizationID string `query:"organization_id" api:"required" format:"uuid" json:"-"`
	// The quota setting to update
	//
	// Any of "allow_pay_as_you_go", "limit_daily_usage_credits",
	// "limit_monthly_usage_credits".
	Setting QuotaManagementNewParamsSetting `json:"setting,omitzero" api:"required"`
	// The value for the setting. For boolean settings, use 1 (enabled) or 0
	// (disabled). For credit limits, the number of credits allowed in the window;
	// delete the setting to remove the limit. For limits denominated in USD cents, a
	// whole number of dollars (a multiple of 100).
	Value int64 `json:"value" api:"required"`
	// Limit this project on its own. Omit to limit the organization as a whole. A
	// project limit does not inherit from the organization limit: both apply, and
	// whichever is reached first stops the work. Credit limits only.
	ProjectID param.Opt[string] `json:"project_id,omitzero" format:"uuid"`
	paramObj
}

func (r QuotaManagementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QuotaManagementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaManagementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [QuotaManagementNewParams]'s query parameters as
// `url.Values`.
func (r QuotaManagementNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The quota setting to update
type QuotaManagementNewParamsSetting string

const (
	QuotaManagementNewParamsSettingAllowPayAsYouGo          QuotaManagementNewParamsSetting = "allow_pay_as_you_go"
	QuotaManagementNewParamsSettingLimitDailyUsageCredits   QuotaManagementNewParamsSetting = "limit_daily_usage_credits"
	QuotaManagementNewParamsSettingLimitMonthlyUsageCredits QuotaManagementNewParamsSetting = "limit_monthly_usage_credits"
)

type QuotaManagementListParams struct {
	SourceID string `query:"source_id" api:"required" json:"-"`
	// Any of "GLOBAL", "organization", "plan_tier", "project".
	SourceType         QuotaManagementListParamsSourceType `query:"source_type,omitzero" api:"required" json:"-"`
	ExcludeSelfService param.Opt[bool]                     `query:"exclude_self_service,omitzero" json:"-"`
	Expand             param.Opt[bool]                     `query:"expand,omitzero" json:"-"`
	Page               param.Opt[int64]                    `query:"page,omitzero" json:"-"`
	PageSize           param.Opt[int64]                    `query:"page_size,omitzero" json:"-"`
	// Any of "allow_pay_as_you_go", "limit_agent_coder_daily_usage_usd",
	// "limit_agent_deployments", "limit_batch_files", "limit_classify_input_tokens",
	// "limit_daily_usage_credits", "limit_directories",
	// "limit_directory_files_per_directory",
	// "limit_directory_ingest_download_size_bytes", "limit_directory_ingest_files",
	// "limit_directory_sync_plan_actions", "limit_embedding_character",
	// "limit_files_per_index", "limit_max_monthly_invoice_total_usd_cents",
	// "limit_monthly_usage_credits", "limit_projects", "limit_split_categories",
	// "limit_total_file_count", "limit_total_file_storage_bytes", "limit_users",
	// "rate_limit_batch_api_creation", "rate_limit_chat_api_message",
	// "rate_limit_classify_api_creation", "rate_limit_classify_api_list",
	// "rate_limit_classify_api_query",
	// "rate_limit_concurrent_jobs_in_execution_default",
	// "rate_limit_concurrent_jobs_in_execution_doc_ingest",
	// "rate_limit_concurrent_jobs_in_execution_metadata_update",
	// "rate_limit_default_api_read", "rate_limit_default_api_write",
	// "rate_limit_directory_file_api_read", "rate_limit_directory_file_api_write",
	// "rate_limit_directory_ingest_project_job_creation",
	// "rate_limit_extract_agent_creation", "rate_limit_extract_api_creation",
	// "rate_limit_extract_api_list", "rate_limit_extract_api_query",
	// "rate_limit_extract_concurrent_default", "rate_limit_file_api_read",
	// "rate_limit_file_api_write", "rate_limit_index_v1_pipeline_concurrent_jobs",
	// "rate_limit_parse_api_creation", "rate_limit_parse_api_list",
	// "rate_limit_parse_api_query", "rate_limit_parse_concurrent_default",
	// "rate_limit_parse_concurrent_pages_agentic",
	// "rate_limit_parse_concurrent_pages_agentic_plus",
	// "rate_limit_parse_concurrent_pages_cost_effective",
	// "rate_limit_parse_concurrent_premium", "rate_limit_parse_token_bucket_agentic",
	// "rate_limit_parse_token_bucket_agentic_plus",
	// "rate_limit_parse_token_bucket_cost_effective",
	// "rate_limit_parse_token_bucket_unknown_tier",
	// "rate_limit_project_concurrent_jobs",
	// "rate_limit_project_concurrent_turbo_jobs", "rate_limit_split_api_creation",
	// "rate_limit_split_api_query", "rate_limit_spreadsheet_api_list",
	// "rate_limit_spreadsheet_api_query", "rate_limit_spreadsheet_creation",
	// "rate_limit_usage_api_query", "rate_limit_verify_api_creation",
	// "rate_limit_verify_api_list", "rate_limit_verify_api_query".
	ConfigurationType QuotaManagementListParamsConfigurationType `query:"configuration_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuotaManagementListParams]'s query parameters as
// `url.Values`.
func (r QuotaManagementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuotaManagementListParamsSourceType string

const (
	QuotaManagementListParamsSourceTypeGlobal       QuotaManagementListParamsSourceType = "GLOBAL"
	QuotaManagementListParamsSourceTypeOrganization QuotaManagementListParamsSourceType = "organization"
	QuotaManagementListParamsSourceTypePlanTier     QuotaManagementListParamsSourceType = "plan_tier"
	QuotaManagementListParamsSourceTypeProject      QuotaManagementListParamsSourceType = "project"
)

type QuotaManagementListParamsConfigurationType string

const (
	QuotaManagementListParamsConfigurationTypeAllowPayAsYouGo                                  QuotaManagementListParamsConfigurationType = "allow_pay_as_you_go"
	QuotaManagementListParamsConfigurationTypeLimitAgentCoderDailyUsageUsd                     QuotaManagementListParamsConfigurationType = "limit_agent_coder_daily_usage_usd"
	QuotaManagementListParamsConfigurationTypeLimitAgentDeployments                            QuotaManagementListParamsConfigurationType = "limit_agent_deployments"
	QuotaManagementListParamsConfigurationTypeLimitBatchFiles                                  QuotaManagementListParamsConfigurationType = "limit_batch_files"
	QuotaManagementListParamsConfigurationTypeLimitClassifyInputTokens                         QuotaManagementListParamsConfigurationType = "limit_classify_input_tokens"
	QuotaManagementListParamsConfigurationTypeLimitDailyUsageCredits                           QuotaManagementListParamsConfigurationType = "limit_daily_usage_credits"
	QuotaManagementListParamsConfigurationTypeLimitDirectories                                 QuotaManagementListParamsConfigurationType = "limit_directories"
	QuotaManagementListParamsConfigurationTypeLimitDirectoryFilesPerDirectory                  QuotaManagementListParamsConfigurationType = "limit_directory_files_per_directory"
	QuotaManagementListParamsConfigurationTypeLimitDirectoryIngestDownloadSizeBytes            QuotaManagementListParamsConfigurationType = "limit_directory_ingest_download_size_bytes"
	QuotaManagementListParamsConfigurationTypeLimitDirectoryIngestFiles                        QuotaManagementListParamsConfigurationType = "limit_directory_ingest_files"
	QuotaManagementListParamsConfigurationTypeLimitDirectorySyncPlanActions                    QuotaManagementListParamsConfigurationType = "limit_directory_sync_plan_actions"
	QuotaManagementListParamsConfigurationTypeLimitEmbeddingCharacter                          QuotaManagementListParamsConfigurationType = "limit_embedding_character"
	QuotaManagementListParamsConfigurationTypeLimitFilesPerIndex                               QuotaManagementListParamsConfigurationType = "limit_files_per_index"
	QuotaManagementListParamsConfigurationTypeLimitMaxMonthlyInvoiceTotalUsdCents              QuotaManagementListParamsConfigurationType = "limit_max_monthly_invoice_total_usd_cents"
	QuotaManagementListParamsConfigurationTypeLimitMonthlyUsageCredits                         QuotaManagementListParamsConfigurationType = "limit_monthly_usage_credits"
	QuotaManagementListParamsConfigurationTypeLimitProjects                                    QuotaManagementListParamsConfigurationType = "limit_projects"
	QuotaManagementListParamsConfigurationTypeLimitSplitCategories                             QuotaManagementListParamsConfigurationType = "limit_split_categories"
	QuotaManagementListParamsConfigurationTypeLimitTotalFileCount                              QuotaManagementListParamsConfigurationType = "limit_total_file_count"
	QuotaManagementListParamsConfigurationTypeLimitTotalFileStorageBytes                       QuotaManagementListParamsConfigurationType = "limit_total_file_storage_bytes"
	QuotaManagementListParamsConfigurationTypeLimitUsers                                       QuotaManagementListParamsConfigurationType = "limit_users"
	QuotaManagementListParamsConfigurationTypeRateLimitBatchAPICreation                        QuotaManagementListParamsConfigurationType = "rate_limit_batch_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitChatAPIMessage                          QuotaManagementListParamsConfigurationType = "rate_limit_chat_api_message"
	QuotaManagementListParamsConfigurationTypeRateLimitClassifyAPICreation                     QuotaManagementListParamsConfigurationType = "rate_limit_classify_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitClassifyAPIList                         QuotaManagementListParamsConfigurationType = "rate_limit_classify_api_list"
	QuotaManagementListParamsConfigurationTypeRateLimitClassifyAPIQuery                        QuotaManagementListParamsConfigurationType = "rate_limit_classify_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitConcurrentJobsInExecutionDefault        QuotaManagementListParamsConfigurationType = "rate_limit_concurrent_jobs_in_execution_default"
	QuotaManagementListParamsConfigurationTypeRateLimitConcurrentJobsInExecutionDocIngest      QuotaManagementListParamsConfigurationType = "rate_limit_concurrent_jobs_in_execution_doc_ingest"
	QuotaManagementListParamsConfigurationTypeRateLimitConcurrentJobsInExecutionMetadataUpdate QuotaManagementListParamsConfigurationType = "rate_limit_concurrent_jobs_in_execution_metadata_update"
	QuotaManagementListParamsConfigurationTypeRateLimitDefaultAPIRead                          QuotaManagementListParamsConfigurationType = "rate_limit_default_api_read"
	QuotaManagementListParamsConfigurationTypeRateLimitDefaultAPIWrite                         QuotaManagementListParamsConfigurationType = "rate_limit_default_api_write"
	QuotaManagementListParamsConfigurationTypeRateLimitDirectoryFileAPIRead                    QuotaManagementListParamsConfigurationType = "rate_limit_directory_file_api_read"
	QuotaManagementListParamsConfigurationTypeRateLimitDirectoryFileAPIWrite                   QuotaManagementListParamsConfigurationType = "rate_limit_directory_file_api_write"
	QuotaManagementListParamsConfigurationTypeRateLimitDirectoryIngestProjectJobCreation       QuotaManagementListParamsConfigurationType = "rate_limit_directory_ingest_project_job_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitExtractAgentCreation                    QuotaManagementListParamsConfigurationType = "rate_limit_extract_agent_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitExtractAPICreation                      QuotaManagementListParamsConfigurationType = "rate_limit_extract_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitExtractAPIList                          QuotaManagementListParamsConfigurationType = "rate_limit_extract_api_list"
	QuotaManagementListParamsConfigurationTypeRateLimitExtractAPIQuery                         QuotaManagementListParamsConfigurationType = "rate_limit_extract_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitExtractConcurrentDefault                QuotaManagementListParamsConfigurationType = "rate_limit_extract_concurrent_default"
	QuotaManagementListParamsConfigurationTypeRateLimitFileAPIRead                             QuotaManagementListParamsConfigurationType = "rate_limit_file_api_read"
	QuotaManagementListParamsConfigurationTypeRateLimitFileAPIWrite                            QuotaManagementListParamsConfigurationType = "rate_limit_file_api_write"
	QuotaManagementListParamsConfigurationTypeRateLimitIndexV1PipelineConcurrentJobs           QuotaManagementListParamsConfigurationType = "rate_limit_index_v1_pipeline_concurrent_jobs"
	QuotaManagementListParamsConfigurationTypeRateLimitParseAPICreation                        QuotaManagementListParamsConfigurationType = "rate_limit_parse_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitParseAPIList                            QuotaManagementListParamsConfigurationType = "rate_limit_parse_api_list"
	QuotaManagementListParamsConfigurationTypeRateLimitParseAPIQuery                           QuotaManagementListParamsConfigurationType = "rate_limit_parse_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitParseConcurrentDefault                  QuotaManagementListParamsConfigurationType = "rate_limit_parse_concurrent_default"
	QuotaManagementListParamsConfigurationTypeRateLimitParseConcurrentPagesAgentic             QuotaManagementListParamsConfigurationType = "rate_limit_parse_concurrent_pages_agentic"
	QuotaManagementListParamsConfigurationTypeRateLimitParseConcurrentPagesAgenticPlus         QuotaManagementListParamsConfigurationType = "rate_limit_parse_concurrent_pages_agentic_plus"
	QuotaManagementListParamsConfigurationTypeRateLimitParseConcurrentPagesCostEffective       QuotaManagementListParamsConfigurationType = "rate_limit_parse_concurrent_pages_cost_effective"
	QuotaManagementListParamsConfigurationTypeRateLimitParseConcurrentPremium                  QuotaManagementListParamsConfigurationType = "rate_limit_parse_concurrent_premium"
	QuotaManagementListParamsConfigurationTypeRateLimitParseTokenBucketAgentic                 QuotaManagementListParamsConfigurationType = "rate_limit_parse_token_bucket_agentic"
	QuotaManagementListParamsConfigurationTypeRateLimitParseTokenBucketAgenticPlus             QuotaManagementListParamsConfigurationType = "rate_limit_parse_token_bucket_agentic_plus"
	QuotaManagementListParamsConfigurationTypeRateLimitParseTokenBucketCostEffective           QuotaManagementListParamsConfigurationType = "rate_limit_parse_token_bucket_cost_effective"
	QuotaManagementListParamsConfigurationTypeRateLimitParseTokenBucketUnknownTier             QuotaManagementListParamsConfigurationType = "rate_limit_parse_token_bucket_unknown_tier"
	QuotaManagementListParamsConfigurationTypeRateLimitProjectConcurrentJobs                   QuotaManagementListParamsConfigurationType = "rate_limit_project_concurrent_jobs"
	QuotaManagementListParamsConfigurationTypeRateLimitProjectConcurrentTurboJobs              QuotaManagementListParamsConfigurationType = "rate_limit_project_concurrent_turbo_jobs"
	QuotaManagementListParamsConfigurationTypeRateLimitSplitAPICreation                        QuotaManagementListParamsConfigurationType = "rate_limit_split_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitSplitAPIQuery                           QuotaManagementListParamsConfigurationType = "rate_limit_split_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitSpreadsheetAPIList                      QuotaManagementListParamsConfigurationType = "rate_limit_spreadsheet_api_list"
	QuotaManagementListParamsConfigurationTypeRateLimitSpreadsheetAPIQuery                     QuotaManagementListParamsConfigurationType = "rate_limit_spreadsheet_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitSpreadsheetCreation                     QuotaManagementListParamsConfigurationType = "rate_limit_spreadsheet_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitUsageAPIQuery                           QuotaManagementListParamsConfigurationType = "rate_limit_usage_api_query"
	QuotaManagementListParamsConfigurationTypeRateLimitVerifyAPICreation                       QuotaManagementListParamsConfigurationType = "rate_limit_verify_api_creation"
	QuotaManagementListParamsConfigurationTypeRateLimitVerifyAPIList                           QuotaManagementListParamsConfigurationType = "rate_limit_verify_api_list"
	QuotaManagementListParamsConfigurationTypeRateLimitVerifyAPIQuery                          QuotaManagementListParamsConfigurationType = "rate_limit_verify_api_query"
)

type QuotaManagementDeleteParams struct {
	OrganizationID string `query:"organization_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [QuotaManagementDeleteParams]'s query parameters as
// `url.Values`.
func (r QuotaManagementDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
