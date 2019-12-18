/*
 *Author: Stefan
 *Date: 12/11/2019
 *Last changes: 12/19/2019 00.00
 *Task:
• TCP Server (S) listen for a connections on a port
• TCP Client (C) connects to a TCP server on a port
• C reads STDIN for any input. On hit enter (‘\n’) C sends input it got to a server
• S reads input (split by ‘\n’) and checks if it’s an int, returns input multiplied by 2
• If it’s not an integer, return uppercased input string(floats will be untouched).
• C shows response from S to STDOUT and waits for another input from STDIN.
• C exits by input `exit`
• S exits by ctrl+C
 * To connection I using PuTTY
**/

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {

	conn, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		//Accept for listening input connection
		connec, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		//concurrent listening
		handleConnection(connec)
	}
}
func handleConnection(con net.Conn) {
	for {
		text, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(" Out > ", string(text))

		txt := strings.TrimSpace(text)

		a, err := strconv.ParseInt(txt, 10, 64)

		if err == nil {
			n := int(a)
			n *= 2
			con.Write([]byte(strconv.Itoa(n) + "\n"))
			//con.Write([]byte(n + "\n\r"))

		} else if err != nil {
			s := strings.ToUpper(text)
			con.Write([]byte(s + "\n\r"))
		}
	}
}
