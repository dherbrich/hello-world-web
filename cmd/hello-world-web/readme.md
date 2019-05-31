# Pack build with local go buildpack to docker image

pack build "dherbrich/hello-world-web" --buildpack ../../go-buildpack --env 'COMMAND=./cmd/hello-world-web' --no-pull