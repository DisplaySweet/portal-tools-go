package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: singular listing endpoint is WIP
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

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	id := string(bytes)

	listing, err := getListingInfo(id, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(listing)
}

func getListingInfo(id string, s *portal.Session) (string, error) {
	temp, err := s.GetListingByID(id)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(temp, "", "\t")
	if err != nil {
		return "", err
	}

	listing := string(str)

	return listing, nil
}
