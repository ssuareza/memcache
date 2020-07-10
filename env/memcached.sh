#!/bin/bash
docker run --name memcached -p 11211:11211 -d memcached memcached -m 1000