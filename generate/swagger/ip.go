package swagger

import (
	"fmt"
	"net"
)

func getlocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetUrl(addr string) string {
	if len(addr) > 0 {
		if addr[0] == ':' {
			ip := getlocalIp()
			if ip == "" {
				ip = "127.0.0.1"
			}
			return ip + addr
		}
	}
	return ""
}
