package artifactory

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		assert.Equal(t, "username", user)
		assert.Equal(t, "password", pass)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "pong")
	}))

	tp := BasicAuthTransport{
		Username: "username",
		Password: "password",
	}

	client, err := NewClient(server.URL, tp.Client())
	assert.Nil(t, err)

	_, _, err = client.System.Ping(context.Background())
	assert.Nil(t, err)
}

func TestTokenAuthTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-JFrog-Art-Api")
		assert.Equal(t, "token", token)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "pong")
	}))

	tp := TokenAuthTransport{
		Token: "token",
	}

	client, err := NewClient(server.URL, tp.Client())
	assert.Nil(t, err)

	_, _, err = client.System.Ping(context.Background())
	assert.Nil(t, err)
}

func TestDo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello World"))
	}))

	client, err := NewClient(server.URL, nil)
	assert.Nil(t, err)

	request, err := client.NewRequest(http.MethodGet, "", nil)
	assert.Nil(t, err)

	resp, err := client.Do(context.TODO(), request, nil)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(body))
	defer resp.Body.Close()
}
