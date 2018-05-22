package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: this tool should be correct, however the backend needs to be updated
//Check this tool at a later date to confirm working status
func main() {
	temp, err := ioutil.ReadFile("./.nv-company.json")
	if err != nil {
		log.Fatalln(err)
	}

	c := &portal.Company{}
	err = json.Unmarshal(temp, c)
	if err != nil {
		log.Fatalln(err)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	permissionLevel, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	//This tool expects the user to have a json formatted
	//user file for the user they desire to create on the portal
	temp, err = ioutil.ReadFile("./.nv-companymaster.json")
	if err != nil {
		log.Fatalln(err)
	}

	u := &portal.User{}
	err = json.Unmarshal(temp, u)
	if err != nil {
		log.Fatalln(err)
	}

	user, err := createUser(u, c, permissionLevel)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}

func createUser(u *portal.User, c *portal.Company, permissionLevel int64) (string, error) {
	result, err := c.AddUser(u, permissionLevel)
	if err != nil {
		return "", err
	}

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", err
	}

	company := string(str)

	return company, nil
}
