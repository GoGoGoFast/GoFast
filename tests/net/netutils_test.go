package netutils_test

import (
	"testing"
	"time"
)

func TestCheckPortOpen(t *testing.T) {
	host := "google.com"
	port := 80
	timeout := 2 * time.Second

	if !netutils.CheckPortOpen(host, port, timeout) {
		t.Errorf("expected port %d on host %s to be open", port, host)
	}
}

func TestGetLocalIP(t *testing.T) {
	ip, err := netutils.GetLocalIP()
	if err != nil || ip == "" {
		t.Errorf("expected a valid local IP address, got error: %v", err)
	}
}

func TestResolveDomain(t *testing.T) {
	domain := "google.com"
	ips, err := netutils.ResolveDomain(domain)
	if err != nil || len(ips) == 0 {
		t.Errorf("expected to resolve domain %s to one or more IP addresses, got error: %v", domain, err)
	}
}

func TestGetMACAddress(t *testing.T) {
	mac, err := netutils.GetMACAddress()
	if err != nil || mac == "" {
		t.Errorf("expected a valid MAC address, got error: %v", err)
	}
}

func TestIsValidIP(t *testing.T) {
	validIP := "192.168.1.1"
	if !netutils.IsValidIP(validIP) {
		t.Errorf("expected %s to be a valid IP address", validIP)
	}

	invalidIP := "999.999.999.999"
	if netutils.IsValidIP(invalidIP) {
		t.Errorf("expected %s to be an invalid IP address", invalidIP)
	}
}

func TestGetHostname(t *testing.T) {
	hostname, err := netutils.GetHostname()
	if hostname == "" || err != nil {
		t.Errorf("expected a valid hostname")
	}
}

func TestGetAllNetworkInterfaces(t *testing.T) {
	interfaces, err := netutils.GetAllNetworkInterfaces()
	if len(interfaces) == 0 || err != nil {
		t.Errorf("expected at least one network interface")
	}
}

func TestGetPublicIP(t *testing.T) {
	ip, err := netutils.GetPublicIP()
	if len(ip) == 0 || err != nil {
		t.Errorf("expected a valid public IP address")
	}
}

func TestCheckPortRangeOpen(t *testing.T) {
	host := "google.com"
	startPort := 79
	endPort := 81
	timeout := 2 * time.Second

	openPorts := netutils.CheckPortRangeOpen(host, startPort, endPort, timeout)
	if len(openPorts) == 0 {
		t.Errorf("expected at least one open port in the range %d-%d on host %s", startPort, endPort, host)
	}
}

func TestPing(t *testing.T) {
	host := "google.com"

	isReachable, err := netutils.Ping(host)
	if !isReachable || err != nil {
		t.Errorf("expected host %s to be reachable", host)
	}
}

func TestLookupMX(t *testing.T) {
	domain := "gmail.com"

	mxRecords, err := netutils.LookupMX(domain)
	if len(mxRecords) == 0 || err != nil {
		t.Errorf("expected at least one MX record for domain %s", domain)
	}
}

func TestLongToIPv4(t *testing.T) {
	longValue := uint32(3232235777) // 192.168.1.1

	expected := "192.168.1.1"
	actual := netutils.LongToIPv4(longValue)

	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestIPv4ToLong(t *testing.T) {
	ip := "192.168.1.1"

	expected := uint32(3232235777) // 192.168.1.1

	actual, _ := netutils.IPv4ToLong(ip)

	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestIsUsableLocalPort(t *testing.T) {
	port := 50000 // 选择一个很少使用的高位端口

	if !netutils.IsUsableLocalPort(port) {
		t.Errorf("expected port %d to be usable", port)
	}
}

func TestIsValidPort(t *testing.T) {
	validPort := 80
	invalidPort := 70000

	if !netutils.IsValidPort(validPort) {
		t.Errorf("expected port %d to be valid", validPort)
	}

	if netutils.IsValidPort(invalidPort) {
		t.Errorf("expected port %d to be invalid", invalidPort)
	}
}

func TestIsInnerIP(t *testing.T) {
	innerIPs := []string{
		"10.0.0.1",
		"172.16.0.1",
		"192.168.1.1",
	}

	outerIPs := []string{
		"8.8.8.8",
		"123.123.123.123",
	}

	for _, ip := range innerIPs {
		if !netutils.IsInnerIP(ip) {
			t.Errorf("expected IP %s to be an inner IP", ip)
		}
	}

	for _, ip := range outerIPs {
		if netutils.IsInnerIP(ip) {
			t.Errorf("expected IP %s to be an outer IP", ip)
		}
	}
}

func TestLocalIPv4s(t *testing.T) {
	ips, _ := netutils.LocalIPv4s()
	if len(ips) == 0 {
		t.Error("expected at least one local IPv4 address ")
	}
}

func TestToAbsoluteURL(t *testing.T) {
	baseURL := "http://example.com"
	relativeURL := "/path/to/resource"
	expected := "http://example.com/path/to/resource"
	actual, _ := netutils.ToAbsoluteURL(baseURL, relativeURL)
	if actual != expected {

		t.Errorf("expected URL %s,but got URL %s", expected, actual)
	}
}

func TestHideIPPart(t *testing.T) {
	ip := "192.168.1.100"
	expected := "192.168.1.*"
	actual, _ := netutils.HideIPPart(ip)
	if actual != expected {

		t.Errorf("expected hidden IP %s,but got hidden IP %s", expected, actual)
	}
}

func TestBuildInetSocketAddress(t *testing.T) {
	host := "localhost"
	port := 8080
	expected := "localhost:8080"
	actual := netutils.BuildInetSocketAddress(host, port)
	if actual != expected {

		t.Errorf("expected InetSocketAddress %s,but got InetSocketAddress %s", expected, actual)
	}
}

func TestGetIpByHost(t *testing.T) {
	host := "google.com"
	ip, _ := netutils.GetIPByHost(host)
	if len(ip) == 0 {
		t.Error("expected a valid IP address ")
	}
}

func TestIsInner(t *testing.T) {
	start := "10 .0 .0 .1"
	end := "10 .255 .255 .255"
	target := "10 .123 .45 .67"

	startLong, _ := netutils.IPv4ToLong(start)
	endLong, _ := netutils.IPv4ToLong(end)
	targetLong, _ := netutils.IPv4ToLong(target)

	if !netutils.IsInner(startLong, endLong, targetLong) {
		t.Error("expected target IP to be within the range ")
	}
}
