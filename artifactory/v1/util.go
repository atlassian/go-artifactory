package v1

import "github.com/go-resty/resty/v2"

func String(v string) *string { return &v }

func NewV1(client *resty.Client) *V1 {
	v := &V1{}
	v.common.client = client

	v.Repositories = (*RepositoriesService)(&v.common)
	v.Security = (*SecurityService)(&v.common)
	v.System = (*SystemService)(&v.common)
	v.Artifacts = (*ArtifactService)(&v.common)

	return v
}
