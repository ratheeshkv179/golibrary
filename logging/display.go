package logging

import (
	"fmt"
)

func PrettyPrint(user, msg string) {
	fmt.Println("Welcome message", user, msg)
}
