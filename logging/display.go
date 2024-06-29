package logging

import (
	"fmt"
)

func PrettyPrint(user, msg string) {
	fmt.Println("Hello", user, msg)
}
