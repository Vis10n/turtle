package structs

type User struct {
	UserID   int    `json:user_id`
	Title    string `json:title`
	IsLeader int    `json:is_leader`
	Name     string `json:name`

	TeamID     int    `json:team_id`
	Department string `json:department`
	Avatar     string `json:avatar`
}
