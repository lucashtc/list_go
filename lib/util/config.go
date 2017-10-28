package util

import (
	"encoding/json"
	"io/ioutil"
	
)

func ConfigJson() (interface{}) {

	type Config struct {
		HotsDB string
		PortDB string
		UserDb string
		PasswordDB string
		NameDB string
	}
	var s Config

	File, err := ioutil.ReadFile("c:/crud/src/github.com/lucashenriqueteixeira/list_go/config.json")
	VerifyErr(err)

	err = json.Unmarshal(File,&s)
	VerifyErr(err)

	return s
}