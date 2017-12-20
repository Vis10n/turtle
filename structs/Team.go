package structs

type TeamID struct {
	TeamID      int    `json:team_id`
	TeamLeader  int    `json:team_leader`
	Description string `json:team_description`
}
