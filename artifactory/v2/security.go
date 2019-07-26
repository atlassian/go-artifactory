package v2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SecurityService Service

// read, write, annotate, delete, manage
const (
	PERM_READ     = "read"
	PERM_WRITE    = "write"
	PERM_ANNOTATE = "annotate"
	PERM_DELETE   = "delete"
	PERM_MANAGE   = "manage"

	PERMISSION_SCHEMA = "application/vnd.org.jfrog.artifactory.security.PermissionTargetV2+json"
)

type PermissionTargetDetails struct {
	Name *string `json:"name,omitempty"`
	Uri  *string `json:"uri,omitempty"`
}

func (r PermissionTargetDetails) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

func (s *SecurityService) ListPermissionTargets(ctx context.Context) (*[]PermissionTargetDetails, *resty.Response, error) {
	permissions := make([]PermissionTargetDetails, 0)
	resp, err := s.client.R().
		SetContext(ctx).
		SetResult(&permissions).
		Get(fmt.Sprintf("/api/v2/security/permissions"))
	return &permissions, resp, err
}

type Entity struct {
	Users  *map[string][]string `json:"users,omitempty"`
	Groups *map[string][]string `json:"groups,omitempty"`
}

type Permission struct {
	IncludePatterns *[]string `json:"include-patterns,omitempty"`
	ExcludePatterns *[]string `json:"exclude-patterns,omitempty"`
	Repositories    *[]string `json:"repositories,omitempty"`
	Actions         *Entity   `json:"actions,omitempty"`
}

type PermissionTarget struct {
	Name  *string     `json:"name,omitempty"` // Optional element in create/replace queries
	Repo  *Permission `json:"repo,omitempty"`
	Build *Permission `json:"build,omitempty"`
}

func (r PermissionTarget) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

func (s *SecurityService) CreatePermissionTarget(ctx context.Context, permissionName string, permissionTargets *PermissionTarget) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		SetBody(permissionTargets).
		Post(fmt.Sprintf("/api/v2/security/permissions/%s", permissionName))
}

func (s *SecurityService) GetPermissionTarget(ctx context.Context, permissionName string) (*PermissionTarget, *resty.Response, error) {
	permission := new(PermissionTarget)
	resp, err := s.client.R().
		SetContext(ctx).
		SetResult(permission).
		Get(fmt.Sprintf("/api/v2/security/permissions/%s", permissionName))
	return permission, resp, err
}

func (s *SecurityService) HasPermissionTarget(ctx context.Context, permissionName string) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		Head(fmt.Sprintf("/api/v2/security/permissions/%s", permissionName))
}

// Missing permission target values will be set to the default values as defined by the consumed type.
// The values defined in the request payload will replace what currently exists in the permission target entity.
// In case the request is missing one of the permission target entities (repo/build), the entity will be deleted.
// This means that if an update request is sent to an entity that contains both repo and build, with only repo,
// the build values will be removed from the entity.
func (s *SecurityService) UpdatePermissionTarget(ctx context.Context, permissionName string, permissionTargets *PermissionTarget) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		SetBody(permissionTargets).
		Put(fmt.Sprintf("/api/v2/security/permissions/%s", permissionName))
}

func (s *SecurityService) DeletePermissionTarget(ctx context.Context, permissionName string) (*resty.Response, error) {
	return s.client.R().
		SetContext(ctx).
		Delete(fmt.Sprintf("/api/v2/security/permissions/%s", permissionName))
}
