package storage

import (
	"encoding/json"
	"os"

	"TaskTracker/internal/schedule"
)

func Save(s *schedule.Schedule) error {
	var data SaveData
	for _, tasks := range s.Dates() {
		for _, t := range tasks {
			data.Tasks = append(
				data.Tasks,
				ToDTO(t),
			)
		}
	}
	bytes, err := json.MarshalIndent(
		data,
		"",
		"    ",
	)
	if err != nil {
		return err
	}
	file, err := dataPath()
	if err != nil {
		return err
	}
	return os.WriteFile(
		file,
		bytes,
		0644,
	)
}
