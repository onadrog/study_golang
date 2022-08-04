package structs

type Visit struct {
	Id          string `json:"id"`
	Page        string `json:"page"`
	Sessionhash string `json:"sessionhash"`
}

type Task struct {
	Date   string
	Visits []Visit
}

type DailyStat struct {
	Date   string         `json:"date"`
	Bypage map[string]int `json:"byPage"`
}
