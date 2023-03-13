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

.PHONY: generate
# generate:
#    Generate the examples code
generate:
	go generate ./_examples/ent
	$(MAKE) fmt
