package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/madwizard/t/timelog"
	"image/color"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

// Config directory and logs path
var path string
var mode os.FileMode = 0770

// Need to take care of program's directory
func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("init: couldn't get current user: %s", err)
	}
	path = usr.HomeDir + "/.timelog"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
}

func main() {

	data := timelog.Log{
		Id:    1,
		Title: "Some type of task",
		Entries: []timelog.Entry{
			timelog.Entry{
				Id:          1,
				Start:       time.Now(),
				End:         time.Now(),
				Title:       "Test",
				Description: "Longer test",
			},
			timelog.Entry{
				Id:          2,
				Start:       time.Now(),
				End:         time.Now(),
				Title:       "Second task",
				Description: "Very long description of second task",
			},
		},
	}

	a := app.New()
	w := a.NewWindow("TimeTrack")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("timetrack", fyne.NewMenuItem("Show", func() {
			w.Show()
		}))
		desk.SetSystemTrayMenu(m)
	}

	logId := canvas.NewText(strconv.Itoa(data.Id), color.White)
	logTitle := canvas.NewText(data.Title, color.White)
	logEntry := canvas.NewText(data.Entries[1].Title, color.White)
	printTimes := widget.NewButton("Show all times", func() { go timelog.ShowAllTimeEntries(a, data) })
	addBttn := widget.NewButton("Add Log Type", func() { go timelog.AddLogType(a) })
	quitBttn := widget.NewButton("Quit", a.Quit)
	content := container.New(layout.NewHBoxLayout(), logId, logTitle, layout.NewSpacer(), logEntry, printTimes, addBttn, quitBttn)
	centered := container.New(layout.NewHBoxLayout())
	w.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	data.Save(path)
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}
