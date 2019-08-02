package artifactory

import (
	"net/http"

	client "github.com/atlassian/go-artifactory/v3/artifactory/client"
	v1 "github.com/atlassian/go-artifactory/v3/artifactory/v1"
	v2 "github.com/atlassian/go-artifactory/v3/artifactory/v2"
	"github.com/go-resty/resty/v2"
)

// Artifactory is the container for all the api methods
type Artifactory struct {
	V1 *v1.V1
	V2 *v2.V2
}

// NewClient creates a Artifactory from a provided base url for an artifactory instance and a service Artifactory
func NewClient(baseURL string, transport http.RoundTripper) *Artifactory {
	c := client.NewClient(baseURL, &http.Client{Transport: transport}).
		SetRetryCount(5).
		SetPreRequestHook(func(client *resty.Client, req *http.Request) error {
			// Workaround as Artifactory does not support charset in Content-Type
			// bug https://www.jfrog.com/jira/browse/RTFACT-14981
			// upstream fix https://github.com/go-resty/resty/issues/258
			if req.Header.Get("Content-Type") == "application/json; charset=utf-8" {
				req.Header.Set("Content-Type", "application/json")
			}
			return nil
		})
	return &Artifactory{v1.NewV1(c), v2.NewV2(c)}
}
