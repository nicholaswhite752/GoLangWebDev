package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

		for {
			conn, err := li.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			defer conn.Close()
			go serve(conn)
			go handle(conn)
		}

}


func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

}


func serve(conn net.Conn) {
	fmt.Fprintln(conn, "Type Anything")
}
