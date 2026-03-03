package model

import "time"

type Metrics struct {
    ServerID  string    `json:"server_id"`
    Timestamp time.Time `json:"timestamp"`
    CPU       float64   `json:"cpu_usage"`
    Memory    float64   `json:"memory_usage"`
    Disk      float64   `json:"disk_usage"`
}