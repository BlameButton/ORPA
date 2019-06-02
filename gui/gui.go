package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/blamebutton/orpa/parser"
	"github.com/dustin/go-humanize"
)

type OrpaWindow struct {
	Replay *parser.ReplayData
	window fyne.Window

	NicknameField   *widget.Entry
	TotalScoreField *widget.Entry
}

func (w *OrpaWindow) buildWindow() {
	a := app.New()

	settings := a.Settings()
	settings.SetTheme(&MaterialTheme{})

	w.window = a.NewWindow("OSU! rawReplay Parser")
	w.window.Resize(fyne.Size{Width: 300, Height: 300})

	w.NicknameField = getReadOnlyEntry("")
	w.TotalScoreField = getReadOnlyEntry("")
	form := widget.NewForm()
	form.Append("Nickname", w.NicknameField)
	form.Append("Total Score", w.TotalScoreField)

	quitButton := widget.NewButton("Quit", func() { a.Quit() })
	w.window.SetContent(widget.NewVBox(form, quitButton))
}

func NewOrpaWindow(replay *parser.ReplayData) *OrpaWindow {
	window := &OrpaWindow{Replay: replay}
	window.buildWindow()
	return window
}

func (w *OrpaWindow) Start() {
	w.Update()
	w.window.ShowAndRun()
}

func (w *OrpaWindow) Update() {
	w.NicknameField.SetText(w.Replay.PlayerName)
	w.TotalScoreField.SetText(humanize.Comma(int64(w.Replay.TotalScore)))
}

func (w *OrpaWindow) SetReplay(replay *parser.ReplayData) {
	w.Replay = replay
	w.Update()
}

func getReadOnlyEntry(value string) *widget.Entry {
	return &widget.Entry{
		Text:     value,
		ReadOnly: true,
	}
}
