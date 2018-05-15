package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	//This tool expects the user to have a json formatted
	//user file for the user they desire to delete on the portal
	//which you likely have retrieved with the getUser tool.
	temp, err := ioutil.ReadFile("./.nv-deluser.json")
	if err != nil {
		log.Fatalln(err)
	}

	u := &portal.User{}
	err = json.Unmarshal(temp, u)
	if err != nil {
		log.Fatalln(err)
	}

	err = deleteUser(u)
	if err != nil {
		log.Fatalln(err)
	}
}

func deleteUser(u *portal.User) error {
	err := u.Delete()
	if err != nil {
		return err
	}

	return nil
}
