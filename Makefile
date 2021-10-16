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

prof:
	go test -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	go tool pprof cpu.profile
	# go tool pprof mem.profile
