.PHONY: build test release clean dist

VERSION ?= 0.1.0
LDFLAGS = -s -w -X main.version=$(VERSION)

build:
	go build -ldflags "$(LDFLAGS)" -o ./dist/ .

# for goreleaser local build artifacts (optional)
dist:
	goreleaser --snapshot --rm-dist

test:
	go test ./...

clean:
	rm -rf dist

release:
	goreleaser release --rm-dist

