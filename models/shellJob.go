package models

type ShellJob struct {
	Id int `json:"id"`
	Name string  `json:"name"`
	Description string `json:"description"`
	Cron string  `json:"cron"`
	Cmd string  `json:"cmd"`
}
