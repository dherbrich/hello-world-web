package hellocontroller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dherbrich/hello-world-pkg"
)

var (
	Healthy = false
)

func Greet(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("PERSON")
	if user == "" {
		user = "webuser"
	}
	response := fmt.Sprintf(hello.SayHello(user)+"at %s", time.Now())
	fmt.Println(response)
	fmt.Fprintf(w, response)

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	if Healthy {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func SwitchHealthzWeb(w http.ResponseWriter, r *http.Request) {
	SwitchHealthz()
}

func SwitchHealthz() {
	Healthy = !Healthy
}
