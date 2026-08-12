// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/run-llama/llama-cloud-admin-go/internal/apijson"
	"github.com/run-llama/llama-cloud-admin-go/internal/apiquery"
	"github.com/run-llama/llama-cloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llama-cloud-admin-go/option"
	"github.com/run-llama/llama-cloud-admin-go/packages/param"
	"github.com/run-llama/llama-cloud-admin-go/packages/respjson"
)

// AdminService contains methods and other services that help with interacting with
// the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	options      []option.RequestOption
	Users        AdminUserService
	UsageMetrics AdminUsageMetricService
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r AdminService) {
	r = AdminService{}
	r.options = opts
	r.Users = NewAdminUserService(opts...)
	r.UsageMetrics = NewAdminUsageMetricService(opts...)
	return
}

// Get File Store Info
func (r *AdminService) GetFilestoresInfo(ctx context.Context, opts ...option.RequestOption) (res *AdminGetFilestoresInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/filestores/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get License Info
func (r *AdminService) GetLicenseInfo(ctx context.Context, query AdminGetLicenseInfoParams, opts ...option.RequestOption) (res *AdminGetLicenseInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/license/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get LlamaExtract feature availability based on available models.
func (r *AdminService) GetLlamaextractFeatures(ctx context.Context, opts ...option.RequestOption) (res *AdminGetLlamaextractFeaturesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/llamaextract/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Llm Info
func (r *AdminService) GetLlmsInfo(ctx context.Context, opts ...option.RequestOption) (res *AdminGetLlmsInfoResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/llms/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get OCR service health status including GPU availability.
func (r *AdminService) GetOcrStatus(ctx context.Context, opts ...option.RequestOption) (res *AdminGetOcrStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/ocr/statusz"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Return resolved S3 configuration and presigned URL signing details.
func (r *AdminService) GetS3Config(ctx context.Context, opts ...option.RequestOption) (res *AdminGetS3ConfigResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/admin/s3/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AdminGetFilestoresInfoResponse struct {
	// Any of "missing_buckets", "missing_credentials", "ok".
	Status             AdminGetFilestoresInfoResponseStatus `json:"status" api:"required"`
	AvailableBuckets   map[string]string                    `json:"available_buckets"`
	UnavailableBuckets map[string]string                    `json:"unavailable_buckets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		AvailableBuckets   respjson.Field
		UnavailableBuckets respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetFilestoresInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetFilestoresInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetFilestoresInfoResponseStatus string

const (
	AdminGetFilestoresInfoResponseStatusMissingBuckets     AdminGetFilestoresInfoResponseStatus = "missing_buckets"
	AdminGetFilestoresInfoResponseStatusMissingCredentials AdminGetFilestoresInfoResponseStatus = "missing_credentials"
	AdminGetFilestoresInfoResponseStatusOk                 AdminGetFilestoresInfoResponseStatus = "ok"
)

type AdminGetLicenseInfoResponse struct {
	// License expiration date
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// License validation status
	Status string `json:"status" api:"required"`
	// License message
	Message string `json:"message" api:"nullable"`
	// License scopes
	Scopes []string `json:"scopes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Status      respjson.Field
		Message     respjson.Field
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLicenseInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLicenseInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetLlamaextractFeaturesResponse struct {
	AvailableModes   []AdminGetLlamaextractFeaturesResponseAvailableMode  `json:"available_modes" api:"required"`
	SchemaGeneration AdminGetLlamaextractFeaturesResponseSchemaGeneration `json:"schema_generation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableModes   respjson.Field
		SchemaGeneration respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLlamaextractFeaturesResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLlamaextractFeaturesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetLlamaextractFeaturesResponseAvailableMode struct {
	Mode      string `json:"mode" api:"required"`
	ParseMode string `json:"parse_mode" api:"required"`
	// Any of "available", "unavailable".
	Status                 string   `json:"status" api:"required"`
	AvailableExtractModels []string `json:"available_extract_models"`
	AvailableParseModels   []string `json:"available_parse_models"`
	MissingExtractModels   []string `json:"missing_extract_models"`
	MissingParseModels     []string `json:"missing_parse_models"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode                   respjson.Field
		ParseMode              respjson.Field
		Status                 respjson.Field
		AvailableExtractModels respjson.Field
		AvailableParseModels   respjson.Field
		MissingExtractModels   respjson.Field
		MissingParseModels     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLlamaextractFeaturesResponseAvailableMode) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLlamaextractFeaturesResponseAvailableMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetLlamaextractFeaturesResponseSchemaGeneration struct {
	Model string `json:"model" api:"required"`
	// Any of "available", "unavailable".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLlamaextractFeaturesResponseSchemaGeneration) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLlamaextractFeaturesResponseSchemaGeneration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetLlmsInfoResponse struct {
	LlmInfo map[string]map[string]AdminGetLlmsInfoResponseLlmInfo `json:"llm_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LlmInfo     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLlmsInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLlmsInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetLlmsInfoResponseLlmInfo struct {
	InternalModelName string    `json:"internal_model_name" api:"required"`
	Valid             bool      `json:"valid" api:"required"`
	ErrorMessage      string    `json:"error_message" api:"nullable"`
	LastValidated     time.Time `json:"last_validated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalModelName respjson.Field
		Valid             respjson.Field
		ErrorMessage      respjson.Field
		LastValidated     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetLlmsInfoResponseLlmInfo) RawJSON() string { return r.JSON.raw }
func (r *AdminGetLlmsInfoResponseLlmInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for OCR service health/GPU status.
type AdminGetOcrStatusResponse struct {
	// Any of "degraded", "ok", "unavailable".
	Status         AdminGetOcrStatusResponseStatus `json:"status" api:"required"`
	Device         string                          `json:"device"`
	ErrorMessage   string                          `json:"error_message" api:"nullable"`
	GPUAvailable   bool                            `json:"gpu_available"`
	GPUDeviceCount int64                           `json:"gpu_device_count" api:"nullable"`
	GPUDeviceName  string                          `json:"gpu_device_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status         respjson.Field
		Device         respjson.Field
		ErrorMessage   respjson.Field
		GPUAvailable   respjson.Field
		GPUDeviceCount respjson.Field
		GPUDeviceName  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetOcrStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetOcrStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetOcrStatusResponseStatus string

const (
	AdminGetOcrStatusResponseStatusDegraded    AdminGetOcrStatusResponseStatus = "degraded"
	AdminGetOcrStatusResponseStatusOk          AdminGetOcrStatusResponseStatus = "ok"
	AdminGetOcrStatusResponseStatusUnavailable AdminGetOcrStatusResponseStatus = "unavailable"
)

type AdminGetS3ConfigResponse struct {
	Buckets AdminGetS3ConfigResponseBuckets `json:"buckets" api:"required"`
	// Whether BYOC mode is enabled
	ByocModeEnabled bool `json:"byoc_mode_enabled" api:"required"`
	// Custom S3 endpoint URL (None = standard AWS)
	EndpointURL string `json:"endpoint_url" api:"required"`
	// Whether a KMS key ID is configured for server-side encryption
	KmsKeyConfigured bool `json:"kms_key_configured" api:"required"`
	// Signature version used when generating presigned URLs. 'unsigned' = s3proxy path
	// (proxy handles auth), 's3v4' = explicit SigV4, 'default' = no override set
	// (botocore default, may produce SigV2 without a region)
	//
	// Any of "default", "s3v4", "unsigned".
	PresignedURLSignatureVersion AdminGetS3ConfigResponsePresignedURLSignatureVersion `json:"presigned_url_signature_version" api:"required"`
	// Resolved value: whether requests are routed through s3proxy
	S3ProxyActive bool `json:"s3_proxy_active" api:"required"`
	// Explicit S3_PROXY_ENABLED override; None means auto-detect
	S3ProxyEnabledOverride bool `json:"s3_proxy_enabled_override" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets                      respjson.Field
		ByocModeEnabled              respjson.Field
		EndpointURL                  respjson.Field
		KmsKeyConfigured             respjson.Field
		PresignedURLSignatureVersion respjson.Field
		S3ProxyActive                respjson.Field
		S3ProxyEnabledOverride       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetS3ConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminGetS3ConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminGetS3ConfigResponseBuckets struct {
	DocumentBucket              string `json:"document_bucket" api:"required"`
	EtlBucket                   string `json:"etl_bucket" api:"required"`
	ExternalComponentsBucket    string `json:"external_components_bucket" api:"required"`
	FileParsingBucket           string `json:"file_parsing_bucket" api:"required"`
	FileScreenshotBucket        string `json:"file_screenshot_bucket" api:"required"`
	LlamaCloudParseOutputBucket string `json:"llama_cloud_parse_output_bucket" api:"required"`
	LlamaExtractOutputBucket    string `json:"llama_extract_output_bucket" api:"required"`
	RawFileBucket               string `json:"raw_file_bucket" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentBucket              respjson.Field
		EtlBucket                   respjson.Field
		ExternalComponentsBucket    respjson.Field
		FileParsingBucket           respjson.Field
		FileScreenshotBucket        respjson.Field
		LlamaCloudParseOutputBucket respjson.Field
		LlamaExtractOutputBucket    respjson.Field
		RawFileBucket               respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminGetS3ConfigResponseBuckets) RawJSON() string { return r.JSON.raw }
func (r *AdminGetS3ConfigResponseBuckets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signature version used when generating presigned URLs. 'unsigned' = s3proxy path
// (proxy handles auth), 's3v4' = explicit SigV4, 'default' = no override set
// (botocore default, may produce SigV2 without a region)
type AdminGetS3ConfigResponsePresignedURLSignatureVersion string

const (
	AdminGetS3ConfigResponsePresignedURLSignatureVersionDefault  AdminGetS3ConfigResponsePresignedURLSignatureVersion = "default"
	AdminGetS3ConfigResponsePresignedURLSignatureVersionS3v4     AdminGetS3ConfigResponsePresignedURLSignatureVersion = "s3v4"
	AdminGetS3ConfigResponsePresignedURLSignatureVersionUnsigned AdminGetS3ConfigResponsePresignedURLSignatureVersion = "unsigned"
)

type AdminGetLicenseInfoParams struct {
	// Whether to include scopes in the response
	IncludeScopes param.Opt[bool] `query:"include_scopes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminGetLicenseInfoParams]'s query parameters as
// `url.Values`.
func (r AdminGetLicenseInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
