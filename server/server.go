package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr}
}

func (s *Server) Listen() error {
	l, err := net.Listen("tcp", s.addr)
	if nil != err {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			return err
		}
		log.Println("one client came")

		session, err := s.handleLogin(conn)
		if nil != err {
			log.Println(err)
			return err
		}
		session.Start()
	}

	return nil
}

func (s *Server) handleLogin(conn net.Conn) (*Session, error) {
	userName, err := bufio.NewReader(conn).ReadString('\n')
	if nil != err {
		log.Println(err)
		return nil, err
	}

	_, err = conn.Write([]byte("hi " + userName))
	if nil != err {
		log.Println(err)
		return nil, err
	}

	session := &Session{
		UserName: strings.TrimRight(userName, "\n"),
		Conn:     conn,
	}

	defaultSessionManager.AddSession(session)

	return session, nil
}
