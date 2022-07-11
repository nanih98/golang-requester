package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/akamensky/argparse"

	//"log"
	"net/http"
	"os"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

// MakeGetRequest execute http get connection to the given domain
func MakeGetRequest(url string) {
	req, err := http.Get(url)

	if err != nil {
		//fmt.Println("Errored when sending request to the server")
		//return
		log.Fatal(err)
	}
	defer req.Body.Close()

	fmt.Printf("Request done to %s. Status code: %d \n", url, req.StatusCode)
}

func main() {
	start := time.Now()

	//Parse
	parser := argparse.NewParser("requester", "Send massive HTTP GET requests to an endpoint")
	domain := parser.String("d", "domain", &argparse.Options{Required: true, Help: "Domain to scan"})
	maxRequests := parser.Int("m", "max-requests", &argparse.Options{Required: false, Help: "Max requests you want to send", Default: 100})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Print(string(colorCyan), parser.Usage((err)), string(colorReset))
	}

	var wg sync.WaitGroup
	wg.Add(*maxRequests)

	//Requester block
	for i := 0; i < *maxRequests; i++ {
		go func() {
			MakeGetRequest(*domain)
			wg.Done()
		}()
	}

	wg.Wait()

	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
