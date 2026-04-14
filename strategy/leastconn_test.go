package strategy_test

import (
	"testing"
	"github.com/SubProblem/tcp-lb/balancer"
	"github.com/SubProblem/tcp-lb/strategy"
	"github.com/stretchr/testify/assert"
)


func TestLeastConn_AllHealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}
	backends[0].ActiveConns.Store(5)
	backends[1].ActiveConns.Store(2)
	backends[2].ActiveConns.Store(8)

	lc := strategy.NewLeastConn()
	result := lc.Next(backends, "")

	assert.Equal(t, "backend2:9000", result.Address)
}

func TestLeastConn_OneUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}
	backends[0].ActiveConns.Store(5)
	backends[1].ActiveConns.Store(2)
	backends[2].ActiveConns.Store(8)
	backends[1].Healthy.Store(false)

	lc := strategy.NewLeastConn()
	result := lc.Next(backends, "")
	assert.Equal(t, "backend1:9000", result.Address)
}

func TestLeastConn_AllUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(false)
	}

	lc := strategy.NewLeastConn()
	result := lc.Next(backends, "")
	assert.Nil(t, result)
}

func TestLeastConn_AllEqualConnections(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}
	backends[0].ActiveConns.Store(3)
	backends[1].ActiveConns.Store(3)
	backends[2].ActiveConns.Store(3)
	
	lc := strategy.NewLeastConn()
	result := lc.Next(backends, "")
	assert.NotNil(t, result)
}