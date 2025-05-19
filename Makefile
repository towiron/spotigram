.PHONY: install-deps lint
.SILENT:



# Dependencies
install-deps:
	@GOBIN=$(CURDIR)/temp/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@GOBIN=$(CURDIR)/temp/bin go install github.com/towiron/import-tidy@latest
	@GOPRIVATE="github.com/towiron" go mod tidy



# Lint
lint:
	@$(CURDIR)/temp/bin/import-tidy --internal-prefix=git.uzinfocom.uz . --fix
	@$(CURDIR)/temp/bin/golangci-lint run -c .golangci.yml --path-prefix . --fix

