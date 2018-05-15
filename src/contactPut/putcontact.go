package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have a json formatted
	//company file for the company they desire to update on the portal
	//which you have retrieved with the getCompany tool.
	//the file must be manually edited with the changes you wish to update
	temp, err := ioutil.ReadFile("./.nv-contact.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Contact{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := updateContact(c)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)

}

func updateContact(c *portal.Contact) (string, error) {
	result, err := c.Update()
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", err
	}

	contact := string(str)

	return contact, nil
}
