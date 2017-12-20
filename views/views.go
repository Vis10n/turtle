package views

import(
	"html/template"
	"net/http"
	"fmt"
)
var templateDir ="./views/templates/"
var templates  = template.Must(template.ParseGlob(templateDir+"*.html"))
var message string
var err error

var loginTemplate = templates.Lookup("login.html")
var homeTemplate = templates.Lookup("home.html")

func ShowLogin (w http.ResponseWriter, r *http.Request){
	loginTemplate.Execute(w, nil)
}
func ShowLoginFailed(w http.ResponseWriter){
	loginTemplate.Execute(w, nil)
	fmt.Fprint(w, "Sai tên đăng nhập hoặc mật khẩu")
}

func ShowHome(w http.ResponseWriter, r *http.Request){
	homeTemplate.Execute(w, nil)
}