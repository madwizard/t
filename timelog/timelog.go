package timelog

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"time"
)

type Entry struct {
	Id          int       `json:"id"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
}

type Log struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Entries []Entry `json:"entries"`
}

func (l Log) Print() {

}

func (l Log) Save(path string) error {
	saveFile := path + "/log.txt"
	save, _ := json.MarshalIndent(l, "", "")
	err := ioutil.WriteFile(saveFile, save, 0640)
	if err != nil {
		log.Fatalf("save: couldn't write to log: %s", err)
	}
	return nil
}

func AddLogType(f fyne.App) {
	win := f.NewWindow("Add Log Type")
	logTitle := widget.NewLabel("Log Title")
	logDesc := widget.NewLabel("Log Description")
	closeBttn := widget.NewButton("Close", func() {
		win.Close()
		return
	})
	content := container.New(layout.NewHBoxLayout(), logTitle, logDesc, closeBttn)

	centered := container.New(layout.NewHBoxLayout())
	win.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	win.Resize(fyne.NewSize(200, 200))
	win.Show()
}

func ShowAllTimeEntries(f fyne.App, l Log) {
	win := f.NewWindow("Print all time entries")
	logTimes := widget.NewLabel("Entered times")
	closeBttn := widget.NewButton("Close", func() {
		win.Close()
		return
	})
	content := container.New(layout.NewHBoxLayout(), logTimes, closeBttn)
	centered := container.New(layout.NewHBoxLayout())
	win.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	win.Resize(fyne.NewSize(200, 200))
	win.Show()
}
