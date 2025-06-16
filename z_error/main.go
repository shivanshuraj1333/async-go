package main

import (
	"errors"
	"fmt"
)

type DBReadError struct {
	ErrorCode int
	Message   string
}

func (e DBReadError) Error() string {
	return fmt.Sprintf("Error with reading DB: %s", e.Message)
}

type DBConnError struct {
	ClientID int
	Message  string
}

func (e DBConnError) Error() string {
	return fmt.Sprintf("Error with connecting DB: %s", e.Message)
}

func connector(action string) error {
	switch action {
	case "Read":
		return DBReadError{5, "generated read error"}
	case "Connect":
		return DBConnError{4, "generated connection error"}
	default:
		return errors.New("new Error Hoss")
	}
}

func main() {
	actions := []string{"Read", "Connect", "Others"}

	for _, action := range actions {
		err := connector(action)

		if err != nil {
			switch t := err.(type) {
			case DBReadError:
				fmt.Println("Found error: ", t)
			case DBConnError:
				fmt.Println("Found error: ", t)
			default:
				fmt.Println("Generic error: ", err)
			}
		}

	}

}
