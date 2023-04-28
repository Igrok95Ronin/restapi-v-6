package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Home struct {
	Title string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db, err := sql.Open("mysql", "newuser:password@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM snippets")
	if err != nil {
		log.Fatal(err)
	}
	var snipTitle []Home
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}

		item := Home{
			title,
		}
		snipTitle = append(snipTitle, item)
	}

	err = ts.Execute(w, snipTitle)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", 405)
		return
	}
	w.Write([]byte("Отображение заметки..."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d: ", id)
	fmt.Fprintf(w, "%s", r.URL.Path)
}

type Abouts struct {
	Id    int
	About string
}

func about(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/about.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db, err := sql.Open("mysql", "newuser:password@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title FROM snippets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sAbout []Abouts
	for rows.Next() {
		var id int
		var title string
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		item := Abouts{
			id,
			title,
		}
		sAbout = append(sAbout, item)

	}

	err = ts.Execute(w, sAbout)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type Contacts struct {
	Id             int
	Title, Content string
}

func contacts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/contacts.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db, err := sql.Open("mysql", "newuser:password@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,title,content FROM snippets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var c []Contacts
	for rows.Next() {
		var id int
		var title, content string

		err = rows.Scan(&id, &title, &content)
		if err != nil {
			log.Fatal(err)
		}

		item := Contacts{
			id,
			title,
			content,
		}
		c = append(c, item)
	}

	err = ts.Execute(w, c)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
