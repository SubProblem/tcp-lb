package strategy

import "github.com/SubProblem/tcp-lb/balancer"

type LeastConn struct {
}

func NewLeastConn() *LeastConn {
	return &LeastConn{}
}

func (lc *LeastConn) Next(backends []balancer.Backend, clientIP string) *balancer.Backend {
	var leastConn *balancer.Backend

	for i := range backends {
		if !backends[i].Healthy.Load() {
			continue
		}
		if leastConn == nil || backends[i].ActiveConns.Load() < leastConn.ActiveConns.Load() {
			leastConn = &backends[i]
		}
	}
	return leastConn
}