package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	portal "portal-go/src"
)

func main() {
	temp, err := ioutil.ReadFile("./.nv-user.json")
	if err != nil {
		log.Fatalln(err)
	}

	u := &portal.User{}
	err = json.Unmarshal(temp, u)
	if err != nil {
		log.Fatalln(err)
	}

	accsconts, err := getUserAccountsContacts(u)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range accsconts {
		fmt.Println(entry)
	}
}

func getUserAccountsContacts(u *portal.User) ([]string, error) {
	accounts, contacts, err := u.GetAccountsContacts()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	list = append(list, "Accounts:")
	for _, acc := range accounts {
		list = append(list, fmt.Sprintf("\t%v, %v", acc.ID, acc.Name))
	}

	list = append(list, "Contacts: ")
	for _, con := range contacts {
		list = append(list, fmt.Sprintf("\t%v, %v", con.ID, con.Email))
	}

	return list, nil
}
