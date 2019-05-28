module github.com/dherbrich/hello-world-web

require (
	github.com/dherbrich/hello-world-pkg v0.0.0-20190413211021-409c828b00ac
	github.com/gorilla/mux v1.7.1
)

go 1.12

//replace github.com/dherbrich/hello-world-pkg => ../hello-world-pkg/
