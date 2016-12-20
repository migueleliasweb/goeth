package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func getDocker0Network() *net.IPNet {
	iface, _ := net.InterfaceByName("docker0")
	addresses, _ := iface.Addrs()

	_, parsedIPNet, _ := net.ParseCIDR(addresses[0].String())

	return parsedIPNet
}

func isLocalIp(ip net.IP) bool {
	//https://en.wikipedia.org/wiki/Private_network#Private_IPv4_address_spaces
	_, classANetwork, _ := net.ParseCIDR("10.0.0.0/8")
	_, classBNetwork, _ := net.ParseCIDR("172.16.0.0/12")
	_, classCNetwork, _ := net.ParseCIDR("192.168.0.0/16")

	return classANetwork.Contains(ip) || classBNetwork.Contains(ip) || classCNetwork.Contains(ip) || ip.IsLoopback()
}

func excludePrivateIPs(ips []string) []string {
	filteredIPs := make([]string, 0)

	for _, ip := range ips {
		if !isLocalIp(net.ParseIP(ip)) {
			continue
		}

		filteredIPs = append(filteredIPs, ip)
	}
	return filteredIPs
}

func excludePublicIPs(ips []string) []string {
	filteredIPs := make([]string, 0)

	for _, ip := range ips {
		if isLocalIp(net.ParseIP(ip)) {
			continue
		}

		filteredIPs = append(filteredIPs, ip)
	}
	return filteredIPs
}

func getIps(
	excludeLocalhost bool,
	excludeDockerNetwork bool,
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
		} else if !allowipv6 && parsedIP.To4() == nil {
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
	allowipv6 := flag.Bool("ipv6", false, "Whether to allow ipv6 results.")
	separator := flag.String("separator", ",", "Separator to use on the output.")

	onlyPublic := flag.Bool("only-public", false, "Whether to return only public addresses.")
	onlyPrivate := flag.Bool("only-private", false, "Whether to return only private addresses.")

	flag.Parse()

	if *onlyPrivate && *onlyPublic {
		fmt.Println("Can't use simultaneously 'only-private' and 'only-public' options !")
		os.Exit(1)
	}

	ips := getIps(
		*excludeLocalhost,
		*excludeDockerNetwork,
		*allowipv6,
	)

	// if *onlyPrivate {
	// 	ips := excludePublicIPs(ips)
	// }

	// if *onlyPublic {
	// 	ips := excludePrivateIPs(ips)
	// }

	output := formatOutput(ips, *separator)

	fmt.Println(output)
}
