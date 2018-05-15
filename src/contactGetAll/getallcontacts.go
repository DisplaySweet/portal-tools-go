package main

import (
	"encoding/json"
	"fmt"
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

	contacts, err := getContacts(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, contact := range contacts {
		fmt.Println(contact)
	}
}

func getContacts(s *portal.Session) ([]string, error) {

	contacts, err := s.GetContacts()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, contact := range contacts {
		str := fmt.Sprintf(
			"%v : %v %v",
			contact.ID,
			contact.Firstname,
			contact.Lastname,
		)
		list = append(list, str)
	}

	return list, nil
}
