package models

import (
	"net"
	"time"
)

// AccessLog struct defintion
type AccessLog struct {
	ID        int
	UserID    int
	Timestamp time.Time
	IPAddress net.IP
	Event     string
}
