package ui

import (
	"TaskTracker/internal/res"
	"TaskTracker/internal/task"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	goldenStar fyne.Resource
	blackStar  fyne.Resource
)

var (
	onOffButton  fyne.Resource
	floralCircle fyne.Resource
	checkMark    fyne.Resource
)

func init() {
	goldenStar = res.ResourceGoldenStarPng
	blackStar = res.ResourceBlackStarPng
	onOffButton = res.ResourceOnOffButtonPng
	floralCircle = res.ResourceFloralCirclePng
	checkMark = res.ResourceCheckMarkPng
}

type starButton struct {
	widget.Icon
	task *task.Task
	app  *App
}

func newStarButton(app *App, t *task.Task) *starButton {
	star := &starButton{
		task: t,
		app:  app,
	}
	star.ExtendBaseWidget(star)
	star.updateIcon()
	return star
}
func (sb *starButton) Tapped(_ *fyne.PointEvent) {
	sb.task.SetFavourite(!sb.task.IsFavourite())
	sb.app.update()
}

func (sb *starButton) updateIcon() {
	if sb.task.IsFavourite() {
		sb.SetResource(goldenStar)
	} else {
		sb.SetResource(blackStar)
	}
}

type stageButton struct {
	widget.Icon
	task *task.Task
	app  *App
}

func newStageButton(app *App, t *task.Task) *stageButton {
	stage := &stageButton{
		task: t,
		app:  app,
	}
	stage.ExtendBaseWidget(stage)
	stage.updateIcon()
	return stage
}

func (sb *stageButton) Tapped(_ *fyne.PointEvent) {
	if sb.task.Stage() == 2 {
		sb.task.SetStage(0)
	} else {
		sb.task.SetStage(
			sb.task.Stage() + 1,
		)
	}
	sb.app.update()
}
func (sb *stageButton) updateIcon() {
	switch sb.task.Stage() {
	case 0:
		sb.SetResource(onOffButton)
	case 1:
		sb.SetResource(floralCircle)
	case 2:
		sb.SetResource(checkMark)
	}
}
