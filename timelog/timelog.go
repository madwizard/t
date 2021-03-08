package timelog

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
)

type Entry struct {
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Id int `json:"id"`
}

func (e Entry) Print() {
	fmt.Println("%+v", e)
}

func (e Entry) Save(path string) error {
	saveFile := path + "/log.txt"
	save, _ := json.MarshalIndent(e, "", "")
	err := ioutil.WriteFile(saveFile, save, 0640)
	if err != nil {
		log.Fatalf("Couldn't write to log. %s", err)
	}
	return nil
}