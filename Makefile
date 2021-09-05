.PHONY: build
build:
	@go build -o openvpnas_exporter ./cmd/main.go 

.PHONY: build-no-c
build-no-c:
	@CGO_ENABLED=0 go build -o openvpnas_exporter ./cmd/main.go 
	

.PHONY: run
run:
	@go run ./cmd/main.go 