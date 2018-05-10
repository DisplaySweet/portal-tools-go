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
	json.Unmarshal(temp, s)
	if err != nil {
		log.Fatalln(err)
	}

	users, err := getUsers(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}

//This tool should really only return a limited set of information,
//whereas an individual getUser(id) should return all.
func getUsers(s *portal.Session) ([]string, error) {

	users, err := s.GetUsers()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, user := range users {
		str := fmt.Sprintf(
			"%v : %v",
			user.ID,
			user.Email)
		list = append(list, str)
	}

	return list, nil
}
