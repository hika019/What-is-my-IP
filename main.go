package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	i, e := net.Interfaces()
	CheckError(e)

	for _, tmp := range i {
		fmt.Println(tmp.Addrs())
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
