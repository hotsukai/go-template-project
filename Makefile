## init
.PHONY: init
init:
	cp -n .env.example .env
	go install golang.org/x/tools/cmd/goimports@v0.12.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
	go install github.com/cweill/gotests/gotests@latest
## fmt
.PHONY: fmt
fmt:
	goimports -w -local "sample" cmd/ pkg/
	gofmt -s -w cmd/ pkg/

## lint
.PHONY: lint
lint:
	golangci-lint run -v cmd/... pkg/...

## generate-test
.PHONY: gen-test
gen-test:
	gotests -w -i -all -excl="New.*" -parallel  ./*


## test
.PHONY: test
test:
	 go test -cover ./pkg/...
