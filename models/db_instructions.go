package models

import(
	"time"
	"context"
	"errors"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	// "fmt" // for debugging
)

type applicationRecordListView struct {
	Id int
	JobTitle, Company *string
	AppDate	*time.Time
}

type applicationView struct {
	Id int
	JobTitle, Description, Url, Company *string
	Resume, CoverLetter []byte
	AppDate, Offer, Rejected, Declined *time.Time
}

type application struct {
	JobTitle, Description, Url, Company string
	Resume, CoverLetter []byte
	AppDate, Offer, Rejected, Declined time.Time
}

type contactRecordListView struct {
	Id int
	Name, Position, Number, Email, Company, Note *string 
}

type contactView struct {
	Id int
	Name, Position, Number, Email, Company, Note *string
}

type contact struct {
	Name, Position, Number, Email, Company, Note string
}

type interviewRecordListView struct {
	Id int
	Appointment *time.Time
	Method, JobTitle, Company string
}

type interview struct {
	Appointment time.Time
	Method string
	JobID int
}

type interviewLink struct {
	Id string
}

func (conn *Db_conn) AllApplications() ([]*applicationRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, job_title, company, app_date FROM application")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]*applicationRecordListView, 0)
	for rows.Next() {
		app := new(applicationRecordListView)
		err := rows.Scan(&app.Id, &app.JobTitle, &app.Company, &app.AppDate)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)		
	}
	
	if err := rows.Err(); err != nil {	
		return nil, err
	}

	return apps, nil
}

func (conn *Db_conn) InsertApplication(req *http.Request) error {
	var err error
	var Buf bytes.Buffer
	req.ParseMultipartForm(32<<20)

	app := application{}
	
	app.JobTitle = req.FormValue("job_title")
	app.Description = req.FormValue("description")
	app.Url = req.FormValue("url")
	app.Company = req.FormValue("company")
	app.AppDate = time.Now()

	if _, ok := req.MultipartForm.File["resume"]; ok {
		resume_file, _, err  := req.FormFile("resume")
		if err != nil {
			return err
		}
		defer resume_file.Close()
		if _, err = io.Copy(&Buf, resume_file); err != nil {
			return err
		}
		app.Resume = Buf.Bytes()
		Buf.Reset()
	}

	if _, ok := req.MultipartForm.File["cvr_letter"]; ok {
		coverletter_file, _, err := req.FormFile("cvr_letter")
		if err != nil {
			return err
		}
		defer coverletter_file.Close()
		if _, err = io.Copy(&Buf, coverletter_file); err != nil{
			return err
		}
		app.CoverLetter = Buf.Bytes()
	}

	if app.JobTitle == "" || app.Company == "" {
		return errors.New("Job Title and Company cannot be left blank.")
	}
	
	_, err = conn.Exec(context.Background(), "INSERT INTO application (job_title, description, url, company, resume, cvr_letter, app_date) VALUES ($1, $2, $3, $4, $5, $6, $7)", app.JobTitle, app.Description, app.Url, app.Company, app.Resume, app.CoverLetter, app.AppDate)
	if err != nil {
		return err
	}
	
	return nil
}

func (conn *Db_conn) ViewApplication(req *http.Request) (*applicationView, error) {
	id := req.FormValue("id")
	if id == "" {
		return new(applicationView), errors.New("Error in retrieving id")
	}

	row := conn.QueryRow(context.Background(), "SELECT * FROM application WHERE id = $1", id)

	app := new(applicationView)
	err := row.Scan(&app.Id, &app.JobTitle, &app.Description, &app.Url, &app.Company, &app.Resume, &app.CoverLetter, &app.AppDate, &app.Offer, &app.Rejected, &app.Declined)
	if err != nil {
		return app, err
	}

	err = ioutil.WriteFile("tmp/resume.pdf", app.Resume, 0644)
	if err != nil {
		return app, err
	}

	err = ioutil.WriteFile("tmp/coverletter.pdf", app.CoverLetter, 0644)
	if err != nil {
		return app, err
	}

	return app, nil
}

