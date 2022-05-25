package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"log"
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

func MakeGetRequest(url string, ch chan<- string) {
	req, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	req.Body.Close()

	ch <- fmt.Sprintf("Request done to %s. Status code: %d", url, req.StatusCode)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	//Parse
	parser := argparse.NewParser("requester", "Send massive HTTP GET requests to an endpoint")
	domain := parser.String("d", "domain", &argparse.Options{Required: true, Help: "Domain to scan"})
	max_requests := parser.Int("m", "max-requests", &argparse.Options{Required: false, Help: "Max requests you want to send", Default: 100})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Print(string(colorCyan), parser.Usage((err)), string(colorReset))
	}

	//Requester block
	for i := 0; i < *max_requests; i++ {
		go MakeGetRequest(*domain, ch)
	}

	for i := 0; i < *max_requests; i++ {
		log.Println(<-ch)
	}

	close(ch)
	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
