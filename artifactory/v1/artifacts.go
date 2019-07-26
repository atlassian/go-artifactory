package v1

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// ArtifactService exposes the Artifact API endpoints from Artifactory
type ArtifactService Service

// SingleReplicationConfig is the model of the Artifactory Replication Config
type SingleReplicationConfig struct {
	RepoKey                *string `json:"repoKey,omitempty"`
	URL                    *string `json:"url,omitempty"`
	SocketTimeoutMillis    *int    `json:"socketTimeoutMillis,omitempty"`
	Username               *string `json:"username,omitempty"`
	Password               *string `json:"password,omitempty"`
	Enabled                *bool   `json:"enabled,omitempty"`
	SyncDeletes            *bool   `json:"syncDeletes,omitempty"`
	SyncProperties         *bool   `json:"syncProperties,omitempty"`
	SyncStatistics         *bool   `json:"syncStatistics,omitempty"`
	PathPrefix             *string `json:"pathPrefix,omitempty"`
	CronExp                *string `json:"cronExp,omitempty"` // Only required when getting list of repositories as C*UD operations will be done through a repConfig obj
	EnableEventReplication *bool   `json:"enableEventReplication,omitempty"`
}

func (r SingleReplicationConfig) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// ReplicationConfig is the model for the multi replication config API endpoints. Its usage is preferred over
// SingleReplicationConfig as it is a more direct mapping of the replicationConfig in the UI
type ReplicationConfig struct {
	RepoKey                *string                    `json:"-"`
	CronExp                *string                    `json:"cronExp,omitempty"`
	EnableEventReplication *bool                      `json:"enableEventReplication,omitempty"`
	Replications           *[]SingleReplicationConfig `json:"replications,omitempty"`
}

func (r ReplicationConfig) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// Creates or replaces a local multi-push replication configuration. Supported by local repositories.
// Notes: Requires an enterprise license
// Security: Requires a privileged user
func (s *ArtifactService) SetRepositoryReplicationConfig(ctx context.Context, repoKey string, config *ReplicationConfig) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		SetBody(config).
		Put("/api/replications/multiple/" + repoKey)
}

// Description: Add or replace replication configuration for given repository key. Supported by local and remote repositories. Accepts the JSON payload returned from Get Repository Replication Configuration for a single and an array of configurations. If the payload is an array of replication configurations, then values for cronExp and enableEventReplication in the first element in the array will determine the corresponding values when setting the repository replication configuration.
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *ArtifactService) SetSingleRepositoryReplicationConfig(ctx context.Context, repoKey string, config *SingleReplicationConfig) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		SetBody(config).
		Put("/api/replications/" + repoKey)
}

// Returns the replication configuration for the given repository key, if found. Supported by local and remote repositories. Note: The 'enableEventReplication' parameter refers to both push and pull replication.
// Notes: Requires Artifactory Pro
// Security: Requires a privileged user
func (s *ArtifactService) GetRepositoryReplicationConfig(ctx context.Context, repoKey string) (*[]SingleReplicationConfig, *resty.Response, error) {
	replications := make([]SingleReplicationConfig, 0)
	resp, err := s.client.R().
		SetContext(ctx).
		SetHeader("Accept", mediaTypeReplicationConfig).
		SetResult(&replications).
		Get("/api/replications/" + repoKey)
	return &replications, resp, err
}
