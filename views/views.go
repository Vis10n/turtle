package views

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateDir = "./views/templates/"
var templates = template.Must(template.ParseGlob(templateDir + "*.html"))
var message string
var err error

var loginTemplate = templates.Lookup("login.html")
var homeTemplate = templates.Lookup("home.html")

//TODO: chia ra xử lý
func ShowLogin(w http.ResponseWriter) {
	loginTemplate.Execute(w, nil)
}
func ShowLoginFailed(w http.ResponseWriter) {
	loginTemplate.Execute(w, nil)
	fmt.Fprint(w, "Sai tên đăng nhập hoặc mật khẩu")
}

func ShowHome(w http.ResponseWriter, username string) {
	//todo: create homepage
	homeTemplate.Execute(w, nil)
	fmt.Fprint(w, username)
}

func ShowSignUpForm(w http.ResponseWriter) {

}

func ShowSignUpFailed(w http.ResponseWriter) {

}

func ShowSignUpSuccess(w http.ResponseWriter) {

}
