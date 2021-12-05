package models

import (
	"gorm.io/gorm"
)

type AuditLogs struct {
	gorm.Model
	RemoteAddress string
	Ip            string
	Port          string
	Network       string
	Status        string
	Description   string
	Location      string
}
