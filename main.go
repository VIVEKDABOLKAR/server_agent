package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	
	Agent "wsms-agent/internal/agent"
)





func main() {
    // Default config path
    configPath := "config.json"
    if len(os.Args) > 1 {
        configPath = os.Args[1]
    }

    // Create agent
    agent, err := Agent.NewAgent(configPath)
    if err != nil {
        fmt.Printf("Failed to create agent: %v\n", err)
        os.Exit(1)
    }

    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigChan
        fmt.Println("\nReceived shutdown signal...")
        agent.Stop()
    }()

    // Start agent
    agent.Start()
}