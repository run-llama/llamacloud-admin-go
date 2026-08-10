// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin_test

import (
	"context"
	"os"
	"testing"

	"github.com/run-llama/llamacloud-admin-go"
	"github.com/run-llama/llamacloud-admin-go/internal/testutil"
	"github.com/run-llama/llamacloud-admin-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Organizations.ListAutoPaging(context.TODO(), llamacloudadmin.OrganizationListParams{
		PageSize: llamacloudadmin.Int(20),
	})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		organization := iter.Current()
		t.Logf("%+v\n", organization.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
