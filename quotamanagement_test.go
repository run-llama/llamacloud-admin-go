// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/run-llama/llamacloud-admin-go"
	"github.com/run-llama/llamacloud-admin-go/internal/testutil"
	"github.com/run-llama/llamacloud-admin-go/option"
)

func TestQuotaManagementNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudadmin.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.QuotaManagement.New(context.TODO(), llamacloudadmin.QuotaManagementNewParams{
		OrganizationID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Setting:        llamacloudadmin.QuotaManagementNewParamsSettingAllowPayAsYouGo,
		Value:          0,
		ProjectID:      llamacloudadmin.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloudadmin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaManagementListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudadmin.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.QuotaManagement.List(context.TODO(), llamacloudadmin.QuotaManagementListParams{
		SourceID:           "source_id",
		SourceType:         llamacloudadmin.QuotaManagementListParamsSourceTypeGlobal,
		ConfigurationType:  llamacloudadmin.QuotaManagementListParamsConfigurationTypeAllowPayAsYouGo,
		ExcludeSelfService: llamacloudadmin.Bool(true),
		Expand:             llamacloudadmin.Bool(true),
		Page:               llamacloudadmin.Int(0),
		PageSize:           llamacloudadmin.Int(1),
	})
	if err != nil {
		var apierr *llamacloudadmin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaManagementDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudadmin.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.QuotaManagement.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudadmin.QuotaManagementDeleteParams{
			OrganizationID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
	)
	if err != nil {
		var apierr *llamacloudadmin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
