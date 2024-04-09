.PHONY: help
# help:
#    Print this help message
help:
	@grep -o '^\#.*' Makefile | cut -d" " -f2-

.PHONY: fmt
# fmt:
#    Format go code
fmt:
	goimports -local github.com/flume -w ./

.PHONY: lint
# lint:
#    lint the code
lint:
	golangci-lint run
	golangci-lint run ./_examples/.

.PHONY: generate
# generate:
#    Generate the examples code
generate:
	(cd _examples && go generate ./... )
	$(MAKE) fmt

.PHONY: upgrade-deps
# upgrade-deps:
#    Upgrade the dependencies
upgrade-deps:
	(cd _examples && go get -u -t ./... && go mod tidy)
	(go get -u -t ./... && go mod tidy)
	go work sync ./...

.PHONY: test
# test:
#    Run the tests
test:
	(cd ./_examples && go test ./.)

tag:
	@echo "creating tag"
	bash ./scripts/tag.sh
