package transport

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ApiKeyAuth exposes a HTTP Client which uses this transport. It authenticates via an Artifactory API token
// It also adds the correct headers to the request
type accessTokenAuth struct {
	accessToken string
	tp          http.RoundTripper
}

func AccessTokenAuth(token string) http.RoundTripper {
	return &accessTokenAuth{
		accessToken: token,
		tp:          http.DefaultTransport,
	}
}

// RoundTrip allows us to add headers to every request
func (t *accessTokenAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of service.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	deepCopyRequest(req, req2)

	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.accessToken))
	req2.Header.Add(HeaderResultDetail, "info, properties")

	if req.Body != nil {
		reader, _ := req.GetBody()
		buf, _ := ioutil.ReadAll(reader)
		chkSum := getSha1(buf)
		req.Header.Add(HeaderChecksumSha1, fmt.Sprintf("%x", chkSum))
	}

	return t.tp.RoundTrip(req2)
}
