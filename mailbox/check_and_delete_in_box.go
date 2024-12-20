package mailbox

func (m *Service) CheckAndDeleteInBox(email string, ConfidenceCode string, operationType string, UserAgent string, IP string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {

			if data.ConfidenceCode == ConfidenceCode && data.Meta.UserAgent == UserAgent && data.Meta.IP == IP {

				delete(operations, operationType)

				if len(operations) == 0 {
					delete(m.listCode, email)
				}

				return true
			}

		}
	}

	return false
}
