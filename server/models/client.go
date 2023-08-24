package models

import (
	"net"
)

type Client struct {
	Username string
	Conn     net.Conn
	GroupId  int
}
