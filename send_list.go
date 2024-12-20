package mailman

import (
	"time"
)

func (m *Mailman) SendList(email string, operationType string, templateList string, userAgent string, IP string, seconds time.Duration) error {
	code, err := m.AddCode(email, operationType, userAgent, IP, seconds)
	if err != nil {
		return err
	}
	err = m.Postman.SendCode(email, *code, templateList)
	if err != nil {
		return err
	}

	return nil
}
