package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (addr IPAddr) String() string {
	var result string = ""

	for index, num := range addr {
		result += strconv.Itoa(int(num))

		if index != len(addr)-1 {
			result += "."
		}
	}

	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
