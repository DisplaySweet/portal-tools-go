package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DisplaySweet/portal-go/src"
)

//TODO: this tool should be correct, however the backend needs to be updated
//Check this tool at a later date to confirm working status
func main() {
	//Initialise the project
	temp, err := ioutil.ReadFile("./.nv-project.json")
	if err != nil {
		log.Fatalln(err)
	}

	p := &portal.Project{}
	err = json.Unmarshal(temp, p)
	if err != nil {
		log.Fatalln(err)
	}

	//Need to pipe in the company you want to add to the current project
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	id := string(bytes)

	success, err := addCompanyToProject(p, id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(success)
}

func addCompanyToProject(p *portal.Project, id string) (string, error) {
	err := p.AddCompany(id)
	if err != nil {
		return "", err
	}

	success := fmt.Sprintf("Successfully added Company: %v to Project: %v", id, p.ID)

	return success, nil
}
