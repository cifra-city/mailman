package postman

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func (p *Postman) GenCode() (string, error) {
	newInt := big.NewInt(900000) // 999999 - 100000 + 1
	n, err := rand.Int(rand.Reader, newInt)
	if err != nil {
		return "", err
	}

	code := 100000 + n.Int64()

	return fmt.Sprintf("%06d", code), nil
}
