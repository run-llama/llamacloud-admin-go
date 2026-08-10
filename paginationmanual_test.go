// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin_test

import (
	"context"
	"os"
	"testing"

	"github.com/run-llama/llama-cloud-admin-go"
	"github.com/run-llama/llama-cloud-admin-go/internal/testutil"
	"github.com/run-llama/llama-cloud-admin-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Organizations.List(context.TODO(), llamacloudadmin.OrganizationListParams{
		PageSize: llamacloudadmin.Int(20),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, organization := range page.Items {
		t.Logf("%+v\n", organization.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, organization := range page.Items {
			t.Logf("%+v\n", organization.ID)
		}
	}
}
