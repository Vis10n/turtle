package culi

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"turtle/structs"
)

func CreateUser(username, password, name, department string) error {
	err := taskQuery("insert into users (username, password, name, department) value(?,?,?,?)", username, password, name, department)
	return err
}

func ValidUser(username, password string)bool{
	var correctPassword string
	usersSQL := "select password from users where username=?"
	log.Print("validating user ", username)
	rows := database.query(usersSQL, username)

	if rows.Next(){
		err:=rows.Scan(&correctPassword)
		if err!= nil{
			return false
		}
	}
	if password == correctPassword{
		return true
	}
	rows.Close()
	return false
}

func GetUserID(username string) (int, error){
	userID := -1
	usersSQL := "select userID from users where username=?"
	rows := database.query(usersSQL, username)
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(&userID)
		if  err!=nil{
			return -1, err
		}
	}
	return userID, nil
}

func GetUser(username string) *structs.User{
	user := structs.User{}
	usersSQL := "select name, avatar, title, department, from users where username=?"
	rows := database.query(usersSQL, username)
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(&user.Name,&user.Avatar,&user.Title,&user.Department)
		if  err!=nil{
			log.Print(err)
			return nil
		}
	}

	return &user
}

func GetUserByID(userID int) *structs.User{
	user := structs.User{}
	usersSQL := "select name, avatar, title, department from users where userID=?"
	rows := database.query(usersSQL, userID)
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(&user.Name,&user.Avatar,&user.Title,&user.Department)
		if  err!=nil{
			log.Print(err)
			return nil
		}
	}

	return &user
}

func GetTeams(teamID string) (usersID []int){
	usersSQL := "select userID from users where teamID=?"
	rows := database.query(usersSQL, teamID)
	defer rows.Close()

	var userID int
	for rows.Next(){
		err := rows.Scan(&userID)
		if  err!=nil{
			log.Print(err)
		}
		usersID = append(usersID, userID)
	}
	return usersID
}