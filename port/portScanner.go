package port

import (
	"net"
	"strconv"
	"time"
)

type Results struct {
	Port   int
	Status string
}

func portScanner(protocol, hostname string, port int) Results {
	result := Results{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.Status = "Closed"
		return result
	}

	defer conn.Close()
	result.Status = "Open"
	return result
}

func InitialScan(protocol, hostname string, startPort, endPort int) []Results {
	var results []Results
	for i := startPort; i <= endPort; i++ {
		results = append(results, portScanner(protocol, hostname, i))
	}
	return results
}
