package models

import(
	"time"
	"context"
	"errors"
	"net/http"
	"bytes"
	"io"
	// "fmt" // for debugging
)

type applicationRecordListView struct {
	JobTitle, Company string
	AppDate	time.Time
}

type contactRecordListView struct {
	Id int
	Name, Position, Number, Email, Company, Note string 
}

type interviewRecordListView struct {
	Appointment time.Time
	Method, JobTitle, Company string
}

type application struct {
	JobTitle, Description, Url, Company string
	Resume, CoverLetter []byte
	AppDate	time.Time
}

func (conn *Db_conn) AllApplications() ([]*applicationRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT job_title, company, app_date FROM application")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]*applicationRecordListView, 0)
	for rows.Next() {
		app := new(applicationRecordListView)
		err := rows.Scan(&app.JobTitle, &app.Company, &app.AppDate)
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
	app.Company= req.FormValue("company")

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

	app.AppDate = time.Now()

	if app.JobTitle == "" || app.Company == "" {
		return errors.New("Job Title and Company cannot be left blank.")
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO application (job_title, description, url, company, resume, cvr_letter, app_date) VALUES ($1, $2, $3, $4, $5, $6, $7)", app.JobTitle, app.Description, app.Url, app.Company, app.Resume, app.CoverLetter, app.AppDate)
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



func (conn *Db_conn) AllInterviews() ([]*interviewRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT I.date, I.method, A.job_title, A.company FROM interview AS I LEFT JOIN application AS A ON I.job_id = A.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interviews := make([]*interviewRecordListView, 0)
	for rows.Next() {
		interview := new(interviewRecordListView)
		err := rows.Scan(&interview.Appointment, &interview.Method, &interview.JobTitle, &interview.Company)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, interview)
	}

	return interviews, nil
}