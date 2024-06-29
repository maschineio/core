.DEFAULT_GOAL := help

VERSION=`git describe --tags`

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## runs `go generate` to build the dynamically generated source files
	@make stringer
	@go generate ./...
	# @go run -mod=mod github.com/google/addlicense -c "Thomas Meitz" -y 2023-present ./


.PHONY: stringer
stringer: ## generate stringer for datatypes
	@stringer -type=TokenType token/token-type.go	
	@stringer -type=DataType token/data-type.go
	@stringer -type=StateType token/state-type.go
	@stringer -type=Type type.go

.PHONY: godoc
godoc: ## godoc server listens on port 6060
	@godoc -http=:6060 -index -links=true

.PHONY: test
test: ## runs all tests
	@go test -v -p 1 ./...

.PHONY: token
token: ## runs parser test
	@go test -v -timeout 30s -run ^TestToken_ maschine.io/core/token

.PHONY: coverage
coverage: ## generate code coverage for this project
	@golangci-lint run --issues-exit-code=0 --out-format checkstyle > coverage/golangci-lint.xml
	@gosec -fmt=sonarqube -no-fail -out=coverage/gosec-report.json ./...
	@go vet ./... > coverage/vet.out
	@go test -coverprofile=coverage/cov.out ./...
	@go test -json ./...  > coverage/report.json

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:	