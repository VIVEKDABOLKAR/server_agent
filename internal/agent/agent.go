package agent

import (
	"fmt"
	"time"

	"wsms-agent/config"
	"wsms-agent/internal/collector"
	"wsms-agent/internal/network"
	"wsms-agent/internal/utils"

	models "wsms-agent/internal/model"
)

type Agent struct {
	config   *config.Config
	logger   *utils.Logger
	cpuColl  *collector.CPUCollector
	memColl  *collector.MemoryCollector
	diskColl *collector.DiskCollector
	monitor  *network.ConnectionMonitor
	stopChan chan struct{}
}

func NewAgent(configPath string) (*Agent, error) {
	//init default config with some values, so that if config file is missing some fields, it will not cause error
	config, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	return &Agent{
		config:   config,
		logger:   utils.NewLogger(),
		cpuColl:  collector.NewCPUCollector(),
		memColl:  collector.NewMemoryCollector(),
		diskColl: collector.NewDiskCollector("/"),
		stopChan: make(chan struct{}),
	}, nil
}

func (a *Agent) Start() {
	a.logger.Info("========================================")
	a.logger.Info("Web Server Management Agent - Starting")
	a.logger.Info("========================================")
	a.logger.Infof("Server ID: %s", a.config.ServerID)
	a.logger.Infof("Collection Interval: %d seconds", (a.config.CollectionInterval / 1000000000))
	a.logger.Info("========================================")

	ticker := time.NewTicker(time.Duration(a.config.CollectionInterval))
	defer ticker.Stop()

	//give user two options , one is to start collecting metrics and print them to console, another is to start monitoring incoming connections on a specified port
	fmt.Println("Choose an option:")
	fmt.Println("1. Start collecting metrics")
	fmt.Println("2. Start monitoring incoming connections")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:

		// Collect immediately on start
		a.collectAndPrintInInterval()

	case 2:

		//ask user for port to monitor
		var port int
		fmt.Print("Enter port to monitor: ")
		fmt.Scanln(&port)
        fmt.Printf("Starting connection monitor on port %d... to the port 5173 \n", port) 

		//for development purpose,
		// instead of going public network we will monoitor localhost,
		//so we will change the port to 8080
		a.monitor = network.NewConnectionMonitor(port, a.logger)
		err := a.monitor.Start()
		if err != nil {
			a.logger.Errorf("Failed to start connection monitor: %v", err)
			return
		}
	}

	//stop and shutdown logic for agent , connection meter if exsists
	<-a.stopChan
	a.logger.Info("Received stop signal...")
	a.logger.Info("Agent stopping...")

}

func (a *Agent) Stop() {
	close(a.stopChan)
}

func (a *Agent) collectAndPrintInInterval() {
	ticker := time.NewTicker(time.Duration(a.config.CollectionInterval))
	defer ticker.Stop()

	for {
		a.logger.Info("Waiting for next collection cycle...")
		select {
		case <-ticker.C:
			a.logger.Info("Starting new collection cycle...")
			a.collectAndPrint()
		case <-a.stopChan:
			a.logger.Info("Received stop signal...")
			a.logger.Info("Agent stopping...")
			return
		}
	}
}

func (a *Agent) collectAndPrint() {
	metrics := &models.Metrics{
		ServerID:  a.config.ServerID,
		Timestamp: time.Now(),
	}

	// Collect CPU
	cpuUsage, err := a.cpuColl.Collect()
	if err != nil {
		a.logger.Errorf("Failed to collect CPU: %v", err)
	} else {
		metrics.CPU = cpuUsage
	}

	// Collect Memory
	memUsage, err := a.memColl.Collect()
	if err != nil {
		a.logger.Errorf("Failed to collect Memory: %v", err)
	} else {
		metrics.Memory = memUsage
	}

	//collect Disk
	diskUsage, err := a.diskColl.Collect()
	if err != nil {
		a.logger.Errorf("Failed to collect Disk: %v", err)
	} else {
		metrics.Disk = diskUsage
	}

	primaryIP, err := network.GetAllPrimaryIPs()
	if err != nil {
		a.logger.Errorf("Failed to get primary IP: %v", err)
	} else {
		a.logger.Infof("Primary IP Address: %s", primaryIP)
	}

	// Print to console
	a.printMetrics(metrics)
}

func (a *Agent) printMetrics(m *models.Metrics) {

	a.logger.Info("========================================")
	a.logger.Infof("Time: %s", m.Timestamp.Format(time.RFC3339))
	a.logger.Infof("Server: %s", m.ServerID)
	a.logger.Infof("CPU Usage: %.2f%%", m.CPU)
	a.logger.Infof("Memory Usage: %.2f%%", m.Memory)
	a.logger.Infof("Disk Usage: %.2f%%", m.Disk)
	a.logger.Info("========================================")
}
