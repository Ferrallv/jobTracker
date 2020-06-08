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
	AllContacts() ([]*contactRecordListView, error)
	AllInterviews() ([]*interviewRecordListView, error)
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