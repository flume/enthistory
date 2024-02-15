.PHONY: help
# help:
#    Print this help message
help:
	@grep -o '^\#.*' Makefile | cut -d" " -f2-

.PHONY: fmt
# fmt:
#    Format go code
fmt:
	(for x in $$(git status -s | awk '$$2 ~ /\.go$$/ { print $$2 }'); do goimports -local github.com/flume -w $$x; done) || true
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
	go generate ./_examples/basic/ent
	go generate ./_examples/custompaths/ent
	go generate ./_examples/updateby_uuid/ent
	go generate ./_examples/graphql
	$(MAKE) fmt


.PHONY: test
# test:
#    Run the tests
test:
	(cd ./_examples && go test ./.)

tag:
	@echo "creating tag"
	bash ./scripts/tag.sh
