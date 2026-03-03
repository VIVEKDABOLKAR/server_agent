package network

import (
    "net"
)

// GetLocalIPs returns all non-loopback IP addresses of the machine
func GetLocalIPs() ([]string, error) {
    var ips []string
    
    // Get all network interfaces
    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    
    for _, iface := range interfaces {
        // Skip down interfaces and loopback
        if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
            continue
        }
        
        // Get addresses for this interface
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }
        
        for _, addr := range addrs {
            // Parse the IP address
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            
            // Skip if not IPv4 (for simplicity)
            if ip == nil || ip.IsLoopback() || ip.To4() == nil {
                continue
            }
            
            ips = append(ips, ip.String())
        }
    }
    
    return ips, nil
}

// GetPrimaryIP returns the first non-loopback IPv4 address
func GetPrimaryIP() (string, error) {
    ips, err := GetLocalIPs()
    if err != nil {
        return "", err
    }
    
    if len(ips) == 0 {
        return "", nil
    }
    
    return ips[0], nil
}

//GetAllPrimaryIPs returns all non-loopback IPv4 addresses
func GetAllPrimaryIPs() ([]string, error) {
	return GetLocalIPs()
}