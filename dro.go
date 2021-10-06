package main

import "time"

type AppDro struct {
	Name string `gorm:"primaryKey"`
}

type CodeDro struct {
	Code     string `gorm:"index"`
	App      string `gorm:"index"`
	Dev      string `gorm:"index"`
	Email    string `gorm:"index"`
	Disabled bool   `gorm:"index"`
	Created  time.Time
	Expires  time.Time
	DevRt    string
}

type TokenDro struct {
	Token    string `gorm:"index"`
	App      string `gorm:"index"`
	Dev      string `gorm:"index"`
	Email    string `gorm:"index"`
	Disabled bool   `gorm:"index"`
	Created  time.Time
	DevRt    string
}
