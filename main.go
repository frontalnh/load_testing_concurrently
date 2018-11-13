package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	calls    = 50
	interval = 70
	url      = "http://localhost:3001/api/v1/stress"
	method   = "GET"
	cookie   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViYTFhMmY5NDYxZmEyMGQxZjQwY2I3YyIsImlhdCI6MTU0MTk5MDkzOCwiZXhwIjoxMDE4MTk5MDkzOH0.3gHE8_7dkK_8zwC2Q9G0m1D8PNZDNQFIAftvucpA3rU"
)

const (
	logHeader = false
	logBody   = true
)

func task(i int) {

	// Make a request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
	start := time.Now()
	ack := make(chan bool, calls)

	for i := 0; i < calls; i++ {
		fmt.Println("🔍   Request", i)
		// time.Sleep(interval * time.Millisecond)
		go func(arg int) {
			task(arg)
			ack <- true
		}(i)
	}

	for i := 0; i < calls; i++ {
		<-ack
	}

	elapsed := time.Since(start)
	fmt.Println("\n\n⌛  >>", elapsed)
}
