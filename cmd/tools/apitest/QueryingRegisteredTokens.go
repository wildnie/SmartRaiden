package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//query registered Token
func QueryingRegisteredTokens(url string) (Tokens []string, Status string, err error) {
	var resp *http.Response
	var count int
	for count = 0; count < MaxTry; count = count + 1 {
		resp, err = http.Get(url + "/api/1/tokens")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if count >= MaxTry {
		Status = "504 TimeOut"
	} else {
		p, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(p, &Tokens)
		Status = resp.Status
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	return
}

//test for querying registered Token
func QueryingRegisteredTokensTest(url string) {
	start := time.Now()
	ShowTime()
	log.Println("Start Querying Registered Tokens")
	_, Status, err := QueryingRegisteredTokens(url)
	ShowError(err)
	//display the details of the error
	ShowQueryingRegisteredTokensMsgDetail(Status)
	switch Status {
	case "200 OK":
		log.Println("Test pass:Querying Registered Tokens Success!")
	default:
		log.Println("Test failed:Querying Registered Tokens Failure! %s", Status)
		if HalfLife {
			log.Fatal("HalfLife,exit")
		}
	}
	duration := time.Since(start)
	ShowTime()
	log.Println("time used:", duration.Nanoseconds()/1000000, " ms")
}

//display the details of the error
func ShowQueryingRegisteredTokensMsgDetail(Status string) {
	switch Status {
	case "200 OK":
		log.Println("Successful query")
	case "404 Not Found":
		log.Println("The token does not exist")
	case "500 Server Error":
		log.Println("Internal Raiden node error")
	case "504 TimeOut":
		log.Println("No response,timeout")
	default:
		fmt.Printf("Unknown error,QueryingRegisteredTokens Failure:%s\n", Status)
	}
}
