package main


import(
	"errors"
	"slices"
)

type Stack struct{
	StackNumbers []int
	lastElementIndex int
}

func NewStack() *Stack{
	var emptyStackNumbers []int = []int{}
	return &Stack{
		emptyStackNumbers,
		0,
	}
}


func NewStackFromNumbers(stackNumbers []int) *Stack{
	return &Stack{
		stackNumbers,
		len(stackNumbers) - 1,
	}
}


// func (stack *Stack) Push() {
	

// }

func (stack *Stack) Pop() (int, error){
	lastElementIndex := stack.lastElementIndex

	if (lastElementIndex >= 0){
		lastElement := stack.StackNumbers[lastElementIndex]
		// stack.StackNumbers = slices.Delete(stack.StackNumbers, lastElementIndex, lastElementIndex)
		stack.StackNumbers = stack.StackNumbers[:lastElementIndex]
		stack.lastElementIndex -= 1
		return lastElement, nil
	}else{
		return -1, errors.New("Cannot remove from empty stack.")
	}

}

func StackMachine(commands string)(int, error) {
return 0, errors.New("")
}

func main() {
// main is unused - run using 
// go test ./...
}