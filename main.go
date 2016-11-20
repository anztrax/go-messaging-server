package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

const (
	port = ":8081"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	pageID := vars["id"];
	fileName := "public/files/" + pageID + ".html";

	//NOTE : why we need OS stat ? because 404 is a situation based on not found a file
	_, err := os.Stat(fileName);
	if err != nil{
		fileName = "public/files/404.html";
	}

	http.ServeFile(w,r,fileName);
}

func serveStatic(w http.ResponseWriter, r *http.Request){
	//NOTE : file location is relative
	http.ServeFile(w,r,"public/json/whoami.json");
}


func main() {
	rtr := mux.NewRouter();
	rtr.HandleFunc("/pages/{id:[0-9]+}",pageHandler);
	rtr.HandleFunc("/homepage",pageHandler);
	rtr.HandleFunc("/contanct",pageHandler);

	http.Handle("/",rtr);

	http.HandleFunc("/whoami",serveStatic);
	http.ListenAndServe(port,nil);
}
