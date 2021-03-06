// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// ClearCacheOption is a non-required ClearCache option that gets applied to an HTTP request.
type ClearCacheOption func(r *transport.Request)

// WithClearCacheIndex - a comma-separated list of index name to limit the operation.
func WithClearCacheIndex(index []string) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithClearCacheAllowNoIndices(allowNoIndices bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// ClearCacheExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type ClearCacheExpandWildcards int

const (
	// ClearCacheExpandWildcardsOpen can be used to set ClearCacheExpandWildcards to "open"
	ClearCacheExpandWildcardsOpen = iota
	// ClearCacheExpandWildcardsClosed can be used to set ClearCacheExpandWildcards to "closed"
	ClearCacheExpandWildcardsClosed = iota
	// ClearCacheExpandWildcardsNone can be used to set ClearCacheExpandWildcards to "none"
	ClearCacheExpandWildcardsNone = iota
	// ClearCacheExpandWildcardsAll can be used to set ClearCacheExpandWildcards to "all"
	ClearCacheExpandWildcardsAll = iota
)

// WithClearCacheExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithClearCacheExpandWildcards(expandWildcards ClearCacheExpandWildcards) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheFieldData - clear field data.
func WithClearCacheFieldData(fieldData bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheFielddata - clear field data.
func WithClearCacheFielddata(fielddata bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheFields - a comma-separated list of fields to clear when using the "field_data" parameter (default: all).
func WithClearCacheFields(fields []string) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithClearCacheIgnoreUnavailable(ignoreUnavailable bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheIndexParam - a comma-separated list of index name to limit the operation.
func WithClearCacheIndexParam(indexParam []string) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheQuery - clear query caches.
func WithClearCacheQuery(query bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheRecycler - clear the recycler cache.
func WithClearCacheRecycler(recycler bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheRequest - clear request cache.
func WithClearCacheRequest(request bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheRequestCache - clear request cache.
func WithClearCacheRequestCache(requestCache bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheErrorTrace - include the stack trace of returned errors.
func WithClearCacheErrorTrace(errorTrace bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheFilterPath - a comma-separated list of filters used to reduce the respone.
func WithClearCacheFilterPath(filterPath []string) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheHuman - return human readable values for statistics.
func WithClearCacheHuman(human bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheIgnore - ignores the specified HTTP status codes.
func WithClearCacheIgnore(ignore []int) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCachePretty - pretty format the returned JSON response.
func WithClearCachePretty(pretty bool) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// WithClearCacheSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithClearCacheSourceParam(sourceParam string) ClearCacheOption {
	return func(r *transport.Request) {
	}
}

// ClearCache - the clear cache API allows to clear either all caches or specific cached associated with one or more indices. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-clearcache.html for more info.
//
// options: optional parameters.
func (i *Indices) ClearCache(options ...ClearCacheOption) (*ClearCacheResponse, error) {
	req := i.transport.NewRequest("POST")
	for _, option := range options {
		option(req)
	}
	resp, err := i.transport.Do(req)
	return &ClearCacheResponse{resp}, err
}

// ClearCacheResponse is the response for ClearCache.
type ClearCacheResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ClearCacheResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
