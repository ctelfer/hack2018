// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// DeleteOption is a non-required Delete option that gets applied to an HTTP request.
type DeleteOption func(r *transport.Request)

// WithDeleteMasterTimeout - explicit operation timeout for connection to master node.
func WithDeleteMasterTimeout(masterTimeout time.Duration) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteErrorTrace - include the stack trace of returned errors.
func WithDeleteErrorTrace(errorTrace bool) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteFilterPath - a comma-separated list of filters used to reduce the respone.
func WithDeleteFilterPath(filterPath []string) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteHuman - return human readable values for statistics.
func WithDeleteHuman(human bool) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteIgnore - ignores the specified HTTP status codes.
func WithDeleteIgnore(ignore []int) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeletePretty - pretty format the returned JSON response.
func WithDeletePretty(pretty bool) DeleteOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithDeleteSourceParam(sourceParam string) DeleteOption {
	return func(r *transport.Request) {
	}
}

// Delete - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a snapshot name.
//
// options: optional parameters.
func (s *Snapshot) Delete(repository string, snapshot string, options ...DeleteOption) (*DeleteResponse, error) {
	req := s.transport.NewRequest("DELETE")
	for _, option := range options {
		option(req)
	}
	resp, err := s.transport.Do(req)
	return &DeleteResponse{resp}, err
}

// DeleteResponse is the response for Delete.
type DeleteResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *DeleteResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
