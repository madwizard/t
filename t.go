package main

import (
	"github.com/madwizard/t/util"
	"github.com/madwizard/t/timelog"
	"time"
)

func main() {

	util.Usage()

	data := timelog.Entry{
		Start: time.Now(),
		End: time.Now(),
		Title: "Test",
		Description: "Longer test",
		Id: 1,
	}

	data.Save()
	data.Print()
}
