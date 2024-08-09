package session

import "gothstarter/internal/store/user"

type Session struct {
	ID        uint
	SessionID string
	UserID    uint
	User      user.User
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*user.User, error)
}
