package main

import "bufio"
import "os"
import "fmt"
import "log"
import "net"
import "sync"
import "time"

// Get current date
var date = time.Now()
var date_string = date.Format("2006-01-02")

var wg sync.WaitGroup
var sem = make(chan struct{}, 8000)

func resolv_domain(domain string) {

    addrs, _ := net.LookupHost(domain)

    if(len(addrs)>0) {
	fmt.Printf("{\"domain\":\"%v\",\"ip\":\"%v\",\"datum\":\"%v\"}\n",domain,addrs[0],date_string)
    } else {
        fmt.Printf("{\"domain\":\"%v\",\"ip\":\"0.0.0.0\",\"datum\":\"%v\"}\n",domain,date_string)
    }

    defer wg.Done()
    <-sem
}

func main() {

	filename := os.Args[1]
        f, err := os.Open(filename)
        if err != nil {
            log.Println(err)
        }
        defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {

		// Add 1 to wg counter
		wg.Add(1)

		// Send Signal into channel
		sem <- struct{}{}

		// Start go routine
		go resolv_domain(s.Text())
	}

	// blocks until the WaitGroup counter is zero.
	wg.Wait()
	close(sem)
}
