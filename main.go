package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	port = ":8080"
)

func testHandler(w http.ResponseWriter, r *http.Request) {

}

func serveDynamic(w http.ResponseWriter, r *http.Request){
	response := "The time is now " + time.Now().String();
	fmt.Fprintln(w, response);
}

func serveStatic(w http.ResponseWriter, r *http.Request){
	//NOTE : file location is relative
	http.ServeFile(w,r,"public/json/whoami.json");
}

func serveError(w http.ResponseWriter, r *http.Request){
	fmt.Println("There's no way way I'll work !");
}

func main() {

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/whoami",serveStatic);
	http.HandleFunc("/",serveDynamic);
	http.HandleFunc("/error",serveError)
	http.ListenAndServe(port,nil);
}
