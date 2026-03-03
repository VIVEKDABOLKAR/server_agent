package model

import "time"

type IPInfo struct {
    ServerID    string    `json:"server_id"`
    IPAddresses []string  `json:"ip_addresses"`
    PrimaryIP   string    `json:"primary_ip"`
    Timestamp   time.Time `json:"timestamp"`
}