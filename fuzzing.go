package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	words := ReadFile()
	url := "https://vpnroulette.com/"

	var wg sync.WaitGroup
	wg.Add(len(words))

	for _, word := range words {
		go func() {
			MakeRequest(url, word)
			wg.Done()
		}()
	}

	wg.Wait()

	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func ReadFile() []string {
	var words []string
	//readFile, err := os.Open("/usr/local/share/SecLists/Discovery/Web-Content/directory-list-2.3-small.txt")
	readFile, err := os.Open("/tmp/fuzzing.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()

	return words
}

func MakeRequest(url string, fuz string) {
	req, err := http.Get(url + fuz)

	if err != nil {
		log.Printf("Request Failed: %s", err)
	}

	defer req.Body.Close()

	fmt.Printf("Request done to %s. Status code: %d \n", url+fuz, req.StatusCode)
}
