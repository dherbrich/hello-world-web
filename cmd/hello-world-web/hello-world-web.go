package main

import "net/http"
import "github.com/dherbric/hello-world-web/internal/hellocontroller"

func main() {
	http.HandleFunc("/", hellocontroller.Greet)
	http.HandleFunc("/healthz", hellocontroller.Healthz)
	http.HandleFunc("/healthz/switch", hellocontroller.SwitchHealthzWeb)

	hellocontroller.SwitchHealthz()
	http.ListenAndServe(":8080", nil)
}
