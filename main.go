package main

import (
	"GameSocketServerClient/cli"
	"bufio"
	"os"
)

func main() {

	c := &cli.CLI{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()
		c.ProcessCommand(cmd)
	}

}
