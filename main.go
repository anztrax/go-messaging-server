package main;

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request){

}

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/test",testHandler);

	http.Handle("/",router);
	fmt.Println("Everything is set up !");
}
