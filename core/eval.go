package core

import (
	"errors"
	"fmt"
	"net"
)

func evalPING(args []string, c net.Conn) error {
	var buf []byte

	if len(args) > 1 {
		return errors.New("ERR wrong number of arguments for 'ping' command")
	}

	if len(args) == 0 {
		buf = Encode("PONG", true)
	} else {
		buf = Encode(args[0], false)
	}

	_, err := c.Write(buf)
	return err
}

func EvalAndResponse(cmd *MemKVCmd, c net.Conn) error {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	}
	return errors.New(fmt.Sprintf("command not found: %s", cmd.Cmd))
}
