.PHONY: build
build: 
	go build -o syntime cmd/syntime/main.go

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build -o syntime.exe cmd/syntime/main.go

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -o syntime cmd/syntime/main.go

.PHONY: run
run:
	go run cmd/syntime/main.go