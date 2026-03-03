package network

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"wsms-agent/internal/utils"
)

type ConnectionMonitor struct {
	port     int
	listener net.Listener
	isActive bool
	logger   *utils.Logger
}

// NewConnectionMonitor creates a monitor on specified port
func NewConnectionMonitor(port int, logger *utils.Logger) *ConnectionMonitor {
	return &ConnectionMonitor{
		port:     port,
		isActive: false,
		logger:   logger,
	}
}

// Start begins listening for connections on the specified port
func (m *ConnectionMonitor) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", m.port))
	if err != nil {
		return fmt.Errorf("failed to start monitor on port %d: %v", m.port, err)
	}

	m.listener = listener
	m.isActive = true

	m.logger.Infof("Connection monitor started on port %d", m.port)
	m.logger.Info("Waiting for incoming connections...")

	go m.acceptConnections()

	return nil
}

// Stop halts the connection monitoring
func (m *ConnectionMonitor) Stop() {
	if m.isActive && m.listener != nil {
		m.listener.Close()
		m.isActive = false
		m.logger.Info("Connection monitor stopped")
	}
}

// acceptConnections handles incoming connections
func (m *ConnectionMonitor) acceptConnections() {
	for m.isActive {
		conn, err := m.listener.Accept()
		if err != nil {
			if m.isActive {
				m.logger.Errorf("Error accepting connection: %v", err)
			}
			continue
		}

		// Handle each connection in a separate goroutine
		go m.handleConnection(conn)
	}
}

// handleConnection processes a single connection
func (m *ConnectionMonitor) handleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	localAddr := conn.LocalAddr().String()
	ip, _, _ := net.SplitHostPort(remoteAddr)

	m.logger.Info("========================================")
	m.logger.Info("INCOMING CONNECTION DETECTED")
	m.logger.Infof("Time: %s", time.Now().Format(time.RFC3339))
	m.logger.Infof("Client IP: %s", ip)
	m.logger.Infof("Destination: %s", localAddr)

	// Read first 4KB to capture HTTP request line
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("❌ Failed to read request: %v\n", err)
		return
	}

	requestData := string(buffer[:n])
	lines := strings.Split(requestData, "\n")
	if len(lines) > 0 {
		requestLine := strings.TrimSpace(lines[0])
		m.logger.Infof("HTTP Request: %s", requestLine)
	}

	// Connect to actual service
	targetConn, err := net.DialTimeout("tcp", "localhost:5173", 5*time.Second)
	if err != nil {
		m.logger.Errorf("Failed to reach localhost:5173: %v", err)
		return
	}
	defer targetConn.Close()

	// Forward already-read data
	targetConn.Write(buffer[:n])

	done := make(chan bool, 2)

	go func() {
		io.Copy(targetConn, conn)
		done <- true
	}()

	go func() {
		io.Copy(conn, targetConn)
		done <- true
	}()

	<-done
	m.logger.Infof("Connection closed for %s", remoteAddr)
	m.logger.Info("========================================")
}
