package main

import (
	"log"
)

type scheduler struct {
	Hosts []net.IP
}

func (s *scheduler) NextHost() net.IP {
	return s.Hosts[0]
}
