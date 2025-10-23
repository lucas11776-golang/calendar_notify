package jobs

import (
	"fmt"

	"github.com/lucas11776-golang/calendar_notify/utils/token"
)

func Run() {
	// Get calender events...

	token, err := token.Get()

	fmt.Println("RESULT", token, err)

	go func() {
		for {

		}
	}()

	// Check check event within one hours
	go func() {
		for {

		}
	}()
}
