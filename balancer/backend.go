package balancer

import "sync/atomic"


type Backend struct {
	Address string
	Healthy atomic.Bool
	ActiveConns atomic.Int32
	TotalRequests atomic.Uint64
}