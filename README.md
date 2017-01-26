# gockertest

[![Build Status](https://travis-ci.org/takecy/gockertest.svg?branch=master)](https://travis-ci.org/takecy/gockertest)
[![Go Report Card](https://goreportcard.com/badge/github.com/takecy/gockertest)](https://goreportcard.com/report/github.com/takecy/gockertest)

![](https://img.shields.io/badge/golang-1.7.4-blue.svg?style=flat-square)
![](https://img.shields.io/badge/docker-1.13.0-blue.svg?style=flat-square)
![](https://img.shields.io/badge/docker--compose-1.10.0-blue.svg?style=flat-square)


Simple tool for testing with docker container on each test by golang.  
Support your private registry.

<br/>

## Prepare
You should run docker on local.  
Check this command.
```
$ docker version
```
If this command occurred error, you should install docker in your PC beforehand.  
[Install Docker](https://docs.docker.com/engine/installation/)

### On CI
ex. travis  
.travis.yml
```yaml
language: go

go: 
    - tip

sudo: false

services:
    - docker # here

script:
    go test -v ./...
```

<br/>

## Usage in your test code
```
$ go get -u github.com/takecy/gockertest
```

<br/>
for public registry.  
[see redis example](./example/redis/redis.go)

```go
args := gockertest.Arguments{
    Ports: map[int]int{6379: 6379},
}
cli := gockertest.Run("redis:3.2-alpine", args)
defer cli.Cleanup()
```

for private registy with basic authentication.  
[see redis example](./example/custom/yours.go)

```go
args := gockertest.Arguments{
    Ports:        map[int]int{6379: 6379},
    RequireLogin: true, // require basic authentication
    Login: gockertest.Login{
        User:     "yourname",          // change to your username
        Password: "pass",              // change to your password
        Registry: "registry.yours.io", // change to your registy domain
    },
}
cli := gockertest.Run("registry.yours.io/redis:3.2-alpine", args)
defer cli.Cleanup()
```

### Supported options of `docker run`

* --net
* -p
* --name
* --rm
* -d

<br/>

### Why not support `--link` ?
Because `--link` is legacy.  
[Official Docs](https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/)  
You should use network features.  

<br/>

## Testing
```
$ go test ./...
```

<br/>

## License
[MIT](./LICENSE)