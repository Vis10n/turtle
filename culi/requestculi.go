package culi

import (
	"log"
	"turtle/structs"
)

func CreateRequest(requesterID int, teamID int, reqSpec string, createdDate string, deadline string, status string) error {
	err := taskQuery("insert into requests (requesterID, teamID, reqSpec, createDate, deadline, status) value(?,?,?,?,?,?)", requesterID, teamID, reqSpec, createdDate, deadline, status)
	return err
}

func GetRequestByID(requestID int) *structs.Request {
	request := structs.Request{}
	requestSQL := "select * from requests where requestID=?"
	rows := database.query(requestSQL, requestID)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&request.RequestID, &request.RequesterID, &request.TeamID, &request.ReqSpec, &request.Priority, &request.Rating, &request.CreatedDate, &request.Deadline, &request.CompletedDate, &request.Status)
		if err != nil {
			log.Print(err)
			return nil
		}
	}
	return &request
}
