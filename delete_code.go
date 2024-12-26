package mailman

func (m *Mailman) DeleteCode(email string, operationType string) {
	m.Mailbox.DeleteFromBox(email, operationType)
}
