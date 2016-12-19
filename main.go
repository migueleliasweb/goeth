package main

import "net"
import "fmt"

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
	separator string) []string {

	adresses, _ := net.InterfaceAddrs()
	ips := make([]string, 0)

	docker0Network := getDocker0Network()

	for _, addr := range adresses {
		parsedIP := net.ParseIP(addr.String())

		fmt.Println(addr.String())

		if excludeLocalhost && parsedIP.IsLoopback() {
			continue
		} else if excludeDockerNetwork && docker0Network.Contains(parsedIP) {
			continue
		} else {
			ips = append(ips, parsedIP.String())
		}

		fmt.Println(parsedIP.String())
	}

	return ips
}

func main() {
	fmt.Println(getIps(true, true, ","))
}
