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
	http.HandleFunc("/applications", env.applicationShow)
	http.HandleFunc("/applications/add", env.applicationAddFormGET)
	http.HandleFunc("/applications/add/execute", env.applicationAddFormPOST)
	http.HandleFunc("/applications/update", env.applicationUpdateFormGET)
	http.HandleFunc("/applications/update/execute", env.applicationUpdateFormPOST)
	http.HandleFunc("/applications/remove/execute", env.applicationRemove)
	http.HandleFunc("/applications/view", env.applicationShowOne)
	http.HandleFunc("/contacts", env.contactShow)
	http.HandleFunc("/contacts/add", env.contactAddFormGET)
	http.HandleFunc("/contacts/add/execute", env.contactAddFormPOST)
	http.HandleFunc("/contacts/update", env.contactUpdateFormGET)
	http.HandleFunc("/contacts/update/execute", env.contactUpdateFormPOST)
	http.HandleFunc("/contacts/remove/execute", env.contactRemove)
	http.HandleFunc("/interviews", env.interviewShow)
	http.HandleFunc("/interviews/add", env.interviewAddFormGET)
	http.HandleFunc("/interviews/add/execute", env.interviewAddFormPOST)
	

	http.Handle("/tmp/", http.StripPrefix("/tmp", http.FileServer(http.Dir("./tmp"))))

	http.ListenAndServe(":8080", nil)
}

func (env *Env) index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (env *Env) applicationShow(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	appsList, err := env.conn.AllApplications()
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}


	tpl.ExecuteTemplate(w, "applications.gohtml", appsList)
}

func (env *Env) applicationShowOne(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	app, err := env.conn.ViewApplication(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "oneApplication.gohtml", app)
}

func (env *Env) applicationAddFormGET(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "applicationAddForm.gohtml", nil)
}

func (env *Env) applicationAddFormPOST(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := env.conn.InsertApplication(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/applications", http.StatusSeeOther)
}

func (env *Env) applicationRemove(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := env.conn.RemoveApplication(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/applications", http.StatusSeeOther)
}

func (env *Env) applicationUpdateFormGET(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	app, err := env.conn.UpdateApplicationGET(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "applicationUpdateForm.gohtml", app)
}

func (env *Env) applicationUpdateFormPOST(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	err := env.conn.UpdateApplicationPOST(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/applications", http.StatusSeeOther)
}

func (env *Env) contactShow(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	contactsList, err := env.conn.AllContacts()
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "contacts.gohtml", contactsList)
}

func (env *Env) contactAddFormGET(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	tpl.ExecuteTemplate(w, "contactAddForm.gohtml", nil)
}

func (env *Env) contactAddFormPOST(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := env.conn.InsertContact(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/contacts", http.StatusSeeOther)
}

func (env *Env) contactRemove(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := env.conn.RemoveContact(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/contacts", http.StatusSeeOther)
}

func (env *Env) contactUpdateFormGET(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	contact, err := env.conn.UpdateContactGET(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "contactUpdateForm.gohtml", contact)
}

func (env *Env) contactUpdateFormPOST(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return	
	}

	err := env.conn.UpdateContactPOST(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/contacts", http.StatusSeeOther)
}


func (env *Env) interviewShow(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	interviewList, err := env.conn.AllInterviews()
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "interviews.gohtml", interviewList)
}

func (env *Env) interviewAddFormGET(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id, err := env.conn.InsertInterviewGET(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return 
	}
	
	tpl.ExecuteTemplate(w, "interviewAddForm.gohtml", id)
}

func (env *Env) interviewAddFormPOST(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return	
	}

	err := env.conn.InsertInterviewPOST(req)
	if err != nil {
		http.Error(w, http.StatusText(500)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/interviews", http.StatusSeeOther)
}