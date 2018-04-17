// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetOption is a non-required Get option that gets applied to an HTTP request.
type GetOption func(r *transport.Request)

// WithGetIgnoreUnavailable - whether to ignore unavailable snapshots, defaults to false which means a SnapshotMissingException is thrown.
func WithGetIgnoreUnavailable(ignoreUnavailable bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetMasterTimeout - explicit operation timeout for connection to master node.
func WithGetMasterTimeout(masterTimeout time.Duration) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetErrorTrace - include the stack trace of returned errors.
func WithGetErrorTrace(errorTrace bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetFilterPath - a comma-separated list of filters used to reduce the respone.
func WithGetFilterPath(filterPath []string) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetHuman - return human readable values for statistics.
func WithGetHuman(human bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetIgnore - ignores the specified HTTP status codes.
func WithGetIgnore(ignore []int) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetPretty - pretty format the returned JSON response.
func WithGetPretty(pretty bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithGetSourceParam(sourceParam string) GetOption {
	return func(r *transport.Request) {
	}
}

// Get - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a comma-separated list of snapshot names.
//
// options: optional parameters.
func (s *Snapshot) Get(repository string, snapshot []string, options ...GetOption) (*GetResponse, error) {
	req := s.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := s.transport.Do(req)
	return &GetResponse{resp}, err
}

// GetResponse is the response for Get.
type GetResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}