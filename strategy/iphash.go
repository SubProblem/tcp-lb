package strategy

import (
	"hash/fnv"
	"github.com/SubProblem/tcp-lb/balancer"
)

type IpHash struct {
}

func NewIpHash() *IpHash {
	return &IpHash{}
}

func (iph *IpHash) Next(backends []balancer.Backend, clientIP string) *balancer.Backend {
	hash := fnv.New32a()
	hash.Write([]byte(clientIP))
	start := hash.Sum32()

	for i := range backends {
		index := (start + uint32(i)) % uint32(len(backends))
		if backends[index].Healthy.Load() {
			return &backends[index]
		}
	}
	return nil
}