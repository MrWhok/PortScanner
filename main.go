package main

import (
	"PortScanner/man"
	"PortScanner/port"
	"fmt"
)

func main() {
	var Protocol, StartPort, EndPort int
	var Address, Continue string

	for {
		man.ShowManual()
		fmt.Println("Press q to coninue")
		fmt.Scanln(&Continue)
		if Continue == "q" {
			break
		}
	}

	print("Choose a protocol type: \n")
	print("1. TCP\n")
	print("2. UDP\n")
	print("=>")
	fmt.Scanln(&Protocol)

	print("Scan a range of ports\n")
	print("Enter the starting port number:\n")
	print("=>")
	fmt.Scanln(&StartPort)
	print("Enter the ending port number:\n")
	print("=>")
	fmt.Scanln(&EndPort)

	print("Enter the IP address:\n")
	print("=>")
	fmt.Scanln(&Address)

	fmt.Printf("Protocol: %d, StartPort: %d, EndPort: %d, Address: %s\n", Protocol, StartPort, EndPort, Address)
	if Protocol == 1 {
		fmt.Println(port.InitialScan("tcp", Address, StartPort, EndPort))
	} else if Protocol == 2 {
		fmt.Println(port.InitialScan("udp", Address, StartPort, EndPort))
	}
}
