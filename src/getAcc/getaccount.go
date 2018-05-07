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

	//This tool expects an ID of an account to be piped | through CLI
	//this can be done be first using getAccounts, finding the acc you want
	//and using it's ID
	bytes, err := ioutil.ReadAll(os.Stdin)
	id := string(bytes)

	listing, err := getAccount(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(listing)
}

func getAccount(id string, s *portal.Session) (string, error) {
	account, err := s.GetAccount(id)
	if err != nil {
		return "", err
	}

	str, err := json.Marshal(account)
	if err != nil {
		return "", err
	}

	acc := string(str)

	return acc, nil
}
