package mailman

func (m *Mailman) DeleteAccess(email string, operationType string) {
	m.AccessBox.DeleteOperationForUser(email, operationType)
}
