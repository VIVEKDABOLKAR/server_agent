package model

import "time"

type RequestLog struct {
	Timestamp time.Time `json:"timestamp"`
	ServerID  string    `json:"server_id"`
	ClientIP  string    `json:"client_ip"`
	Method    string    `json:"method"`
	URL       string    `json:"url"`
	Port      int       `json:"port"`
}