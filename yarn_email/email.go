package main

import (
	"errors"
	"log"

	"github.com/wneessen/go-mail"
)

func SendUserEmail(email string) error {
	message := mail.NewMsg()
	message.Subject("Order Confirmation")
	message.SetBodyString(mail.TypeTextPlain, "Your Order is Being Processed.")
	err := sendEmail(message, "kurt.heaney@ethereal.email")
	if err != nil {
		return err
	}
	return nil
}

func sendEmail(message *mail.Msg, recipient string) error {
	if err := message.From("gamesapi@automated.com"); err != nil {
		return errors.New("failed to set sender: " + err.Error())
	}
	c, err := mail.NewClient("smtp.ethereal.email", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("kurt.heaney@ethereal.email"), mail.WithPassword("2g8QsnyNdt2zhChpQ8"))
	if err != nil {
		return errors.New("failed to create mail client: " + err.Error())
	}
	if err := message.To(recipient); err != nil {
		log.Fatalf("failed to set recipient: {%s}", err)
		return errors.New("failed to set recipient: " + err.Error())
	}
	if err := c.DialAndSend(message); err != nil {
		return errors.New("failed to send mail: " + err.Error())
	}
	return nil
}