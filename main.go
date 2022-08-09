package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Tool struct
type Tool struct {
	Id       int
	Task     string
	Assignee string
	Deadline string
	Status   string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbServer := os.Getenv("DATABASE_SERVER")
	dbPort := os.Getenv("DATABASE_PORT")
	log.Println("Database host: " + dbServer)
	log.Println("Database port: " + dbPort)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbServer+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM pegawai ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}
	res := []Tool{}

	for selDB.Next() {
		var id int
		var task, assignee, deadline, status string
		err := selDB.Scan(&id, &task, &assignee, &deadline, &status)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | task " + task + " | assignee " + assignee + " | deadline " + deadline + " | status " + status)

		tool.Id = id
		tool.Task = task
		tool.Assignee = assignee
		tool.Deadline = deadline
		tool.Status = status
		res = append(res, tool)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM pegawai WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id int
		var task, assignee, deadline string
		err := selDB.Scan(&id, &task, &assignee, &deadline)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | task " + task + " | assignee " + assignee + " | deadline " + deadline)

		tool.Id = id
		tool.Task = task
		tool.Assignee = assignee
		tool.Deadline = deadline
	}
	tmpl.ExecuteTemplate(w, "Show", tool)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM pegawai WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id int
		var task, assignee, deadline, status string
		err := selDB.Scan(&id, &task, &assignee, &deadline, &status)
		if err != nil {
			panic(err.Error())
		}

		tool.Id = id
		tool.Task = task
		tool.Assignee = assignee
		tool.Deadline = deadline
		tool.Status = status
	}

	tmpl.ExecuteTemplate(w, "Edit", tool)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		assignee := r.FormValue("assignee")
		deadline := r.FormValue("deadline")
		insForm, err := db.Prepare("INSERT INTO pegawai (task, assignee, deadline) VALUES (?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(task, assignee, deadline)
		log.Println("Insert Data: task " + task + " | assignee " + assignee + " | deadline " + deadline)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		assignee := r.FormValue("assignee")
		deadline := r.FormValue("deadline")
		status := r.FormValue("status")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE pegawai SET task=?, assignee=?, deadline=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(task, assignee, deadline, id)
		log.Println("UPDATE Data: task " + task + " | assignee " + assignee + " | deadline " + deadline + " | status " + status)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	tool := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM pegawai WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tool)
	log.Println("DELETE " + tool)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		// task := r.FormValue("task")
		// assignee := r.FormValue("assignee")
		// deadline := r.FormValue("deadline")
		id_pegawai := r.FormValue("id_pegawai")
		status := "True"
		// id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE pegawai SET status=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(status, id_pegawai)
		log.Println("UPDATE Data: status " + status + " | id_pegawai " + id_pegawai)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/updateStatus", UpdateStatus)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
