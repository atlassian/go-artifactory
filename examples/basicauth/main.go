package main

import (
	"context"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/artifactory"
	"github.com/atlassian/go-artifactory/v2/artifactory/transport"
)

func main() {
	tp := transport.BasicAuth{
		Username: "admin",
		Password: "password",
	}

	rt, err := artifactory.NewClient("http://localhost:8080/artifactory", tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	_, _, err = rt.V1.System.Ping(context.Background())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	} else {
		fmt.Println("OK")
	}
}
