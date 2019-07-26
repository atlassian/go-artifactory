package artifactory

import (
	client "github.com/atlassian/go-artifactory/v3/artifactory/client"
	v1 "github.com/atlassian/go-artifactory/v3/artifactory/v1"
	v2 "github.com/atlassian/go-artifactory/v3/artifactory/v2"

	"net/http"
)

// Artifactory is the container for all the api methods
type Artifactory struct {
	V1 *v1.V1
	V2 *v2.V2
}

// NewClient creates a Artifactory from a provided base url for an artifactory instance and a service Artifactory
func NewClient(baseURL string, transport http.RoundTripper) *Artifactory {
	c := client.NewClient(baseURL, &http.Client{Transport: transport}).SetRetryCount(5)
	return &Artifactory{v1.NewV1(c), v2.NewV2(c)}
}
