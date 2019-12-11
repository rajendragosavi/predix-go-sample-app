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
	
	var vcapServices map[string]interface{}
	// instead of loadconfig and reading from config.json we will read from env variables.
	json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcapServices)
	log.Println("---- VCAP services ---", vcapServices)
	postgres := vcapServices["postgres-2.0"].([]interface{})[0].(map[string]interface{})["credentials"].(map[string]interface{})

	log.Println("Reading from the VCAP services .....")

	config.Postgres.DBName = postgres["database"].(string)
	config.Postgres.Hostname = postgres["hostname"].(string)
	config.Postgres.Password = postgres["password"].(string)
	portNumber := postgres["port"].(float64)
	p := int(portNumber)
	config.Postgres.Port = p
	config.Postgres.User = postgres["username"].(string)
	//config.WebServer.Port = "8585"

	log.Println("-------------DBName------------", config.Postgres.DBName)
	log.Println("--------------- Hostname------------", config.Postgres.Hostname)
	log.Println("-------------- Password----------", config.Postgres.Password)
	log.Println("----------------Port------------", config.Postgres.Port)
	log.Println("---------- User --------------", config.Postgres.User)
	fmt.Println("*********** RAJENDRA *************", vcapServices)
	
	/*
	err := loadConfig("config.json", &config)
	checkErr(err)
	*/
	initDB(&config)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v", config.WebServer.Port)
		port = config.WebServer.Port
	}

	http.Handle("/", &indexHandler{conf: &config})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicf("ListenAndServe error: %v\n", err)
	}
}
