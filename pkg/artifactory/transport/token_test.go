package transport

import (
	"context"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenAuthTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-JFrog-Art-Api")
		assert.Equal(t, "token", token)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "pong")
	}))

	tp := TokenAuth{
		Token: "token",
	}

	rt, err := artifactory.NewClient(server.URL, tp.Client())
	assert.Nil(t, err)

	_, _, err = rt.System.Ping(context.Background())
	assert.Nil(t, err)
}
