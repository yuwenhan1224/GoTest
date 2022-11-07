package grpcserver

import "net"

// Option -.
type Option func(*Server)

// Port -.
func Port(port string) Option {
	return func(s *Server) {
		s.Addr = net.JoinHostPort("", port)
	}
}
