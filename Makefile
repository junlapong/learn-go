test:
	go test -v ./...

cover:
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html

bench:
	go test -bench=. atoi_test.go
