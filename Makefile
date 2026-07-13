GO ?= go
GOVULNCHECK ?= govulncheck
CMD ?=
.PHONY: test benchmark generate-orm tarantool-integration-test update-vulnerable-dependencies verify-dependency-security run vulncheck

test:
	$(GO) test ./...

benchmark:
	RUN_DB_BENCHMARKS=1 $(GO) test -bench=. -benchmem .

generate-orm:
	RUN_TARANTOOL_GENERATOR=1 $(GO) test ./mTest -run TestGenerateOrm

tarantool-integration-test:
	RUN_TARANTOOL_INTEGRATION=1 $(GO) test ./mTest/...

update-vulnerable-dependencies:
	bash ./scripts/update-vulnerable-dependencies.sh

verify-dependency-security:
	GOFLAGS=-mod=mod bash ./scripts/verify-dependency-security.sh

vulncheck:
	$(GOVULNCHECK) ./...

run:
	@test -n "$(CMD)" || (echo "usage: make run CMD='go test ./...'" >&2; exit 2)
	$(CMD)
