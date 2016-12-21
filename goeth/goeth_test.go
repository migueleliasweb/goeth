package goeth

import (
	"net"
	"testing"
)

func TestGetDocker0Network(t *testing.T) {
	resultNetwork := getDocker0Network()

	if !resultNetwork.Contains(resultNetwork.IP) {
		t.Errorf("Wrong IP configuration for Docker0")
	}
}

func TestIsLocalIp(t *testing.T) {
	privateIPs := []string{
		"127.0.0.1",
		"172.17.0.1",
		"10.168.0.1",
		"192.168.0.1",
		getDocker0Network().IP.String(),
	}

	for _, localIP := range privateIPs {
		if !isLocalIp(net.ParseIP(localIP)) {
			t.Errorf("127.0.0.1 must be an local IP address")
		}
	}
}

func TestExcludePrivateIPs(t *testing.T) {
	IPs := []string{
		"127.0.0.1",
		"172.17.0.1",
		"10.168.0.1",
		"192.168.0.1",
		getDocker0Network().IP.String(),
		"50.50.50.50",
	}

	filteredIPs := excludePrivateIPs(IPs)

	if len(filteredIPs) != 1 {
		t.Errorf("Only one IP should be public")
	}
}

func TestExcludePublicIPs(t *testing.T) {
	IPs := []string{
		"127.0.0.1",
		"172.17.0.1",
		"10.168.0.1",
		"192.168.0.1",
		getDocker0Network().IP.String(),
		"50.50.50.50",
	}

	filteredIPs := excludePublicIPs(IPs)

	if len(filteredIPs) != 5 {
		t.Errorf("Only one IP should have been filtered")
	}
}

func TestGetIPs(t *testing.T) {
	result := getIps(false, false, true)

	if len(result) == 0 {
		t.Errorf("Could not find any IP")
	}

	for _, ip := range result {
		if net.ParseIP(ip) == nil {
			t.Errorf("Could not parse IP [%s]", ip)
		}
	}
}

func TestFormatOutput(t *testing.T) {
	IPs := []string{
		"127.0.0.1",
		"50.50.50.50",
	}

	output := formatOutput(IPs, "|")
	correctOutput := "127.0.0.1|50.50.50.50"

	if output != correctOutput {
		t.Errorf("Wrong output format. Should have been [%s], but was [%s]",
			correctOutput,
			output,
		)
	}
}

// ????
// func TestRun(t *testing.T) {
// 	buffer := bytes.Buffer{}
// 	Run(buffer)
// }
