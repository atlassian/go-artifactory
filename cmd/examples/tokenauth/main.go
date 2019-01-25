package main

import (
	"context"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory"
	"github.com/atlassian/go-artifactory/v2/pkg/artifactory/transport"
	"os"
)

func main() {
	tp := transport.TokenAuth{
		Token: os.Getenv("ARTIFACTORY_TOKEN"),
	}

	client, err := artifactory.NewClient(os.Getenv("ARTIFACTORY_URL"), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	_, _, err = client.System.Ping(context.Background())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
	} else {
		fmt.Println("OK")
	}
}
