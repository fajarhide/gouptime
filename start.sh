#!/usr/bin/env bash
# Start with running docker by @fajarhide
service=$1
# Building Docker Images for Static Go Binaries
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go
docker build --rm -t $service -f Dockerfile .
# replace service running
foundService=$(docker ps | grep $service | wc -l)
if [[ $foundService != 0 ]]
then
        docker rm -f $service
fi
docker run --name $service --rm -d -P $service
sleep 10
# remove <none>
foundImage=$(docker images | grep '<none>' | wc -l)
if [[ $foundImage != 0 ]]
then
    docker rmi -f $(docker images -f 'dangling=true' -q)
fi
