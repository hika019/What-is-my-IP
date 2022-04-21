package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println(localIP())
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func localIP() string {
	interfaces, e := net.Interfaces()
	CheckError(e)

	for _, inter := range interfaces {
		addr, err := inter.Addrs()
		CheckError(err)

		for _, a := range addr {
			if ipnet, ok := a.(*net.IPNet); ok {
				if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "Sorry I don't know"
}
