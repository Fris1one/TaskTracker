package ui

import (
	"TaskTracker/internal/schedule"
	"TaskTracker/internal/storage"
	"TaskTracker/internal/task"
"TaskTracker/internal/res"
	"fyne.io/fyne/v2"
)

type Date struct {
	task.TaskDeadLine
}

type App struct {
	app        fyne.App
	window     fyne.Window
	schedule   *schedule.Schedule
	chosenDate Date
}

func New(a fyne.App, s *schedule.Schedule) *App {
	w := a.NewWindow("TaskTracker")
	w.Resize(fyne.NewSize(1200, 675))
	Logo := res.ResourceLogoJpg
	w.SetIcon(Logo)
	a.SetIcon(Logo)
	return &App{
		app:      a,
		window:   w,
		schedule: s,
		chosenDate: Date{
			TaskDeadLine: *task.NewDeadLine(),
		},
	}
}

func (a *App) Run() {
	a.window.SetContent(a.buildMainContent())
	a.window.ShowAndRun()
}

func (a *App) update() {
	err := storage.Save(a.schedule)
	if err != nil {
		println(err.Error())
	}
	a.window.SetContent(a.buildMainContent())
}
