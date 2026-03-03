package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	ServerID           string        `json:"server_id"`
	ServerName         string        `json:"server_name"`
	AuthToken          string        `json:"auth_token"`
	BackendURL         string        `json:"backend_url"`
	WebSocketURL       string        `json:"websocket_url"`
	CollectionInterval time.Duration `json:"collection_interval"`

	// Feature flags
	EnableCPU        bool `json:"enable_cpu"`
	EnableMemory     bool `json:"enable_memory"`
	EnableDisk       bool `json:"enable_disk"`
	EnableNetwork    bool `json:"enable_network"`
	EnableWebServer  bool `json:"enable_web_server"`
	EnableAlerts     bool `json:"enable_alerts"`
	EnableIPBlocking bool `json:"enable_ip_blocking"`
	EnableBackup     bool `json:"enable_backup"`

	// Thresholds
	CPUThreshold    float64 `json:"cpu_threshold"`
	MemoryThreshold float64 `json:"memory_threshold"`
	DiskThreshold   float64 `json:"disk_threshold"`

	// IP Blocking
	BlockListRefreshInterval time.Duration `json:"block_list_refresh_interval"`

	// Logging
	LogLevel string `json:"log_level"`
	LogFile  string `json:"log_file"`

	// Web Server specific
	WebServerType string `json:"web_server_type"` // apache, nginx, iis
	WebServerPort int    `json:"web_server_port"`
}

func Load(path string) (*Config, error) {
	// file, err := os.Open(path)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()

	config := &Config{
		ServerID: "kali-vm-01",
		CollectionInterval:       5 * time.Second,
		BlockListRefreshInterval: 60 * time.Second,
		CPUThreshold:             80.0,
		MemoryThreshold:          85.0,
		DiskThreshold:            90.0,
		EnableCPU:                true,
		EnableMemory:             true,
		LogLevel:                 "info",
	}

	// decoder := json.NewDecoder(file)
	// if err := decoder.Decode(config); err != nil {
	// 	return nil, err
	// }

	return config, nil
}

func (c *Config) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}
