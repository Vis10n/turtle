package structs

type Request struct {
	RequestID     int    `json:request_id`
	RequesterID   int    `json:requester_id`
	TeamID        int    `json:team_id`
	Status        string `json:status`
	ReqSpec       string `json:req_spec`
	CreatedDate   string `json:created_date`
	Deadline      string `json:deadline`
	Rating        string `json:rating`
	Priority      string `json:priority`
	CompletedDate string `json:completed_date`
}
