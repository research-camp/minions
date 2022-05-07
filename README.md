# xerox

Xerox is a proxy server implemented in Golang.

## How does it work?
By creating a reverse proxy server, user request will go through
our proxy server and the server will send that request to a base url
that we set on the application configs, after that, server responses with
the response from origin server.

## Run application
You can run the application by docker:
```shell
docker-compose up -d
```

You can set the base url address in the docker-compose file:
```yaml
environment:
      BASE_URL: "[Your base url]"
```

The default base url is **localhost:8081** which is the test server.

## Configs
If you want to run the application, make sure to create the config.yaml file.
You can check some application configs, in configs directory.

```shell
cp ./configs/[your file].yaml config.yaml
```

## Metrics
You can check the application prometheus metrics on port _1220_ route _/metric_.
