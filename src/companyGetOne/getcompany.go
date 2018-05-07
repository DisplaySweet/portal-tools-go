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
	if err != nil {
		log.Fatalln(err)
	}

	//This tool expects an ID of an company to be piped | through CLI
	//this can be done be first using getCompanies, finding the company you want
	//and using it's ID
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	id := string(bytes)

	company, err := getCompanyData(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(company)
}

func getCompanyData(id string, s *portal.Session) (string, error) {
	temp, err := s.GetCompany(id)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(temp, "", "\t")
	if err != nil {
		return "", err
	}

	company := string(str)

	return company, nil
}
