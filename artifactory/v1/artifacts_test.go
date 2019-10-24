package v1

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/atlassian/go-artifactory/v2/artifactory/client"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"strings"
)

func TestFileInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/storage/arbitrary-repository/path/to/an/existing/artifact", r.RequestURI)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
    	
		dummyRes := `{
  "repo" : "arbitrary-repository",
  "path" : "/path/to/an/existing/artifact",
  "created" : "2019-10-22T07:12:08.538+02:00",
  "createdBy" : "jondoe",
  "lastModified" : "2019-10-22T09:38:55.713+02:00",
  "modifiedBy" : "janedoe",
  "lastUpdated" : "2019-10-22T09:38:55.731+02:00",
  "downloadUri" : "http://%s/arbitrary-repository/path/to/an/existing/artifact",
  "mimeType" : "application/zip",
  "size" : "13400",
  "checksums" : {
    "sha1" : "1bc68542d65869e38eece7cfb1b038104ba7a5fb",
    "md5" : "ccb552c5b0714ced4852c8d696da3387",
    "sha256" : "3a4d369251cdd78d616873e1eb7352f83997949969b020397fabc6e2d18801b9"
  },
  "originalChecksums" : {
    "sha1" : "1bc68542d65869e38eece7cfb1b038104ba7a5fb",
    "md5" : "ccb552c5b0714ced4852c8d696da3387",
    "sha256" : "3a4d369251cdd78d616873e1eb7352f83997949969b020397fabc6e2d18801b9"
  },
  "uri" : "%s"
}`

		_, _ = fmt.Fprint(w, fmt.Sprintf(dummyRes, r.Host, r.RequestURI))
	}))

	c, _ := client.NewClient(server.URL, http.DefaultClient);
	v := NewV1(c);

	fileInfo, _, err := v.Artifacts.FileInfo(context.Background(), "arbitrary-repository", "/path/to/an/existing/artifact");
	assert.Nil(t, err)

	assert.Equal(t, "arbitrary-repository",                      *fileInfo.Repo)
	assert.Equal(t, "/path/to/an/existing/artifact",             *fileInfo.Path)
	assert.Equal(t, "2019-10-22T07:12:08.538+02:00",             *fileInfo.Created)
	assert.Equal(t, "jondoe",                                    *fileInfo.CreatedBy)
	assert.Equal(t, "2019-10-22T09:38:55.713+02:00",             *fileInfo.LastModified)
	assert.Equal(t, "janedoe",                                   *fileInfo.ModifiedBy)
	assert.Equal(t, "2019-10-22T09:38:55.731+02:00",             *fileInfo.LastUpdated)
	assert.Equal(t, fmt.Sprintf("%s/arbitrary-repository/path/to/an/existing/artifact", server.URL), *fileInfo.DownloadUri)
	assert.Equal(t, "application/zip",                           *fileInfo.MimeType)
	assert.Equal(t, 13400,                                       *fileInfo.Size)
	assert.Equal(t, "1bc68542d65869e38eece7cfb1b038104ba7a5fb",  *fileInfo.Checksums.Sha1)
	assert.Equal(t, "1bc68542d65869e38eece7cfb1b038104ba7a5fb",  *fileInfo.OriginalChecksums.Sha1)
	assert.Equal(t, "ccb552c5b0714ced4852c8d696da3387",          *fileInfo.Checksums.Md5)
	assert.Equal(t, "ccb552c5b0714ced4852c8d696da3387",          *fileInfo.OriginalChecksums.Md5)
	assert.Equal(t, "3a4d369251cdd78d616873e1eb7352f83997949969b020397fabc6e2d18801b9", *fileInfo.Checksums.Sha256)
	assert.Equal(t, "3a4d369251cdd78d616873e1eb7352f83997949969b020397fabc6e2d18801b9", *fileInfo.OriginalChecksums.Sha256)
	assert.Equal(t, "/api/storage/arbitrary-repository/path/to/an/existing/artifact", *fileInfo.Uri)
}

func TestFileContents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res := ""

		if strings.HasSuffix(r.RequestURI, "/content") {
			w.Header().Set("Content-Type", "text/plain")
			res = "dummy content"
		} else {
			w.Header().Set("Content-Type", "application/json")

			res = fmt.Sprintf(`{ "downloadUri" : "http://%s%s/content" }`, r.Host, r.RequestURI)
		}

		_, _ = fmt.Fprint(w, res)
	}))

	c, _ := client.NewClient(server.URL, http.DefaultClient);
	v := NewV1(c);

	target := bytes.NewBufferString("")
	fileInfo, _, err := v.Artifacts.FileContents(context.Background(), "arbitrary-repository", "/path/to/an/existing/artifact", target);

	assert.Equal(t, "dummy content", target.String())
	assert.NotNil(t, fileInfo)
	assert.Nil(t, err)
}