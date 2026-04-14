package strategy_test

import (
	"testing"
	"github.com/SubProblem/tcp-lb/balancer"
	"github.com/SubProblem/tcp-lb/strategy"
	"github.com/stretchr/testify/assert"
)

func TestRoundRobin_AllHealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}

	rr := strategy.NewRoundRobin()

	first := rr.Next(backends, "")
	second := rr.Next(backends, "")
	third := rr.Next(backends, "")

	assert.Equal(t, "backend2:9000", first.Address)
	assert.Equal(t, "backend3:9000", second.Address)
	assert.Equal(t, "backend1:9000", third.Address)
}


func TestRoundRobin_OneUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}

	backends[1].Healthy.Store(false)

	rr := strategy.NewRoundRobin()

	for i := range 6 {
		result := rr.Next(backends, "")
		assert.NotEqual(t, "backend2:9000", result.Address, "call %d: returned unhealthy backend", i+1)
	}
}

func TestRoundRobin_AllUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(false)
	}

	rr := strategy.NewRoundRobin()
	result := rr.Next(backends, "")
	assert.Nil(t, result)
}