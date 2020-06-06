package main

import (
	"net/http"
	"html/template"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"log"
	"jobtracker/models"
)

var tpl *template.Template 

type Env struct {
	conn models.Db_conn_funcs
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	conn, err := models.NewConn(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	env := &Env{conn}

	http.HandleFunc("/", env.index)
	http.HandleFunc("/applications", env.applicationListView)

	http.ListenAndServe(":8080", nil)
}

func (env *Env) index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (env *Env) applicationListView(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	appsList, err := env.conn.AllApplications()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return 
	}


	tpl.ExecuteTemplate(w, "applications.gohtml", appsList)
}