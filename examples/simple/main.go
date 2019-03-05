// Retrieves list of all repositories for an artifactory instance
package main

import (
	"context"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/artifactory"
	"github.com/atlassian/go-artifactory/v2/artifactory/transport"
	"github.com/atlassian/go-artifactory/v2/artifactory/v1"
	"os"
)

func main() {
	tp := transport.BasicAuth{
		Username: os.Getenv("ARTIFACTORY_USERNAME"),
		Password: os.Getenv("ARTIFACTORY_PASSWORD"),
	}

	client, err := artifactory.NewClient(os.Getenv("ARTIFACTORY_URL"), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	opts := v1.RepositoryListOptions{
		Type: "local",
	}
	repos, _, err := client.V1.Repositories.ListRepositories(context.Background(), &opts)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	} else if repos == nil {
		fmt.Printf("\nerror: repos cannot be nil\n")
		return
	}

	fmt.Println("Found these local repos:")
	for _, repo := range *repos {
		fmt.Println(repo.Key)
	}
}
