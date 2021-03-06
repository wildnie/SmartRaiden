package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(c *http.Client, req *http.Request) (Status string, body []byte, err error) {
	var buf [4096]byte
	var n = 0
	if req.Body != nil {
		n, err = req.Body.Read(buf[:])
		if err != nil {
			log.Printf("req %s body err %s", req.URL.String(), err)
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	}
	if n > 0 {
		log.Printf("send -> %s %s:\n%s\n", req.Method, req.URL.String(), string(buf[:n]))
	} else {
		log.Printf("send -> %s %s\n", req.Method, req.URL.String())
	}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	Status = resp.Status
	n, err = resp.Body.Read(buf[:])
	if err != nil {
		log.Printf("req % read err %s", req.URL.String(), err)
	}
	body = buf[:n]
	if len(body) > 0 {
		log.Printf("receive <- :\n%s\n", string(body))
	}
	req.Body.Close()
	resp.Body.Close()
	return
}
