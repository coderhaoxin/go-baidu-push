GO_FILES = ./push/*.go

test:
	@go test $(GO_FILES)

test-util:
	@go test ./push/util*.go

fmt:
	@go fmt $(GO_FILES)

.PHONY: test
