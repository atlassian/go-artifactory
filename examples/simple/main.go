// Retrieves list of all repositories for an artifactory instance
package main

import (
	"context"
	"fmt"

	"github.com/atlassian/go-artifactory/v3/artifactory"
	"github.com/atlassian/go-artifactory/v3/artifactory/transport"
	v1 "github.com/atlassian/go-artifactory/v3/artifactory/v1"
)

func main() {
	client := artifactory.NewClient("http://localhost:8080/artifactory", transport.BasicAuth("admin", "password"))
	repos, resp, err := client.V1.Repositories.ListRepositories(context.Background(), &v1.RepositoryListOptions{
		Type: "local",
	})

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	} else if resp.IsError() {
		fmt.Printf("\nerror: %s - %s\n", resp.Status(), resp)
		return
	}

	fmt.Println("Found these local repos:")
	for _, repo := range *repos {
		fmt.Println(repo.Key)
	}
}
