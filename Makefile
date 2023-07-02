.PHONY: gen
gen:
	@echo "Generating..."
	protoc --go_out=pb --go-grpc_out=pb proto/*.proto

.PHONY: run
run:
	@echo "Running..."
	go run main.go

.PHONY: test
test:
	go test ./...