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
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("input > ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(text)) == "exit" {
			fmt.Println("Bye!")
			return
		}
		fmt.Fprintf(conn, text+"\n")
		tex, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Answer > " + tex)
	}
}
