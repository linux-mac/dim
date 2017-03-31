package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Session struct {
	Conn     net.Conn
	UserName string
}

func NewSession(name string, conn net.Conn) *Session {
	return &Session{
		Conn:     conn,
		UserName: name,
	}
}

func (s *Session) Start() {
	go s.readloop()
}

func (s *Session) Send(b []byte) (int, error) {
	return s.Conn.Write(b)
}

func (s *Session) Read(b []byte) (int, error) {
	return s.Conn.Read(b)
}

func (s *Session) readloop() {
	reader := bufio.NewReader(s.Conn)
	for {
		rec, err := reader.ReadString('\n')
		if nil != err {
			log.Println(err)
			return
		}
		baz := strings.Split(rec, "#")
		if 2 != len(baz) {
			log.Println("wrong msg format")
			return
		}

		userName := baz[0]
		msg := baz[1]

		_, err = defaultSessionManager.GetSession(userName).Send([]byte(s.UserName + ": " + msg))
		if nil != err {
			log.Println(err)
			return
		}
	}
}
