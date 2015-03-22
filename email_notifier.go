package main

import (
	"bytes"
	"net/smtp"
)

type EmailNotifier struct {
	Smtp      string
	Sender    string
	Recipient string
}

func (notifier EmailNotifier) Notify(text string) {
	c, err := smtp.Dial(notifier.Smtp)
	check(err)

	c.Mail(notifier.Sender)
	c.Rcpt(notifier.Recipient)

	wc, err := c.Data()
	check(err)

	defer wc.Close()

	buf := bytes.NewBufferString(text)
	_, err = buf.WriteTo(wc)
	check(err)
}
