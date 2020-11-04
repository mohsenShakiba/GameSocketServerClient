package cli

import (
	"GameSocketServerClient/client"
	"fmt"
	"strings"
)

type CLI struct {
	serverInitialized bool
	serverConnected bool
	c *client.TestClient
}

func (c *CLI) ProcessCommand(cmd string)  {

	// parse the command
	cmdParts := strings.Split(cmd, " ")
	mainCmd := cmdParts[0]

	switch strings.ToLower(mainCmd) {
	case "connect":
		c.connect(cmdParts[1:])
		break
	case "create":
		c.createRoom(cmdParts[1:])
		break
	case "delete":
		c.deleteRoom(cmdParts[1:])
		break
	case "join":
		c.joinRoom(cmdParts[1:])
		break
	case "leave":
		c.leaveRoom(cmdParts[1:])
		break
	case "broadcast":
		c.broadcast(cmdParts[1:])
		break
	default:
		c.printInvalidValue()
		c.printCommands()
		break
	}
}

func (c *CLI) connect(args []string) {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c != nil {
		fmt.Println("the client is already created")
		return
	}

	url := args[0]

	c.c = &client.TestClient{}
	go c.c.Start(url)

	fmt.Println("connected to server")
}

func (c *CLI) createRoom(args []string)  {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c == nil {
		fmt.Println("the client hasn't been connected to server yet, use connect first")
		return
	}

	roomName := args[0]
	cmd := fmt.Sprintf("CREATE_ROOM %s", roomName)
	err := c.c.Send(cmd)

	if err != nil {
		fmt.Printf("failed to send command, err: %s", err)
	}
}

func (c *CLI) deleteRoom(args []string)  {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c == nil {
		fmt.Println("the client hasn't been connected to server yet, use connect first")
		return
	}

	roomName := args[0]
	cmd := fmt.Sprintf("DELETE_ROOM %s", roomName)
	err := c.c.Send(cmd)

	if err != nil {
		fmt.Printf("failed to send command, err: %s", err)
	}
}

func (c *CLI) joinRoom(args []string)  {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c == nil {
		fmt.Println("the client hasn't been connected to server yet, use connect first")
		return
	}

	roomName := args[0]
	cmd := fmt.Sprintf("JOIN_ROOM %s", roomName)
	err := c.c.Send(cmd)

	if err != nil {
		fmt.Printf("failed to send command, err: %s", err)
	}
}

func (c *CLI) leaveRoom(args []string)  {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c == nil {
		fmt.Println("the client hasn't been connected to server yet, use connect first")
		return
	}

	roomName := args[0]
	cmd := fmt.Sprintf("LEAVE_ROOM %s", roomName)
	err := c.c.Send(cmd)

	if err != nil {
		fmt.Printf("failed to send command, err: %s", err)
	}
}

func (c *CLI) broadcast(args []string)  {
	if len(args) != 1 {
		c.printInvalidArgument()
		return
	}

	if c.c == nil {
		fmt.Println("the client hasn't been connected to server yet, use connect first")
		return
	}

	roomName := args[0]
	cmd := fmt.Sprintf("BROADCAST %s", roomName)
	err := c.c.Send(cmd)

	if err != nil {
		fmt.Printf("failed to send command, err: %s", err)
	}
}



func (c *CLI) printInvalidArgument()  {
	fmt.Printf("invalid command! \n")
	c.printCommands()
}

func (c *CLI) printInvalidValue() {
	fmt.Printf("invalid command! \n")
}

func (c *CLI) printCommands()  {
	fmt.Printf("Command List: \n")
	fmt.Printf("1. connect	${uri}\n")
	fmt.Printf("2. create	${room name}\n")
	fmt.Printf("3. delete	${room name}\n")
	fmt.Printf("4. join 		${room name}\n")
	fmt.Printf("5. leave 	${room name}\n")
	fmt.Printf("6. broadcase	${message}\n")
	fmt.Printf("7. notify 	${message}\n")
}