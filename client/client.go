package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Client struct {
	addr string
	Conn net.Conn
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (s *Client) Run() {
	conn, err := net.Dial("tcp", s.addr)
	if nil != err {
		log.Fatalln(err)
	}
	s.Conn = conn

	err = s.login()
	if nil != err {
		log.Fatal(err)
	}

	go s.writeloop()
	s.readloop()
}

func (s *Client) login() error {
	_, err := s.Conn.Write([]byte(*userName + "\n"))
	if nil != err {
		log.Println(err)
		return err
	}

	return nil
}
func (s *Client) readloop() error {
	for {
		bin := make([]byte, 10, 10)
		n, err := s.Conn.Read(bin)
		if nil != err {
			log.Println(err)

			err = s.Conn.Close()
			if nil != err {
				log.Println(err)
				return err
			}

			return err
		}
		fmt.Print(string(bin[:n]))
	}

	return nil
}
func (s *Client) writeloop() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if nil != err {
			log.Println(err)
			return err
		}

		_, err = s.Conn.Write([]byte(input))
		if nil != err {
			log.Println(err)
			err = s.Conn.Close()
			if nil != err {
				log.Println(err)
				return err
			}

			return err
		}
	}

	return nil
}
