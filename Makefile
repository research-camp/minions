# build
build:
	go build -o /main
# run
run:
	go run main.go
# test server
server-test:
	curl -i localhost:8081
# test proxy server
proxy-test:
	curl -i localhost:8080