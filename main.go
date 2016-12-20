package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func getDocker0Network() *net.IPNet {
	iface, _ := net.InterfaceByName("docker0")
	addresses, _ := iface.Addrs()

	_, parsedIPNet, _ := net.ParseCIDR(addresses[0].String())

	return parsedIPNet
}

func getIps(
	excludeLocalhost bool,
	excludeDockerNetwork bool,
	allowpv4 bool,
	allowipv6 bool) []string {

	adresses, _ := net.InterfaceAddrs()
	ips := make([]string, 0)

	docker0Network := getDocker0Network()

	for _, addr := range adresses {
		parsedIP, _, _ := net.ParseCIDR(addr.String())

		if excludeLocalhost && parsedIP.IsLoopback() {
			continue
		} else if excludeDockerNetwork && docker0Network.Contains(parsedIP) {
			continue
		} else if allowpv4 && parsedIP.To4() == nil {
			continue
		} else if allowipv6 && parsedIP.To4() != nil {
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
	excludeLocalhost := flag.Bool("exclude-localhost", false, "Whether to exclude localhost from the result.")
	excludeDockerNetwork := flag.Bool("exclude-docker-network", false, "Whether to exclude the docker network from the result.")
	allowipv4 := flag.Bool("allow-ipv4", true, "Whether to allow ipv4 results.")
	allowipv6 := flag.Bool("allow-ipv6", false, "Whether to allow ipv6 results.")
	separator := flag.String("separator", ",", "Separator to use on the output.")

	flag.Parse()

	ips := getIps(
		*excludeLocalhost,
		*excludeDockerNetwork,
		*allowipv4,
		*allowipv6,
	)

	output := formatOutput(ips, *separator)

	fmt.Println(output)
}
