package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have first run the initSession tool
	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	s := &portal.Session{}
	err = json.Unmarshal(temp, s)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v. Payload:\n%v", err, string(temp))
	}

	companyNames, err := getCompanyNamesIDs(s)
	if err != nil {
		log.Fatalf("Error getting company names and IDs: %v", err)
	}

	for _, company := range companyNames {
		fmt.Println(company)
	}
}

func getCompanyNamesIDs(s *portal.Session) ([]string, error) {

	companies, err := s.GetAllCompanies()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, company := range companies {
		str := fmt.Sprintf(
			"%v : %v",
			company.ID,
			company.Name)
		list = append(list, str)
	}

	return list, nil
}
