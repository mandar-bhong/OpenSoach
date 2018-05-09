package processor

import (
	"fmt"
)

func TaskController(msg string) (string, error) {

	fmt.Println("Task received by server")
	fmt.Println(msg)

	return "", nil
}
