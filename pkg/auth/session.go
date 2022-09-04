package auth

import "github.com/gps/pkg/pwd"


var SESSION_LENGTH = 12

/*
Create new session id for the user
*/
func NewSession() string {
	sessionID := pwd.GenerateNumber(SESSION_LENGTH)
	return sessionID
}