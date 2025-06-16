package main

import "fmt"

type Notifier interface {
	Send() string
}

type Email struct {
	EMailID string
	Message string
}

func (e Email) Send() string {
	return fmt.Sprintf("Sent Email: %s", e.EMailID)
}

type SMS struct {
	Number int
}

func (s SMS) Send() string {
	return fmt.Sprintf("Sent SMS: %d", s.Number)
}

type PushNotif struct {
	Channel string
}

func (p PushNotif) Send() string {
	return fmt.Sprintf("Sent message to channel: %s", p.Channel)
}

func main() {
	email := Email{"abc@xyz.com", "boom"}
	sms := SMS{242531}
	pn := PushNotif{"slack"}

	notifications := []Notifier{email, sms, pn}

	for _, s := range notifications {
		response := s.Send()
		fmt.Println(response)
	}
}
