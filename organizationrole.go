// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudadmin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/run-llama/llamacloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llamacloud-admin-go/option"
)

// OrganizationRoleService contains methods and other services that help with
// interacting with the llama-cloud-admin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationRoleService] method instead.
type OrganizationRoleService struct {
	options []option.RequestOption
}

// NewOrganizationRoleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationRoleService(opts ...option.RequestOption) (r OrganizationRoleService) {
	r = OrganizationRoleService{}
	r.options = opts
	return
}

// List all roles in an organization.
func (r *OrganizationRoleService) List(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *[]Role, err error) {
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/organizations/%s/roles", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
