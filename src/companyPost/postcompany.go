package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: this tool should be correct, however the backend needs to be updated
//Check this tool at a later date to confirm working status
func main() {
	//This tool expects the user to have first run the initSession tool
	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}

	s := &portal.Session{}
	err = json.Unmarshal(temp, s)
	if err != nil {
		log.Fatalln(err)
	}

	//This tool expects the user to have a json formatted
	//company file for the company they desire to create on the portal
	temp, err = ioutil.ReadFile("./.nv-createcompany.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Company{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	company, err := createCompany(c, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(company)
}

func createCompany(c *portal.Company, s *portal.Session) (string, error) {
	result, err := s.CreateCompany(c)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", err
	}

	company := string(str)

	return company, nil
}
