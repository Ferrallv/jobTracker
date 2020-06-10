package models

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"net/http"
)

type Db_conn_funcs interface {
	AllApplications() ([]*applicationRecordListView, error)
	InsertApplication(req *http.Request) error
	UpdateApplicationGET(req *http.Request) (*applicationView, error)
	UpdateApplicationPOST(req *http.Request) error
	ViewApplication(req *http.Request) (*applicationView, error)
	RemoveApplication(req *http.Request) error
	AllContacts() ([]*contactRecordListView, error)
	InsertContact(req *http.Request) error
	UpdateContactGET(req *http.Request) (*contactView, error)
	UpdateContactPOST(req *http.Request) error
	RemoveContact(req *http.Request) error
	AllInterviews() ([]*interviewRecordListView, error)
	InsertInterviewGET(req *http.Request) (*interviewLink, error)
	InsertInterviewPOST(req *http.Request) error
}

type Db_conn struct {
	*pgxpool.Pool
}

func NewConn(SourceURL string) (*Db_conn, error) {
	conn, err := pgxpool.Connect(context.Background(), SourceURL)
	if err != nil {
		return nil, err
	}
	fmt.Println("You connected to your database.")

	return &Db_conn{conn}, nil
}