package config

import (
	"fmt"
)

type Server struct {
	Host string
	Port int
}

func (c *Server) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
