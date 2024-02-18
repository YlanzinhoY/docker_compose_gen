build:
	GOOS=linux GOARCH=amd64 go build -o dockercomposegen/linux/dockercomposegen-linux
	GOOS=darwin GOARCH=amd64 go build -o dockercomposegen/macos/dockercomposegen-darwin

test:
	GOOS=linux GOARCH=amd64 go build -o dockercomposegen/dockercomposegen_test