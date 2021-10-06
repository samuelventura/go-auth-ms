package main

import "time"

type AppDro struct {
	Name string `gorm:"primaryKey"`
}

type CodeDro struct {
	Code     string `gorm:"index"`
	App      string `gorm:"index:code_owner"`
	Dev      string `gorm:"index:code_owner"`
	Email    string `gorm:"index:code_owner"`
	Disabled bool   `gorm:"index"`
	Created  time.Time
	Expires  time.Time
}

type TokenDro struct {
	Token    string `gorm:"index"`
	App      string `gorm:"index:token_owner"`
	Dev      string `gorm:"index:token_owner"`
	Email    string `gorm:"index:token_owner"`
	Disabled bool   `gorm:"index"`
	Created  time.Time
}
