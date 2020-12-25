package console

import (
	"fmt"
)

type consoleViewer struct{}

func New() *consoleViewer {
	return &consoleViewer{}
}

func (c *consoleViewer) Print(message string) error {
	_, err := fmt.Println(message)

	return err
}

func (c *consoleViewer) GetNumber() (int, error) {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return 0, err
	}

	return number, nil
}
