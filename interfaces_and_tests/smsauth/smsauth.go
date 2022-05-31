package smsauth

import "log"

type Challenger interface {
	Challenge() (code string, phoneNumber string)
}

type Prompter interface {
	Prompt(phoneNumber string) string
}

func VerifySMS(c Challenger, p Prompter) bool {
	code, phoneNumber := c.Challenge()
	userAnswer := p.Prompt(phoneNumber)
	log.Println(userAnswer, code)
	return userAnswer == code
}
