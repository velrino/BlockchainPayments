package app

import (
	"os"
)

func Run() {
	var cmd string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "listen"
	}
	switch cmd {
	case "migrations":
		Migrations()
	default:
		Routes()
	}
}

