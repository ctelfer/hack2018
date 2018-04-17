// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package nodes

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// InfoOption is a non-required Info option that gets applied to an HTTP request.
type InfoOption func(r *transport.Request)

// WithInfoMetric - a comma-separated list of metrics you wish returned. Leave empty to return all.
func WithInfoMetric(metric []string) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoNodeID - a comma-separated list of node IDs or names to limit the returned information; use "_local" to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithInfoNodeID(nodeID []string) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoFlatSettings - return settings in flat format (default: false).
func WithInfoFlatSettings(flatSettings bool) InfoOption {
	return func(r *transport.Request) {
	}
}

// WithInfoTimeout - explicit operation timeout.
func WithInfoTimeout(timeout time.Duration) InfoOption {
	return func(r *transport.Request) {
	}
}

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

// Info - the cluster nodes info API allows to retrieve one or more (or all) of the cluster nodes information. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-nodes-info.html for more info.
//
// options: optional parameters.
func (n *Nodes) Info(options ...InfoOption) (*InfoResponse, error) {
	req := n.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := n.transport.Do(req)
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
