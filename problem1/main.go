package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	urls, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("error while opening file: ", err)
		return
	}
	defer urls.Close()
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(urls)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err := client.Get(u)
			if err != nil {
				fmt.Println("[ERROR] Ulaşılamadı", u, err)
				return
			}

			fmt.Printf("[%s] %s\n", resp.Status, u)
			resp.Body.Close()

		}(url)
	}
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scanning file: ", err)
		return
	}

}
