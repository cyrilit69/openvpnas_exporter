.PHONY: build
build:
	@go build -o openvpnas_exporter ./cmd/main.go 

.PHONY: run
run:
	@go run ./cmd/main.go 