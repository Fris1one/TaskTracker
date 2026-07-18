package main

import (
	"TaskTracker/internal/schedule"
	"TaskTracker/internal/storage"
	"TaskTracker/internal/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	s, err := storage.Load()
	if err != nil {
		s = schedule.New()
	}
	a := app.NewWithID("TaskTracker")
	
	ui.New(a, s).Run()
}
