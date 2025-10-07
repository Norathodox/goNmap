package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var portRangeFlag string
	flag.StringVar(&portRangeFlag, "p", "80-80", "Port range to scan (default 80)") //port range default and formatting
	flag.Parse()
	if ip := flag.Arg(0); ip != "" { //error case
		dt := time.Now()
		fmt.Printf("Starting network scan at %s\n", dt.Format(time.UnixDate))
		fmt.Printf("Network scan report for %s\n", ip)
		fmt.Printf("PORT\tSTATE\n") //headers

		portRange := strings.Split(portRangeFlag, "-")
		startPort, _ := strconv.Atoi(portRange[0])
		endPort, _ := strconv.Atoi(portRange[len(portRange)-1]) //port range parsing and bounds declaration 

		for port := startPort; port <= endPort; port++ {
			addr := fmt.Sprintf("%s:%d", ip, port)
			_, err := net.Dial("tcp", addr)//actual tcp connection attempt
			if err == nil {
				fmt.Printf("%d\t%s\t\n", port, "open")
			} else {
				fmt.Printf("%d\t%s\t\n", port, "closed")
			} //prints for the port
		}
	} else {
		fmt.Println("Error, please provide an IP address.") 
		os.Exit(1)
	} //^ printing the information about the provided ip, and labels for collumns
}
