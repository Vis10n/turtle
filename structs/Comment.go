package structs

type Comments struct {
	RequestID int    `json:request_id`
	UserID    int    `json:user_id`
	UserName  string `json:user_name`
	Comment   string `json:comment`
	CreatAt   string `json:create_at`
}
