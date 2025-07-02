package models

import "time"

type SessionLog struct {
	Username  string    `json:"username"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	IPAddress string    `json:"ip_address"`
}
