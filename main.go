package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	schema := `
    CREATE TABLE jobs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        company TEXT NOT NULL,
        location TEXT,
        description TEXT,
        status TEXT DEFAULT 'open',
        date_posted DATE DEFAULT CURRENT_DATE
    );

    CREATE TABLE applications (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        job_id INTEGER NOT NULL,
        applicant_name TEXT NOT NULL,
        applicant_email TEXT NOT NULL,
        status TEXT DEFAULT 'applied',
        date_applied DATE DEFAULT CURRENT_DATE,
        FOREIGN KEY(job_id) REFERENCES jobs(id)
    );

    INSERT INTO jobs (title, company, location, description)
    VALUES
      ('Backend Engineer', 'Acme Corp', 'Remote', 'Work on distributed APIs'),
      ('Frontend Developer', 'Globex', 'New York, NY', 'React + TypeScript UI development'),
      ('Data Analyst', 'Initech', 'Austin, TX', 'SQL + Python analytics role');

    INSERT INTO applications (job_id, applicant_name, applicant_email, status)
    VALUES
      (1, 'Dior Shelton', 'dior@example.com', 'applied'),
      (2, 'Dior Shelton', 'dior@example.com', 'interviewing');
    `
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	// Simple query: list jobs
	rows, _ := db.Query("SELECT j.id, j.title, j.company, a.status FROM jobs j LEFT JOIN applications a ON j.id = a.job_id")

	defer rows.Close()
	//snippet to load jobs
	for rows.Next() {
		var id int
		var title, company, status string
		rows.Scan(&id, &title, &company, &status)
		fmt.Printf("Job #%d: %s at %s -> %s\n", id, title, company, status)
	}

}
