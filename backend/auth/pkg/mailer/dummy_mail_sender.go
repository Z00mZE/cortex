package mailer

import (
	"context"
	"fmt"
	"math/rand"
)

type DummyMailSender struct {
}

func NewDummyMailSender() *DummyMailSender {
	return new(DummyMailSender)
}

// SendRegistrationMail TODO implement me
func (d DummyMailSender) SendRegistrationMail(ctx context.Context, s string) error {
	if rand.Int()%2 == 1 {
		return fmt.Errorf("implement me")
	}
	return nil
}
