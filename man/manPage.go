package man

import "fmt"

func ShowManual() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("          Port Scanner Usage Manual                 ")
	fmt.Println("----------------------------------------------------")
	fmt.Println("This program performs port scanning on a given IP address.")
	fmt.Println("You can scan a range of TCP or UDP ports.")
	fmt.Println("\nUsage Instructions:")
	fmt.Println("1. Choose a protocol:")
	fmt.Println("   - Type '1' for TCP scanning.")
	fmt.Println("   - Type '2' for UDP scanning.")
	fmt.Println("2. Enter the range of ports you want to scan:")
	fmt.Println("   - Provide a starting port number.")
	fmt.Println("   - Provide an ending port number.")
	fmt.Println("3. Enter the IP address to scan.")
	fmt.Println("\nExample:")
	fmt.Println("   Protocol: 1 (for TCP)")
	fmt.Println("   Start Port: 5500")
	fmt.Println("   End Port: 5500")
	fmt.Println("   Address: 127.0.0.1")
	fmt.Println("\nThe program will then scan the specified range of ports on the given IP address.")
	fmt.Println("----------------------------------------------------")
}
