package strategy_test

import (
	"testing"

	"github.com/SubProblem/tcp-lb/balancer"
	"github.com/SubProblem/tcp-lb/strategy"
	"github.com/stretchr/testify/assert"
)

func TestIpHash_SameIpReturnsSameBackend(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}

	iph := strategy.NewIpHash()
	ip := "1.2.3.4"

	first := iph.Next(backends, ip)
	assert.NotNil(t, first)

	for range 5 {
		result := iph.Next(backends, ip)
		assert.Equal(t, first.Address, result.Address)
	}
}

func TestIpHash_AllUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(false)
	}

	iph := strategy.NewIpHash()
	ip := "1.2.3.4"
	result := iph.Next(backends, ip)
	assert.Nil(t, result)
}

func TestIpHash_OneUnhealthy(t *testing.T) {
	backends := []balancer.Backend{
		{Address: "backend1:9000"},
		{Address: "backend2:9000"},
		{Address: "backend3:9000"},
	}
	for i := range backends {
		backends[i].Healthy.Store(true)
	}

	iph := strategy.NewIpHash()
	ip := "1.2.3.4"
	target := iph.Next(backends, ip)
	target.Healthy.Store(false)

	result := iph.Next(backends, ip)
	assert.NotNil(t, result)
	assert.NotEqual(t, target.Address, result.Address)
}