#!/usr/bin/env bash

rm mTest/*.sqlite
docker kill $(docker ps -q)
docker container prune -f
docker network prune -f
docker volume prune -f
docker-compose up

# go test -bench=. -benchmem -benchtime=1s .