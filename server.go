package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	fmt.Println("Running")
	http.ListenAndServe(":8090", nil)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	matchTime := duration.Seconds() < 10

	fmt.Println("match: ", matchTime, "duration: ", duration.Seconds())

	if matchTime {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/family/members.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}

func Secret(httpWriter http.ResponseWriter, request *http.Request) {

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")

	fmt.Fprintf(httpWriter, "User: %s, Pass: %s", string(user), string(pass))
}

func Hello(httpWriter http.ResponseWriter, request *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	resp := "Xablau " + name + " you're " + age + " fucking years old"

	headerSize, error := httpWriter.Write([]byte(resp))
	if error != nil {
		fmt.Printf("error in request: %v", error)
	}
	fmt.Println("Request completed, header size: ", headerSize)
}
