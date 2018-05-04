package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DisplaySweet/portal-go/src"
)

type Config struct {
	APIt string `json:"Key"`
	APIc string `json:"Company"`
	APIp string `json:"Project"`
	URL  string `json:"TargetEnv"`
}

func main() {
	//main expects a creds.json file to be piped into it
	// 	cat creds.json | initSession.go
	//You should then output the stdout of the package to a file
	// 	cat creds.json | initSession.go > .nv-session
	bytes, err := ioutil.ReadAll(os.Stdin)

	creds := &Config{}
	err = json.Unmarshal(bytes, creds)
	if err != nil {
		log.Fatalln(err)
	}

	a := &portal.Auth{
		APIKey:         creds.APIt,
		Company:        "",
		PortalEndpoint: creds.URL,
	}

	s := &portal.Session{
		Project: portal.Project{
			ID: creds.APIp,
		},
		Company: portal.Company{
			ID: creds.APIc,
		},
		Auth: *a,
	}

	minifiedSession, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(minifiedSession))
}
