/*
 *Author: Stefan
 *Date: 12/11/2019
 *Last changes: 12/11/2019 03.38
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

func handleConnection(conn net.Conn) {
	user := conn.RemoteAddr().String()

	fmt.Println("Connected " + user)             //log on server
	conn.Write([]byte("Hello " + user + "\n\r")) // answer to client

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		text := scanner.Text()

		if text == "exit" {
			conn.Write([]byte("Goodbye " + user + "\n\r")) //log
			fmt.Println("Bye")                             //answer
			break
		} else if text == "" {
			fmt.Println("Is the empty string")
			conn.Write([]byte("Roll" + "\n\r"))
			continue
		}

		a, err := strconv.ParseInt(text, 10, 64)

		if err == nil {
			n := strconv.Itoa(int(a * int64(2)))
			fmt.Println("Your number * 2 = ", n)
			conn.Write([]byte("Success Multiply" + "\n\r"))

		} else if err != nil {
			s := strings.ToUpper(text)
			fmt.Println(" Your string: ", s)
			conn.Write([]byte("Success Upper" + "\n\r"))
		}
	}
}

func main() {
	//Create listener using method Listen from net
	//Choose protocol and port for listening

	//To connection I using PuTTY

	listene, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	//Deamon
	for {
		//Accept for listening input connection
		connec, err := listene.Accept()
		if err != nil {
			panic(err)
		}
		//concurrent listening
		go handleConnection(connec)
	}
}
