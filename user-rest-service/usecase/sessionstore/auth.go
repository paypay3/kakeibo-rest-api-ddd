package sessionstore

import "github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/domain/userdomain"

type SessionStore interface {
	StoreUserBySessionID(sessionID string, userID userdomain.UserID) error
	DeleteUserBySessionID(sessionID string) error
	FetchUserByUserID(sessionID string) (userdomain.UserID, error)
}
