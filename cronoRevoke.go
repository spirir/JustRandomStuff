package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	deletes := [10]string{"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXX"} 
	// This array holds the access tokens or refresh tokens we want to revoke
	// any other method of supplying the refresh token or access token can also be used instead of this array
	// a database call could also be used, or reading from a file.
}
	fmt.Println(len(deletes)) //most of the of println is not actually necessary, but it can provide some information during the excution
	//for production these should either be left out or at the most write to an appropriate log

	for i := 0; i <= len(deletes); i++ { 	//this for loop iterates through the array of tokens we want to remove
		fmt.Println(i)						//it builds up the endpoint and payload everytime and can be optimized for production use.
		fmt.Println(deletes[i])

		url := "https://api.cronofy.com/oauth/token/revoke/"
		method := "POST" //must use post for this call

		payload := strings.NewReader(`{
	    "client_id": "XXXXXXXXXXXXX",
	    "client_secret": "XXXXXXXXXXXXXXXXXX",  
	    "token": "` + deletes[i] + `"}
	`)//Client_id and Client_secret can be found in the webinterface

		fmt.Println(payload) // included as a help, not really required

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload) 

		if err != nil {
			fmt.Println(err)	//print the error if any
			return
		}
		req.Header.Add("Content-Type", "application/json")
		//req.Header.Add("Authorization", "Bearer XXXXXXXXXXXXXXXXXXX") //Bearer is not required but can be optained if needed through other API calls

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)	//print the error if any
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)	//print the error if any
			return
		}
		fmt.Println(string(body)) // included as a help, not really required
	}
}
