package port

import (
	"fmt"
	"strconv"
)

func Parser(input string) (int, error) {
	portNumber, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	if portNumber < 1 || portNumber > 65535 {
		return 0, fmt.Errorf("Port: %d, out of range", portNumber)
	}
	return portNumber, err
}
