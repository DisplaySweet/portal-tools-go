package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

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

	info, err := getInfo(s)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range info {
		fmt.Println(entry)
	}
}

func getInfo(s *portal.Session) ([]string, error) {

	accounts, err := s.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	acccompanyprojlist := make([]string, 0, 0)

	//for every agent, first add a line AGENTID : COMPANYID : COMPANYNAME
	for _, account := range accounts {
		str := fmt.Sprintf("%v : %v : %v", account.ID, account.OwnerID, account.Owner.Name)
		acccompanyprojlist = append(acccompanyprojlist, str)
		projlist := account.Owner.Projects
		//list every project that belongs to those usercomanies
		for _, project := range projlist {
			str := fmt.Sprintf(
				"\t%v : %v",
				project.ID,
				project.Name,
			)
			acccompanyprojlist = append(acccompanyprojlist, str)
		}
	}

	return acccompanyprojlist, nil
}

//output should read:
//AGENTID : COMPANYID : COMPANYNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME
//AGENTID : COMPANYID : COMPANYNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME
//	PROJECTID : PROJECTNAME etc...
