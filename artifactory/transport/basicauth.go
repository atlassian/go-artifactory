package transport

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// basicAuth allows the construction of a HTTP Client that authenticates with basic auth
// It also adds the correct headers to the request
type basicAuth struct {
	username string
	password string
	tp       http.RoundTripper
}

func BasicAuth(username string, password string) http.RoundTripper {
	return &basicAuth{
		username: username,
		password: password,
		tp:       http.DefaultTransport,
	}
}

// RoundTrip allows us to add headers to every request
func (t *basicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of service.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	deepCopyRequest(req, req2)

	req2.SetBasicAuth(t.username, t.password)
	req2.Header.Add(HeaderResultDetail, "info, properties")

	if req.Body != nil {
		reader, _ := req.GetBody()
		buf, _ := ioutil.ReadAll(reader)
		chkSum := getSha1(buf)
		req.Header.Add(HeaderChecksumSha1, fmt.Sprintf("%x", chkSum))
	}

	return t.tp.RoundTrip(req2)
}
