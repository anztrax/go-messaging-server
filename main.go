package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	port = ":8081"
	DBHost = "localhost"
	DBPort = ":8890"
	DBUser = "root"
	DBPass = "root"
	DBDbase = "cms"
)

type Page struct{
	Title string
	Content string
	Date string
}

var database *sql.DB

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	pageID := vars["id"];

	thisPage := Page{};

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date);
	if err != nil{
		log.Println("couldn't get page with id : ",pageID);
		log.Println(err)
	}

	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title + `</h1><div>` + thisPage.Content + `</div></body></html>`;
	fmt.Fprintln(w,html);
}

func serveStatic(w http.ResponseWriter, r *http.Request){
	//NOTE : file location is relative
	http.ServeFile(w,r,"public/json/whoami.json");
}

func connectToDB(){
	dbConn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s",DBUser, DBPass, DBHost,DBPort, DBDbase)
	fmt.Println(dbConn);
	db, err := sql.Open("mysql",dbConn)

	if err != nil{
		log.Println("Couln't connect");
		log.Fatal(err);
	}else{
		log.Println("success connect to db")
	}
	database = db;
}

func main() {
	connectToDB();

	rtr := mux.NewRouter();
	rtr.HandleFunc("/pages/{id:[0-9]+}",pageHandler);
	http.Handle("/",rtr);

	http.HandleFunc("/whoami",serveStatic);
	http.ListenAndServe(port,nil);
}
