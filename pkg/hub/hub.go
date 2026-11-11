package hub

// Add registers c in the hub.
//
// The c parameter must identify the client through its Id field. A nil client
// is ignored. If another client already has the same ID, it is replaced.
func (h *Hub) Add(c *Client) {
	if c == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.clientes[c.Id] = c
}

// Remove unregisters c from the hub.
//
// The c parameter is matched by its Id field. A nil client is ignored, and
// removing a client that is not registered has no effect.
func (h *Hub) Remove(c *Client) {
	if c == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clientes, c.Id)
}

// HostDelete unregisters c from the hub and closes its WebTransport session.
//
// The c parameter is matched by its Id field. A nil client is ignored. The
// session is closed after the client is removed from the hub; this prevents
// session shutdown from blocking access to the hub's client registry.
func (h *Hub) HostDelete(c *Client) {
	if c == nil {
		return
	}

	h.mu.Lock()
	delete(h.clientes, c.Id)
	h.mu.Unlock()

	if c.session != nil {
		_ = c.session.CloseWithError(0, "host deleted")
	}
}

// List returns a snapshot of the clients currently registered in the hub.
//
// The returned map is a copy and can be modified by the caller without
// changing the hub's registry. The Client values themselves are shared.
func (h *Hub) List() map[string]*Client {
	h.mu.Lock()
	defer h.mu.Unlock()

	clients := make(map[string]*Client, len(h.clientes))
	for id, client := range h.clientes {
		clients[id] = client
	}

	return clients
}
