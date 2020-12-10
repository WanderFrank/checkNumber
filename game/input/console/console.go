package console

import "fmt"

type console struct{}

func New() *console {
	return new(console)
}

func (c *console) GetNumber() int {
	var number int
	fmt.Scanf("%d\n", &number)

	return number
}
