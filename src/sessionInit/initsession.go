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
	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
	}

	creds := &Config{}
	err = json.Unmarshal(bytes, creds)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v. Payload was:\n%v", err, string(bytes))
	}

	a := &portal.Auth{
		APIKey:         creds.APIt,
		UserAPIKey:     creds.APIt,
		Company:        "",
		PortalEndpoint: creds.URL,
	}

	s := &portal.Session{
		ProjectID:         creds.APIp,
		CompanyID:         creds.APIc,
		Auth:              *a,
		DumpErrorPayloads: false,
	}

	minifiedSession, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatalf("Error when beautifying output: %v", err)
	}

	fmt.Print(string(minifiedSession))
}
