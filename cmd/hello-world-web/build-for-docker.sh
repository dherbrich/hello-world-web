#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world-web .
docker build -t dherbrich/hello-world-web -f DOCKERFILE . 

docker push dherbrich/hello-world-web:latest
docker tag dherbrich/hello-world-web:latest dherbrich/hello-world-private:latest 
docker push dherbrich/hello-world-private:latest