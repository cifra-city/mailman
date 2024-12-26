package mailbox

func (m *Service) DeleteFromBox(email string, operationType string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if _, opExists := operations[operationType]; opExists {
			delete(operations, operationType)

			if len(operations) == 0 {
				delete(m.listCode, email)
			}
		}
	}
}
