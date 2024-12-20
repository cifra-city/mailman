package mailbox

func (m *Service) CheckUserInBox(email string, operationType string, ConfidenceCode string) *Data {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {
			if data.ConfidenceCode == ConfidenceCode {
				return &data
			}
		}
	}

	return nil
}
