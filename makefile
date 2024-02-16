build:
	GOOS=linux GOARCH=amd64 go build -o dockercomposegen/dockercomposegen

test:
	GOOS=linux GOARCH=amd64 go build -o dockercomposegen/dockercomposegen_test