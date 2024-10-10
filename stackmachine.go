package main


import(
	"errors"
	"slices"
)

type Stack struct{
	StackNumbers []int
}

func (stack *Stack) Pop() int{
	lastElementIndex := len(stack.StackNumbers) -1
	lastElement := stack.StackNumbers[lastElementIndex]
	stack.StackNumbers = slices.Delete(stack.StackNumbers, lastElementIndex, lastElementIndex)

	return lastElement
}

func StackMachine(commands string)(int, error) {
	return 0, errors.New("")
}

func main() {
	// main is unused - run using 
	// go test ./...
}