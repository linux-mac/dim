package main

var (
	defaultSessionManager = &SessionManager{
		Sessions: map[string]*Session{},
	}
)

type SessionManager struct {
	Sessions map[string]*Session
}

func (sm *SessionManager) AddSession(s *Session) {
	sm.Sessions[s.UserName] = s
}

func (sm *SessionManager) GetSession(userName string) *Session {
	return sm.Sessions[userName]
}
