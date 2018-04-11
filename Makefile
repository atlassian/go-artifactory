all: install-hooks test

install-hooks:
	@misc/scripts/install-hooks

dep:
	@misc/scripts/deps-ensure
	@dep ensure -v

test:
	@go test -v ./pkg/...
