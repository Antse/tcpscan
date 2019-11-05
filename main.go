package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 100)
	//fmt.Printf("Testing port %d ...\n", port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func main() {
	ports := []string{}
	sites := []string{"www.ut-capitole.fr",
		"www.google.com",
		"www.lemonde.fr",
		"www.neocasesoftware.com",
		"www.youtube.com",
		"www.apple.com",
		"www.apple.fr"}

	wg := &sync.WaitGroup{}
	timeout := time.Millisecond * 200

	for _, s := range sites {
		wg.Add(1)
		go func(s string) {
			opened := isOpen(s, 443, timeout)
			if opened {
				ports = append(ports, s)
			}
			wg.Done()
		}(s)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}
