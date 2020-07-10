# About
A simple API where you can insert values in in memcached. Files supported.

**Note**: the API will be running at port 8090.

# Build
To build the API run:
```sh
make build
````
This will leave the binaries in the "build" directory.

## Start
Just run the binary specific for your platform (mac or linux). Example:
```sh
build/mapi-darwin-amd64
```
**Note**: The application will be running at port 8090.

## Dependencies
The API interact with a **memcached** server running at "localhost:11211". Make sure to have it before using it :-)

# Usage

## Set key
```sh
$ curl -s -X POST http://localhost:8090/set/mykey?value=random
STORED
```

## Set key with file
```sh
$ curl -s -X POST -F "file=@file.dat" http://localhost:8090/set/mykey
STORED
```

## Get key
```sh
$ curl -s http://localhost:8090/get/mykey
dago
```

## Flush cache
```sh
$ curl -s -X POST http://localhost:8090/flush
FLUSHED
```
