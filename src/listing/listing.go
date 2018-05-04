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
	temp, err := ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}
	s := &portal.Session{}
	json.Unmarshal(temp, s)

	bytes, err := ioutil.ReadAll(os.Stdin)
	id := string(bytes)

	listing, err := getListingInfo(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(listing)
}

func getListingInfo(id string, s *portal.Session) (string, error) {
	listing, err := s.GetListingByID(id)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf(
		"%v : %v, %v, $%v",
		listing.ID,
		listing.Name,
		listing.Availability,
		listing.Price)

	return str, nil
}
