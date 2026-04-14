package balancer

type Strategy interface {
	Next(backends []Backend, clientIP string) *Backend
}