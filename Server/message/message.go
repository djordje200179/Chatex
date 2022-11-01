package message

import (
	"net"
	"time"
)

type Message struct {
	Text      string
	Sender    net.Addr
	Timestamp time.Time
}
