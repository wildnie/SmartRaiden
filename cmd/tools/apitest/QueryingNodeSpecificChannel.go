package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//query a node specified channel
func QueryingNodeSpecificChannel(url string, channel string) (Channel NodeChannel, Status string, err error) {
	var resp *http.Response
	var count int
	for count = 0; count < MaxTry; count = count + 1 {
		resp, err = http.Get(url + "/api/1/channels/" + channel)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if count >= MaxTry {
		Status = "504 TimeOut"
	}
	if resp != nil {
		p, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(p, &Channel)
		Status = resp.Status
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	return
}

//test for querying a node specified channel
func QueryingNodeSpecificChannelTest(url string) {

	var existedchannel string
	existedchannel = ""
	start := time.Now()
	ShowTime()
	log.Println("Start Querying Node Existed Specific Channel")
	Channels, _, _ := QueryingNodeAllChannels(url)
	if Channels != nil {
		if len(Channels) >= 1 {
			existedchannel = Channels[0].ChannelAddress
		}
	}
	//fmt.Printf("Existed Specific Channel:%s\n", existedchannel)
	_, Status, err := QueryingNodeSpecificChannel(url, existedchannel)
	ShowError(err)
	ShowQueryingNodeSpecificChannelMsgDetail(Status)
	switch Status {
	case "200 OK":
		log.Println("Test pass:querying node existed Specific channel Success!")
	default:
		log.Println("Test failed:querying node1 existed channels ", existedchannel, " Failure:", Status)
		if HalfLife {
			//log.Fatal("HalfLife,exit")
			time.Sleep(80 * time.Second)
		}
	}
	log.Println("Start Querying Node  not existed Specific Channel")
	_, Status, err = QueryingNodeSpecificChannel(url, "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	ShowError(err)
	ShowQueryingNodeSpecificChannelMsgDetail(Status)
	switch Status {
	case "404 Not Found":
		log.Println("Test pass:querying node not existed Specific channel Success!")
	default:
		log.Println("Test failed:querying node1 not existed channels Failure:", Status)
		if HalfLife {
			log.Fatal("HalfLife,exit")
		}
	}
	duration := time.Since(start)
	ShowTime()
	log.Println("time used:", duration.Nanoseconds()/1000000, " ms")
}

//display the details of the error
func ShowQueryingNodeSpecificChannelMsgDetail(Status string) {
	switch Status {
	case "200 OK":
		log.Println("Successful Query")
	case "404 Not Found":
		log.Println("The channel does not exist")
	case "500 Server Error":
		log.Println("Internal Raiden node error")
	case "504 TimeOut":
		log.Println("No response,timeout")
	default:
		log.Println("Unknown error,QueryingNodeSpecificChannel:", Status)
	}
}
