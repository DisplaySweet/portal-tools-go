package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	temp, err := ioutil.ReadFile("./.user_creds.json")
	if err != nil {
		log.Fatalln(err)
	}
	ua := &portal.UserAuth{}
	json.Unmarshal(temp, ua)

	temp, err = ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}
	s := &portal.Session{}
	json.Unmarshal(temp, s)

	headers, err = s.AuthAsUser(ua)
	if err != nil {
		log.Fatalln(err)
	}
}
