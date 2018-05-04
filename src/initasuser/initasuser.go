package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

//To be used like so: go run ./initasuser/initasuser.go >./.nv-auth.json
func main() {
	temp, err := ioutil.ReadFile("./.user_creds.json")
	if err != nil {
		log.Fatalln(err)
	}
	ua := &portal.UserAuth{}
	json.Unmarshal(temp, ua)

	temp, err = ioutil.ReadFile("./.nv-session.json")
	if err != nil {
		log.Fatalln(err)
	}
	s := &portal.Session{}
	json.Unmarshal(temp, s)

	rh := &portal.ResponseHeaders{}
	rh, err = s.AuthAsUser(ua)
	if err != nil {
		log.Fatalln(err)
	}

	minifiedAuth, err := json.Marshal(rh)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(minifiedAuth))
}
