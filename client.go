package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	nick    string
	room    *room
	commans chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commans <- command{id: CMD_NICK, clint: c, args: args}
		case "/join":
			c.commans <- command{id: CMD_JOIN, clint: c, args: args}
		case "/rooms":
			c.commans <- command{id: CMD_ROOMS, clint: c, args: args}
		case "/msg":
			c.commans <- command{id: CMD_MESSAGE, clint: c, args: args}
		case "/quit":
			c.commans <- command{id: CMD_QUIT, clint: c, args: args}
		default:
			fmt.Println("unknown command.")
			fmt.Println("command list:")
			fmt.Print("/nick /join /rooms /msg /quit ")
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("-> " + msg + "\n"))
}
