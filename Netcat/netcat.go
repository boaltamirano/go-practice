package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3091, "port")
	host = flag.String("h", "localhost", "host")
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) //os.Stdout -> nos permite escribir en la consola; conn -> es el lector de los que se escribe en la consola
		done <- struct{}{}       // Notificar atraves del canal que se termino el proceso assignado
	}()

	//
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(destinatario io.Writer, lector io.Reader) {
	_, err := io.Copy(destinatario, lector)

	if err != nil {
		log.Fatal(err)
	}
}
