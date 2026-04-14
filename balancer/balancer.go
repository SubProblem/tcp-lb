package balancer

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/SubProblem/tcp-lb/config"
)



func Start(cfg *config.Config, strategy Strategy) {
	var wg sync.WaitGroup
	backends := make([]Backend, len(cfg.Backends))
	for i, addr := range cfg.Backends {
		backends[i] = Backend{Address: addr}
	}

	StartHealthChecker(backends, 10*time.Second)

	listener, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Error accepting: ", err)
				return
			}
			clientIP, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
			backend := strategy.Next(backends, clientIP)
			if backend == nil {
				log.Println("No healthy backends available")
				conn.Close()
				continue
			}
			wg.Add(1)
			go handleConnection(conn, backend, &wg)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	
	log.Println("Shutting down...")
	listener.Close()
	wg.Wait()
}

func handleConnection(listener net.Conn, backend *Backend, wg *sync.WaitGroup) {
	defer wg.Done()
	defer listener.Close()
	sender, err := net.Dial("tcp", backend.Address)
	if err != nil {
		log.Println("Error dialing:", err)
		return
	}
	backend.ActiveConns.Add(1)
	defer backend.ActiveConns.Add(-1)
	defer sender.Close()

	go io.Copy(sender, listener)
	io.Copy(listener, sender)
}