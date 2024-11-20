package netutils

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

// CheckPortOpen checks if the specified host and port are open
func CheckPortOpen(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// GetLocalIP gets the first non-loopback IP address of the local machine
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("no non-loopback IP address found")
}

// ResolveDomain resolves a domain name and returns the corresponding list of IP addresses
func ResolveDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	var ipStrs []string
	for _, ip := range ips {
		ipStrs = append(ipStrs, ip.String())
	}
	return ipStrs, nil
}

// GetMACAddress gets the first MAC address of the local machine
func GetMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String(), nil
		}
	}
	return "", fmt.Errorf("no MAC address found")
}

// IsValidIP checks if the given string is a valid IP address
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// GetHostname gets the local hostname
func GetHostname() (string, error) {
	name, err := net.LookupAddr("127.0.0.1")
	if err != nil || len(name) == 0 {
		return "", fmt.Errorf("could not determine hostname")
	}
	return name[0], nil
}

// GetAllNetworkInterfaces gets information about all network interfaces
func GetAllNetworkInterfaces() ([]net.Interface, error) {
	return net.Interfaces()
}

// GetPublicIP gets the public IP address
func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip := make([]byte, 16)
	n, err := resp.Body.Read(ip)
	if err != nil && n == 0 {
		return "", err
	}

	return string(ip[:n]), nil
}

// CheckPortRangeOpen checks which ports are open in the specified host and port range
func CheckPortRangeOpen(host string, startPort int, endPort int, timeout time.Duration) []int {
	openPorts := []int{}
	for port := startPort; port <= endPort; port++ {
		if CheckPortOpen(host, port, timeout) {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}

// Ping performs a ping operation to check if the target host is reachable
func Ping(host string) (bool, error) {
	cmd := exec.Command("ping", "-c", "1", host)
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}

// LookupMX looks up the MX records for the specified domain name
func LookupMX(domain string) ([]*net.MX, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	return mxRecords, nil
}

// LongToIPv4 converts a long value to an IPv4 address
func LongToIPv4(ipLong uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ipLong>>24),
		byte(ipLong>>16),
		byte(ipLong>>8),
		byte(ipLong))
}

// IPv4ToLong converts an IP address to a long value
func IPv4ToLong(ip string) (uint32, error) {
	var long uint32
	ipParts := strings.Split(ip, ".")
	if len(ipParts) != 4 {
		return 0, fmt.Errorf("invalid IPv4 address: %s", ip)
	}
	for i := 0; i < 4; i++ {
		var part uint8
		_, err := fmt.Sscanf(ipParts[i], "%d", &part)
		if err != nil {
			return 0, err
		}
		long |= uint32(part) << uint(8*(3-i))
	}
	return long, nil
}

// IsUsableLocalPort checks if a local port is usable
func IsUsableLocalPort(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer ln.Close()
	return true
}

// IsValidPort checks if it is a valid port number
func IsValidPort(port int) bool {
	return port >= 0 && port <= 65535
}

// IsInnerIP determines whether it is an internal IP
func IsInnerIP(ip string) bool {
	innerIPBlocks := []*net.IPNet{}
	for _, cidr := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"} {
		_, block, _ := net.ParseCIDR(cidr)
		innerIPBlocks = append(innerIPBlocks, block)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, block := range innerIPBlocks {
		if block.Contains(parsedIP) {
			return true
		}
	}

	return false
}

// LocalIPv4s gets a list of IP addresses of the local machine
func LocalIPv4s() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

// ToAbsoluteURL converts a relative URL to an absolute URL
func ToAbsoluteURL(baseURL string, relativeURL string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	rel, err := url.Parse(relativeURL)
	if err != nil {
		return "", err
	}

	absURL := base.ResolveReference(rel)
	return absURL.String(), nil
}

// HideIPPart hides the last part of an IP address and replaces it with *
func HideIPPart(ip string) (string, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return "", fmt.Errorf("invalid IPv4 address: %s", ip)
	}

	parts[3] = "*"
	return strings.Join(parts, "."), nil
}

// BuildInetSocketAddress builds InetSocketAddress
func BuildInetSocketAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// GetIPByHost gets IP by domain name
func GetIPByHost(host string) (string, error) {
	ips, err := net.LookupHost(host)
	if len(ips) == 0 || err != nil {
		return "", fmt.Errorf("could not resolve host: %s", host)
	}

	return ips[0], nil
}

// IsInner determines whether the long value of an IP is within the specified range
func IsInner(start uint32, end uint32, target uint32) bool {
	return target >= start && target <= end
}
