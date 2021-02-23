package timelog

import (
	"fmt"
	"time"
)

type Entry struct {
	Start time.Time
	End time.Time
	Title string
	Description string
	Id int
}

type Log struct {
	Series []Entry
	Id int
}

func (log Log) print() {
	fmt.Print("Here be printing log")
}