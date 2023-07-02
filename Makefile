.PHONY: gen
gen:
	@echo "Generating..."
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: run
run:
	@echo "Running..."
	go run main.go