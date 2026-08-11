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

func TestAdminUsageMetricAggregateWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.UsageMetrics.Aggregate(context.TODO(), llamacloudadmin.AdminUsageMetricAggregateParams{
		DayOnOrAfter:   "day_on_or_after",
		DayOnOrBefore:  "day_on_or_before",
		GroupBy:        []string{"string"},
		EventTypes:     []string{"audio_seconds_parsed", "chart_parsing_agentic"},
		OrganizationID: llamacloudadmin.String("organization_id"),
		ProjectID:      llamacloudadmin.String("project_id"),
		UserID:         llamacloudadmin.String("user_id"),
	})
	if err != nil {
		var apierr *llamacloudadmin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
