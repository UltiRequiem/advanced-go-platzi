package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

func PortOpen(port int, site string, wg sync.WaitGroup) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", site, port))

	if err != nil {
		return
	}

	wg.Done()
	conn.Close()

	fmt.Printf("Port %d is open.\n", port)

}

const PORTS_TO_SCAN = 10000

var site = flag.String("site", "scanme.nmap.org", "URL to scan.")

func main() {
	flag.Parse()

	var wg sync.WaitGroup

	for i := 0; i < PORTS_TO_SCAN; i++ {
		wg.Add(1)
		go PortOpen(i, *site, wg)
	}

	wg.Wait()
}
