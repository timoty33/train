package hub

import (
	"sync"
	"time"

	wt "github.com/quic-go/webtransport-go"
)

type Client struct {
	Id          string    `json:"id"`
	Addr        string    `json:"addr"`
	ConnectedAt time.Time `json:"connected_at"`
	session     *wt.Session
}

type Hub struct {
	id       string
	mu       sync.Mutex
	clientes map[string]*Client
}
