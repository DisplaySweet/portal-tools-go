package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have a json formatted
	//company file for the company they desire to update on the portal
	//which you have retrieved with the getCompany tool.
	//the file must be manually edited with the changes you wish to update
	temp, err := ioutil.ReadFile("./.nv-company.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Company{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	err = updateCompany(c)
	if err != nil {
		log.Fatalln(err)
	}
}

func updateCompany(c *portal.Company) error {
	err := c.Update()
	if err != nil {
		return err
	}

	return nil
}
