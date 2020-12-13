package console

import (
	"fmt"
)

type consoleViewer struct {}


func New() *consoleViewer {
	return &consoleViewer{}
}

func (c *consoleViewer) Print(message string) {
	fmt.Println(message)
}

func (c *consoleViewer) GetNumber() int {
	var number int
	fmt.Scanf("%d\n", &number)

	return number
}
