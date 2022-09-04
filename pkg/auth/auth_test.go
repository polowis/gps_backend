package auth

import "testing"

func TestPWDRegisterTexture(t *testing.T) {
	authService := NewAuth()
	authService.SetFolder("../../storage/sp")
	authService.RegisterPWD()
}

func TestPWDRemoveTexture(t *testing.T) {
	authService := NewAuth()
	authService.SetFolder("../../storage/sp")
	authService.ClearSessionTexture(authService.Session())
}