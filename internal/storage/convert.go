package storage

import "TaskTracker/internal/task"

func ToDTO(t *task.Task) TaskDTO {
	deadline := t.DeadLine()
	return TaskDTO{
		Name:      string(t.Name()),
		Priority:  int(t.Priority()),
		Day:       deadline.Day(),
		Month:     deadline.Month(),
		Year:      deadline.Year(),
		Favourite: t.IsFavourite(),
		Stage:     t.Stage(),
	}
}

func FromDTO(dto TaskDTO) (*task.Task, error) {
	t, err := task.NewWithPriority(
		dto.Name,
		dto.Priority,
	)
	if err != nil {
		return nil, err
	}
	t.SetDeadLine(
		dto.Day,
		dto.Month,
		dto.Year,
	)
	t.SetFavourite(dto.Favourite)
	err = t.SetStage(dto.Stage)
	if err != nil {
		return nil, err
	}
	return t, nil
}
