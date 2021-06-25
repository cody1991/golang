package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	render := bufio.NewReader(os.Stdin)
	input, err := render.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, nil
	}

	return number, nil
}
