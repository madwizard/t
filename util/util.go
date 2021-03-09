package util

import ("fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("t add --task TASKNAME --start STARTTIME --end ENDTIME --comment COMMENT")
	fmt.Println("\tEither --start or --end is mandatory. Both can be specified at the same call")
	fmt.Println("\t--comment is optional")
	fmt.Println("t print")
	fmt.Println("\tPrints all log from current month.")
}

func ReadConfig() error {
	return nil
}


func AddLogType(f fyne.App){
	win := f.NewWindow("Add Log Type")
	logTitle := widget.NewLabel("Log Title")
	logDesc := widget.NewLabel("Log Description")
	closeBttn := widget.NewButton("Close", func() {
		win.Hide()
		return
	})
	content := container.New(layout.NewHBoxLayout(), logTitle, logDesc,closeBttn)

	centered := container.New(layout.NewHBoxLayout())
	win.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	win.Resize(fyne.NewSize(200, 200))
	win.Show()

}