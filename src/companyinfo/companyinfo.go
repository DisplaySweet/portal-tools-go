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

	//This tool expects the user to pipe in an id through CLI
	bytes, err := ioutil.ReadAll(os.Stdin)
	id := string(bytes)

	company, err := getCompanyData(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(company)
}

func getCompanyData(id string, s *portal.Session) (*portal.Company, error) {
	company, err := s.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return company, nil
}
