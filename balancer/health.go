package balancer

import (
	"log"
	"net"
	"time"
)

func StartHealthChecker(backends []Backend, interval time.Duration) {
	checkAll(backends)
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			checkAll(backends)
		}
	}()
}

func checkAll(backends []Backend) {
	for i := range backends {
		conn, err := net.DialTimeout("tcp", backends[i].Address, 2*time.Second)
		if err != nil {
			log.Printf("Health check failed for %s: %v", backends[i].Address, err)
			backends[i].Healthy.Store(false)
		} else {
			conn.Close()
			backends[i].Healthy.Store(true)
		}
	}
}