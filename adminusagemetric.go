// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/run-llama/llama-cloud-admin-go/internal/apijson"
	"github.com/run-llama/llama-cloud-admin-go/internal/apiquery"
	"github.com/run-llama/llama-cloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llama-cloud-admin-go/option"
	"github.com/run-llama/llama-cloud-admin-go/packages/param"
	"github.com/run-llama/llama-cloud-admin-go/packages/respjson"
)

// AdminUsageMetricService contains methods and other services that help with
// interacting with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUsageMetricService] method instead.
type AdminUsageMetricService struct {
	options []option.RequestOption
}

// NewAdminUsageMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminUsageMetricService(opts ...option.RequestOption) (r AdminUsageMetricService) {
	r = AdminUsageMetricService{}
	r.options = opts
	return
}

// Aggregate usage metrics by one or more dimensions, reporting total credits used.
// Global admin only.
//
// A date range is required, which bounds the scan via the `day`-leading index.
// Supplying `organization_id` narrows it further via the `(organization_id, day)`
// index.
//
// Supported `group_by` dimensions: `day`, `organization_id`, `project_id`,
// `event_type`, `user_id`. Buckets are ordered by total credits descending.
func (r *AdminUsageMetricService) Aggregate(ctx context.Context, query AdminUsageMetricAggregateParams, opts ...option.RequestOption) (res *AdminUsageMetricAggregateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/usage-metrics/aggregate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Export usage metrics line by line as CSV over a date range. Global admin only.
//
// Each row is a single usage metric. Use the optional filters to scope the export
// to an organization, project, user, or set of event types.
func (r *AdminUsageMetricService) Export(ctx context.Context, query AdminUsageMetricExportParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "api/v1/admin/usage-metrics/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Response containing usage metrics aggregated by one or more dimensions.
type AdminUsageMetricAggregateResponse struct {
	// The aggregation buckets, ordered by total credits descending
	Buckets []AdminUsageMetricAggregateResponseBucket `json:"buckets" api:"required"`
	// The dimensions the metrics were grouped by
	//
	// Any of "day", "event_type", "organization_id", "project_id", "user_id".
	GroupBy []string `json:"group_by" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		GroupBy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminUsageMetricAggregateResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminUsageMetricAggregateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single aggregation bucket grouped by the requested dimensions.
type AdminUsageMetricAggregateResponseBucket struct {
	// The dimension values that define this bucket
	Dimensions map[string]string `json:"dimensions" api:"required"`
	// Number of metric rows in this bucket
	MetricCount int64 `json:"metric_count" api:"required"`
	// Total credits consumed by metrics in this bucket
	TotalCredits AdminUsageMetricAggregateResponseBucketTotalCreditsUnion `json:"total_credits" api:"required"`
	// Total of the metric `value` field in this bucket
	TotalValue int64 `json:"total_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions   respjson.Field
		MetricCount  respjson.Field
		TotalCredits respjson.Field
		TotalValue   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminUsageMetricAggregateResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *AdminUsageMetricAggregateResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdminUsageMetricAggregateResponseBucketTotalCreditsUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type AdminUsageMetricAggregateResponseBucketTotalCreditsUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AdminUsageMetricAggregateResponseBucketTotalCreditsUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdminUsageMetricAggregateResponseBucketTotalCreditsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AdminUsageMetricAggregateResponseBucketTotalCreditsUnion) RawJSON() string { return u.JSON.raw }

func (r *AdminUsageMetricAggregateResponseBucketTotalCreditsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminUsageMetricAggregateParams struct {
	// Inclusive lower bound on the day (YYYY-MM-DD, UTC)
	DayOnOrAfter string `query:"day_on_or_after" api:"required" json:"-"`
	// Inclusive upper bound on the day (YYYY-MM-DD, UTC)
	DayOnOrBefore string `query:"day_on_or_before" api:"required" json:"-"`
	// Dimensions to group by: day, organization_id, project_id, event_type, user_id
	GroupBy []string `query:"group_by,omitzero" api:"required" json:"-"`
	// Filter by organization ID
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	// Filter by project ID
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	// Filter by user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Filter by event types
	//
	// Any of "audio_seconds_parsed", "chart_parsing_agentic",
	// "chart_parsing_efficient", "chart_parsing_plus", "chat_message_sent",
	// "confidence_score_high", "directory_count_snapshot",
	// "directory_file_count_snapshot", "directory_files_exported",
	// "directory_files_ingested", "directory_pages_exported", "extraction_num_pages",
	// "extraction_num_pages_parsed", "form_parsing_pages", "image_classified",
	// "index_retrieve_query", "layout_aware_chart_extraction", "layout_aware_parsing",
	// "layout_extracted", "pages_classified", "pages_embedded", "pages_indexed",
	// "pages_parsed", "pages_split", "pages_verified", "precise_bbox_extraction",
	// "set_total_indexes", "set_total_pages_indexed", "spreadsheet_regions_extracted",
	// "stored_file_count", "stored_file_mb".
	EventTypes []string `query:"event_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminUsageMetricAggregateParams]'s query parameters as
// `url.Values`.
func (r AdminUsageMetricAggregateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminUsageMetricExportParams struct {
	// Inclusive lower bound on the day (YYYY-MM-DD, UTC)
	DayOnOrAfter string `query:"day_on_or_after" api:"required" json:"-"`
	// Inclusive upper bound on the day (YYYY-MM-DD, UTC)
	DayOnOrBefore string `query:"day_on_or_before" api:"required" json:"-"`
	// Filter by organization ID
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	// Filter by project ID
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	// Filter by user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Filter by event types
	//
	// Any of "audio_seconds_parsed", "chart_parsing_agentic",
	// "chart_parsing_efficient", "chart_parsing_plus", "chat_message_sent",
	// "confidence_score_high", "directory_count_snapshot",
	// "directory_file_count_snapshot", "directory_files_exported",
	// "directory_files_ingested", "directory_pages_exported", "extraction_num_pages",
	// "extraction_num_pages_parsed", "form_parsing_pages", "image_classified",
	// "index_retrieve_query", "layout_aware_chart_extraction", "layout_aware_parsing",
	// "layout_extracted", "pages_classified", "pages_embedded", "pages_indexed",
	// "pages_parsed", "pages_split", "pages_verified", "precise_bbox_extraction",
	// "set_total_indexes", "set_total_pages_indexed", "spreadsheet_regions_extracted",
	// "stored_file_count", "stored_file_mb".
	EventTypes []string `query:"event_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminUsageMetricExportParams]'s query parameters as
// `url.Values`.
func (r AdminUsageMetricExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
