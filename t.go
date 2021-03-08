package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/madwizard/t/timelog"
	"github.com/madwizard/t/util"
	"image/color"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
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
		Id:      1,
		Title:   "Some type of task",
		Entries: []timelog.Entry{
			timelog.Entry{
				Id:          1,
				Start:       time.Now(),
				End:         time.Now(),
				Title:       "Test",
				Description: "Longer test",
			},
			timelog.Entry{
				Id: 		 2,
				Start: 		 time.Now(),
				End: 		 time.Now(),
				Title: 		 "Second task",
				Description: "Very long description of second task",
			},
		},

	}

	myApp := app.New()
	myWindow := myApp.NewWindow("TimeTrack")
	text1 := canvas.NewText(strconv.Itoa(data.Id), color.White)
	text2 := canvas.NewText(data.Title, color.White)
	text3 := canvas.NewText(data.Entries[1].Title, color.White)
	addBttn := widget.NewButton("Add Log Type", func() { go util.AddLogType(myApp)})
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3, addBttn)
	centered := container.New(layout.NewHBoxLayout())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	data.Save(path)

	myWindow.Show()
	myApp.Run()
	myApp.Quit()
}
