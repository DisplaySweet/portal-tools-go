package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	//an id needs to be piped in
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	id := string(bytes)

	user, err := getUser(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}

func getUser(id string, s *portal.Session) (string, error) {
	temp, err := s.GetUserByID(id)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(temp, "", "\t")
	if err != nil {
		return "", err
	}

	user := string(str)

	return user, nil
}
