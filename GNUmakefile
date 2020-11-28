BINARY := terraform-provider-jsonschema
SOURCES := $(wildcard *.go) $(wildcard jsonschema/*.go)

default: test $(BINARY)

$(BINARY): $(SOURCES)
	go build -o $(BINARY)

test: $(SOURCES)
	TF_ACC=1 go test ./...

release:
	goreleaser release --rm-dist

.PHONY: test release
