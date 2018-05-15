package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have a json formatted
	//user file for the user they desire to update on the portal
	//which you have retrieved with the getUser tool.
	//the file must be manually edited with the changes you wish to update
	temp, err := ioutil.ReadFile("./.nv-user.json")
	if err != nil {
		log.Fatalln(err)
	}

	u := &portal.User{}
	err = json.Unmarshal(temp, u)
	if err != nil {
		log.Fatalln(err)
	}

	err = updateUser(u)
	if err != nil {
		log.Fatalln(err)
	}
}

func updateUser(u *portal.User) error {
	err := u.Update()
	if err != nil {
		return err
	}

	return nil
}
