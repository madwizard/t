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
	fmt.Print("%+v\n", e)
}

func (e Entry) Save() error {
	save, _ := json.MarshalIndent(e, "", "")
	err := ioutil.WriteFile("log.txt", save, 0640)
	if err != nil {
		log.Fatalf("Couldn't write to log. %s", err)
	}
	return nil
}