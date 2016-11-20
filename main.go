package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"html/template"
	"os"
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
	guID := vars["guID"];

	thisPage := Page{};
	fmt.Println("guID :",guID);
	currWd , _ := os.Getwd();
	fmt.Println("current working directory : ",currWd);

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", guID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date);
	if err != nil{
		http.Error(w, http.StatusText(404), http.StatusNotFound);
		log.Println("couldn't get page with id : ",guID);
		log.Println(err);
		return;
	}

	//NOTE : template directory is absolute to project root
	currTmplDir := "src/github.com/anztrax/messaging-server/public/templates/blog.html";
	t , _ := template.ParseFiles(currTmplDir);
	t.Execute(w,thisPage);
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
	rtr.HandleFunc("/pages/{guID:[0-9a-zA\\-]+}",pageHandler);
	http.Handle("/",rtr);

	http.HandleFunc("/whoami",serveStatic);
	http.ListenAndServe(port,nil);
}
