go_mod:
	go mod download
	go mod tidy

test_environments:
	go test -v ./tests/environments_test.go

test_contracts:
	go test -v ./tests/contracts_test.go
