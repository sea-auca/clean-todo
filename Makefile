build:
	go build -o bin/clean-todo cmd/apiserver/main.go
	
run:
	go run cmd/apiserver/main.go

test_all:
	echo "☹️ tests are not implemented"