package strategy

import (
	"sync/atomic"
	"github.com/SubProblem/tcp-lb/balancer"
)

type RoundRobin struct {
	counter uint64
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (rr *RoundRobin) Next(backends []balancer.Backend, clientIP string) *balancer.Backend {
	n := atomic.AddUint64(&rr.counter, 1)
	for i := uint64(0); i < uint64(len(backends)); i++ {
		backend := &backends[(n+i) % uint64(len(backends))]
		if backend.Healthy.Load() {
			return backend
		}
	}
	return nil
}