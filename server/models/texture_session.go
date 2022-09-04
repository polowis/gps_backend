package models

import "gorm.io/gorm"

type TextureSession struct {
	gorm.Model
	SessionID    string
	ShuffleOrder string  // store id shuffle order
}