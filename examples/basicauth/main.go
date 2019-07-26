package main

import (
	"context"
	"fmt"

	"github.com/atlassian/go-artifactory/v3/artifactory"
	"github.com/atlassian/go-artifactory/v3/artifactory/transport"
)

func main() {
	rt := artifactory.NewClient("http://localhost:8080/artifactory", transport.BasicAuth("admin", "password"))

	resp, err := rt.V1.System.Ping(context.Background())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	} else {
		fmt.Println(resp.String())
	}
}
