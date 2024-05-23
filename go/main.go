package main

import (
	"example/hello/initializers"
	"fmt"
	"log"
	"net/http"

	"github.com/zc2638/swag"
)

const port = ":3000"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	handle := swag.UIHandler("/swagger/ui", "", false)
	patterns := swag.UIPatterns("/swagger/ui")
	for _, pattern := range patterns {
		http.DefaultServeMux.Handle(pattern, handle)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