func (conn *Db_conn) RemoveApplication(req *http.Request) error {
	id := req.FormValue("id")
	if id == "" {
		return errors.New("Error in retrieving id")
	}

	_, err := conn.Exec(context.Background(), "DELETE FROM interview WHERE job_id = $1;", id)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "DELETE FROM application WHERE id = $1;", id)
	if err != nil {
		return err
	}

	return nil
}
func (conn *Db_conn) UpdateApplicationGET(req *http.Request) (*applicationView, error) {
	id := req.FormValue("id")
	if id == "" {
		return new(applicationView), errors.New("Error in retrieving id")
	}

	row := conn.QueryRow(context.Background(), "SELECT * FROM application WHERE id = $1", id)

	app := new(applicationView)
	err := row.Scan(&app.Id, &app.JobTitle, &app.Description, &app.Url, &app.Company, &app.Resume, &app.CoverLetter, &app.AppDate, &app.Offer, &app.Rejected, &app.Declined)
	if err != nil {
		return app, err
	}

	return app, nil
}

func (conn *Db_conn) UpdateApplicationPOST(req *http.Request) error {
	var err error
	var Buf bytes.Buffer
	req.ParseMultipartForm(32<<20)

	time_layout := "2006-01-02 15:04:05.000000 -0700 MST"

	id := req.FormValue("id")
	if id == "" {
		return errors.New("Error in retrieving id")
	}
	
	app := application{}
	
	app.JobTitle = req.FormValue("job_title")
	app.Description = req.FormValue("description")
	app.Url = req.FormValue("url")
	app.Company = req.FormValue("company")
	
	app.AppDate, err = time.Parse(time_layout, req.FormValue("appdate"))
	if err != nil {
		return err
	}

	if req.FormValue("offer") != "" {
		app.Offer, err = time.Parse(time_layout, req.FormValue("offer"))
		if err != nil {
			return err
		}
	}
	if req.FormValue("rejected") != "" {
		app.Rejected, err = time.Parse(time_layout, req.FormValue("rejected"))
		if err != nil {
			return err
		}
	}
	if req.FormValue("declined") != "" {
		app.Declined, err = time.Parse(time_layout, req.FormValue("declined"))
		if err != nil {
			return err
		}
	}

	if _, ok := req.MultipartForm.File["resume"]; ok {
		resume_file, _, err  := req.FormFile("resume")
		if err != nil {
			return err
		}
		defer resume_file.Close()
		if _, err = io.Copy(&Buf, resume_file); err != nil {
			return err
		}
		app.Resume = Buf.Bytes()
		Buf.Reset()
	}

	if _, ok := req.MultipartForm.File["cvr_letter"]; ok {
		coverletter_file, _, err := req.FormFile("cvr_letter")
		if err != nil {
			return err
		}
		defer coverletter_file.Close()
		if _, err = io.Copy(&Buf, coverletter_file); err != nil{
			return err
		}
		app.CoverLetter = Buf.Bytes()
	}

	if app.JobTitle == "" || app.Company == "" {
		return errors.New("Job Title and Company cannot be left blank.")
	}
	_, err = conn.Exec(context.Background(), "UPDATE application SET (job_title, description, url, company, resume, cvr_letter, app_date, offer, rejected, declined) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) WHERE id = $11", app.JobTitle, app.Description, app.Url, app.Company, app.Resume, app.CoverLetter, app.AppDate, app.Offer, app.Rejected, app.Declined, id)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) AllContacts() ([]*contactRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contacts := make([]*contactRecordListView, 0)
	for rows.Next() {
		contact := new(contactRecordListView)
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Position, &contact.Number, &contact.Email, &contact.Company, &contact.Note)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (conn *Db_conn) InsertContact(req *http.Request) error {
	req.ParseMultipartForm(32<<20)

	contact := contact{
		Name: req.FormValue("name"),
		Position: req.FormValue("position"),
		Number: req.FormValue("number"),
		Email: req.FormValue("email"),
		Company: req.FormValue("company"),
		Note: req.FormValue("note"),
	}

	_, err := conn.Exec(context.Background(), "INSERT INTO contacts (name, position, number, email, company, note) VALUES ($1, $2, $3, $4, $5, $6)", contact.Name, contact.Position, contact.Number, contact.Email, contact.Company, contact.Note)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) RemoveContact(req *http.Request) error {
	id := req.FormValue("id")
	if id == "" {
		return errors.New("Error in retrieving id")
	}

	_, err := conn.Exec(context.Background(), "DELETE FROM contacts WHERE id = $1;", id)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) UpdateContactGET(req *http.Request) (*contactView, error) {
	id := req.FormValue("id")
	if id == "" {
		return new(contactView), errors.New("Error in retrieving id")
	}
	
	row := conn.QueryRow(context.Background(), "SELECT * FROM contacts WHERE id = $1", id)

	contact := new(contactView)
	err := row.Scan(&contact.Id, &contact.Name, &contact.Position, &contact.Number, &contact.Email, &contact.Company, &contact.Note)
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (conn *Db_conn) UpdateContactPOST(req *http.Request) error {
	id := req.FormValue("id")
	if id == "" {
		return errors.New("Error in retrieving id")
	}
	
	req.ParseMultipartForm(32<<20)

	contact := contact{
		Name: req.FormValue("name"),
		Position: req.FormValue("position"),
		Number: req.FormValue("number"),
		Email: req.FormValue("email"),
		Company: req.FormValue("company"),
		Note: req.FormValue("note"),
	}

	_, err := conn.Exec(context.Background(), "UPDATE contacts SET (name, position, number, email, company, note) = ($1, $2, $3, $4, $5, $6) WHERE id = $7", contact.Name, contact.Position, contact.Number, contact.Email, contact.Company, contact.Note, id)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) AllInterviews() ([]*interviewRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT I.id, I.date, I.method, A.job_title, A.company FROM interview AS I LEFT JOIN application AS A ON I.job_id = A.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interviews := make([]*interviewRecordListView, 0)
	for rows.Next() {
		interview := new(interviewRecordListView)
		err := rows.Scan(&interview.Id, &interview.Appointment, &interview.Method, &interview.JobTitle, &interview.Company)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, interview)
	}
	return interviews, nil
}

func (conn *Db_conn) InsertInterviewGET(req *http.Request) (*interviewLink, error) {
	id, ok := req.URL.Query()["id"]
	if !ok {
		return new(interviewLink), errors.New("Error in retrieving id.")
	}
	link := new(interviewLink)
	link.Id = id[0]
	
	return link, nil
}

func (conn *Db_conn) InsertInterviewPOST(req *http.Request) error {
	var err error

	time_layout := "2006-01-02T15:04"
	req.ParseMultipartForm(32<<20)
	interview := interview{}
	
	interview.Appointment, err = time.Parse(time_layout, req.FormValue("date"))
	if err != nil {
		return err
	}
	interview.Method = req.FormValue("method")
	interview.JobID, err = strconv.Atoi(req.FormValue("id"))
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO interview (date, method, job_id) VALUES ($1, $2, $3)", interview.Appointment, interview.Method, interview.JobID)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) UpdateInterviewPOST(req *http.Request) error {
	var err error

	id_slice, ok := req.URL.Query()["id"]
	if !ok {
		errors.New("Error in retrieving id.")
	}

	time_layout := "2006-01-02T15:04"
	req.ParseMultipartForm(32<<20)
	interview := interview{}

	interview.Appointment, err = time.Parse(time_layout, req.FormValue("date"))
	if err != nil {
		return err
	}
	interview.Method = req.FormValue("method")

	_, err = conn.Exec(context.Background(), "UPDATE interview SET (date, method) = ($1, $2) WHERE id = $3", interview.Appointment, interview.Method, id_slice[0])
	if err != nil {
		return err
	}

	return nil
}

func (conn *Db_conn) RemoveInterview(req *http.Request) error {
	id_slice, ok := req.URL.Query()["id"]
	if !ok {
		errors.New("Error in retrieving id.")
	}

	_, err := conn.Exec(context.Background(), "DELETE FROM interview WHERE id = $1;", id_slice[0])
	if err != nil {
		return err
	}

	return nil
}