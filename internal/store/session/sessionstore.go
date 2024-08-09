package session

import (
	"fmt"
	"gothstarter/internal/store/user"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const UserSchema = `
CREATE TABLE sessions (
	id INTEGER PRIMARY KEY,
	session_id TEXT UNIQUE,
	user_id INTEGER,
	FOREIFN KEY(user_id) REFERENCES users(id)
);`

type SQLSessionStore struct {
	db *sqlx.DB
}

type NewSessionStoreParams struct {
	DB *sqlx.DB
}

func NewSessionStore(params NewSessionStoreParams) *SQLSessionStore {
	return &SQLSessionStore{
		db: params.DB,
	}
}

func (s *SQLSessionStore) CreateSession(session *Session) (*Session, error) {

	session.SessionID = uuid.New().String()

	_, err := s.db.NamedExec(`INSERT INTO sessions (session_id, user_id) VALUES (:session_id, :user_id);`, map[string]interface{}{
		"session_id": session.SessionID,
		"user_id":    session.UserID,
	})

	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SQLSessionStore) GetUserFromSession(sessionID string, userID string) (*user.User, error) {
	user := user.User{}
	rows, err := s.db.NamedQuery(`
		SELECT u.id as id, u.email as email, u.password as password
		FROM users as u
		JOIN sessions as s ON u.id = s.user_id
		WHERE s.id=:session_id AND u.id=:user_id`,
		map[string]interface{}{
			"session_id": sessionID,
			"user_id":    userID,
		})

	if rows.Next() {
		err = rows.StructScan(&user)
	}

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("no user associated with the session")
	}

	return &user, nil
}
