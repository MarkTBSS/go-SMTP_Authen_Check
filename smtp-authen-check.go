package main

import (
	"fmt"
	"net"
	"net/textproto"
)

func main() {
	smtpHost := "email.pea.co.th"
	smtpPort := "587"

	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		fmt.Printf("Failed to connect to SMTP server: %v\n", err)
		return
	}
	defer conn.Close()

	textConn := textproto.NewConn(conn)
	_, _, _ = textConn.ReadResponse(220) // Read server's greeting

	// Send EHLO command
	err = textConn.PrintfLine("EHLO localhost")
	if err != nil {
		fmt.Printf("Failed to send EHLO command: %v\n", err)
		return
	}

	// Read and print the server's response
	for {
		line, err := textConn.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(line)
	}
}
