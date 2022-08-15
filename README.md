<p align="center">
    <img src="assets/logo.jpeg" width="300" alt="logo" />
</p>

<h1 align="center">
    Xerox
</h1>

<p align="center">
    A library for creating a proxy server.
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    <img src="https://img.shields.io/badge/Version-1.0.1-informational?style=for-the-badge&logo=none" alt="version" />
</p>

## How to use?
You can use **Xerox** in Golang application or using it with Docker.

### Golang 
Get the repository:
```shell
go get github.com/amirhnajafiz/xerox
```

In your code, you have to set a target for proxy server and an address
for your proxy server to get start on:
```go
package main

import "github.com/amirhnajafiz/xerox"

func main() {
	// creating a proxy server on port 8080 and
	// bind to localhost:8081
	proxy := xerox.NewProxyServer("localhost:8081", "8080")

	// starting the proxy server
	proxy.Start()
}
```

### Docker 
To use **Xerox** in docker container:
```shell
docker run -p 8080:8080 --env SERVER_PORT=8080 --env SERVER_TARGET=www.google.com amirhossein21/xerox:v1.0.1
```
