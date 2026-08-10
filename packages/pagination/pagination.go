// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/run-llama/llamacloud-admin-go/internal/apijson"
	"github.com/run-llama/llamacloud-admin-go/internal/requestconfig"
	"github.com/run-llama/llamacloud-admin-go/option"
	"github.com/run-llama/llamacloud-admin-go/packages/param"
	"github.com/run-llama/llamacloud-admin-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginatedCursor[T any] struct {
	Items         []T    `json:"items"`
	NextPageToken string `json:"next_page_token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedCursor[T]) GetNextPage() (res *PaginatedCursor[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedCursorAutoPager[T any] struct {
	page *PaginatedCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedCursorAutoPager[T any](page *PaginatedCursor[T], err error) *PaginatedCursorAutoPager[T] {
	return &PaginatedCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedCursorAutoPager[T]) Index() int {
	return r.run
}
