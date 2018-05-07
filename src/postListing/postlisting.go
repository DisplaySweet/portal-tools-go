package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//Create listing doesn't yet function, need a better understanding
//of how to manage listing structure and creation in the backend
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

	//This tool expects the user to have created a json formatted
	//listing file for the listing they desire to create on the portal
	temp, err = ioutil.ReadFile("./.listing.json")
	if err != nil {
		log.Fatalln(err)
	}

	tempListing := map[string]portal.Listing{}
	err = json.Unmarshal(temp, &tempListing)
	if err != nil {
		log.Fatalln(err)
	}

	listing := &portal.Listing{}

	for id, obj := range tempListing {
		obj.ID = id
		*listing = obj
	}

	err = createListing(listing, s)
	if err != nil {
		log.Fatalln(err)
	}

}

func createListing(listing *portal.Listing, s *portal.Session) error {
	err := s.CreateListing(listing)
	if err != nil {
		return err
	}

	return nil
}
