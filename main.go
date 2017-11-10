package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type RowResult struct {
	EquipType   string
	Color       string
	Location    string
	InstallDate time.Time
}

type Page struct {
	Date    string
	Time    string
	Results []RowResult
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func add(x, y int) int {
	return x + y
}

type indexHandler struct {
	conf *AppConfig
}

func (i *indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	now := time.Now()

	HomePageVars := Page{
		Date:    now.Format("02-01-2006"),
		Time:    now.Format("15:04:05"),
		Results: retrieveAll(),
	}
	funcs := template.FuncMap{"add": add}
	t := template.Must(template.New("index.html").Funcs(funcs).ParseFiles("index.html"))
	err := t.Execute(w, HomePageVars)

	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	config := AppConfig{}
	err := loadConfig("config.json", &config)
	checkErr(err)

	initDB(&config)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v", config.WebServer.Port)
		port = config.WebServer.Port
	}

	http.Handle("/", &indexHandler{conf: &config})
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicf("ListenAndServe error: %v\n", err)
	}
}
