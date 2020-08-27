# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test -v -mod=vendor

test:
	CGO_ENABLED=0 $(GOTEST) ./...
