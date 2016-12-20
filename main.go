package main

import (
	"fmt"
	"net"
	"strings"
)

func getDocker0Network() net.IPNet {
	iface, _ := net.InterfaceByName("docker0")
	addresses, _ := iface.Addrs()
	docker0IP := net.ParseIP(addresses[0].String())

	return net.IPNet{
		IP:   docker0IP,
		Mask: docker0IP.DefaultMask(),
	}
}

func getIps(
	excludeLocalhost bool,
	excludeDockerNetwork bool,
	onlyipv4 bool,
	onlyipv6 bool) []string {

	adresses, _ := net.InterfaceAddrs()
	ips := make([]string, 0)

	docker0Network := getDocker0Network()

	for _, addr := range adresses {
		parsedIP, _, _ := net.ParseCIDR(addr.String())

		if excludeLocalhost && parsedIP.IsLoopback() {
			continue
		} else if excludeDockerNetwork && docker0Network.Contains(parsedIP) {
			continue
		} else if onlyipv4 && parsedIP.To4() == nil {
			continue
		} else if onlyipv6 && parsedIP.To4() != nil {
			continue
		}

		ips = append(ips, parsedIP.String())
	}

	return ips
}

func formatOutput(ips []string, separator string) string {
	return strings.Join(ips, separator)
}

func main() {
	excludeLocalhost := false
	excludeDockerNetwork := true
	onlyipv4 := true
	onlyipv6 := false
	separator := ","

	ips := getIps(
		excludeLocalhost,
		excludeDockerNetwork,
		onlyipv4,
		onlyipv6,
	)

	output := formatOutput(ips, separator)

	fmt.Println(output)
}
