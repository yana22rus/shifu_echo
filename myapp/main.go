package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}



func main() {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS data_user (id INTEGER PRIMARY KEY, login TEXT UNIQUE NOT NULL,email TEXT UNIQUE NOT NULL,pswd TEXT NOT NULL)")
	statement.Exec()

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.GET("/", index)
	e.POST("/registration",registration)
	e.Logger.Fatal(e.Start(":5000"))
}

func registration(c echo.Context) error {

	login := c.FormValue("login")
	fmt.Println(login)
	return c.Redirect(302,"/")
}

func index(c echo.Context) error {

	return c.Render(http.StatusOK, "index","")
}