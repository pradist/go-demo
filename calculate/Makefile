banch:
	go test -bench=.

test:
	go test ./... -v -coverprofile=coverage.out

test-cover:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out
