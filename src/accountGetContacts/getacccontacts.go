package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//As there is no functionality to get one specific account contact
//this tool must therefore return a complete list of all the account-contact
//data, rather than just id-name pairs ----- unless the ID maps to a contact,
//where we can then grep out this ID and feed it into getContact tool??
func main() {
	// temp, err := ioutil.ReadFile("./.nv-session.json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// s := &portal.Session{}
	// json.Unmarshal(temp, s)

	temp, err := ioutil.ReadFile("./.nv-account.json")
	if err != nil {
		log.Fatalln(err)
	}
	a := &portal.Account{}
	err = json.Unmarshal(temp, a)
	if err != nil {
		log.Fatalln(err)
	}

	accountcontacts, err := getAccountContacts(a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(accountcontacts)
}

//This tool requires an account object, thus the get account tool needs to be run first
//
func getAccountContacts(a *portal.Account) (string, error) {
	acs, err := a.GetOwnedContacts()
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(acs, "", "\t")

	aclist := string(str)

	return aclist, nil
}
