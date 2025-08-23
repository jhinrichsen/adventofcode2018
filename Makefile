.PHONY: all
all: lint test

.PHONY: bench
bench:
	go test -bench=.

.PHONY: lint
lint:
	golint
	go vet
	staticcheck

.PHONY: test
test:
	go test

.PHONY: gopls
gopls:
	# go run golang.org/x/tools/gopls@latest serve -listen 127.0.0.1:4389
	go run golang.org/x/tools/gopls@latest serve -listen 127.0.0.1:4389 -rpc.trace -logfile gopls.log

prof:
	go test -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	go tool pprof cpu.profile
	# go tool pprof mem.profile
