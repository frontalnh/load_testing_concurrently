package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	calls    = 10 // How many calls
	interval = 50 // Requests interval
	url      = "http://localhost:3001/api/v1/transactions/bp/pay"
	method   = "POST"
	cookie   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViYTFhMmY5NDYxZmEyMGQxZjQwY2I3YyIsImlhdCI6MTU0MzE5ODQ2NywiZXhwIjoxMDE4MzE5ODQ2N30.NfYsLDiJHZi6kgJRT5mfAebHwIy6fUirX1kixBWVPws"
)

const (
	logHeader = false
	logBody   = true
)

func task(i int) {

	var jsonStr = []byte(`{
		"buyer": "5ba1a2f9461fa20d1f40cb7c",
		"amount":12,
		"password": "test1234",
		"business": "5bc6a30704f3050023171285"
	}`)

	// Make a request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	_cookie := http.Cookie{Name: "jwtToken", Value: cookie}
	req.Header.Set("X-Powered-By", "Express")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
	req.Header.Set("Expires", "-1")
	req.Header.Set("Pragma", "no-cache")
	req.AddCookie(&_cookie)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if logHeader {
		fmt.Println("✅   [Header] >> ", resp)
		println()
	}

	if logBody {
		ret, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return
		}

		fmt.Println("✅   [BODY] ", string(ret))
		println()
	}

}
func main() {
	swarmCalls := 0
	taskCalls := 0
	for {
		println("==============================================\n")
		taskStart := time.Now()
		ack := make(chan bool, calls)
		swarmCalls++

		for i := 0; i < calls; i++ {
			taskCalls++
			fmt.Println("🔍  Request", taskCalls)
			// Uncomment to make an interval
			time.Sleep(time.Duration(rand.Intn(2)+interval) * time.Millisecond)
			go func(arg int) {
				task(arg)
				ack <- true
			}(i)
		}

		println()

		for i := 0; i < calls; i++ {
			<-ack
		}

		taskElapsed := time.Since(taskStart)
		fmt.Printf("\n⌛  [Swarm #%v] \tElapsed: %v\n", swarmCalls, taskElapsed)
		time.Sleep(1 * time.Second)
	}
}
