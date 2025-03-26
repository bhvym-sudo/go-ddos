package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	".s"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to bhvym cybersecurity tool! What do you want to do!")
	fmt.Println("Choose attack type:")
	fmt.Println("[1] HTTP Flood")
	fmt.Println("[2] TCP Flood")
	fmt.Print("Enter your choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Print("Enter Target URL: ")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		fmt.Println("[-] Invalid URL format.")
		os.Exit(1)
	}

	fmt.Print("Enter Number of Threads: ")
	var threads int
	fmt.Scanln(&threads)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	switch choice {
	case "1":
		fmt.Println("[+] HTTP Flood Attack Selected.")
		go attack.HttpFlood(target, threads)
	case "2":
		fmt.Println("[+] TCP Flood Not Implemented Yet") // Placeholder for TCP attack
	default:
		fmt.Println("[-] Invalid choice.")
		os.Exit(1)
	}

	fmt.Println("[+] Press CTRL+C to stop.")
	<-stop
	fmt.Println("\n[+] Stopping attack...")
}
