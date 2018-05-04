package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {

	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}
	s := &portal.Session{}
	json.Unmarshal(temp, s)
	companyNames, err := returnCompanyNamesIDs(s)
	if err != nil {
		log.Fatalln(err)
	}
	for _, company := range companyNames {
		fmt.Println(company)
	}
}

func returnCompanyNamesIDs(s *portal.Session) ([]string, error) {

	companies, err := s.GetAllCompanies()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, company := range companies {
		str := fmt.Sprintf("%v : %v", company.ID, company.Name)
		list = append(list, str)
	}

	return list, nil
}
