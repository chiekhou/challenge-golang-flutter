package main

import (
	"example/hello/initializers"
	"fmt"
	"net/http"
)

const port = ":3000"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/", helloWorld)
	http.ListenAndServe(port, s)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
