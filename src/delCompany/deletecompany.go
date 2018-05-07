package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: still need to test this tool once the create
//tool successfully posts a new test company
func main() {
	//This tool expects the user to have a json formatted
	//company file for the company they desire to delete on the portal
	//which you likely have retrieved with the getCompany tool.
	temp, err := ioutil.ReadFile("./.nv-delcompany.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Company{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	err = deleteCompany(c)
	if err != nil {
		log.Fatalln(err)
	}
}

func deleteCompany(c *portal.Company) error {
	err := c.Delete()
	if err != nil {
		return err
	}

	return nil
}
