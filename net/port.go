package main

import (
	"fmt"
	"net"
)

func main() {

	// Escanear cada puerto del  1 al 100 de scanme.nmap.org y hacer una conexión
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}

		conn.Close()
		fmt.Println("Port", i, "is open")
	}
}
