package client

import (
	"encoding/binary"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

type TestClient struct {
	conn net.Conn
}

func (c *TestClient) Start(url string) error {

	client, err := net.Dial("tcp", url)

	if err != nil {
		return err
	}

	defer client.Close()

	c.conn = client

	c.ProcessIncomingMessages()

	return nil
}

func (c *TestClient) ProcessIncomingMessages() {
	for {
		msg, err := c.Read()

		if err != nil {
			log.Errorf("failed to read message, error: %s", err)
			continue
		}

		fmt.Printf("MSG: %s\n", msg)
	}
}

func (c *TestClient) Send(cmd string) error {

	bSize := make([]byte, 4)

	binary.BigEndian.PutUint32(bSize, uint32(len(cmd)))

	_, err := c.conn.Write(bSize)

	if err != nil {
		return err
	}

	_, err = c.conn.Write([]byte(cmd))

	return nil
}

func (c *TestClient) Read() ([]byte, error) {
	msgSizeB := make([]byte, 4)

	_, err := io.ReadFull(c.conn, msgSizeB)

	if err != nil {
		return nil, err
	}

	// read the size of message
	msgSize := binary.BigEndian.Uint32(msgSizeB)

	// read exactly the size of message from connection
	msgB := make([]byte, msgSize)
	_, err = io.ReadFull(c.conn, msgB)

	if err != nil {
		return nil, err
	}

	return msgB, nil
}