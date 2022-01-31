## ----------------------------------------------------------------------
## 		A little manual for using this Makefile.
## ----------------------------------------------------------------------


.PHONY: test
test:  ## Run golang tests
	go test --short -race  -coverprofile=coverage.out -cover `go list ./... | grep -v mocks `


.PHONY: linter
linter:	## Run linter for *.go files
	revive -config .linter.toml  -exclude ./vendor/... -formatter unix ./...


.PHONY: help
help:     ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
