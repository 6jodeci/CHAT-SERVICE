package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MESSAGE
	CMD_QUIT
)

type command struct {
	id    commandID
	clint *client
	args  []string
}
