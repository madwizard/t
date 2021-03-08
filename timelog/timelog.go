package timelog

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Entry struct {
	Id int `json:"id"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Title string `json:"title"`
	Description string `json:"desc"`
}

type Log struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Entries []Entry `json:"entries"`
}

func (l Log) Print() {

}

func (l Log) Save(path string) error {
	saveFile := path + "/log.txt"
	save, _ := json.MarshalIndent(l, "", "")
	err := ioutil.WriteFile(saveFile, save, 0640)
	if err != nil {
		log.Fatalf("Couldn't write to log. %s", err)
	}
	return nil
}