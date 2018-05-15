package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	portal "portal-go/src"
)

func main() {
	//This tool expects the user to have first run the initSession tool
	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}

	s := &portal.Session{}
	json.Unmarshal(temp, s)
	if err != nil {
		log.Fatalln(err)
	}

	users, err := getEOIUsers(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}

func getEOIUsers(s *portal.Session) ([]string, error) {
	users, err := s.GetEOIUsers()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)
	list = append(list, "EOI USERS")

	for _, user := range users {
		str := fmt.Sprintf(
			"\t%v : %v",
			user.ID,
			user.Email,
		)
		list = append(list, str)
	}

	return list, nil
}
