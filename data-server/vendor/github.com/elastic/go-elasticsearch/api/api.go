// Package api is the root API package.
package api

// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

import (
	"github.com/elastic/go-elasticsearch/api/cat"
	"github.com/elastic/go-elasticsearch/api/cluster"
	"github.com/elastic/go-elasticsearch/api/indices"
	"github.com/elastic/go-elasticsearch/api/ingest"
	"github.com/elastic/go-elasticsearch/api/nodes"
	"github.com/elastic/go-elasticsearch/api/remote"
	"github.com/elastic/go-elasticsearch/api/snapshot"
	"github.com/elastic/go-elasticsearch/api/tasks"
	"github.com/elastic/go-elasticsearch/transport"
)

// API is the root API client.
type API struct {
	// Cat is the cat client.
	Cat *cat.Cat

	// Cluster is the cluster client.
	Cluster *cluster.Cluster

	// Indices is the indices client.
	Indices *indices.Indices

	// Ingest is the ingest client.
	Ingest *ingest.Ingest

	// Nodes is the nodes client.
	Nodes *nodes.Nodes

	// Remote is the remote client.
	Remote *remote.Remote

	// Snapshot is the snapshot client.
	Snapshot *snapshot.Snapshot

	// Tasks is the tasks client.
	Tasks *tasks.Tasks

	// transport is the transport client.
	transport *transport.Transport
}

// New is the constructor for API. Note that this is automatically invoked by the client.Client type.
func New(transport *transport.Transport) *API {
	return &API{
		Cat:      cat.New(transport),
		Cluster:  cluster.New(transport),
		Indices:  indices.New(transport),
		Ingest:   ingest.New(transport),
		Nodes:    nodes.New(transport),
		Remote:   remote.New(transport),
		Snapshot: snapshot.New(transport),
		Tasks:    tasks.New(transport),

		transport: transport,
	}
}
