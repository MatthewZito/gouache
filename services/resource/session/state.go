package session

// var sessions *SessionManager

var providers = make(map[string]SessionProvider)

// func init() {
// 	sessions, _ = NewSessionManager("memory", "gosessionid", 3600)
// 	go sessions.FinalizeSessions()
// 	s := sessions.NewSession()

// 	s.Get()
// }
