package balancer

import "sync/atomic"


type Backend struct {
	Address string
	Healthy atomic.Bool
	ActiveConns atomic.Int32
}