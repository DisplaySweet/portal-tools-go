package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: this tool should be correct, however the backend needs to be updated
//Check this tool at a later date to confirm working status
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
	//user file for the user they desire to create on the portal
	temp, err = ioutil.ReadFile("./.nv-createuser.json")
	if err != nil {
		log.Fatalln(err)
	}

	u := &portal.User{}
	err = json.Unmarshal(temp, u)
	if err != nil {
		log.Fatalln(err)
	}

	user, err := createUser(u, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}

func createUser(u *portal.User, s *portal.Session) (string, error) {
	result, err := s.CreateUser(u)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", err
	}

	user := string(str)

	return user, nil
}
