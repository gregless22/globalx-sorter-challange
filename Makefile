# Make file reference https://kodfabrik.com/journal/a-good-makefile-for-go

test:
	echo "Running the tests"
	go test -v ./...

compile:
	echo "Compiling for every OS and Platform"
	go mod tidy
	go build -o bin/name-sorter cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/name-sorter-freebsd-386 cmd/main.go
	GOOS=linux GOARCH=386 go build -o bin/name-sorter-linux-386 cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/name-sorter-windows-386 cmd/main.go