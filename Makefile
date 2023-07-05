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
	go generate ./_examples/basic/ent
	go generate ./_examples/custompaths/ent
	$(MAKE) fmt

.PHONY: test
# test:
#    Run the tests
test:
	go test ./_examples/.

validate-tag-arg:
ifeq ("", "$(v)")
	@echo "version arg (v) must be used with the 'tag' target"
	@exit 1;
endif
ifneq ("v", "$(shell echo $(v) | head -c 1)")
	@echo "version arg (v) must begin with v"
	@exit 1;
endif

# ex: make tag v=v0.1.0
tag: validate-tag-arg
	@echo "creating tag $(v)"
	git tag $(v)
	git push origin $(v)
