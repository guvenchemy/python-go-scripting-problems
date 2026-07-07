package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func worker(ports chan int, wg *sync.WaitGroup, ip string, timeout time.Duration) {
	defer wg.Done() // İşçi işini bitirdiğinde WaitGroup'tan düşülür.
	// ports kanalından veri geldikçe bu döngü döner. Kanal close() edilirse döngü biter.
	for port := range ports {
		address := fmt.Sprintf("%s:%d", ip, port)
		con, err := net.DialTimeout("tcp", address, timeout)

		if err == nil {
			fmt.Printf("Port %d açık!\n", port)
			con.Close()
		}
	}
}

func main() {
	timeout := 500 * time.Millisecond
	scanner := bufio.NewScanner(os.Stdin)
	var startPort, endPort int
	var ip string

	for {
		fmt.Print("Lütfen taranacak IP adresini giriniz: ")
		if scanner.Scan() {
			ip = strings.TrimSpace(scanner.Text())
			if net.ParseIP(ip) != nil {
				break
			}
			fmt.Println(" Geçersiz IP adresi! Lütfen tekrar deneyin (Örn: 192.168.1.1).")
		}
	}
	fmt.Println(ip, ": IP adresi onaylandı.")

	for {
		fmt.Print("Lütfen taranacak port aralığını giriniz (örn.: 20-1000): ")
		if scanner.Scan() {
			portInput := strings.TrimSpace(scanner.Text())
			parts := strings.Split(portInput, "-")
			if len(parts) == 1 {
				p, err := strconv.Atoi(parts[0])
				if err == nil && p >= 1 && p <= 65535 {
					startPort = p
					endPort = p
					break
				}
			} else if len(parts) == 2 {
				pStart, err1 := strconv.Atoi(parts[0])
				pEnd, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil && pStart >= 1 && pEnd <= 65535 && pStart <= pEnd {
					startPort = pStart
					endPort = pEnd
					break
				}
			}
			fmt.Println(" Geçersiz port aralığı! Lütfen tekrar deneyin (Örn: 20-1000).")
			continue
		}

	}
	fmt.Println("Port aralığı onaylandı: ", startPort, "-", endPort)

	numWorkers := 100
	portsChan := make(chan int, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)                              // Her işçi için WaitGroup'ı 1 artırıyoruz.
		go worker(portsChan, &wg, ip, timeout) // İşçi go-routine'lerini başlatıyoruz.
	}

	for port := startPort; port <= endPort; port++ {
		portsChan <- port
	}
	close(portsChan) // 1. Önce kanalı kapatıyoruz
	wg.Wait()        // 2. Sonra tüm işçilerin (100 tane Done) bitmesini bekliyoruz

}
