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

func ValidUser(username, password string) bool {
	var correctPassword string
	usersSQL := "select password from users where username=?"
	log.Print("validating user ", username)
	rows := database.query(usersSQL, username)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&correctPassword)
		if err != nil {
			return false
		}
	}
	if password == correctPassword {
		log.Print(username + " validated")
		return true
	}

	return false
}

func GetUserID(username string) (int, error) {
	userID := -1
	usersSQL := "select userID from users where username=?"
	rows := database.query(usersSQL, username)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			return -1, err
		}
	}
	return userID, nil
}

//trả về toàn bộ đối tượng User
func GetUser(username string) *structs.User {
	user := structs.User{}
	usersSQL := "select userID, title, isLeader, name, teamID, department, avatar from users where username=?"
	rows := database.query(usersSQL, username)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.Title, &user.IsLeader, &user.Name, &user.TeamID, &user.Department, &user.Avatar)
		if err != nil {
			log.Print(err)
			return nil
		}
	}

	return &user
}

func GetUserByID(userID int) *structs.User {
	user := structs.User{}
	usersSQL := "select userID, title, isLeader, name, teamID, department, avatar from users where userID=?"
	rows := database.query(usersSQL, userID)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.Title, &user.IsLeader, &user.Name, &user.TeamID, &user.Department, &user.Avatar)
		if err != nil {
			log.Print(err)
			return nil
		}
	}
	return &user
}

func GetTeams(teamID int) (usersID []int) {
	usersSQL := "select userID from users where teamID=?"
	rows := database.query(usersSQL, teamID)
	defer rows.Close()

	var userID int
	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			log.Print(err)
		}
		usersID = append(usersID, userID)
	}
	return usersID
}

func GetAllUserName() (AllUserName []string) {
	usersSQL := "select username from users"
	rows := database.query(usersSQL)
	defer rows.Close()
	for rows.Next() {
		temp := ""
		rows.Scan(&temp)
		AllUserName = append(AllUserName, temp)
	}
	return AllUserName
}
