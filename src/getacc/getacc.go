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

	bytes, err := ioutil.ReadAll(os.Stdin)
	id := string(bytes)

	listing, err := getAccount(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(listing)
}

func getAccount(id string, s *portal.Session) (string, error) {
	account, err := s.GetAccount(id)
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf(
		"%v : %v, Email: %v, Owner: %v",
		account.ID,
		account.Name,
		account.Email,
		account.OwnerID)

	return str, nil
}
