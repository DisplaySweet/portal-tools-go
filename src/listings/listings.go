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

	listings, err := getListingsInfo(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, listing := range listings {
		fmt.Println(listing)
	}
}

func getListingsInfo(s *portal.Session) ([]string, error) {

	listings, err := s.GetListings()
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 0)

	for _, listing := range listings {
		str := fmt.Sprintf(
			"%v : %v, %v, $%v",
			listing.ID,
			listing.Name,
			listing.Availability,
			listing.Price)

		list = append(list, str)
	}

	return list, nil
}
