package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	calls    = 10 // How many calls
	interval = 50 // Requests interval
	url      = "http://localhost:3001/api/v1/stress"
	method   = "POST"
)

const (
	logHeader = false
	logBody   = false
)

func task(i int) {

	// Make a request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("X-Powered-By", "Express")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
	req.Header.Set("Expires", "-1")
	req.Header.Set("Pragma", "no-cache")

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
			time.Sleep(interval * time.Millisecond)
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
