package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have first run the initSession tool
	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}

	s := &portal.Session{}
	err = json.Unmarshal(temp, s)
	if err != nil {
		log.Fatalln(err)
	}

	//This tool expects the user to have a json formatted
	//contact file for the contact they desire to create on the portal
	temp, err = ioutil.ReadFile("./.nv-createcontact.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Contact{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	err = createContact(c, s)
	if err != nil {
		log.Fatalln(err)
	}
}

func createContact(c *portal.Contact, s *portal.Session) error {
	err := s.CreateContact(c)
	if err != nil {
		return err
	}

	return nil
}
