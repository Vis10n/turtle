package structs

type User struct {
	Name       string `json:name`
	UserID     int    `json:user_id`
	TeamID     int    `json:team_id`
	Title      string `json:title`
	IsLeader   int    `json:is_leader`
	Department string `json:department`
	Avatar     string `json:avatar`
}
