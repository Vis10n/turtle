package structs

type Task struct {
	TaskName        string `json:task_name`
	UserID          int    `json:user_id`
	TaskID          int    `json:task_id`
	TaskDescription string `json:task_description`
	Deadline        string `json:deadline`
	CreatedDate     string `json:created_date`
	CompletedDate   string `json:completed_date`
	Status          string `json:status`
}
