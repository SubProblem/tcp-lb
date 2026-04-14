package main

import (
	"log"

	"github.com/SubProblem/tcp-lb/balancer"
	"github.com/SubProblem/tcp-lb/config"
	"github.com/SubProblem/tcp-lb/strategy"
)

func main() {
	cfg := config.Config{}
	if err := cfg.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting %s on %s", cfg.Strategy, cfg.ListenAddr)
	switch cfg.Strategy {

	case "roundrobin":
		balancer.Start(&cfg, strategy.NewRoundRobin())
	case "leastconnection":
		balancer.Start(&cfg, strategy.NewLeastConn())
	case "iphash":
		balancer.Start(&cfg, strategy.NewIpHash())
	default:
		log.Fatal("Incorrect strategy was chosen")
	}

}
