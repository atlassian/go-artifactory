package artifactory

import (
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory/client"
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory/v1"
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory/v2"
	"net/http"
)

// Artifactory is the container for all the api methods
type Artifactory struct {
	// HTTP Artifactory used to communicate with the API.
	client *client.Client

	// Reuse a single struct instead of allocating one for each Service on the heap.
	common client.Service

	V2 *v2.V2

	// Services used for talking to different parts of the Artifactory API.
	Repositories *v1.RepositoriesService
	Security     *v1.SecurityService
	System       *v1.SystemService
	Artifacts    *v1.ArtifactService
}

// NewClient creates a Artifactory from a provided base url for an artifactory instance and a service Artifactory
func NewClient(baseURL string, httpClient *http.Client) (*Artifactory, error) {
	c, err := client.NewClient(baseURL, httpClient)

	if err != nil {
		return nil, err
	}

	rt := &Artifactory{
		client: c,
	}

	rt.common.Client = c

	rt.Repositories = (*v1.RepositoriesService)(&rt.common)
	rt.Security = (*v1.SecurityService)(&rt.common)
	rt.System = (*v1.SystemService)(&rt.common)
	rt.Artifacts = (*v1.ArtifactService)(&rt.common)

	rt.V2 = &v2.V2{
		Security: (*v2.SecurityService)(&rt.common),
	}

	return rt, nil
}
