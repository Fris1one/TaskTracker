package storage

import (
	"encoding/json"
	"errors"
	"os"

	"TaskTracker/internal/schedule"
)

func Load() (*schedule.Schedule, error) {
	s := schedule.New()
	file, err := dataPath()
	if err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, err
	}
	var data SaveData
	err = json.Unmarshal(
		bytes,
		&data,
	)
	if err != nil {
		return nil, err
	}
	for _, dto := range data.Tasks {
		t, err := FromDTO(dto)
		if err != nil {
			continue
		}
		s.AddTask(t)
	}
	return s, nil
}
