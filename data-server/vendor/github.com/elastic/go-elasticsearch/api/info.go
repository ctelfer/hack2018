// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// InfoOption is a non-required Info option that gets applied to an HTTP request.
type InfoOption func(r *transport.Request)

// WithInfoErrorTrace - include the stack trace of returned errors.
func WithInfoErrorTrace(errorTrace bool) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoFilterPath - a comma-separated list of filters used to reduce the respone.
func WithInfoFilterPath(filterPath []string) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoHuman - return human readable values for statistics.
func WithInfoHuman(human bool) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoIgnore - ignores the specified HTTP status codes.
func WithInfoIgnore(ignore []int) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoPretty - pretty format the returned JSON response.
func WithInfoPretty(pretty bool) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithInfoSourceParam(sourceParam string) InfoOption {
	return func(r *transport.Request) {
	}
}

// Info - see https://www.elastic.co/guide/ for more info.
//
// options: optional parameters.
func (a *API) Info(options ...InfoOption) (*InfoResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &InfoResponse{resp}, err
}

// InfoResponse is the response for Info.
type InfoResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *InfoResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}