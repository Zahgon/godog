package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type server struct {
	db *sql.DB
}

type user struct {
	ID       int64  `json:"-"`
	Username string `json:"username"`
	Email    string `json:"-"`
}

func (s *server) users(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// an empty array in this case

func main() {
	db, err := sql.Open("mysql", "root@/godog")
	if err != nil {
		panic(err)
	}
	s := &server{db: db}
	http.HandleFunc("/users", s.users)
	http.ListenAndServe(":8080", nil)
}

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) { _ = "STUB: not implemented"; return }

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) { _ = "STUB: not implemented"; return }
