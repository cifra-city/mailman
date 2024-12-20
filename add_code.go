package mailman

import (
	"errors"
	"time"

	"github.com/cifra-city/mailman/meta"
)

func (m *Mailman) AddCode(email string, operationType string, userAgent string, IP string, seconds time.Duration) (*string, error) {
	code, err := m.Postman.GenCode()
	if err != nil {
		return nil, errors.New("failed to generate code")
	}
	err = m.Mailbox.AddToBox(email, code, operationType, meta.Data{
		UserAgent: userAgent,
		IP:        IP,
	}, seconds)
	if err != nil {
		return nil, errors.New("failed to add to mailbox")
	}
	return &code, nil
}
