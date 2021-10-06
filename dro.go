package main

import "time"

type AppDro struct {
	Name string `gorm:"primaryKey"`
}

type MessageDro struct {
	Mid     string `gorm:"primaryKey"`
	From    string
	To      string
	Subject string
	Mime    string
	Body    string
	Created time.Time
}

type AttemptDro struct {
	Mid     string
	Created time.Time
	Result  string
}
