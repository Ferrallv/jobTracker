package models

import(
	"time"
	"context"
)

type applicationRecordListView struct {
	JobTitle string
	Company string
	AppDate	time.Time
}

func (conn *Db_conn) AllApplications() ([]*applicationRecordListView, error) {
	rows, err := conn.Query(context.Background(), "SELECT job_title, company, app_date FROM application")
	if err != nil {
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	apps := make([]*applicationRecordListView, 0)
	for rows.Next() {
		app := new(applicationRecordListView)
		err := rows.Scan(&app.JobTitle, &app.Company, &app.AppDate)
		if err != nil {
			// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return nil, err
		}
		apps = append(apps, app)		
	}
	
	if err := rows.Err(); err != nil {	
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return nil, err
	}

	return apps, nil
}