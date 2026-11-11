package hub

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	wt "github.com/quic-go/webtransport-go"
)

func aleatorioHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}

func NewHub() *Hub {
	return &Hub{id: aleatorioHex(4), clientes: map[string]*Client{}}
}

func NewClient(h *Hub, s *wt.Session, a string) *Client {
	client := &Client{Id: aleatorioHex(4), session: s, Addr: a, ConnectedAt: time.Now()}

	h.Add(client)

	return client
}
