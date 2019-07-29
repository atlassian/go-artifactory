# go-artifactory #
go-artifactory is a Go client library for accessing the [Artifactory API](https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API)

[![Build Status](https://travis-ci.org/atlassian/go-artifactory.svg?branch=master)](https://travis-ci.org/atlassian/go-artifactory)

## V3 Breaking Changes ##
go-artifactory now depends on [resty](https://github.com/go-resty/resty/). This
was done to consolidate all of the request logic throughout the client. There 
was significant inconsistencies with the behaviour of the return objects, some
methods returned client.Status, others returned errors on 4/5xx. The new 
behaviour is now all methods return `*resty.Response` and `error`. error will be
non-nil if the request fails. If err is nil than the response is non-nil but may
still be 4/5xx. If response is 2xx than methods that return a body will return a
non-nil body.

Methods that used to parse the body as a string now do not return anything. The
body can be fetched with `resp.String()`.

Response checking should look like:
```go
repo, resp, err := client.V1.Repositories.GetRepository(context.Background(), "my-repo")
if err != nil {
	fmt.Printf("request failed: %s\n", err)
} else if resp.IsError() { // do on 4/5xx
	fmt.Printf("response non-successful: %s\n", resp.Status())
} else {
	fmt.Println("got body: %v", repo)
}
```

## Requirements ##
- Go version 1.12+
- Go mod is used for dependency management

## Usage ##
```go
import "github.com/atlassian/go-artifactory/v3/artifactory"
```

Construct a new Artifactory client, then use the various services on the client to
access different parts of the Artifactory API. For example:

```go
client := artifactory.NewClient(client, nil)

// list all repositories
repos, resp, err := client.V1.Repositories.List(context.Background(), nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
client := artifactroy.NewClient(client, nil)

// list all public local repositories
opt := &artifactory.RepositoryListOptions{Type: "local"}
client.V1.Repositories.ListRepositories(ctx, opt)
```

The services of a client divide the API into logical chunks and correspond to
the structure of the Artifactory API documentation at
[https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API](https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API).

NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

### Authentication ###

The go-artifactory library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. 

For API methods that require HTTP Basic Authentication, use the BasicAuthTransport or TokenTransport

```go
package main

import (
	"github.com/atlassian/go-artifactory/pkg/artifactory"
	"fmt"
	"context"
)

func main() {
	tp := artifactory.BasicAuthTransport{
		Username: "<YOUR_USERNAME>",
		Password: "<YOUR_PASSWORD>",
	}
	
	client, err := artifactory.NewClient("https://localhost/artifactory", tp.Client())
	if err != nil {
		fmt.Println(err.Error())
	}

	repos, resp, err := client.V1.Repositories.ListRepositories(context.Background(), nil)
}
```

### Creating and Updating Resources ###
All structs for GitHub resources use pointer values for all non-repeated fields.
This allows distinguishing between unset fields and those set to a zero-value.
Helper functions have been provided to easily create these pointers for string,
bool, and int values. For example:

```go
    // create a new local repository named "lib-releases"
    repo := artifactory.LocalRepository{
		Key:             artifactory.String("lib-releases"),
		RClass:          artifactory.String("local"),
		PackageType:     artifactory.String("maven"),
		HandleSnapshots: artifactory.Bool(false);
	}

	client.V1.Repositories.CreateLocal(context.Background(), &repo)
```

Users who have worked with protocol buffers should find this pattern familiar.

## Roadmap ##

This library is being initially developed for an internal application at
Atlassian, so API methods will likely be implemented in the order that they are
needed by that application. Eventually, it would be ideal to cover the entire
Artifactory API, so contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward.

## Versioning ##

In general, go-artifactory follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package. For self-contained libraries, the
application of semantic versioning is relatively straightforward and generally
understood. But because go-artifactory is a client library for the Artifactory API 
we've adopted the following versioning policy:

* We increment the **major version** with any incompatible change to
	functionality, including changes to the exported Go API surface
	or behavior of the API.
* We increment the **minor version** with any backwards-compatible changes to
	functionality.
* We increment the **patch version** with any backwards-compatible bug fixes.

Generally methods will be annotated with a since version.

## Reporting issues ##

We believe in open contributions and the power of a strong development community. Please read our 
[Contributing guidelines](.github/CONTRIBUTING.md) on how to contribute back and report issues to go-artifactory.

## Contributors ##

Pull requests, issues and comments are welcomed. For pull requests:

* Add tests for new features and bug fixes
* Follow the existing style
* Separate unrelated changes into multiple pull requests
* Read [Contributing guidelines](.github/CONTRIBUTING.md) for more details

See the existing issues for things to start contributing.

For bigger changes, make sure you start a discussion first by creating
an issue and explaining the intended change.

Atlassian requires contributors to sign a Contributor License Agreement,
known as a CLA. This serves as a record stating that the contributor is
entitled to contribute the code/documentation/translation to the project
and is willing to have it used in distributions and derivative works
(or is willing to transfer ownership).

Prior to accepting your contributions we ask that you please follow the appropriate
link below to digitally sign the CLA. The Corporate CLA is for those who are
contributing as a member of an organization and the individual CLA is for
those contributing as an individual.

* [CLA for corporate contributors](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b)
* [CLA for individuals](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d)


## License ##
Copyright (c) 2019 Atlassian and others. Apache 2.0 licensed, see [LICENSE][LICENSE] file.


[CONTRIBUTING]: .github/CONTRIBUTING.md
[LICENSE]: ./LICENSE.txt
