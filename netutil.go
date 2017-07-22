package netutil

import (
	"fmt"
	"net"
)

func Example() {
	ips := FindHostIPs()
	fmt.Printf("ipv4=%v\n", ips)

}

func FindHostIPs() []string {
	return FindIPs([]string{"eth0", "wlan0", "eth1", "em0", "em1"})
}

// return IPv4 addresses of given interfaces
func FindIPs(ifnames []string) (ips []string) {
	for _, ifname := range ifnames {
		iface, err := net.InterfaceByName(ifname)
		if err != nil {
			continue
		}
		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok {
					if ipnet.IP.To4() != nil {
						ips = append(ips, ipnet.IP.To4().String())
					}
				}
			}
		}
	}
	return ips
}
