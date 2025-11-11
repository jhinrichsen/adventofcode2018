GO ?= CGO_ENABLED=0 go
BENCH_FILE ?= benches/$(shell go env GOOS)-$(shell go env GOARCH)-$(shell lscpu | grep "Model name:" | cut -d: -f2 | xargs | sed 's/ \(CPU\|@\|w\/\).*//' | tr -d '()' | sed 's/ /_/g').txt

.PHONY: all
all: clean lint test

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	rm -f \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		govulncheck.sarif \
		junit.xml \
		README.html \
		staticcheck.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -run=^$ -bench=. -benchmem

$(BENCH_FILE):
	@echo "Running benchmarks and saving to $@..."
ifeq ($(shell go env GOOS),linux)
	@if [ -d /sys/devices/system/cpu/cpu0/cpufreq ]; then \
		SAVED_GOV=$$(cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor); \
		echo "Saving current CPU governor: $$SAVED_GOV"; \
		echo "Setting CPU governor to performance mode..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo performance | sudo tee $$cpu > /dev/null; \
		done; \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
		echo "Restoring CPU governor to $$SAVED_GOV..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo $$SAVED_GOV | sudo tee $$cpu > /dev/null; \
		done; \
		echo "CPU governor restored to $$SAVED_GOV"; \
	else \
		echo "CPU frequency scaling not available, running benchmarks normally..."; \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
	fi
else
	@$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@
endif

.PHONY: lint
lint:
	test -z $(gofmt -l .)
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2 run ./...
	$(GO) run github.com/client9/misspell/cmd/misspell@latest *

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

# Gitlab test report
junit.xml: test.log
	$(GO) run github.com/jstemmer/go-junit-report/v2@latest < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	$(GO) run github.com/boumenot/gocover-cobertura@latest -cobertura < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: staticcheck.json
	$(GO) run github.com/banyansecurity/golint-convert@latest -convert > $@

staticcheck.json:
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -format=sarif ./... > $@

.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

peg: grammar.go

.PHONY: update-logo
update-logo:
	@command -v glab >/dev/null 2>&1 || { echo "glab not found. Install glab: https://gitlab.com/gitlab-org/cli" >&2; exit 1; }
	@[ -f static/logo.png ] || { echo "static/logo.png not found; nothing to upload" >&2; exit 1; }
	@proj_id=$$(glab repo view --json id --jq .id); \
	  echo "Updating GitLab project $$proj_id avatar from static/logo.png..."; \
	  glab api "projects/$$proj_id" -X PUT -F "avatar=@static/logo.png" >/dev/null; \
	  echo "Project avatar updated."

README.html: README.adoc
	asciidoc $^

.PHONY: total
total: $(BENCH_FILE)
	awk -f total.awk < $(BENCH_FILE)
