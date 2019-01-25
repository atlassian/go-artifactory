package transport

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// TokenAuth exposes a HTTP Client which uses this transport. It authenticates via an Artifactory API token
// It also adds the correct headers to the request
type TokenAuth struct {
	Token     string
	Transport http.RoundTripper
}

// Client returns a HTTP Client and injects the token auth transport
func (t *TokenAuth) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenAuth) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// RoundTrip allows us to add headers to every request
func (t *TokenAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of service.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	deepCopyRequest(req, req2)

	req2.Header.Set(HeaderApiToken, t.Token)
	req2.Header.Add(HeaderResultDetail, "info, properties")

	if req.Body != nil {
		reader, _ := req.GetBody()
		buf, _ := ioutil.ReadAll(reader)
		chkSum := getSha1(buf)
		req.Header.Add(HeaderChecksumSha1, fmt.Sprintf("%x", chkSum))
	}

	return t.transport().RoundTrip(req2)
}
