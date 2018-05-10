package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DisplaySweet/portal-go/src"
)

func main() {
	temp, err := ioutil.ReadFile("./.nv-account.json")
	if err != nil {
		log.Fatalln(err)
	}
	a := &portal.Account{}
	err = json.Unmarshal(temp, a)
	if err != nil {
		log.Fatalln(err)
	}

	accountdeposits, err := getAccountDeposits(a)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range accountdeposits {
		fmt.Println(entry)
	}
}

//This tool requires an account object, thus the get account tool needs to be run first
func getAccountDeposits(a *portal.Account) ([]string, error) {
	temp, err := a.GetOwnedDeposits()
	if err != nil {
		return nil, err
	}

	deposits := make([]string, 0, 0)
	deposits = append(deposits, fmt.Sprintf("Account: %v", a.ID))

	for _, deposit := range temp {
		str := fmt.Sprintf("DepositID: %v", deposit.ID)
		deposits = append(deposits, str)
	}

	return deposits, nil
}
