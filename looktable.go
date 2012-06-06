// postfix tcp_table
// 基本完成
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func check(c rune) bool {
	if c == '\r' || c == '\n' {
		return true
	}
	return false
}

func parse(line string) (cmd, req string) {
	cmd, req = "any", ""
	line = strings.TrimRightFunc(line, check)
	if strings.HasPrefix(line, "quit") {
		cmd, req = "quit", ""
		return cmd, req
	}

	if strings.HasPrefix(strings.ToLower(line), "get") {
		switch index := strings.Index(line, " "); index {
		case 3:
			str := strings.SplitN(line, " ", 2)
			cmd, req = "get", str[1]
			if req != "" {
				cmd = "get"
			} else {
				cmd = "bad"
			}
		case -1:
			if strings.EqualFold(line, "get") {
				cmd = "bad"
			} else {
				cmd = "any"
			}
		}
	}

	return
}

var pool = []string{"a.com", "b.com", "c.com"}

func reply(cmd, req string) (res string) {
	switch cmd {
	case "bad":
		res = "400 get invalid\n"
	case "get":
		for _, v := range pool {
			if req == v {
				res = "200 " + v + "\n"
				return
			}
		}
		res = "500 data does not exist\n"
	case "quit":
		res = "220 bye\n"
	default:
		res = "400 command not found\n"
	}
	return
}

func handler(conn net.Conn, timeout time.Duration, id int64) {
	log.Printf("-- %d --\n", id)
	var buf = bufio.NewReader(conn)
	conn.SetDeadline(time.Now().Add(timeout))
	for {
		line, err := buf.ReadString('\n')
		if err == nil {
			conn.SetDeadline(time.Now().Add(timeout))
			command, req := parse(line)
			res := reply(command, req)
			if command == "quit" {
				fmt.Fprint(conn, res)
				break
			}
			log.Printf("%d Response: %s", id, res)
			fmt.Fprint(conn, res)
		} else {
			break
		}
	}
	if err := conn.Close(); err != nil {
		log.Println("%d conn close error: %s.\n", id, err)
	}
	log.Printf("%d Connection closed.\n", id)
	return
}

var timeout = 10 * time.Second

func main() {
	ln, err := net.Listen("tcp", ":10020")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("conn err:", err)
		}
		id := time.Now().UnixNano()
		go handler(conn, timeout, id)
	}
}
