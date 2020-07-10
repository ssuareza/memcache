# About
A simple API where you can insert values in in memcached. Files supported.

**Note**: the API will be running at port 8090.

# Usage

## Set key
```sh
$ curl -s -X POST http://localhost:8090/set/mykey?value=random
stored
```

## Set key with file
$ curl -s -X POST -F "file=@file.dat" http://localhost:8090/set/mykey

## Get key
```sh
$ curl -s http://localhost:8090/get/mykey
dago
```

## Flush cache
```sh
$ curl -s -X POST http://localhost:8090/flush
```
