package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	portal "portal-go/src"
)

func main() {
	temp, err := ioutil.ReadFile("./.nv-company.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Company{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	accsconts, err := getCompanyAccountsContacts(c)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range accsconts {
		fmt.Println(entry)
	}
}

func getCompanyAccountsContacts(c *portal.Company) ([]string, error) {
	accounts, contacts, err := c.GetAccountsContacts()
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
