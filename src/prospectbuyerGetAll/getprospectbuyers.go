package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: ProspectBuyers endpoints are WIP
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

	companyNames, err := returnProspectBuyers(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, company := range companyNames {
		fmt.Println(company)
	}
}

func returnProspectBuyers(s *portal.Session) ([]string, error) {

	pb, err := s.GetProspectBuyers()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, buyer := range pb {
		str := fmt.Sprintf(
			"%v : %v %v",
			buyer.ID,
			buyer.Firstname,
			buyer.Lastname)
		list = append(list, str)
	}

	return list, nil
}
