# Web Service

This is the Web service

Generated with

```
micro new xiaoshijie.com/micro/hello/web --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.web
- Type: web
- Alias: web

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./web-web
```

Build a docker image
```
make docker
```

 REGISTER_AUTH_USERNAME=root;REGISTER_AUTH_PASSWORD=mm123456 REGISTER_AUTH_USERNAME=root REGISTER_AUTH_PASSWORD=mm123456 go run main.go plugin.go --registry=etcdv3  --registry_address 0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379


Go Proxy Bug

```cassandraql
replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.3.0
```