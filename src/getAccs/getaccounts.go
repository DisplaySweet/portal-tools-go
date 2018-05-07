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

	listings, err := getAccounts(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, listing := range listings {
		fmt.Println(listing)
	}
}

//This tool should really only return a limited set of information,
//whereas an individual getAccount(id) should return all.
func getAccounts(s *portal.Session) ([]string, error) {

	accounts, err := s.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, account := range accounts {
		str := fmt.Sprintf(
			"%v : %v, Email: %v, Owner: %v",
			account.ID,
			account.Name,
			account.Email,
			account.OwnerID)

		list = append(list, str)
	}

	return list, nil
}
