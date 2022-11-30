package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("SERVER_NAME")
	fmt.Fprintf(w, "<h1>Hello World, from %s</h1>", name)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "<h1>Hello World, from %s</h1>", password)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Printf(w, "My family: %s", string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() > 25 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("ok")))
	}

}
