package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
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

	temp, err = ioutil.ReadFile("./.nv-companyuserlist.json")
	if err != nil {
		log.Fatalln(err)
	}

	var a []*portal.UserAdd
	err = json.Unmarshal(temp, &a)
	if err != nil {
		log.Fatalln(err)
	}

	err = addCompanyUsers(c, a)
	if err != nil {
		log.Fatalln(err)
	}
}

func addCompanyUsers(c *portal.Company, a []*portal.UserAdd) error {
	err := c.AddUsers(a)
	if err != nil {
		return err
	}

	return nil
}
