package main

import (
	"jobtracker/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockDB struct{}

var the_time *int64

func (mdb *mockDB) AllApplications() ([]*models.applicationRecordListView, error) {
	apps := make([]*models.applicationRecordListView, 0)
	*the_time = time.Now().Unix()
	apps = append(apps, &models.applicationRecordListView{1, "jobTitle1", "Company1", the_time})
	apps = append(apps, &models.applicationRecordListView{2, "jobTitle2", "Company2", the_time})
	return apps, nil
}

func testApplicationShow(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/applications", nil)

	env := Env{conn: &mockDB{}}
	http.HandlerFunc(env.applicationShow).ServeHTTP(rec, req)

	expected := "1, jobTitle1, Company1, " + the_time + "\n2, jobTitle2, Company2, " + the_time
	if expected != req.Body.String() {
		t.Errorf("\nExpected %v\n Got %v\n", expected, rec.Body.String())
	}
}
