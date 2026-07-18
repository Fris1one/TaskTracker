package storage

type TaskDTO struct {
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Favourite bool   `json:"favourite"`
	Stage     int    `json:"stage"`
}

type SaveData struct {
	Tasks []TaskDTO `json:"tasks"`
}
