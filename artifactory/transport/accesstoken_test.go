package transport

import (
	"context"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/artifactory"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccessTokenTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		assert.Equal(t, "Bearer sometoken", token)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "pong")
	}))

	tp := AccessTokenAuth{
		AccessToken: "sometoken",
	}

	rt, err := artifactory.NewClient(server.URL, tp.Client())
	assert.Nil(t, err)

	_, _, err = rt.V1.System.Ping(context.Background())
	assert.Nil(t, err)
}
