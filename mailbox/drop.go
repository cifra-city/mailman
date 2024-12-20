package mailbox

func (m *Service) Drop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.listCode = make(map[string]map[string]Data)
}
