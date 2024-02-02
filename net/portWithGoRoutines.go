package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()

	// Escanear cada puerto del  1 al 100 de scanme.nmap.org y hacer una conexión

	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}

			conn.Close()
			fmt.Println("Port", port, "is open")
		}(i)
	}
	wg.Wait()
}
