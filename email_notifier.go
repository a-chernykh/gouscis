package main

import (
	"bytes"
	"net/smtp"
)

type emailNotifier struct {
	SMTP      string
	Sender    string
	Recipient string
}

func (notifier emailNotifier) Notify(text string) {
	c, err := smtp.Dial(notifier.SMTP)
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
